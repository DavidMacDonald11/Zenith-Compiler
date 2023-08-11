package zenith.lexer

import zenith.*

class Result(val tokens: List<Token>, val faults: List<Fault>) {
    constructor(token: Token? = null): this(listOfNotNull(token), listOf())
    constructor(fault: Fault): this(listOf(fault.obj as Token), listOf(fault))

    val failed get() = faults.any { it is Fault.Failure }
    val errored get() = failed || faults.any { it is Fault.Error }
    val isEmpty get() = tokens.isEmpty() && faults.isEmpty()

    operator fun plus(result: Result) =
        Result(tokens + result.tokens, faults + result.faults)

    fun lastTokenHas(string: String) =
        tokens.lastOrNull()?.has(string) ?: false
}

fun tokenizeFile(filePath: String): Result {
    var result = Result()
    val file = SourceFile(filePath)

    while(!file.atEnd) {
        val newResult = makeToken(file, result.lastTokenHas("\n"))
        if(newResult.failed) return result + newResult
        result += newResult
    }

    val EOFToken = newToken(file, Token.Type.PUNC, Grammar.EOF)
    return (result + Result(EOFToken)).also { file.close() }
}

private const val FAULT_LABEL = "Lexing"
private const val MULTI_QUOTE = "\"\"\""

private fun Warning(token: Token, message: String) =
    Fault.Warning(FAULT_LABEL, token, message)

private fun Error(token: Token, message: String) =
    Fault.Error(FAULT_LABEL, token, message)

private fun Failure(token: Token, message: String) =
    Fault.Failure(FAULT_LABEL, token, message)

private fun newToken(file: SourceFile, type: Token.Type, string: String) =
    Token(type, string, UNum.new(file.charPos - string.length.toUInt()))

private fun newFaultToken(file: SourceFile, string: String) =
    newToken(file, Token.Type.NONE, string)

private fun makeToken(file: SourceFile, lastWasNewline: Boolean): Result {
    val result = readSpacesAndComments(file)
    if(result.failed || file.atEnd) return result

    return when(file.nextChar) {
        '\n' -> makeNewline(file, lastWasNewline)
        in Grammar.NUM_START_SYMS -> makeNum(file)
        in Grammar.PUNC_SYMS -> makePunc(file)
        in Grammar.ID_START_SYMS -> makeIdOrKey(file)
        '\'' -> makeChar(file)
        '"' -> makeStr(file)
        else -> {
            val token = newFaultToken(file, file.readChar())
            return Result(Failure(token, "Unrecognized symbol"))
        }
    }
}

private fun readSpacesAndComments(file: SourceFile): Result {
    while(file.nextChar.isWhitespace()) {
        if(file.nextChar == '\n') return Result()
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
            return Result(Failure(token, "Unterminated multiline comment"))
        }
    }

    return Result()
}

private fun makeNewline(file: SourceFile, lastWasNewline: Boolean): Result {
    val token = newToken(file, Token.Type.PUNC, file.readChar())
    return Result(if(lastWasNewline) null else token)
}

private fun makeNum(file: SourceFile): Result {
    if(file.peek(2).matches(Grammar.NON_NUM_PATTERN)) return makePunc(file)

    val string = file.readTheseChars(Grammar.NUM_SYMS)
    val token = newToken(file, Token.Type.NUM, string)
    val num = string.replace("_", "")

    if(num.matches(Grammar.NUM_PATTERN)) return Result(token)
    return Result(Error(token, "Invalid number"))
}

private fun makePunc(file: SourceFile): Result {
    for(puncSize in Grammar.LONGEST_PUNC_SIZE downTo 1) {
        if(file.peek(puncSize) in Grammar.PUNCS) {
            val string = file.readChars(puncSize)
            return Result(newToken(file, Token.Type.PUNC, string))
        }
    }

    val token = newFaultToken(file, file.readChar())
    return Result(Error(token, "Incomplete punctuator"))
}

private fun makeIdOrKey(file: SourceFile): Result {
    val string = file.readTheseChars(Grammar.ID_SYMS)
    val type = if(string !in Grammar.KEYS) Token.Type.ID else Token.Type.KEY
    return Result(newToken(file, type, string))
}

private fun makeChar(file: SourceFile): Result {
    val builder = StringBuilder(file.readChar())

    when(file.nextChar) {
        '\'' -> {
            builder.append(file.readChar())
            val token = newFaultToken(file, "$builder")
            return Result(Error(token, "Empty character"))
        }
        '\\' -> {
            builder.append(file.readChars(2))
            builder.append(file.readCharsUntil("'"))
        }
        else -> builder.append(file.readChar())
    }

    builder.append(file.readChar())
    val charIsClosed = (builder.last() == '\'')
    val token = newToken(file, Token.Type.CHAR, "$builder")

    if(!charIsClosed) return Result(Failure(token, "Unterminated character"))
    if(builder.matches(Grammar.CHAR_PATTERN)) return Result(token)
    return Result(Error(token, "Invaid Character"))
}

private fun makeStr(file: SourceFile): Result {
    val isMulti = (file.peek(3) == MULTI_QUOTE)
    val quoteLength = if(isMulti) 3 else 1
    val stopPattern = if(isMulti) "$\"" else "$\"\\\n"

    var result = Result(newToken(file, Token.Type.STR_START, ""))
    val builder = StringBuilder(file.readChars(quoteLength))

    while(!file.atEnd) {
        builder.append(file.readCharsUntil(stopPattern))

        when(file.nextChar) {
            '$' -> {
                val token = newToken(file, Token.Type.STR, "$builder")
                val templateResult = makeStrTemplate(file)

                if(result.isEmpty) {
                    builder.append(file.readChar())
                    continue
                }

                builder.clear()
                result += Result(token) + templateResult
                if(result.failed) return result
            }
            '"' -> if(!isMulti || file.peek(3) == MULTI_QUOTE) {
                builder.append(file.readChars(quoteLength))

                val tokens = listOf(
                    newToken(file, Token.Type.STR, "$builder"),
                    newToken(file, Token.Type.STR_END, "")
                )

                return result + Result(tokens, listOf())
            }
            '\\' -> {
                val escapeLength = if(file.peek(2) == "\\u") 6 else 2
                val escape = file.readChars(escapeLength)
                builder.append(escape)

                if(escape.matches(Grammar.ESCAPE_PATTERN)) continue
                val token = newFaultToken(file, escape)
                result += Result(Error(token, "Invalid escape sequence"))
            }
            '\n' -> break
        }
    }

    val kind = if(isMulti) " multiline" else ""
    val token = newToken(file, Token.Type.STR, "$builder")
    return result + Result(Failure(token, "Unterminated$kind string literal"))
}

private fun makeStrTemplate(file: SourceFile): Result {
    return when(file.peek(2).last()) {
        in Grammar.ID_START_SYMS -> {
            file.readChar()
            makeIdOrKey(file)
        }
        '{' -> makeStrExprTemplate(file)
        else -> Result()
    }
}

private fun makeStrExprTemplate(file: SourceFile): Result {
    var result = Result()
    val faultToken = newFaultToken(file, file.readChars(2))

    while(!file.atEnd) {
        val newResult = makeToken(file, result.lastTokenHas("\n"))

        when {
            newResult.failed -> return result + newResult
            newResult.lastTokenHas("}") -> return result
            else -> result += newResult
        }
    }

    return result + Result(Failure(faultToken, "Unterminated string template"))
}
