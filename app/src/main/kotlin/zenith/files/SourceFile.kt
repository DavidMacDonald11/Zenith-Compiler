package zenith.files

import java.io.File
import java.io.BufferedReader
import zenith.Grammar

private val BUF_SIZE = maxOf(3, Grammar.LONGEST_PUNC_SIZE)

class SourceFile(val path: String) {
    private val reader = File(path).bufferedReader()
    private val buffer = ArrayDeque<Char>(BUF_SIZE)
    private val readBuffer = StringBuilder()

    val nextChar get() = peek(1).lastOrNull() ?: '\u0000'
    val atEnd get() = peek(1).isEmpty()
    var charPos = 0u; private set

    fun close() = reader.close()

    fun peek(n: Int): String {
        while(buffer.count() < n) {
            val char = reader.readChar()
            if(char == null) break else buffer.addLast(char)
        }

        return buffer.take(n).joinToString("")
    }

    fun readChar() = readChars(1)
    fun readChars(n: Int) = readWhile { readBuffer.length < n }
    fun readTheseChars(these: String) = readWhile { nextChar in these }
    fun readCharsUntil(until: String) = readWhile { nextChar !in until }

    private fun readWhile(condition: () -> Boolean): String {
        while(!atEnd && condition()) {
            val char = buffer.removeFirstOrNull() ?: reader.readChar()
            if(char == null) break else readBuffer.append(char)
        }

        val string = "$readBuffer"
        charPos += readBuffer.length.toUInt()
        readBuffer.clear()

        return string
    }
}

private fun BufferedReader.readChar(): Char? {
    val char = read()
    return if(char == -1) null else char.toChar()
}
