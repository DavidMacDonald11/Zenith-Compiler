package zenith.parser

import zenith.*

internal fun parseDecl(ctx: Context): NodeResult? {
    return parseVarDecl(ctx) ?: parseFunDecl(ctx)
}

/*
funDecl : "fun" ID "(" NL? [funParams] NL? ")" ["->" NL? type] NL? multiStat
        | "fun" ID "(" NL? [funParams] NL? ")" ["->" NL? type] "=" NL? expr
        ;

funParams : [pFunParams "," NL?] dFunParams
          | pFunParams
          ;

pFunParams : [pFunParams "," NL?] ID ":" NL? type ;
dFunParams : [dFunParams "," NL?] ID [":" NL? type] "=" NL? expr ;
*/

internal fun parseFunDecl(ctx: Context): NodeResult? {
    if(!ctx.next.has("fun")) return null
    val result = MutableNodeResult("FunDecl")

    result.add(ctx.take())
    if(result.add(ctx.expectingOf(Token.Type.ID))) return result.toNodeResult()

    var allowPParams = true

    if(result.add(ctx.expectingHas("("))) return result.toNodeResult()
    ctx.skipNewline()

    while(!ctx.next.has(")")) {
        val param = parseFunParam(ctx, allowPParams)
        if(result.add(param.result)) return result.toNodeResult()
        allowPParams = param.allowPParams

        if(!ctx.next.has(",")) break
        ctx.take()
        ctx.skipNewline()
    }

    ctx.skipNewline()
    if(result.add(ctx.expectingHas(")"))) return result.toNodeResult()

    if(ctx.next.has("->")) {
        result.add(ctx.take())
        ctx.skipNewline()
        if(result.add(parseType(ctx))) return result.toNodeResult()
    } else if(!ctx.next.has("=")) {
        result.error("Function requires expression value or type annotation")
        return result.toNodeResult()
    }

    if(ctx.next.has("=")) {
        result.add(ctx.take())
        ctx.skipNewline()
        result.add(parseExpr(ctx))
        return result.toNodeResult()
    }

    result.add(parseMultiStat(ctx))
    return result.toNodeResult()
}

private class FunParamResult(val result: NodeResult, val allowPParams: Boolean)

private fun parseFunParam(ctx: Context, allowPParams: Boolean): FunParamResult {
    val result = MutableNodeResult("FunParam")

    if(result.add(ctx.expectingOf(Token.Type.ID)))
        return FunParamResult(result.toNodeResult(), allowPParams)

    if(ctx.next.has(":")) {
        result.add(ctx.take())
        ctx.skipNewline()

        if(result.add(parseType(ctx)))
            return FunParamResult(result.toNodeResult(), allowPParams)
    } else if(!ctx.next.has("=")) {
        result.error("Parameter requires default value or type annotation")
        return FunParamResult(result.toNodeResult(), allowPParams)
    }

    if(result.add(ctx.expectingHas("=")))
        return FunParamResult(result.toNodeResult(), allowPParams)

    ctx.skipNewline()
    result.add(parseExpr(ctx))

    return FunParamResult(result.toNodeResult(), false)
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
