package zenith.lexer

import zenith.*
import zenith.files.SourceFile

private typealias TokenResult = Result<List<Token>>
private fun TokenResult(token: Token? = null) = Result(listOfNotNull(token))
private fun TokenResult(fault: Fault) = Result(listOf(fault.obj as Token), fault)

fun tokenizeFile(filePath: String): TokenResult {
    val tokens = mutableListOf<Token>()
    val faults = mutableListOf<Fault>()
    val file = SourceFile(filePath)

    while(!file.atEnd) {
        val saveNextNewline = !tokens.isEmpty() && !lastHas(tokens, "\n")
        val result = makeToken(file, saveNextNewline)

        tokens += result.value
        faults += result.faults

        if(result.failed) return result
    }

    tokens += newToken(file, Token.Type.PUNC, Grammar.EOF)
    file.close()

    return Result(tokens, faults)
}

private const val LABEL = "Lexing"
private fun Warning(obj: Token, msg: String) = Fault(LABEL, 'W', obj, msg)
private fun Error(obj: Token, msg: String) = Fault(LABEL, 'E', obj, msg)
private fun Failure(obj: Token, msg: String) = Fault(LABEL, 'F', obj, msg)

private fun newToken(file: SourceFile, type: Token.Type, string: String) =
    Token(type, string, file.charPos - string.length.toUInt())

private fun newFaultToken(file: SourceFile, string: String) =
    newToken(file, Token.Type.NONE, string)

private fun lastHas(tokens: List<Token>, string: String): Boolean {
    return tokens.lastOrNull()?.has(string) ?: false
}

private fun makeToken(file: SourceFile, saveNextNewline: Boolean): TokenResult {
    val result = readSpacesAndComments(file)
    if(result.failed || file.atEnd) return result

    return when(file.nextChar) {
        '\n' -> makeNewline(file, saveNextNewline)
        in Grammar.NUM_START_SYMS -> makeNum(file)
        in Grammar.PUNC_SYMS -> makePunc(file)
        in Grammar.ID_START_SYMS -> makeIdOrKey(file)
        '\'' -> makeChar(file)
        '"' -> makeStr(file)
        else -> {
            val token = newFaultToken(file, file.readChar())
            return TokenResult(Failure(token, "Unrecognized symbol"))
        }
    }
}

private fun readSpacesAndComments(file: SourceFile): TokenResult {
    while(file.nextChar.isWhitespace()) {
        if(file.nextChar == '\n') return TokenResult()
        file.readChar()
    }

    when(file.peek(2)) {
        "//" -> file.readCharsUntil("\n")
        "/*" -> {
            val builder = StringBuilder(file.readChars(2))

            while(!file.atEnd) {
                if(file.peek(2) == "*/") {
                    file.readChars(2)
                    return readSpacesAndComments(file)
                }

                builder.append(file.readChar())
                builder.append(file.readCharsUntil("*"))
            }

            val token = newFaultToken(file, "$builder")
            return TokenResult(Failure(token, "Unterminated multiline comment"))
        }
    }

    return TokenResult()
}

private fun makeNewline(file: SourceFile, saveNext: Boolean): TokenResult {
    val token = newToken(file, Token.Type.PUNC, file.readChar())
    return TokenResult(if(saveNext) token else null)
}

private fun makeNum(file: SourceFile): TokenResult {
    if(file.peek(2).matches(Grammar.NON_NUM_PATTERN)) return makePunc(file)

    val string = file.readTheseChars(Grammar.NUM_SYMS)
    val token = newToken(file, Token.Type.NUM, string)
    val num = string.replace("_", "")

    if(num.matches(Grammar.NUM_PATTERN)) return TokenResult(token)
    return TokenResult(Error(token, "Invalid number"))
}

private fun makePunc(file: SourceFile): TokenResult {
    for(puncSize in Grammar.LONGEST_PUNC_SIZE downTo 1) {
        if(file.peek(puncSize) in Grammar.PUNCS) {
            val string = file.readChars(puncSize)
            return TokenResult(newToken(file, Token.Type.PUNC, string))
        }
    }

    val token = newFaultToken(file, file.readChar())
    return TokenResult(Error(token, "Incomplete punctuator"))
}

