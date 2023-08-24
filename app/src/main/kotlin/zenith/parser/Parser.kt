package zenith.parser

import zenith.*

internal typealias NodeResult = Result<Node>

data class ParserResult(val node: Node, val faults: Faults)

fun parseTokens(tokens: List<Token>): ParserResult {
    val result = parseFileStat(Context(tokens))
    return ParserResult(result.value, result.faults)
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

internal class MutableNodeResult(private val type: String) {
    private val children = mutableListOf<Faultable>()
    private val faults = mutableListOf<Fault>()
    val failed get() = faults.any { it.type == 'F' }

    constructor(type: String, obj: Faultable): this(type) { add(obj) }
    constructor(type: String, result: NodeResult): this(type) { add(result) }

    constructor(result: NodeResult): this(result.value.type) {
        children += result.value.children
        faults += result.faults
    }

    fun addAll(faults: Collection<Fault>) { this.faults += faults }
    fun add(fault: Fault) { faults += fault }
    fun add(obj: Faultable) { children += obj }

    @JvmName("addNodeResult")
    fun add(result: NodeResult): Boolean {
        children += result.value
        faults += result.faults
        return result.failed
    }

    @JvmName("addTokenResult")
    fun add(result: Result<Token>): Boolean {
        children += result.value
        faults += result.faults
        return result.failed
    }

    fun error(message: String) = add(Error(Node(type, children), message))
    fun toNodeResult() = NodeResult(Node(type, children), faults)
}
