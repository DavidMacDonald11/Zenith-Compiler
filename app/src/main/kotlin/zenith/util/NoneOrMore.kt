package zenith

sealed interface NoneOrMore<T> {
    class None<T>(): NoneOrMore<T>
    class One<T>(val value: T): NoneOrMore<T>
    class More<T>(val values: List<T>): NoneOrMore<T>

    fun any(predicate: (T) -> Boolean): Boolean = when(this) {
        is None -> false
        is One -> predicate(value)
        is More -> values.any(predicate)
    }

    fun lastOrNull() = when(this) {
        is None -> null
        is One -> value
        is More -> values.last()
    }

    operator fun plus(obj: NoneOrMore<T>): NoneOrMore<T> {
        return when {
            obj is None -> this
            this is None -> obj
            this is One && obj is One -> More(listOf(this.value, obj.value))
            this is One && obj is More -> More(listOf(this.value) + obj.values)
            this is More && obj is One -> More(this.values + obj.value)
            this is More && obj is More -> More(this.values + obj.values)
            else -> None()
        }
    }
}

typealias None<T> = NoneOrMore.None<T>
typealias One<T> = NoneOrMore.One<T>
typealias More<T> = NoneOrMore.More<T>