private fun makeIdOrKey(file: SourceFile): TokenResult {
    val string = file.readTheseChars(Grammar.ID_SYMS)
    val type = if(string !in Grammar.KEYS) Token.Type.ID else Token.Type.KEY
    return TokenResult(newToken(file, type, string))
}

private fun makeChar(file: SourceFile): TokenResult {
    val builder = StringBuilder(file.readChar())

    when(file.nextChar) {
        '\'' -> {
            builder.append(file.readChar())
            val token = newFaultToken(file, "$builder")
            return TokenResult(Error(token, "Empty character"))
        }
        '\\' -> {
            builder.append(file.readChars(2))
            builder.append(file.readCharsUntil("'"))
        }
        else -> builder.append(file.readChar())
    }

    builder.append(file.readChar())
    val charEnded = (builder.last() == '\'')
    val token = newToken(file, Token.Type.CHAR, "$builder")

    if(!charEnded) return TokenResult(Failure(token, "Unterminated character"))
    if(builder.matches(Grammar.CHAR_PATTERN)) return TokenResult(token)
    return TokenResult(Error(token, "Invaid Character"))
}

private const val MULTI_QUOTE = "\"\"\""

private fun makeStr(file: SourceFile): TokenResult {
    val isMulti = (file.peek(3) == MULTI_QUOTE)
    val quoteLength = if(isMulti) 3 else 1
    val stopPattern = if(isMulti) "$\"" else "$\"\\\n"

    val tokens = mutableListOf(newToken(file, Token.Type.STR_START, "<S>"))
    val faults = mutableListOf<Fault>()
    val builder = StringBuilder(file.readChars(quoteLength))

    while(!file.atEnd) {
        builder.append(file.readCharsUntil(stopPattern))

        when(file.nextChar) {
            '$' -> {
                val token = newToken(file, Token.Type.STR, "$builder")
                val result = makeStrTemplate(file)

                if(result.value.isEmpty() && result.faults.isEmpty()) {
                    builder.append(file.readChar())
                    continue
                }

                builder.clear()
                tokens += token
                tokens += result.value
                faults += result.faults

                if(result.failed) return Result(tokens, faults)
            }
            '"' -> if(!isMulti || file.peek(3) == MULTI_QUOTE) {
                builder.append(file.readChars(quoteLength))
                tokens += newToken(file, Token.Type.STR, "$builder")
                tokens += newToken(file, Token.Type.STR_END, "</S>")

                return Result(tokens, faults)
            }
            '\\' -> {
                val escapeLength = if(file.peek(2) == "\\u") 6 else 2
                val escape = file.readChars(escapeLength)
                builder.append(escape)

                if(escape.matches(Grammar.ESCAPE_PATTERN)) continue
                tokens += newFaultToken(file, escape)
                faults += Error(tokens.last(), "Invalid escape sequence")
            }
            '\n' -> break
        }
    }

    val kind = if(isMulti) " multiline" else ""
    tokens += newToken(file, Token.Type.STR, "$builder")
    faults += Failure(tokens.last(), "Unterminated$kind string literal")

    return Result(tokens, faults)
}

private fun makeStrTemplate(file: SourceFile): TokenResult {
    return when(file.peek(2).last()) {
        in Grammar.ID_START_SYMS -> {
            file.readChar()
            makeIdOrKey(file)
        }
        '{' -> makeStrExprTemplate(file)
        else -> TokenResult()
    }
}

private fun makeStrExprTemplate(file: SourceFile): TokenResult {
    val tokens = mutableListOf<Token>()
    val faults = mutableListOf<Fault>()
    val faultToken = newFaultToken(file, file.readChars(2))

    while(!file.atEnd) {
        val result = makeToken(file, false)

        if(lastHas(result.value, "}")) {
            if(tokens.isEmpty())
                faults += Warning(faultToken, "Empty string template")

            return Result(tokens, faults)
        }

        tokens += result.value
        faults += result.faults

        if(result.failed) return Result(tokens, faults)
    }

    tokens += faultToken
    faults += Failure(faultToken, "Unterminated string template")
    return Result(tokens, faults)
}
