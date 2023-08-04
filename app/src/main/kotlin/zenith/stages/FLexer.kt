package zenith

sealed interface FLexerResult {
    data class Valid(val tokens: Tokens, val faults: Faults): FLexerResult
    data class Invalid(val tokens: Tokens, val faults: Faults): FLexerResult
}

fun tokenize(filePath: String): FLexerResult {
    val tokens = mutableListOf<Token>()
    val faults = mutableListOf<Fault>()
    val file = SourceFile(filePath)

    while(!file.atEnd) {
        makeToken(file, tokens.lastOrNull()).let { result ->
            when(result) {
                is Result.None -> Unit
                is Result.One -> tokens.add(result.token)
                is Result.OneFault -> {
                    tokens.add(result.token)
                    faults.add(result.fault)
                }
                is Result.Many -> {
                    tokens.addAll(result.tokens)
                    faults.addAll(result.faults)
                }
                is Result.Failure -> {
                    faults.add(result.fault)
                    return FLexerResult.Invalid(tokens, faults)
                }
            }
        }
    }

    tokens.add(Token(Token.Type.PUNC, Grammar.EOF, UNum.new(file.charPos)))
    file.close()

    if(faults.any { it is Fault.Error })
        return FLexerResult.Invalid(tokens, faults)

    return FLexerResult.Valid(tokens, faults)
}

private sealed interface Result {
    class None(): Result
    class One(val token: Token): Result
    class OneFault(val fault: Fault): Result {
        val token get() = fault.obj as Token
    }

    class Many(val tokens: Tokens, val faults: Faults = emptyList()): Result
    class Failure(token: Token, message: String): Result {
        val fault = Fault.Failure(LABEL, token, message)
    }
}

private fun makeToken(file: SourceFile, lastToken: Token?): Result {
    readSpacesAndComments(file).let { result ->
        when(result) {
            is Result.Failure -> return result
            else -> if(file.atEnd) return Result.None()
        }
    }

    return when(file.nextChar) {
        '\n' -> makeNewline(file, lastToken)
        in Grammar.NUM_START_SYMS -> makeNum(file)
        in Grammar.PUNC_SYMS -> makePunc(file)
        in Grammar.ID_START_SYMS -> makeIdOrKey(file)
        else -> {
            val token = newErrorToken(file, file.readChar())
            return Result.Failure(token, "Unrecognized symbol")
        }
    }
}

private fun readSpacesAndComments(file: SourceFile): Result {
    while(file.nextChar.isWhitespace()) {
        if(file.nextChar == '\n') return Result.None()
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

            val result = newErrorToken(file, "$builder")
            return Result.Failure(result, "Unterminated multiline comment")
        }
    }

    return Result.None()
}

private fun makeNewline(file: SourceFile, lastToken: Token?): Result {
    val string = file.readChar()
    if(lastToken?.has("\n") ?: false) return Result.None()
    return Result.One(newToken(file, Token.Type.PUNC, string))
}

private fun makeNum(file: SourceFile): Result {
    if(file.peek(2).matches(Grammar.NON_NUM_PATTERN)) return makePunc(file)

    val (token, string) = file.readTheseChars(Grammar.NUM_SYMS).let {
        Pair(newToken(file, Token.Type.NUM, it), it.replace("_", ""))
    }

    if(string.matches(Grammar.NUM_PATTERN)) return Result.One(token)
    return Result.OneFault(Error(token, "Invalid number"))
}

private fun makePunc(file: SourceFile): Result {
    for(puncSize in Grammar.LONGEST_PUNC_SIZE downTo 1) {
        if(file.peek(puncSize) in Grammar.PUNCS) {
            val string = file.readChars(puncSize)
            return Result.One(newToken(file, Token.Type.PUNC, string))
        }
    }

    val errorToken = newErrorToken(file, file.readChar())
    return Result.OneFault(Error(errorToken, "Incomplete punctuator"))
}

private fun makeIdOrKey(file: SourceFile): Result {
    val string = file.readTheseChars(Grammar.ID_SYMS)
    val type = if(string in Grammar.KEYS) Token.Type.KEY else Token.Type.ID
    return Result.One(newToken(file, type, string))
}

