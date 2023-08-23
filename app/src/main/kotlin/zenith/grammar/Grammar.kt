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
        "~", "#", "!", ".",
        "^", "*", "/", "%", "+", "-", "<<", ">>", "&", "$", "|",
        ";", "(", ")", "{", "}", "[", "]", ",", ":", "?",
        "="
    )
    val PUNC_SYMS = PUNCS.flatMap{ it.asIterable() }.toSet()
    val LONGEST_PUNC_SIZE = PUNCS.maxBy{ it.length }.length

    private val LETTERS = ('A'..'Z').joinToString("")
    val ID_START_SYMS = "_" + LETTERS + LETTERS.lowercase()
    val ID_SYMS = ID_START_SYMS + "0123456789"

    val TYPE_KEYS = setOf(
        "Bool", "Char", "String",
        "Int8", "Int16", "Int32",   "Int64",   "IntMax",
        "Nat8", "Nat16", "Nat32",   "Nat64",   "NatMax",
                         "Real32",  "Real64",  "RealMax",
                         "Cplex32", "Cplex64", "CplexMax",
        "Size"
    )
    val STORAGE_KEYS = setOf("static", "auto", "dynamic")
    val PRIMARY_KEYS = setOf("true", "false", "none", "undefined")
    val KEYS = TYPE_KEYS + STORAGE_KEYS + PRIMARY_KEYS + setOf(
        "not", "in", "is", "and", "xor", "or", "if", "else"
    )

    val CHAR_PATTERN = """
        '(\\[btnr'"\\$] | \\u[a-fA-F0-9]{4} | [^\\'])'
    """.replace("\\s".toRegex(), "").toRegex()

    val ESCAPE_PATTERN = """
        \\[btnr'"\\$] | \\u[a-fA-F0-9]{4}
    """.replace("\\s".toRegex(), "").toRegex()

    val EOF = ""
    val LINE_END = setOf(";", "\n", EOF)

    /*
        Three reference types: immutable($), mutable(&), and exclusive(#)
        Vars are deallocated when block ends unless exclusivity has been given
            by fun call or assignment
        Only dynamic vars can give exclusive references
        Refs cannot be assigned to static vars

        val x = 5
        var y = 6
        val xRef: &Int32? = &x
        val xRef = auto<T>()

        val a_raw = 10
        val a = &a_raw
        val b = &a
        &a = &b

        &Int32?

        var z: &Int32? = none
        var z2 = none<&Int32>
    */
}
