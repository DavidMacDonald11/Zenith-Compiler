package zenith

import java.io.File

interface Faultable { val faultPosition: UIntRange }

typealias Faults = List<Fault>
typealias MutableFaults = MutableList<Fault>

class Fault(label: String, val type: Char, val obj: Faultable, msg: String) {
    val message = "$label ${mapType()}: $msg"

    private fun mapType() = when(type) {
        'W' -> "Warning"
        'E' -> "Error"
        'F' -> "Failure"
        else -> throw IllegalArgumentException()
    }
}

data class Result<T>(val value: T, val faults: List<Fault> = listOf()) {
    constructor(value: T, fault: Fault): this(value, listOf(fault))

    val failed get() = faults.any { it.type == 'F' }
    val errored get() = faults.any { it.type in listOf('F', 'E') }
    val hasFaults get() = faults.isNotEmpty()
}

fun printFaults(
    filePath: String,
    faults: Faults,
    printWarnings: Boolean = false
): Boolean {
    if(faults.isEmpty()) return false
    if(!printWarnings && faults.all { it.type == 'W' }) return false

    File(filePath).useLines {
        val lines = it.toList().map { "$it\n" }
        println("Faults in $filePath:")

        for(fault in faults) {
            println(fault.message)
            printFileLines(lines, fault.obj.faultPosition)
        }
    }

    return true
}

private fun printFileLines(lines: List<String>, position: UIntRange) {
    var charCount = 0u
    var lineNum = 0

    for(line in lines) {
        if(charCount + line.length.toUInt() > position.start) break
        charCount += line.length.toUInt()
        lineNum += 1
    }

    do {
        if(lineNum == lines.size) {
            println("${"%4d".format(lineNum + 1)}|EOF")
            println("    |^^^")
            return
        }

        val line = lines[lineNum]
        val marks = StringBuilder()

        for(i in line.indices) {
            marks.append(if(charCount in position) "^" else " ")
            charCount += 1u
        }

        print("${"%4d".format(lineNum + 1)}|$line")
        println("    |$marks")

        lineNum += 1
    } while(charCount < position.endInclusive)
}