private fun makeChar(file: SourceFile): Result {
    val builder = StringBuilder(file.readChar())

    when(file.nextChar) {
        '\'' -> {
            builder.append(file.readChar())
            val token = newErrorToken(file, "$builder")
            return Result.OneFault(Error(token, "Empty character"))
        }
        '\\' -> {
            builder.append(file.readChars(2))
            builder.append(file.readCharsUntil("'"))
        }
        else -> builder.append(file.readChar())
    }

    val lastChar = file.readChar()
    builder.append(lastChar)
    val token = newToken(file, Token.Type.CHAR, "$builder")

    if(lastChar != "'") return Result.Failure(token, "Unterminated character")
    if("$builder".matches(Grammar.CHAR_PATTERN)) return Result.One(token)
    return Result.OneFault(Error(token, "Invalid character"))
}

private fun makeStr(file: SourceFile): Result {
    val isMulti = (file.peek(3) == MULTI_QUOTE)
    val quoteLength = if(isMulti) 3 else 1
    val stopPattern = if(isMulti) "$\"" else "$\"\\\n"

    val tokens = mutableListOf<Token>()
    val faults = mutableListOf<Fault>()
    val builder = StringBuilder(file.readChars(quoteLength))
    var isStart = true
    var charPos = file.charPos

    while(!file.atEnd) {
        builder.append(file.readCharsUntil(stopPattern))

        when(file.nextChar) {
            '$' -> makeStrTemplate(file).let { result ->
                when(result) {
                    is Result.None -> builder.append(file.readChar())
                    is Result.Many -> {
                        if(builder.isNotEmpty()) {
                            tokens.add(newStrToken("$builder", charPos, isStart))
                            isStart = false
                            charPos = file.charPos
                        }

                        tokens.addAll(result.tokens)
                        faults.addAll(result.faults)
                    }
                    else -> throw IllegalArgumentException()
                }
            }
            '"' -> if(!isMulti || file.peek(3) == MULTI_QUOTE) {
                builder.append(file.readChars(quoteLength))
                tokens.add(newStrToken(file, "$builder", isStart, true))
                return Result.Many(tokens, faults)
            }
            '\\' -> {
                val escapeSize = if(file.peek(2) == "\\u") 6 else 2
                val escape = file.readChars(escapeSize)
                builder.append(escape)

                if(escape.matches(Grammar.ESCAPE_PATTERN)) continue

                val token = newErrorToken(file, escape)
                faults.add(Error(token, "Invalid escape sequence"))
            }
            '\n' -> break
        }
    }

    val token = newErrorToken("$builder", charPos)
    val kind = if(isMulti) "multiline" else ""
    return Result.Failure(token, "Unterminated $kind string literal")
}

private fun makeStrTemplate(file: SourceFile): Result {
    return Result.None()
}

private typealias Tokens = List<Token>
private typealias Faults = List<Fault>

private const val LABEL = "Lexing"
private const val MULTI_QUOTE = "\"\"\""

private fun newErrorToken(string: String, position: UInt) =
    Token(Token.Type.NONE, string, UNum.new(position))

private fun newErrorToken(file: SourceFile, string: String) =
    newToken(file, Token.Type.NONE, string)

private fun newToken(type: Token.Type, string: String, position: UInt) =
    Token(type, string, UNum.new(position))

private fun newToken(file: SourceFile, type: Token.Type, string: String) =
    Token(type, string, UNum.new(file.charPos - string.length.toUInt()))

private fun newStrToken(
    string: String,
    position: UInt,
    isStart: Boolean = false,
    isEnd: Boolean = false
): Token {
    val type = when {
        isStart && isEnd -> Token.Type.STR_FULL
        isStart -> Token.Type.STR_START
        isEnd -> Token.Type.STR_END
        else -> Token.Type.STR_MIDDLE
    }

    return newToken(type, string, position)
}

private fun newStrToken(
    file: SourceFile,
    string: String,
    isStart: Boolean = false,
    isEnd: Boolean = false
) = newStrToken(string, file.charPos - string.length.toUInt(), isStart, isEnd)

private fun Warning(token: Token, message: String) =
    Fault.Warning(LABEL, token, message)

private fun Error(token: Token, message: String) =
    Fault.Error(LABEL, token, message)
