package zenith

private const val label = "Lexing"

class Lexer(faults: Faults) {
    private val state = LexerState(faults)

    fun makeTokens(): List<Token> {
        ignoreSpacesAndComments()

        while(!state.atEOF) {
            makeToken()
            ignoreSpacesAndComments()
        }

        return state.finish()
    }

    private fun ignoreSpacesAndComments() {
        while(state.nextChar.isWhitespace()) {
            if(state.nextChar == '\n') return
            state.skipChars(1)
        }

        when(state.nextStr(2)) {
            "//" -> state.skipCharsUntil("\n")
            "/*" -> {
                state.initToken()
                state.takeChars(2)

                while(!state.atEOF) {
                    if(state.nextStr(2) == "*/") {
                        state.takeChars(2)
                        return ignoreSpacesAndComments()
                    }

                    state.takeChars(1)
                    state.takeCharsUntil("*")
                }

                state.finishToken()
                throw state.fail("Unterminated multiline comment")
            }
        }
    }

    private fun makeToken() {
        when(state.nextChar) {
            '\n' -> makeNewline()
            in Grammar.NUM_START_SYMS -> makeNum()
            in Grammar.PUNC_SYMS -> makePunc()
            in Grammar.ID_START_SYMS -> makeKeyOrId()
            '\'' -> makeChar()
            '"' -> makeStr()
            else -> {
                state.addErrorToken("${state.nextChar}", state.filePos)
                state.fail("Unrecognized symbol '${state.nextChar}'")
            }
        }
    }

    private fun makeNewline() {
        state.initToken()
        state.takeChars(1)

        if(state.lastToken?.has("\n") ?: false) return
        state.finishToken(Token.Type.PUNC)
    }

    private fun makeNum() {
        if(state.nextStr(2).matches(Grammar.NON_NUM_PATTERN)) return makePunc()

        state.initToken()
        state.takeTheseChars(Grammar.NUM_SYMS)
        state.finishToken(Token.Type.NUM, state.takenStr.replace("_", ""))

        if(state.lastToken!!.string.matches(Grammar.NUM_PATTERN)) return
        state.error("Invalid number")
    }

    private fun makePunc() {
        state.initToken()

        for(puncSize in Grammar.LONGEST_PUNC_SIZE downTo 1) {
            if(state.nextStr(puncSize) in Grammar.PUNCS) {
                state.takeChars(puncSize)
                break
            }
        }

        state.finishToken(Token.Type.PUNC)
    }

    private fun makeKeyOrId() {
        state.initToken()
        state.takeTheseChars(Grammar.ID_SYMS)

        val type = when(state.takenStr) {
            in Grammar.KEYS -> Token.Type.KEY
            else -> Token.Type.ID
        }

        state.finishToken(type)
    }

    private fun makeChar() {
        state.initToken()
        state.takeChars(1)

        when(state.nextChar) {
            '\'' -> {
                state.takeChars(1)
                state.finishToken(Token.Type.CHAR)
                state.error("Empty character")
                return
            }
            '\\' -> {
               state.takeChars(2)
               state.takeCharsUntil("'")
            }
            else -> state.takeChars(1)
        }

        val quote = state.takeChars(1)
        state.finishToken(Token.Type.CHAR)

        if(quote != "'") throw state.fail("Unterminated characrer")
        if(state.takenStr.matches(Grammar.CHAR_PATTERN)) return
        state.error("Invalid character")
    }

    private fun makeStr() {
        val multi = (state.nextStr(3) == "\"\"\"")
        val quotes = if(multi) 3 else 1

        state.initLeadingStrToken()
        state.takeChars(quotes)

        while(!state.atEOF) {
            state.takeCharsUntil(if(multi) "$\"" else "$\"\\\n")

            when(state.nextChar) {
                '$' -> makeStrTemplate()
                '"' -> {
                    if(!multi || state.nextStr(3) == "\"\"\"") {
                        state.takeChars(quotes)
                        state.finishStrToken(true)
                        return
                    }
                }
                '\\' -> if(!multi) makeStrEscapeSeq()
                '\n' -> if(!multi) break
            }
        }

        state.finishStrToken(true)
        val kind = if(multi) "multiline" else ""
        throw state.fail("Unterminated $kind string literal")
    }

    private fun makeStrEscapeSeq() {
        val filePos = state.filePos
        var escape = state.takeChars(2)

        if(escape == "\\u") escape += state.takeChars(4)
        if(escape.matches(Grammar.ESCAPE_PATTERN)) return

        state.addErrorToken(escape, filePos)
        state.error("Invalid escape sequence")
    }

    private fun makeStrTemplate() {
        when(state.nextStr(2).last()) {
            in Grammar.ID_START_SYMS -> {
                state.finishStrToken()
                state.skipChars(1)
                makeKeyOrId()
                state.initToken()
            }
            '{' -> {
                state.finishStrToken()

                val filePos = state.filePos
                state.skipChars(2)
                ignoreSpacesAndComments()

                while(!state.atEOF && state.nextChar != '}') {
                    makeToken()
                    ignoreSpacesAndComments()
                }

                if(state.nextChar != '}') {
                    state.addErrorToken("${'$'}{", filePos)
                    throw state.fail("Unterminated string template")
                }

                state.skipChars(1)
                state.initToken()
            }
            else -> state.takeChars(1)
        }
    }
}

private class LexerState(private val faults: Faults) {
    private val file = SourceFile(faults.filePath)
    private val tokens = mutableListOf<Token>()
    private val builder = StringBuilder()
    private var tokenPos = filePos
    private var atStringStart = false

    val takenStr get() = "$builder"
    val atEOF get() = file.atEnd
    val filePos get() = UNum.new(file.charPos)
    val nextChar get() = file.nextChar
    val lastToken get() = tokens.lastOrNull()

    fun nextStr(n: Int) = file.nextStr(n)

    fun initToken() {
        builder.clear()
        tokenPos = UNum.new(file.charPos)
    }

    fun initLeadingStrToken() {
        atStringStart = true
        initToken()
    }

    fun skipChars(n: Int = -1) = file.readChars(n)
    fun skipTheseChars(these: String) = file.readTheseChars(these)
    fun skipCharsUntil(until: String) = file.readCharsUntil(until)

    fun takeChars(n: Int = -1) = take(file.readChars(n))
    fun takeTheseChars(these: String) = take(file.readTheseChars(these))
    fun takeCharsUntil(until: String) = take(file.readCharsUntil(until))

    fun finishToken(
        type: Token.Type = Token.Type.NONE,
        string: String = takenStr
    ) = tokens.add(Token(type, string, tokenPos))

    fun finishStrToken(isEnd: Boolean = false) {
        if(takenStr.isEmpty()) return

        val type = when {
            atStringStart && isEnd -> Token.Type.STR_FULL
            atStringStart -> Token.Type.STR_START
            isEnd -> Token.Type.STR_END
            else -> Token.Type.STR_MIDDLE
        }

        atStringStart = false
        finishToken(type)
    }

    fun addErrorToken(string: String, position: UNum) {
        tokens.add(Token(Token.Type.NONE, string, position))
    }

    fun warn(message: String) = faults.warn(label, lastToken!!, message)
    fun error(message: String) = faults.error(label, lastToken!!, message)
    fun fail(message: String) = faults.fail(label, lastToken!!, message)

    fun finish(): List<Token> {
        tokens.add(Token(Token.Type.PUNC, Grammar.EOF, filePos))
        file.close()
        return tokens
    }

    private fun take(string: String): String {
        builder.append(string)
        return string
    }
}
