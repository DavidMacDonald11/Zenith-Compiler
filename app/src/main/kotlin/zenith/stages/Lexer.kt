package zenith

private const val LABEL = "Lexing"
private const val MULTI_QUOTE = "\"\"\""

class Lexer(private val faults: Faults) {
    private val file = SourceFile(faults.filePath)
    private val newToken = TokenBuilder()
    private val tokens = mutableListOf<Token>()

    private val lastToken get() = tokens.lastOrNull()

    fun makeTokens(): List<Token> {
        ignoreSpacesAndComments()

        while(!file.atEnd) {
            makeToken()
            ignoreSpacesAndComments()
        }

        tokens.add(Token(Token.Type.PUNC, Grammar.EOF, UNum.new(file.charPos)))
        file.close()

        return tokens
    }

    private fun ignoreSpacesAndComments() {
        while(file.nextChar.isWhitespace()) {
            if(file.nextChar == '\n') return
            file.readChar()
        }

        when(file.peek(2)) {
            "//" -> file.readCharsUntil("\n")
            "/*" -> {
                newToken.init(Token.Type.NONE, file.charPos)
                newToken.add(file.readChars(2))

                while(!file.atEnd) {
                    if(file.peek(2) == "*/") {
                        file.readChars(2)
                        return ignoreSpacesAndComments()
                    }

                    newToken.add(file.readChar())
                    newToken.add(file.readCharsUntil("*"))
                }

                add(newToken.finish())
                throw fail("Unterminated multiline comment")
            }
        }
    }

    private fun makeToken() {
        when(file.nextChar) {
            '\n' -> makeNewline()
            in Grammar.NUM_START_SYMS -> makeNum()
            in Grammar.PUNC_SYMS -> makePunc()
            in Grammar.ID_START_SYMS -> makeIdOrKey()
            '\'' -> makeChar()
            '"' -> makeStr()
            else -> {
                addNewErrorToken(file.readChar(), file.charPos)
                throw fail("Unrecognized symbol")
            }
        }
    }

    private fun makeNewline() {
        newToken.init(Token.Type.PUNC, file.charPos)
        newToken.add(file.readChar())

        if(lastToken?.has("\n") ?: false) return
        add(newToken.finish())
    }

    private fun makeNum() {
        if(file.peek(2).matches(Grammar.NON_NUM_PATTERN)) return makePunc()

        newToken.init(Token.Type.NUM, file.charPos)
        newToken.add(file.readTheseChars(Grammar.NUM_SYMS))
        add(newToken.finish())

        val replacedString = lastToken!!.string.replace("_", "")
        if(replacedString.matches(Grammar.NUM_PATTERN)) return
        error("Invalid number")
    }

    private fun makePunc() {
        newToken.init(Token.Type.PUNC, file.charPos)

        for(puncSize in Grammar.LONGEST_PUNC_SIZE downTo 1) {
            if(file.peek(puncSize) in Grammar.PUNCS) {
                newToken.add(file.readChars(puncSize))
                break
            }
        }

        if(newToken.addedChars.isEmpty()) {
            addNewErrorToken(file.readChar(), file.charPos - 1u)
            error("Incomplete punctuator")
            return
        }

        add(newToken.finish())
    }

    private fun makeIdOrKey() {
        newToken.init(Token.Type.ID, file.charPos)
        newToken.add(file.readTheseChars(Grammar.ID_SYMS))

        if(newToken.addedChars !in Grammar.KEYS) add(newToken.finish())
        else add(newToken.finish(Token.Type.KEY))
    }

    private fun makeChar() {
        newToken.init(Token.Type.CHAR, file.charPos)
        newToken.add(file.readChar())

        when(file.nextChar) {
            '\'' -> {
                newToken.add(file.readChar())
                add(newToken.finish())
                error("Empty character")
                return
            }
            '\\' -> {
                newToken.add(file.readChars(2))
                newToken.add(file.readCharsUntil("'"))
            }
            else -> newToken.add(file.readChar())
        }

        val lastChar = file.readChar()
        newToken.add(lastChar)
        add(newToken.finish())

        if(lastChar != "'") throw fail("Unterminated character")
        if(newToken.addedChars.matches(Grammar.CHAR_PATTERN)) return
        error("Invalid character")
    }

