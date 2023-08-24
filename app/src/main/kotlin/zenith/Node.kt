package zenith

class Node(val type: String, val children: List<Faultable>): Faultable {
    override val faultPosition get(): UIntRange {
        val start = children.first().faultPosition.start
        val end = children.last().faultPosition.endInclusive
        return start..end
    }

    override fun toString() = treeString("")

    fun treeString(prefix: String): String {
        val builder = StringBuilder(type)

        for((i, child) in children.withIndex()) {
            val isLastChild = (i == children.size - 1)

            val childStr = when(child) {
                is Node -> {
                    val branch = if(isLastChild) " " else "│"
                    child.treeString("$prefix$branch   ")
                }
                else -> child.toString()
            }

            val branch = if(isLastChild) "└──" else "├──"
            builder.append("\n$prefix$branch $childStr")
        }

        return "$builder"
    }
}
