package zenith.lexer

import zenith.*
import zenith.io.SourceFile

interface Result: FaultResult {
    val tokens: List<Token>
    val isNone get() = tokens.isEmpty()

    fun lastHas(string: String) = tokens.lastOrNull()?.has(string) ?: false

    private class FinalResult(
        override val tokens: List<Token>,
        override val faults: List<Fault>
    ): Result {
        constructor(token: Token? = null, fault: Fault? = null):
            this(listOfNotNull(token), listOfNotNull(fault))
    }

    companion object {
        internal val LABEL = "Lexing"

        fun none(): Result = FinalResult()
        fun token(token: Token): Result = FinalResult(token)
        fun warn(token: Token, message: String) = fault('W', token, message)
        fun error(token: Token, message: String) = fault('E', token, message)
        fun fail(token: Token, message: String) = fault('F', token, message)

        private fun fault(type: Char, token: Token, message: String): Result =
            FinalResult(token, Fault(LABEL, type, token, message))
    }
}

internal class MutableResult(
    override val tokens: MutableList<Token> = mutableListOf<Token>(),
    override val faults: MutableList<Fault> = mutableListOf<Fault>()
): Result {
    operator fun plusAssign(result: Result) {
        tokens += result.tokens
        faults += result.faults
    }

    operator fun plusAssign(token: Token) { tokens += token }
    fun addWarning(token: Token, message: String) { fault('W', token, message) }
    fun addError(token: Token, message: String) { fault('E', token, message) }
    fun addFailure(token: Token, message: String) { fault('F', token, message) }

    private fun fault(type: Char, token: Token, message: String) {
        tokens += token
        faults += Fault(Result.LABEL, type, token, message)
    }
}

fun tokenizeFile(filePath: String): Result {
    val result = MutableResult()
    val file = SourceFile(filePath)

    while(!file.atEnd) {
        val saveNextNewline = !result.isNone && !result.lastHas("\n")
        result += makeToken(file, saveNextNewline)
        if(result.failed) return result
    }

    result += newToken(file, Token.Type.PUNC, Grammar.EOF)
    file.close()

    return result
}

internal fun newToken(file: SourceFile, type: Token.Type, string: String) =
    Token(type, string, file.charPos - string.length.toUInt())

internal fun newFaultToken(file: SourceFile, string: String) =
    newToken(file, Token.Type.NONE, string)

internal fun makeToken(file: SourceFile, saveNextNewline: Boolean): Result {
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
            return Result.fail(token, "Unrecognized symbol")
        }
    }
}

internal fun readSpacesAndComments(file: SourceFile): Result {
    while(file.nextChar.isWhitespace()) {
        if(file.nextChar == '\n') return Result.none()
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
            return Result.fail(token, "Unterminated multiline comment")
        }
    }

    return Result.none()
}

internal fun makeNewline(file: SourceFile, saveNext: Boolean): Result {
    val token = newToken(file, Token.Type.PUNC, file.readChar())
    return if(saveNext) Result.token(token) else Result.none()
}

internal fun makeNum(file: SourceFile): Result {
    if(file.peek(2).matches(Grammar.NON_NUM_PATTERN)) return makePunc(file)

    val string = file.readTheseChars(Grammar.NUM_SYMS)
    val token = newToken(file, Token.Type.NUM, string)
    val num = string.replace("_", "")

    if(num.matches(Grammar.NUM_PATTERN)) return Result.token(token)
    return Result.error(token, "Invalid number")
}

internal fun makePunc(file: SourceFile): Result {
    for(puncSize in Grammar.LONGEST_PUNC_SIZE downTo 1) {
        if(file.peek(puncSize) in Grammar.PUNCS) {
            val string = file.readChars(puncSize)
            return Result.token(newToken(file, Token.Type.PUNC, string))
        }
    }

    val token = newFaultToken(file, file.readChar())
    return Result.error(token, "Incomplete punctuator")
}

internal fun makeIdOrKey(file: SourceFile): Result {
    val string = file.readTheseChars(Grammar.ID_SYMS)
    val type = if(string !in Grammar.KEYS) Token.Type.ID else Token.Type.KEY
    return Result.token(newToken(file, type, string))
}

internal fun makeChar(file: SourceFile): Result {
    val builder = StringBuilder(file.readChar())

    when(file.nextChar) {
        '\'' -> {
            builder.append(file.readChar())
            val token = newFaultToken(file, "$builder")
            return Result.error(token, "Empty character")
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

    if(!charEnded) return Result.fail(token, "Unterminated character")
    if(builder.matches(Grammar.CHAR_PATTERN)) return Result.token(token)
    return Result.error(token, "Invaid Character")
}

private const val MULTI_QUOTE = "\"\"\""

internal fun makeStr(file: SourceFile): Result {
    val isMulti = (file.peek(3) == MULTI_QUOTE)
    val quoteLength = if(isMulti) 3 else 1
    val stopPattern = if(isMulti) "$\"" else "$\"\\\n"

    val builder = StringBuilder(file.readChars(quoteLength))
    val result = MutableResult()
    result += newToken(file, Token.Type.STR_START, "<S>")

    while(!file.atEnd) {
        builder.append(file.readCharsUntil(stopPattern))

        when(file.nextChar) {
            '$' -> {
                val token = newToken(file, Token.Type.STR, "$builder")
                val template = makeStrTemplate(file)

                if(result.isNone) {
                    builder.append(file.readChar())
                    continue
                }

                builder.clear()
                result += token
                result += template

                if(result.failed) return result
            }
            '"' -> if(!isMulti || file.peek(3) == MULTI_QUOTE) {
                builder.append(file.readChars(quoteLength))
                result += newToken(file, Token.Type.STR, "$builder")
                result += newToken(file, Token.Type.STR_END, "</S>")

                return result
            }
            '\\' -> {
                val escapeLength = if(file.peek(2) == "\\u") 6 else 2
                val escape = file.readChars(escapeLength)
                builder.append(escape)

                if(escape.matches(Grammar.ESCAPE_PATTERN)) continue

                val token = newFaultToken(file, escape)
                result.addError(token, "Invalid escape sequence")
            }
            '\n' -> break
        }
    }

    val token = newToken(file, Token.Type.STR, "$builder")
    val kind = if(isMulti) " multiline" else ""
    result.addFailure(token, "Unterminated$kind string literal")

    return result
}

internal fun makeStrTemplate(file: SourceFile): Result {
    return when(file.peek(2).last()) {
        in Grammar.ID_START_SYMS -> {
            file.readChar()
            makeIdOrKey(file)
        }
        '{' -> makeStrExprTemplate(file)
        else -> Result.none()
    }
}

internal fun makeStrExprTemplate(file: SourceFile): Result {
    val result = MutableResult()
    val faultToken = newFaultToken(file, file.readChars(2))

    while(!file.atEnd) {
        val newTokens = makeToken(file, false)

        if(!newTokens.lastHas("}")) {
            result += newTokens
            if(result.failed) return result
            continue
        }

        if(result.isNone) result.addWarning(faultToken, "Empty string template")
        return result
    }

    result.addFailure(faultToken, "Unterminated string template")
    return result
}
