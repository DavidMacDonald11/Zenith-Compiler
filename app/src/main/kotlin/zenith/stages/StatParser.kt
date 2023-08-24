package zenith.parser

import zenith.*

internal fun parseFileStat(ctx: Context): NodeResult {
    val result = MutableNodeResult("FileStat")

    while(!ctx.next.has(Grammar.EOF)) {
        if(ctx.next.has(";")) {
            ctx.take()
            ctx.skipNewline()
            continue
        }

        if(result.add(parseStat(ctx))) return result.toNodeResult()
        if(ctx.next.has(Grammar.EOF)) break

        val end = ctx.expectingHas(";", "\n")
        ctx.skipNewline()
        result.addAll(end.faults)

        if(end.failed) return result.toNodeResult()
    }

    return result.toNodeResult()
}

internal fun parseStat(ctx: Context): NodeResult {
    if(ctx.next.has("{")) return parseMultiStat(ctx)
    return parseDecl(ctx) ?: parseExpr(ctx)
}

internal fun parseMultiStat(ctx: Context): NodeResult {
    val result = MutableNodeResult("MultiStat")

    if(result.add(ctx.expectingHas("{"))) return result.toNodeResult()
    ctx.skipNewline()

    while(!ctx.next.has("}", Grammar.EOF)) {
        if(ctx.next.has(";")) {
            ctx.take()
            ctx.skipNewline()
            continue
        }

        if(result.add(parseStat(ctx))) return result.toNodeResult()
        if(ctx.next.has("}")) break

        val end = ctx.expectingHas(";", "\n")
        ctx.skipNewline()
        result.addAll(end.faults)

        if(end.failed) return result.toNodeResult()
    }

    result.add(ctx.expectingHas("}"))
    return result.toNodeResult()
}
