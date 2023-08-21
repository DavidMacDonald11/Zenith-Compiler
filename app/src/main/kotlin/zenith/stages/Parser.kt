package zenith.parser

import zenith.*

internal typealias NodeResult = Result<Node>

fun parseTokens(tokens: List<Token>): NodeResult {
    val ctx = Context(tokens)
    val stats = mutableListOf<Faultable>()
    val faults = mutableListOf<Fault>()

    while(!ctx.next.has(Grammar.EOF)) {
        val result = parseExpr(ctx)
        ctx.skipNewline()

        stats += result.value
        faults += result.faults

        if(result.failed) break
    }

    return NodeResult(Node("FileStat", stats), faults)
}

internal const val LABEL = "Parsing"
internal fun Warning(obj: Faultable, msg: String) = Fault(LABEL, 'W', obj, msg)
internal fun Error(obj: Faultable, msg: String) = Fault(LABEL, 'E', obj, msg)
internal fun Failure(obj: Faultable, msg: String) = Fault(LABEL, 'F', obj, msg)

internal class Context(private val tokens: List<Token>) {
    private var n = 0
    val next get() = peek(0)

    fun peek(n: Int) = tokens[this.n + n]
    fun take() = next.also { n += 1 }
    fun untake() { n -= 1 }
    fun skipNewline() = if(next.has("\n")) take().let { true } else false

    fun expectingHas(vararg strings: String): Result<Token> {
        val next = take()
        if(next.has(*strings)) return Result(next)

        val want = strings.map {
            if(it == Grammar.EOF) "EOF" else it.replace("\n", "\\n")
        }

        return Result(next, Failure(next, "Expecting token has one of $want"))
    }

    fun expectingOf(vararg types: Token.Type): Result<Token> {
        val next = take()
        if(next.of(*types)) return Result(next)
        return Result(next, Failure(next, "Expecting token type of $types"))
    }

    fun expectingHas(c: Collection<String>) = expectingHas(*c.toTypedArray())
    fun expectingOf(c: Collection<Token.Type>) = expectingOf(*c.toTypedArray())
}
