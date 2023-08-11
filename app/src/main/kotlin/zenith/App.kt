package zenith

import zenith.lexer.*

fun main() {
    val filePath = "../test.zen"
    val faults = mutableListOf<Fault>()

    val tokens = tokenizeFile(filePath).let { result ->
        faults.addAll(result.faults)

        if(result.errored) {
            printTokens(result.tokens)
            printFaults(filePath, faults)
            return
        }

        result.tokens
    }

    printTokens(tokens)
    printFaults(filePath, faults)
}

fun printTokens(tokens: List<Token>) {
    println(tokens.joinToString(", ", "[", "]"))
}
