package Day9

import java.io.File

class App {

    fun solve(lines: List<String>) {
        println("Puzzle solution part one: ${solvePartOne(lines)}")
        println("Puzzle solution part two: ${solvePartTwo(lines)}")
    }
    
    private fun solvePartOne(lines: List<String>): Int {
        val parsedLines = lines.map { line -> 
            line.split(" ").map { lineElement -> lineElement.toInt() } 
        }
        // parsedLines.forEach{ println("Next for line $it is ${findNextForLine(it)}") }
        return parsedLines.sumOf { findNextForLine(it) }
    }

    private fun solvePartTwo(lines: List<String>): Int {
        val parsedLines = lines.map { line -> 
            line.split(" ").map { lineElement -> lineElement.toInt() } 
        }
        // parsedLines.forEach{ println("Previous for line $it is ${findPreviousForLine(it)}") }
        return parsedLines.sumOf { findPreviousForLine(it) }
    }

    private fun findNextForLine(line: List<Int>): Int {
        val diff = mutableListOf<Int>()
        for(i in line.indices) {
            if(i != 0) {
                diff.add(line[i] - line[i-1]) 
            }
        }
        if (diff.all { it == 0 }) {
            return line.last()
        }
        return line.last() + findNextForLine(diff)
    }

    private fun findPreviousForLine(line: List<Int>): Int {
        val diff = mutableListOf<Int>()
        for(i in line.indices) {
            if(i != 0) {
                diff.add(line[i] - line[i-1]) 
            }
        }
        if (diff.all { it == 0 }) {
            return line.first()
        }
        return line.first() - findPreviousForLine(diff)
    }
}


fun main(args: Array<String>) {
  // access first argument
  val filename = args[0]

  val app = App()

  // read and print each line
  File(filename).useLines { lines ->
    app.solve(lines.toList())
  }
}
