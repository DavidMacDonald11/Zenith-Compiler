package zenith.io

import java.io.File
import kotlin.system.measureTimeMillis

enum class Status(val code: Int) {
    GOOD(0),
    BAD_COMMAND(1),
    BAD_ARGUMENT(2),
    DIR_ALREADY_EXISTS(3),
    NOT_IN_PROJECT_DIR(4),
    COMPILER_ERROR(5)
}

fun noCommand(): Status {
    println("Error: A command is required")
    println("Run 'zenith help' for usage information")
    return Status.BAD_COMMAND
}

private const val USAGE_LINE_MAX = 50

fun helpCommand(): Status {
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

fun initCommand(name: String?): Status {
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

fun buildCommand(buildFile: (String) -> Boolean): Status {
    if(!inProject()) return Status.NOT_IN_PROJECT_DIR
    var errored = false

    val time = measureTimeMillis {
        for(file in File("src").walk()) {
            if(!file.path.endsWith(".zen")) continue

            if(buildFile(file.path)) {
                errored = true
                break
            }
        }
    }

    println(" Build took ${"%,d".format(time)} ms")
    return if(errored) Status.COMPILER_ERROR else Status.GOOD
}

fun cleanCommand(): Status {
    if(!inProject()) return Status.NOT_IN_PROJECT_DIR

    for(file in File("bin").walk()) {
        if(file.path == "bin") continue
        file.delete()
    }

    return Status.GOOD
}

fun unknownCommand(command: String): Status {
    println("Error: There is no such command '$command'")
    println("Run 'zenith help' for usage information")
    return Status.BAD_COMMAND
}

private fun inProject(): Boolean {
    val inside = File("zenith.config").exists()
    if(!inside) println("Error: The current directory is not a zenith project")
    return inside
}
