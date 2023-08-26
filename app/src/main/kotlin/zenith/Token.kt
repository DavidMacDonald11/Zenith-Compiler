package zenith

data class Token(
    val type: Type,
    val string: String,
    val position: UInt
): Faultable {
    enum class Type {
        PUNC, ID, KEY, NUM, CHAR, STR, STR_START, STR_END, RAW, NONE;
        override fun toString() = if(this == NONE) "?" else "${name[0]}"
    }

    override val faultPosition get(): UIntRange {
        val start = position
        if(of(Type.STR_START, Type.STR_END)) return UIntRange(start, start)

        val end = start + string.length.toUInt() - 1u
        return UIntRange(start, end)
    }

    override fun toString() = when {
        of(Type.PUNC) && has("") -> "$type'EOF'"
        of(Type.STR_START) -> "<S>"
        of(Type.STR_END) -> "</S>"
        else -> "$type'${string.replace("\n", "\\n")}'"
    }

    fun of(vararg types: Type) = type in types
    fun of(c: Collection<Type>) = of(*c.toTypedArray())
    fun has(vararg strings: String) = string in strings
    fun has(c: Collection<String>) = has(*c.toTypedArray())
}