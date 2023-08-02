package zenith

import java.io.File

class Faults(val filePath: String) {
    data class Position(val start: UInt, val end: UInt) {
        operator fun contains(value: UInt) = value >= start && value < end
    }

    class Failure: Exception()
    interface Faultable { val faultPosition: Position }
    data class Fault(val message: String, val position: Position)

    private val warnings = mutableListOf<Fault>()
    private val errors = mutableListOf<Fault>()

    fun warn(label: String, fault: Faultable, message: String) {
        warnings += Fault("$label Warning: $message", fault.faultPosition)
    }

    fun error(label: String, fault: Faultable, message: String) {
        errors += Fault("$label Error: $message", fault.faultPosition)
    }

    fun fail(label: String, fault: Faultable, message: String): Failure {
        errors += Fault("$label Failure: $message", fault.faultPosition)
        return Failure()
    }

    fun throwIfErrors() { if(errors.size > 0) throw Failure() }

    fun print() = File(filePath).useLines {
        println("\nFaults in $filePath:")

        val lines = it.toList().map { "$it\n" }

        for(fault in warnings + errors) {
            println(fault.message)
            printFileLines(lines, fault.position)
        }
    }

    private fun printFileLines(lines: List<String>, position: Position) {
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
        } while(charCount + 1u < position.end)
    }
}
