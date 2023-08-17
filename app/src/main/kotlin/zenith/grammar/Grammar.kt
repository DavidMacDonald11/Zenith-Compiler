package zenith

object Grammar {
    val NUM_START_SYMS = "0123456789."
    val NUM_SYMS = "box_abcdefABCDEF_ijINRX$NUM_START_SYMS"

    private val BIN = "[01]"
    private val OCT = "[0-7]"
    private val HEX = "[0-9a-fA-F]"
    val NUM_PATTERN = """
        (
            (\d+(\.\d*)?) | (\d*\.\d+) |
            (0b($BIN+(\.$BIN*)? | ($BIN*\.$BIN+))) |
            (0o($OCT+(\.$OCT*)? | ($OCT*\.$OCT+))) |
            (0x($HEX+(\.$HEX*)? | ($HEX*\.$HEX+)))
        ) [ij]? (
            I | I8 | I16 | I32 | I64 | IMax |
            N | N8 | N16 | N32 | N64 | NMax |
            R |            R32 | R64 | RMax |
            X |            X32 | X64 | XMax
        )?
    """.replace("\\s".toRegex(), "").toRegex()

    val NON_NUM_PATTERN = """\.[^\d]""".toRegex()

    val COMPARE_OPS = setOf("<", ">", "<=", ">=", "==", "!=")
    val PUNCS = COMPARE_OPS + setOf(
        "^", "*", "/", "%", "+", "-", "<<", ">>", "&", "$", "|",
        "&&", "$$", "||",
        ";", "(", ")", "{", "}", "[", "]", ",", ":"
    )
    val PUNC_SYMS = PUNCS.flatMap{ it.asIterable() }.toSet()
    val LONGEST_PUNC_SIZE = PUNCS.maxBy{ it.length }.length

    private val LETTERS = ('A'..'Z').joinToString("")
    val ID_START_SYMS = "_" + LETTERS + LETTERS.lowercase()
    val ID_SYMS = ID_START_SYMS + "0123456789"

    val KEYS = setOf<String>()

    val CHAR_PATTERN = """
        '(\\[btnr'"\\$] | \\u[a-fA-F0-9]{4} | [^\\'])'
    """.replace("\\s".toRegex(), "").toRegex()

    val ESCAPE_PATTERN = """
        \\[btnr'"\\$] | \\u[a-fA-F0-9]{4}
    """.replace("\\s".toRegex(), "").toRegex()

    val EOF = ""
    val LINE_END = setOf(";", "\n", EOF)
}
