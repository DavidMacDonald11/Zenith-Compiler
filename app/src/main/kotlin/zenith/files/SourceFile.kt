package zenith

import java.io.File
import java.io.BufferedReader

class SourceFile(val path: String) {
    private val source = File(path).bufferedReader()
    private var buffer = StringBuilder()
    private var charPos = 0u

    init { addToBuffer() }

    val atEnd get() = buffer.isEmpty()
    val next get() = buffer.getOrNull(0)?:'\u0000'
    val position get() = Position(path, charPos)

    var builder = StringBuilder()

    fun tokenInit(): Pair<Position, StringBuilder> {
        builder = StringBuilder()
        return Pair(position, builder)
    }

    fun close() = source.close()
    fun nextString(n: Int) = buffer.take(n)

    fun take(num: Int = -1, these: String = "", until: String = ""): String {
        if(atEnd || num == 0) return ""
        val tempBuilder = StringBuilder()

        while(!atEnd && tempBuilder.length != num) {
            if(these != "" && next !in these || next in until) break

            tempBuilder.append(next)
            charPos++

            buffer.deleteCharAt(0)
            if(buffer.length < 2) addToBuffer()
        }

        builder.append(tempBuilder)
        return "$tempBuilder"
    }

    private fun addToBuffer() {
        buffer.append(source.readLine()?.plus("\n")?:"")
    }

    class Position(val file: String, charPos: UInt) {
        private val _charPos: Num
        val charPos get() = _charPos.toUInt()

        init {
            val maxUByte = UByte.MAX_VALUE.toUInt()
            val maxUShort = UShort.MAX_VALUE.toUInt()

            _charPos = when {
                charPos <= maxUByte -> Num.UByte(charPos.toUByte())
                charPos <= maxUShort -> Num.UShort(charPos.toUShort())
                else -> Num.UInt(charPos)
            }
        }

        private sealed interface Num {
            fun toUInt(): kotlin.UInt

            data class UByte(val value: kotlin.UByte): Num {
                override fun toUInt() = value.toUInt()
            }

            data class UShort(val value: kotlin.UShort): Num {
                override fun toUInt() = value.toUInt()
            }

            data class UInt(val value: kotlin.UInt): Num {
                override fun toUInt() = value
            }
        }
    }
}
