package zenith

import kotlin.system.exitProcess
import zenith.lexer.tokenizeFile
import zenith.parser.parseTokens
import zenith.io.*

fun main(args: Array<String>) {
    val command = args.getOrNull(0)

    val status = when(command) {
        null -> noCommand()
        in listOf("help", "--help", "-h") -> helpCommand()
        "init" -> initCommand(args.getOrNull(1))
        "build" -> buildCommand(::buildFile)
        "clean" -> cleanCommand()
        else -> unknownCommand(command)
    }

    exitProcess(status.code)
}

private fun Runtime.usedMemoryKB(): Long {
    gc(); return (totalMemory() - freeMemory()) / 1024
}

private fun buildFile(path: String): Boolean {
    val startingMem = Runtime.getRuntime().usedMemoryKB()
    val faults = mutableListOf<Fault>()

    val tokens = tokenizeFile(path).let { faults += it.faults; it.tokens }
    println("Tokens: \n${tokens.joinToString(", ", "[", "]")}\n")
    if(printFaults(path, faults)) return endBuildFile(startingMem, true)

    val tree = parseTokens(tokens).let { faults += it.faults; it.node }
    println("AST: \n${tree}\n")
    if(printFaults(path, faults)) return endBuildFile(startingMem, true)

    printFaults(path, faults, printWarnings = true)
    return endBuildFile(startingMem, false)
}

private fun endBuildFile(startingMem: Long, early: Boolean): Boolean {
    val finalMem = Runtime.getRuntime().usedMemoryKB()
    print("\nUsed ${"%,d".format(finalMem - startingMem)} KB")
    print(" (${"%,d".format(finalMem)} KB total) of memory.")

    return early
}
