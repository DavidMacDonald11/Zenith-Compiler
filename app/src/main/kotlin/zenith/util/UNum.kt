package zenith

private val maxUByte = UByte.MAX_VALUE.toUInt()
private val maxUShort = UShort.MAX_VALUE.toUInt()

sealed interface UNum {
    class UByte(val n: kotlin.UByte): UNum
    class UShort(val n: kotlin.UShort): UNum
    class UInt(val n: kotlin.UInt): UNum

    fun toUInt() = when(this) {
        is UByte -> this.n.toUInt()
        is UShort -> this.n.toUInt()
        is UInt -> this.n
    }

    operator fun plus(n: UNum) = new(toUInt() + n.toUInt())
    operator fun plus(n: Int) = new(toUInt() + n.toUInt())

    companion object {
        fun new(n: kotlin.UInt) = when {
            n <= maxUByte -> UByte(n.toUByte())
            n <= maxUShort -> UShort(n.toUShort())
            else -> UInt(n)
        }
    }
}
