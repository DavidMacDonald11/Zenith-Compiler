package zenith.parser

import zenith.*

class Result(val node: Node, val faults: List<Fault>) {
    val failed get() = faults.any { it is Fault.Failure }
    val errored get() = faults.any { it is Fault.Failure || it is Fault.Error }
}

fun parseTokens(tokens: List<Token>): Result {
    val stats = mutableListOf<Faultable>()
    val faults = mutableListOf<Fault>()
    var ctx = Context(tokens)

    while(!ctx.tokens[0].has(Grammar.EOF, "\n")) {
        val (context, result) = parseExpr(ctx)
        ctx = context

        stats += result.node
        faults += result.faults

        if(result.failed) break
    }

    val node = Node("FileStat", stats)
    return Result(node, faults)
}

private class Context(val tokens: List<Token>) {
    private var n = 0
    val next get() = tokens[n]

    fun peek(n: Int) = tokens[this.n + n]
    fun take() = next.also { n += 1 }
    fun untake() { n -= 1 }
    fun skipNewline() = if(next.has("\n")) take() else null
    fun done() = Context(tokens.drop(n))
}

private data class ContextResult(val ctx: Context, val result: Result) {
    constructor(ctx: Context, node: Node, faults: List<Fault> = listOf()):
        this(ctx, Result(node, faults))
}

private const val FAULT_LABEL = "Parsing"

private fun Warning(node: Node, message: String) =
    Fault.Warning(FAULT_LABEL, node, message)

private fun Error(node: Node, message: String) =
    Fault.Error(FAULT_LABEL, node, message)

private fun Failure(node: Node, message: String) =
    Fault.Failure(FAULT_LABEL, node, message)

private tailrec fun parseLeftRecBinaryExpr(
    type: String,
    parseChild: (Context) -> ContextResult,
    opList: List<String>,
    ctx: Context,
    left: ContextResult = parseChild(ctx)
): ContextResult {
    if(left.result.failed) return left
    if(!left.ctx.next.has(opList)) return left

    val op = left.ctx.take()
    left.ctx.skipNewline()
    val right = parseChild(left.ctx.done())
    val faults = left.result.faults + right.result.faults

    val node = Node(type, listOf(left.result.node, op, right.result.node))
    val ctxResult = ContextResult(right.ctx.done(), node, faults)

    return parseLeftRecBinaryExpr(type, parseChild, opList, ctx, ctxResult)
}

private fun parseExpr(ctx: Context): ContextResult {
    return parseBitOrExpr(ctx)
}

private fun parseBitOrExpr(ctx: Context): ContextResult {
    val type = "BitOrExpr"
    return parseLeftRecBinaryExpr(type, ::parseBitXorExpr, listOf("|"), ctx)
}

private fun parseBitXorExpr(ctx: Context): ContextResult {
    val type = "BitXorExpr"
    return parseLeftRecBinaryExpr(type, ::parseBitAndExpr, listOf("$"), ctx)
}

private fun parseBitAndExpr(ctx: Context): ContextResult {
    val type = "BitAndExpr"
    return parseLeftRecBinaryExpr(type, ::parseBitShiftExpr, listOf("&"), ctx)
}

private fun parseBitShiftExpr(ctx: Context): ContextResult {
    val type = "BitShiftExpr"
    val opList = listOf("<<", ">>")
    return parseLeftRecBinaryExpr(type, ::parseAdditiveExpr, opList, ctx)
}

private fun parseAdditiveExpr(ctx: Context): ContextResult {
    val type = "AdditiveExpr"
    val opList = listOf("+", "-")
    return parseLeftRecBinaryExpr(type, ::parseMultiplicativeExpr, opList, ctx)
}

private fun parseMultiplicativeExpr(ctx: Context): ContextResult {
    val type = "MultiplicativeExpr"
    val opList = listOf("*", "/", "%")
    return parseLeftRecBinaryExpr(type, ::parsePrimaryExpr, opList, ctx)
}

private fun parsePrefixExpr(ctx: Context): ContextResult {
    val opList = listOf("+", "-", "~", "$", "&", "#", "not")
    if(!ctx.next.has(opList)) return parsePostfixExpr(ctx)

    val op = ctx.take()
    val right = parsePostfixExpr(ctx.done())
    val node = Node("PrefixExpr", listOf(op, right.result.node))

    return ContextResult(right.ctx.done(), node, right.result.faults)
}

private fun parsePostfixExpr(ctx: Context): ContextResult {
    return parsePrimaryExpr(ctx)
}

private fun parsePrimaryExpr(ctx: Context): ContextResult {
    val type = "PrimaryExpr"

    if(ctx.next.of(Token.Type.NUM)) {
        val node = Node(type, listOf(ctx.take()))
        return ContextResult(ctx.done(), node)
    }

    val node = Node(type, listOf(ctx.take()))
    val fault = Failure(node, "Unexpected token in primary expression")
    return ContextResult(ctx.done(), node, listOf(fault))
}
