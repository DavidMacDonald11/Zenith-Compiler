package zenith

import java.io.File

interface Faultable { val faultPosition: UIntRange }

sealed class Fault(val obj: Faultable, val message: String) {
    class Warning(label: String, obj: Faultable, message: String)
        : Fault(obj, "$label Warning: $message")

    class Error(label: String, obj: Faultable, message: String)
        : Fault(obj, "$label Error: $message")

    class Failure(label: String, obj: Faultable, message: String)
        : Fault(obj, "$label Failure: $message")
}

fun printFaults(filePath: String, faults: List<Fault>) {
    if(faults.isEmpty()) return

    File(filePath).useLines {
        val lines = it.toList().map { "$it\n" }
        println("Faults in $filePath:")

        for(fault in faults) {
            println(fault.message)
            printFileLines(lines, fault.obj.faultPosition)
        }
    }
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
