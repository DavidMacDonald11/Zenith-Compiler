package zenith.parser

import zenith.*

internal fun parseExpr(ctx: Context) = parseConditionExpr(ctx)

internal fun parseConditionExpr(ctx: Context): NodeResult {
    val left = parseLogicOrExpr(ctx)
    if(left.failed || !ctx.next.has("if")) return left

    val children = mutableListOf<Faultable>(left.value)
    val faults = left.faults.toMutableList()
    val type = "ConditionExpr"

    children += ctx.take()
    ctx.skipNewline()

    val condition = parseLogicOrExpr(ctx)

    children += condition.value
    faults += condition.faults

    if(condition.failed) return NodeResult(Node(type, children), faults)

    ctx.skipNewline()
    val elseResult = ctx.expectingHas("else")

    children += elseResult.value
    faults += elseResult.faults

    if(elseResult.failed) return NodeResult(Node(type, children), faults)

    ctx.skipNewline()
    val right = parseConditionExpr(ctx)

    children += right.value
    faults += right.faults

    return NodeResult(Node(type, children), faults)
}

internal fun parseLogicOrExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "LogicOrExpr", ::parseLogicXorExpr, listOf("or")
)

internal fun parseLogicXorExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "LogicXorExpr", ::parseLogicAndExpr, listOf("xor")
)

internal fun parseLogicAndExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "LogicAndExpr", ::parseCompareExpr, listOf("and")
)

internal fun parseCompareExpr(ctx: Context): NodeResult {
    val children = mutableListOf<Faultable>()
    val faults = mutableListOf<Fault>()

    val left = parseBitOrExpr(ctx)
    if(left.failed || !ctx.next.has(Grammar.COMPARE_OPS)) return left

    children += left.value
    faults += left.faults

    while(ctx.next.has(Grammar.COMPARE_OPS)) {
        children += ctx.take()
        ctx.skipNewline()

        val next = parseBitOrExpr(ctx)

        children += next.value
        faults += next.faults

        if(next.failed) break
    }

    return NodeResult(Node("CompareExpr", children), faults)
}

internal fun parseBitOrExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "BitOrExpr", ::parseBitXorExpr, listOf("|")
)

internal fun parseBitXorExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "BitXorExpr", ::parseBitAndExpr, listOf("$")
)

internal fun parseBitAndExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "BitAndExpr", ::parseBitShiftExpr, listOf("&")
)

internal fun parseBitShiftExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "BitShiftExpr", ::parseAddExpr, listOf("<<", ">>")
)

internal fun parseAddExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "AddExpr", ::parseMultiplyExpr, listOf("+", "-")
)

internal fun parseMultiplyExpr(ctx: Context) = parseLeftBinaryExpr(
    ctx, "MultiplyExpr", ::parseExponentExpr, listOf("*", "/", "%")
)

internal fun parseExponentExpr(ctx: Context): NodeResult {
    val left = parsePrefixExpr(ctx)
    if(left.failed || !ctx.next.has("^")) return left

    val op = ctx.take()
    ctx.skipNewline()

    val right = parseExponentExpr(ctx)
    val node = Node("ExponentExpr", listOf(left.value, op, right.value))

    return NodeResult(node, left.faults + right.faults)
}

internal fun parsePrefixExpr(ctx: Context): NodeResult {
    val opList = listOf("+", "-", "~", "$", "&", "#", "not")
    if(!ctx.next.has(opList)) return parsePostfixExpr(ctx)

    val op = ctx.take()
    val right = parsePostfixExpr(ctx)
    val node = Node("PrefixExpr", listOf(op, right.value))

    return NodeResult(node, right.faults)
}

internal fun parsePostfixExpr(
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
        val (id, faults) = ctx.expectingOf(Token.Type.ID)
        val node = Node(type, listOf(left.value, op, id))

        return parsePostfixExpr(ctx, NodeResult(node, left.faults + faults))
    }

    if(skipped) ctx.untake()
    return left
}

internal fun parsePrimaryExpr(ctx: Context): NodeResult {
    val type = "PrimaryExpr"

    if(ctx.next.of(Token.Type.NUM, Token.Type.ID, Token.Type.CHAR)) {
        return NodeResult(Node(type, listOf(ctx.take())))
    }

    if(ctx.next.has(Grammar.PRIMARY_KEYS)) {
        return NodeResult(Node(type, listOf(ctx.take())))
    }

    if(ctx.next.has("(")) return parseParenExpr(ctx)
    if(ctx.next.of(Token.Type.STR_START)) return parseStringExpr(ctx)

    val node = Node(type, listOf(ctx.take()))
    val fault = Failure(node, "Unexpected token in primary expression")
    return NodeResult(node, listOf(fault))
}

internal fun parseParenExpr(ctx: Context): NodeResult {
    val type = "ParenExpr"

    val left = ctx.expectingHas("(")
    if(left.failed) return NodeResult(Node(type, listOf(left.value)))

    ctx.skipNewline()
    val expr = parseExpr(ctx)
    ctx.skipNewline()

    val right = ctx.expectingHas(")")
    val faults = expr.faults + right.faults

    val node = Node(type, listOf(left.value, expr.value, right.value))
    return NodeResult(node, faults)
}

internal fun parseStringExpr(ctx: Context): NodeResult {
    val children = mutableListOf<Faultable>()
    val faults = mutableListOf<Fault>()

    val start = ctx.expectingOf(Token.Type.STR_START)
    children += start.value
    faults += start.faults

    while(!ctx.next.of(Token.Type.STR_END)) {
        if(ctx.next.of(Token.Type.STR)) {
            children += ctx.take()
            continue
        }

        val expr = parseExpr(ctx)

        children += expr.value
        faults += expr.faults

        if(expr.failed) break
    }

    if(ctx.next.of(Token.Type.STR_END)) children += ctx.take()
    return NodeResult(Node("StringExpr", children), faults)
}

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
