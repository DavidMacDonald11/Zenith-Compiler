package zenith.parser

import zenith.*

internal fun parseDataType(ctx: Context): NodeResult {
    val type = ctx.expectingHas(Grammar.TYPE_KEYS)
    return NodeResult(Node("DataType", listOf(type.value)), type.faults)
}
