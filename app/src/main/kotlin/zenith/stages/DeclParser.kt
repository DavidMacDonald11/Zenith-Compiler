package zenith.parser

import zenith.*

internal fun parseDecl(ctx: Context): NodeResult? {
    if(!ctx.next.has(Grammar.STORAGE_KEYS) && !ctx.next.has("val", "var"))
        return null

    return parseVarDecl(ctx)
}

internal fun parseVarDecl(ctx: Context): NodeResult? {
    val result = MutableNodeResult("VarDecl")

    val hasStorage = ctx.next.has(Grammar.STORAGE_KEYS)
    if(hasStorage) result.add(ctx.take())

    if(!ctx.next.has("val", "var")) {
        if(hasStorage) ctx.untake()
        return null
    }

    result.add(ctx.take())
    if(result.add(ctx.expectingOf(Token.Type.ID))) return result.toNodeResult()

    if(ctx.next.has(":")) {
        result.add(ctx.take())
        ctx.skipNewline()
        if(result.add(parseType(ctx))) return result.toNodeResult()
    }

    if(result.add(ctx.expectingHas("="))) return result.toNodeResult()
    ctx.skipNewline()

    result.add(parseExpr(ctx))
    return result.toNodeResult()
}

internal fun parseType(ctx: Context) =
    if(ctx.next.has(Grammar.REF_OPS)) parseRefType(ctx) else parseDataType(ctx)

internal fun parseRefType(ctx: Context): NodeResult {
    val result = MutableNodeResult("RefType")

    val op = ctx.expectingHas(Grammar.REF_OPS)
    if(result.add(op)) return result.toNodeResult()

    if(result.add(parseDataType(ctx))) return result.toNodeResult()
    if(ctx.next.has("?")) result.add(ctx.take())

    return result.toNodeResult()
}

internal fun parseDataType(ctx: Context): NodeResult {
    val result = MutableNodeResult("DataType")
    result.add(ctx.expectingOf(Token.Type.ID))
    return result.toNodeResult()
}
