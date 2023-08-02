package zenith

import java.io.File
import java.io.BufferedReader

class SourceFile(val path: String) {
    private val source = File(path).bufferedReader()
    private val builder = StringBuilder()
    private val buffer = StringBuilder(80)
    var charPos = 0u
        private set

    init { growBuffer() }

    val atEnd get() = buffer.isEmpty()
    val nextChar get() = buffer.getOrNull(0) ?: '\u0000'

    fun close() = source.close()
    fun nextStr(n: Int) = buffer.take(n)

    fun readChars(n: Int = -1) = readWhile { builder.length < n }
    fun readTheseChars(these: String) = readWhile { nextChar in these }
    fun readCharsUntil(until: String) = readWhile { nextChar !in until }

    private fun readWhile(condition: () -> Boolean): String {
        while(!atEnd && condition()) {
            builder.append(nextChar)
            buffer.deleteCharAt(0)
            if(buffer.length == 1) growBuffer()
        }

        charPos += builder.length.toUInt()

        val string = "$builder"
        builder.clear()
        return string
    }

    private fun growBuffer() {
        buffer.append(source.readLine()?.plus("\n") ?: "")
    }
}
