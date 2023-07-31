package zenith

private const val label = "Lexing"
private fun fail(token: Token, message: String) = fail(label, token, message)
private fun error(token: Token, message: String) = error(label, token, message)

private data class LexerCtx(val file: SourceFile) {
    private var builder = StringBuilder()
    private var position = file.position
    var atStringStart = false

    val atEnd get() = file.atEnd
    val next get() = file.next
    fun nextString(n: Int) = file.nextString(n)

    fun take(num: Int = -1, these: String = "", until: String = ""): String {
        val taken = file.take(num, these, until)
        builder.append(taken)
        return taken
    }

    fun skip(num: Int = -1, these: String = "", until: String = ""): String {
        return file.take(num, these, until)
    }

    fun initToken() {
        builder.clear()
        position = file.position
    }

    fun newToken(
        type: Token.Type = Token.Type.NONE,
        string: String = "$builder",
        isStringEnd: Boolean = false
    ): Token {
        val realType = when {
            atStringStart && isStringEnd -> Token.Type.STR
            atStringStart -> Token.Type.STR_START
            isStringEnd -> Token.Type.STR_END
            else -> type
        }

        return Token(realType, string, position)
    }
}

fun tokenize(file: SourceFile): List<Token> {
    val ctx = LexerCtx(file)
    val tokens = mutableListOf<Token>()

    while(true) {
        ignoreSpacesAndComments(ctx)
        if(file.atEnd) break

        val token = makeToken(file, tokens.lastOrNull())
        token?.let { tokens += it }
    }

    tokens += Token(Token.Type.PUNC, Token.EOF, file.position)
    return tokens
}

private fun ignoreSpacesAndComments(ctx: LexerCtx) {
    while(ctx.next.isWhitespace()) {
        if(ctx.next == '\n') return
        ctx.skip(1)
    }

    when(file.nextString(2)) {
        "//" -> file.take(until = "\n")
        "/*" -> {
            val (position, builder) = file.tokenInit()

            file.take(2)

            while(!file.atEnd) {
                if(file.nextString(2) != "*/") {
                    file.take(1)
                    file.take(until = "*")
                    continue
                }

                file.take(2)
                return ignoreSpacesAndComments(file)
            }

            val token = Token(Token.Type.NONE, "$builder", position)
            throw fail(token, "Unterminated multiline comment")
        }
        else -> return
    }
}

private fun makeToken(file: SourceFile, lastToken: Token?): Token? {
    return when(file.next) {
        '\n' -> makeNewline(file, lastToken)
        in Token.NUM_START_SYMS -> makeNum(file)
        in Token.PUNC_SYMS -> makePunc(file)
        in Token.ID_SYMS -> makeIdOrKey(file)
        '\'' -> makeChar(file)
        else -> {
            val (position, builder) = file.tokenInit()
            file.take(1)

            val token = Token(Token.Type.NONE, "$builder", position)
            throw fail(token, "Unrecognized symbol '${token.string}'")
        }
    }
}

private fun makeNewline(file: SourceFile, lastToken: Token?): Token? {
    val (position, builder) = file.tokenInit()
    file.take(1)

    if(lastToken != null && lastToken.has("\n")) return null
    return Token(Token.Type.PUNC, "$builder", position)
}

private fun makePunc(file: SourceFile): Token {
    val (position, builder) = file.tokenInit()

    for(i in Token.LONGEST_PUNC downTo 1) {
        if(file.nextString(i) in Token.PUNCS) {
            file.take(i)
            break
        }
    }

    return Token(Token.Type.PUNC, "$builder", position)
}

private fun makeNum(file: SourceFile): Token {
    val (position, builder) = file.tokenInit()

    if(file.nextString(2).matches("\\.[^\\d]".toRegex()))
        return makePunc(file)

    file.take(these = Token.NUM_SYMS)
    val token = Token(Token.Type.NUM, "$builder".replace("_", ""), position)

    if(!token.string.matches(Token.NUM_PATTERN))
        error(token, "Invalid number")

    return token
}

private fun makeIdOrKey(file: SourceFile): Token {
    val (position, builder) = file.tokenInit()
    file.take(these = Token.ID_SYMS)

    val string = "$builder"
    val type = if(string in Token.KEYS) Token.Type.KEY else Token.Type.ID

    return Token(type, string, position)
}

private fun makeChar(file: SourceFile): Token {
    val (position, builder) = file.tokenInit()
    file.take(1)

    if(file.next == '\'') {
        file.take(1)
        val token = Token(Token.Type.CHAR, "$builder", position)

        error(token, "Empty character")
        return token
    }

    if(file.next == '\\') {
        file.take(2)
        file.take(until = "'")
    } else file.take(1)

    val quote = file.take(1)
    val token = Token(Token.Type.CHAR, "$builder", position)

    if(quote != "'") throw fail(token, "Unterminated character")

    if(!token.string.matches(Token.CHAR_PATTERN))
        error(token, "Invalid character")

    return token
}

private data class StrCtx(val file: SourceFile) {
    val (position, builder) = file.tokenInit()
    val firstToken = true
}
