package zenith

data class Token(
    val type: Type,
    val string: String,
    val position: UNum
): Faults.Faultable {
    enum class Type {
        PUNC, ID, KEY, NUM, CHAR, RAW, NONE,
        STR_START, STR_MIDDLE, STR_END, STR_FULL;
        override fun toString() = if(this == NONE) "?" else "${name[0]}"
    }

    override val faultPosition get(): Faults.Position {
        val start = position.toUInt()
        val end = start + string.length.toUInt()
        return Faults.Position(start, end)
    }

    override fun toString() =
        if(string == "") "$type'EOF'"
        else "$type'${string.replace("\n", "\\n")}'"

    fun of(vararg types: Type) = type in types
    fun of(c: Collection<Type>) = of(*c.toTypedArray())
    fun has(vararg strings: String) = string in strings
    fun has(c: Collection<String>) = has(*c.toTypedArray())
}
