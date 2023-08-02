package zenith

data class Token(
    val type: Type,
    val string: String,
    val position: UNum
): Faults.Faultable {
    enum class Type {
        PUNC, ID, KEY, NUM, CHAR, RAW, NONE,
        STR_START, STR_MIDDLE, STR_END, STR_FULL;

        override fun toString() = when(this) {
            NONE -> "?"
            STR_START -> "S{"
            STR_FULL -> "S{"
            else -> "${name[0]}"
        }
    }

    override val faultPosition get(): Faults.Position {
        val start = position.toUInt()
        val end = start + string.length.toUInt()
        return Faults.Position(start, end)
    }

    override fun toString(): String {
        if(string == "") return "$type'EOF'"

        var str = "$type'${string.replace("\n", "\\n")}'"
        if(type in listOf(Token.Type.STR_END, Token.Type.STR_FULL)) str += '}'

        return str
    }

    fun of(vararg types: Type) = type in types
    fun of(c: Collection<Type>) = of(*c.toTypedArray())
    fun has(vararg strings: String) = string in strings
    fun has(c: Collection<String>) = has(*c.toTypedArray())
}
