package zenith

import java.io.File

class CompilerFailure: Exception() {}

private typealias Fault = Pair<String, Faultable.Position>
private val warnings = mutableListOf<Fault>()
private val errors = mutableListOf<Fault>()

interface Faultable {
    data class Position(val file: String, val start: UInt, val end: UInt) {
        constructor(start: SourceFile.Position, strLen: UInt) : this(
            file = start.file,
            start = start.charPos,
            end = start.charPos + strLen
        )
    }

    val faultablePosition: Position
}

fun warn(label: String, fault: Faultable, message: String) {
    warnings += Pair("$label Warning: $message", fault.faultablePosition)
}

fun error(label: String, fault: Faultable, message: String) {
    errors += Pair("$label Error: $message", fault.faultablePosition)
}

fun fail(label: String, fault: Faultable, message: String): CompilerFailure {
    errors += Pair("$label Failure: $message", fault.faultablePosition)
    return CompilerFailure()
}

fun checkFaults() { if(errors.size > 0) throw CompilerFailure() }

fun handleFaults() {
    val fileFaultMap = (warnings + errors).groupBy { it.second.file }

    for((filePath, faults) in fileFaultMap) {
        File(filePath).useLines {
            println("Faults in \"$filePath\":")
            handleFileFaults(it, faults)
            println()
        }
    }
}

private fun handleFileFaults(linesSeq: Sequence<String>, faults: List<Fault>) {
    val lines = linesSeq.toList().map{ "$it\n" }

    for(fault in faults) {
        println(fault.first)
        printLines(lines, fault.second)
    }
}

private fun printLines(lines: List<String>, position: Faultable.Position) {
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
            val inside = charCount >= position.start && charCount < position.end
            marks.append(if(inside) "^" else " ")
            charCount += 1u
        }

        print("${"%4d".format(lineNum + 1)}|$line")
        println("    |$marks")

        lineNum += 1
    } while(charCount + 1u < position.end)
}
