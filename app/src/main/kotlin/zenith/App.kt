package zenith

fun main() {
    val file = SourceFile("../test.zen")
    var tokens: List<Token>? = null

    try {
        tokens = tokenize(file)
        checkFaults()

        handleFaults()
    } catch(e: CompilerFailure) {
        handleFaults()
    }

    println((tokens?:listOf<Token>()).joinToString(", ", "[", "]"))
    file.close()
}
