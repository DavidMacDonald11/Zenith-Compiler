package zenith

fun main() {
    val filePath = "../test.zen"
    val faults = Faults(filePath)
    var tokens: List<Token>? = null

    try {
        val lexer = Lexer(faults)
        tokens = lexer.makeTokens()
        faults.throwIfErrors()
    } catch(e: Faults.Failure) {
        faults.print()
    }

    tokens?.let { println(it.joinToString(", ", "[", "]")) }
}