    private fun makeStr() {
        val isMulti = (file.peek(3) == MULTI_QUOTE)
        val quoteLength = if(isMulti) 3 else 1
        val stopPattern = if(isMulti) "$\"" else "$\"\\\n"

        newToken.init(Token.Type.STR_START, file.charPos)
        newToken.add(file.readChars(quoteLength))

        while(!file.atEnd) {
            newToken.add(file.readCharsUntil(stopPattern))

            when(file.nextChar) {
                '$' -> makeStrTemplate()
                '"' -> {
                    if(!isMulti || file.peek(3) == MULTI_QUOTE) {
                        newToken.add(file.readChars(quoteLength))
                        add(newToken.finishStr(isEnd = true))
                        return
                    }
                }
                '\\' -> {
                    val charPos = file.charPos
                    val escapeSize = if(file.peek(2) == "\\u") 6 else 2
                    val escape = file.readChars(escapeSize)

                    newToken.add(escape)
                    if(escape.matches(Grammar.ESCAPE_PATTERN)) continue

                    addNewErrorToken(escape, charPos)
                    error("Invalid escape sequence")
                }
                '\n' -> break
            }
        }

        add(newToken.finishStr(isEnd = true))
        val kind = if(isMulti) "multiline" else ""
        throw fail("Unterminated $kind string literal")
    }

    private fun makeStrTemplate() {
        when(file.peek(2).last()) {
            in Grammar.ID_START_SYMS -> {
                if(!newToken.atInit) add(newToken.finishStr())
                file.readChar()
                makeIdOrKey()
                newToken.init(Token.Type.STR_MIDDLE, file.charPos)
            }
            '{' -> {
                if(!newToken.atInit) add(newToken.finishStr())

                val charPos = file.charPos
                file.readChars(2)
                ignoreSpacesAndComments()

                while(!file.atEnd && file.nextChar != '}') {
                    makeToken()
                    ignoreSpacesAndComments()
                }

                if(file.nextChar != '}') {
                    addNewErrorToken("\${", charPos)
                    throw fail("Unterminated string template")
                }

                file.readChar()
                newToken.init(Token.Type.STR_MIDDLE, file.charPos)
            }
            else -> newToken.add(file.readChar())
        }
    }

    private fun add(token: Token) = tokens.add(token)
    private fun addNewErrorToken(string: String, position: UInt) =
        add(Token(Token.Type.NONE, string, UNum.new(position)))

    private fun warn(msg: String) = faults.warn(LABEL, lastToken!!, msg)
    private fun error(msg: String) = faults.error(LABEL, lastToken!!, msg)
    private fun fail(msg: String) = faults.fail(LABEL, lastToken!!, msg)
}

private class TokenBuilder() {
    private val builder = StringBuilder()
    private var position = 0u
    private lateinit var initType: Token.Type

    val addedChars get() = "$builder"
    val atInit get() = builder.isEmpty()

    fun init(type: Token.Type, filePos: UInt) {
        initType = type
        position = filePos
        builder.clear()
    }

    fun add(str: String) { builder.append(str) }

    fun finish(type: Token.Type? = null): Token {
        val token = Token(type ?: initType, addedChars, UNum.new(position))
        return token
    }

    fun finishStr(isEnd: Boolean = false): Token {
        val isStart = (initType == Token.Type.STR_START)
        val type = when {
            isStart && isEnd -> Token.Type.STR_FULL
            isStart -> Token.Type.STR_START
            isEnd -> Token.Type.STR_END
            else -> Token.Type.STR_MIDDLE
        }

        return finish(type)
    }
}
