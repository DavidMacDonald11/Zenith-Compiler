package zenith.parser

import zenith.*

internal fun parseFileStat(ctx: Context): NodeResult {
    val result = MutableNodeResult("FileStat")

    while(!ctx.next.has(Grammar.EOF)) {
        val stat = parseDecl(ctx) ?: parseExpr(ctx)
        if(result.add(stat)) return result.toNodeResult()

        ctx.skipNewline()
    }

    return result.toNodeResult()
}
