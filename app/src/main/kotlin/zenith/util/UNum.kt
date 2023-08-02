package zenith

private val maxUByte = UByte.MAX_VALUE.toUInt()
private val maxUShort = UShort.MAX_VALUE.toUInt()

sealed interface UNum {
    fun toUInt(): kotlin.UInt
    operator fun plus(n: UNum) = new(toUInt() + n.toUInt())
    operator fun plus(n: Int) = new(toUInt() + n.toUInt())

    companion object {
        fun new(n: kotlin.UInt) = when {
            n <= maxUByte -> UByte(n.toUByte())
            n <= maxUShort -> UShort(n.toUShort())
            else -> UInt(n)
        }
    }

    data class UByte(private val n: kotlin.UByte): UNum {
        override fun toUInt() = n.toUInt()
    }

    data class UShort(private val n: kotlin.UShort): UNum {
        override fun toUInt() = n.toUInt()
    }

    data class UInt(private val n: kotlin.UInt): UNum {
        override fun toUInt() = n
    }
}
