package zenith.parser

import zenith.*

private typealias NodeResult = Result<Node>

fun parseTokens(tokens: List<Token>): NodeResult {
    val ctx = Context(tokens)
    val stats = mutableListOf<Faultable>()
    val faults = mutableListOf<Fault>()

    while(!ctx.next.has(Grammar.EOF, "\n")) {
        val result = parseExpr(ctx)

        stats += result.value
        faults += result.faults

        if(result.failed) break
    }

    val node = Node("FileStat", stats)
    return NodeResult(node, faults)
}

private class Context(private val tokens: List<Token>) {
    private var n = 0
    val next get() = peek(0)

    fun peek(n: Int) = tokens[this.n + n]
    fun take() = next.also { n += 1 }
    fun untake() { n -= 1 }
    fun skipNewline() = if(next.has("\n")) take().let { true } else false

    fun expectingHas(vararg strings: String) {

    }
}

private const val LABEL = "Parsing"
private fun Warning(obj: Faultable, msg: String) = Fault(LABEL, 'W', obj, msg)
private fun Error(obj: Faultable, msg: String) = Fault(LABEL, 'E', obj, msg)
private fun Failure(obj: Faultable, msg: String) = Fault(LABEL, 'F', obj, msg)

private tailrec fun parseLeftBinaryExpr(
    ctx: Context,
    type: String,
    parseChild: (Context) -> NodeResult,
    opList: List<String>,
    left: NodeResult = parseChild(ctx)
): NodeResult {
    if(left.failed || !ctx.next.has(opList)) return left

    val op = ctx.take()
    ctx.skipNewline()

    val right = parseChild(ctx)
    val node = Node(type, listOf(left.value, op, right.value))
    val result = NodeResult(node, left.faults + right.faults)

    return parseLeftBinaryExpr(ctx, type, parseChild, opList, result)
}

private fun parseExpr(ctx: Context): NodeResult {
    return parseBitOrExpr(ctx)
}

private fun parseBitOrExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "BitOrExpr", ::parseBitXorExpr, listOf("|")
)

private fun parseBitXorExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "BitXorExpr", ::parseBitAndExpr, listOf("$")
)

private fun parseBitAndExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "BitAndExpr", ::parseBitShiftExpr, listOf("&")
)

private fun parseBitShiftExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "BitShiftExpr", ::parseAddExpr, listOf("<<", ">>")
)

private fun parseAddExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "AddExpr", ::parseMultiplyExpr, listOf("+", "-")
)

private fun parseMultiplyExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "MultiplyExpr", ::parsePrefixExpr, listOf("*", "/", "%")
)

private fun parseExponentExpr(ctx: Context): NodeResult {
    val left = parsePrefixExpr(ctx)
    if(left.failed || !ctx.next.has("^")) return left

    val op = ctx.take()
    ctx.skipNewline()

    val right = parseExponentExpr(ctx)
    val node = Node("ExponentExpr", listOf(left.value, op, right.value))

    return NodeResult(node, left.faults + right.faults)
}

private fun parsePrefixExpr(ctx: Context): NodeResult {
    val opList = listOf("+", "-", "~", "$", "&", "#", "not")
    if(!ctx.next.has(opList)) return parsePostfixExpr(ctx)

    val op = ctx.take()
    val right = parsePostfixExpr(ctx)
    val node = Node("PrefixExpr", listOf(op, right.value))

    return NodeResult(node, right.faults)
}

private fun parsePostfixExpr(
    ctx: Context,
    left: NodeResult = parsePrimaryExpr(ctx)
): NodeResult {
    if(left.failed) return left
    val type = "PostfixExpr"

    if(ctx.next.has("!")) {
        val node = Node(type, listOf(left.value, ctx.take()))
        return parsePostfixExpr(ctx, NodeResult(node, left.faults))
    }

    val skipped = ctx.skipNewline()

    if(ctx.next.has(".")) {
        val op = ctx.take()
        ctx.skipNewline()

        // expecting

        return left
    }

    if(skipped) ctx.untake()
    return left
}

private fun parsePrimaryExpr(ctx: Context): NodeResult {
    val type = "PrimaryExpr"

    if(ctx.next.of(Token.Type.NUM)) {
        val node = Node(type, listOf(ctx.take()))
        return NodeResult(node)
    }

    val node = Node(type, listOf(ctx.take()))
    val fault = Failure(node, "Unexpected token in primary expression")
    return NodeResult(node, listOf(fault))
}
