package zenith.parser

import zenith.*

internal fun parseExpr(ctx: Context) = parseConditionExpr(ctx)

internal fun parseConditionExpr(ctx: Context): NodeResult {
    val left = parseLogicOrExpr(ctx)
    if(left.failed || !ctx.next.has("if")) return left

    val result = MutableNodeResult("ConditionExpr", left)
    result.add(ctx.take())
    ctx.skipNewline()

    if(result.add(parseLogicOrExpr(ctx))) return result.toNodeResult()
    ctx.skipNewline()

    if(result.add(ctx.expectingHas("else"))) return result.toNodeResult()
    ctx.skipNewline()

    result.add(parseConditionExpr(ctx))
    return result.toNodeResult()
}

internal fun parseLogicOrExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "LogicOrExpr", ::parseLogicXorExpr, "or")
)

internal fun parseLogicXorExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "LogicXorExpr", ::parseLogicAndExpr, "xor")
)

internal fun parseLogicAndExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "LogicAndExpr", ::parseCompareExpr, "and")
)

internal fun parseCompareExpr(ctx: Context): NodeResult {
    val left = parseBitOrExpr(ctx)
    if(left.failed || !ctx.next.has(Grammar.COMPARE_OPS)) return left

    val result = MutableNodeResult("CompareExpr", left)

    while(ctx.next.has(Grammar.COMPARE_OPS)) {
        result.add(ctx.take())
        ctx.skipNewline()
        if(result.add(parseBitOrExpr(ctx))) break
    }

    return result.toNodeResult()
}

internal fun parseBitOrExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "BitOrExpr", ::parseBitXorExpr, "|")
)

internal fun parseBitXorExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "BitXorExpr", ::parseBitAndExpr, "$")
)

internal fun parseBitAndExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "BitAndExpr", ::parseBitShiftExpr, "&")
)

internal fun parseBitShiftExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "BitShiftExpr", ::parseAddExpr, "<<", ">>")
)

internal fun parseAddExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "AddExpr", ::parseMultiplyExpr, "+", "-")
)

internal fun parseMultiplyExpr(ctx: Context) = parseLeftBinaryExpr(
    BinaryArgs(ctx, "MultiplyExpr", ::parseExponentExpr, "*", "/", "%")
)

internal fun parseExponentExpr(ctx: Context): NodeResult {
    val left = parsePrefixExpr(ctx)
    if(left.failed || !ctx.next.has("^")) return left

    val result = MutableNodeResult("ExponentialExpr", left)

    result.add(ctx.take())
    ctx.skipNewline()

    result.add(parseExponentExpr(ctx))
    return result.toNodeResult()
}

internal fun parsePrefixExpr(ctx: Context): NodeResult {
    val opList = listOf("+", "-", "~", "$", "&", "#", "not")
    if(!ctx.next.has(opList)) return parsePostfixExpr(ctx)

    val result = MutableNodeResult("PrefixExpr", ctx.take())
    result.add(parsePostfixExpr(ctx))
    return result.toNodeResult()
}

internal fun parsePostfixExpr(
    ctx: Context,
    left: NodeResult = parsePrimaryExpr(ctx)
): NodeResult {
    if(left.failed) return left
    val result = MutableNodeResult("PostfixExpr", left)

    if(ctx.next.has("!")) {
        result.add(ctx.take())
        return result.toNodeResult()
    }

    val skipped = ctx.skipNewline()

    if(ctx.next.has(".")) {
        result.add(ctx.take())
        ctx.skipNewline()

        result.add(ctx.expectingOf(Token.Type.ID))
        return parsePostfixExpr(ctx, result.toNodeResult())
    }

    if(skipped) ctx.untake()
    return left
}

internal fun parsePrimaryExpr(ctx: Context): NodeResult {
    val type = "PrimaryExpr"
    val primaryTypes = listOf(Token.Type.NUM, Token.Type.ID, Token.Type.CHAR)

    if(ctx.next.of(primaryTypes) || ctx.next.has(Grammar.PRIMARY_KEYS)) {
        return NodeResult(Node(type, listOf(ctx.take())))
    }

    if(ctx.next.has("(")) return parseParenExpr(ctx)
    if(ctx.next.has(Grammar.STORAGE_KEYS)) return parseAllocExpr(ctx)
    if(ctx.next.of(Token.Type.STR_START)) return parseStringExpr(ctx)

    val node = Node(type, listOf(ctx.take()))
    val fault = Failure(node, "Unexpected token in primary expression")
    return NodeResult(node, listOf(fault))
}

internal fun parseParenExpr(ctx: Context): NodeResult {
    val result = MutableNodeResult("ParenExpr")

    if(result.add(ctx.expectingHas("("))) return result.toNodeResult()
    ctx.skipNewline()

    if(result.add(parseExpr(ctx))) return result.toNodeResult()
    ctx.skipNewline()

    result.add(ctx.expectingHas(")"))
    return result.toNodeResult()
}

internal fun parseAllocExpr(ctx: Context): NodeResult {
    val result = MutableNodeResult("AllocExpr")

    val storage = ctx.expectingHas(Grammar.STORAGE_KEYS)
    if(result.add(storage)) return result.toNodeResult()

    result.add(parseParenExpr(ctx))
    return result.toNodeResult()
}

internal fun parseStringExpr(ctx: Context): NodeResult {
    val result = MutableNodeResult("StringExpr")

    val start = ctx.expectingOf(Token.Type.STR_START)
    if(result.add(start)) return result.toNodeResult()

    while(!ctx.next.of(Token.Type.STR_END)) {
        if(ctx.next.of(Token.Type.STR)) {
            result.add(ctx.take())
            continue
        }

        if(result.add(parseExpr(ctx))) return result.toNodeResult()
    }

    result.add(ctx.expectingOf(Token.Type.STR_END))
    return result.toNodeResult()
}

private tailrec fun parseLeftBinaryExpr(
    args: BinaryArgs,
    left: NodeResult = args.parseChild(args.ctx)
): NodeResult {
    if(left.failed || !args.ctx.next.has(args.opList)) return left

    val result = MutableNodeResult(args.type, args.ctx.take())
    args.ctx.skipNewline()

    result.add(args.parseChild(args.ctx))
    return parseLeftBinaryExpr(args, result.toNodeResult())
}

private typealias parseFun = (Context) -> NodeResult

private class BinaryArgs(
    val ctx: Context,
    val type: String,
    val parseChild: parseFun,
    val opList: List<String>
) {
    constructor(ctx: Context, type: String, parse: parseFun, vararg op: String):
        this(ctx, type, parse, op.toList())
}
