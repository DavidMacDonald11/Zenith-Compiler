package zenith

import java.io.File
import kotlin.system.exitProcess
import zenith.args.*
import zenith.lexer.*

private enum class Status(val code: Int) {
    GOOD(0),
    BAD_COMMAND(1),
    BAD_ARGUMENT(2),
    DIR_ALREADY_EXISTS(3),
    NOT_IN_PROJECT_DIR(4),
    COMPILER_ERROR(5)
}

fun main(args: Array<String>) {
    val command = args.getOrNull(0)

    val status = when(command) {
        null -> noCommand()
        in listOf("help", "--help", "-h") -> helpCommand()
        "init" -> initCommand(args.getOrNull(1))
        "build" -> buildCommand()
        "clean" -> cleanCommand()
        else -> {
            println("Error: There is no such command '${args[0]}'")
            println("Run 'zenith help' for usage information")
            Status.BAD_COMMAND
        }
    }

    exitProcess(status.code)
}

private fun noCommand(): Status {
    println("Error: A command is required")
    println("Run 'zenith help' for usage information")
    return Status.BAD_COMMAND
}

private const val USAGE_LINE_MAX = 50

private fun helpCommand(): Status {
    fun uprintln(arg: String, info: String) {
        var infoLength = 0

        for(word in info.split(" ")) {
            val oldLength = infoLength
            infoLength += word.length + 1

            if(infoLength > USAGE_LINE_MAX) {
                uprintln(arg, info.substring(0, oldLength))
                uprintln("", info.substring(oldLength))
                return
            }
        }

        print("  %-16s".format(arg))
        println("\t%-${USAGE_LINE_MAX}s".format(info))
    }

    println("Usage: zenith COMMAND [ARGS]")
    println("COMMAND:")
    uprintln("help", "shows this usage information")
    uprintln("init NAME", "creates a project folder named NAME")
    uprintln("build", "compiles all .zen files from src into bin")
    uprintln("clean", "deletes all compiled files from bin")

    return Status.GOOD
}

private fun initCommand(name: String?): Status {
    if(name == null) {
        println("Error: The 'init' command requires NAME argument")
        return Status.BAD_ARGUMENT
    }

    val dir = File(name)

    if(dir.exists()) {
        println("Error: Directory '${dir.path}' already exists")
        return Status.DIR_ALREADY_EXISTS
    }

    dir.mkdir()

    File(dir, "src").mkdir()
    File(dir, "lib").mkdir()
    File(dir, "bin").mkdir()
    File(dir, "zenith.config").writeText("")

    File(dir, ".gitignore").writeText("""
        |# Ignore binaries
        |bin
        |
    """.trimMargin())

    println("Created '${dir.path}' project folder")
    return Status.GOOD
}

private fun buildCommand(): Status {
    if(!inProject()) return Status.NOT_IN_PROJECT_DIR

    for(file in File("src").walk()) {
        if(!file.path.endsWith(".zen")) continue
        val faults = mutableListOf<Fault>()

        val tokens = tokenizeFile(file.path).let { result ->
            when(result.faults) {
                is None -> Unit
                is One -> faults.add(result.faults.value)
                is More -> faults.addAll(result.faults.values)
            }

            if(result.errored) {
                printFaults(file.path, faults)
                return Status.COMPILER_ERROR
            }

            when(result.tokens) {
                is None -> emptyList()
                is One -> listOf(result.tokens.value)
                is More -> result.tokens.values
            }
        }

        println("Created tokens: \n${tokens.joinToString(",", "[", "]")}")
        printFaults(file.path, faults)
    }

    return Status.GOOD
}

private fun cleanCommand(): Status {
    if(!inProject()) return Status.NOT_IN_PROJECT_DIR

    for(file in File("bin").walk()) {
        if(file.path == "bin") continue
        file.delete()
    }

    return Status.GOOD
}

private fun inProject(): Boolean {
    val inside = File("zenith.config").exists()
    if(!inside) println("Error: The current directory is not a zenith project")
    return inside
}
