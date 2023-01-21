import AllChallenges.CamelCase.*

import java.math.*
import scala.util.Random

// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class CamelCaseSuite extends munit.FunSuite {

  test("example test that succeeds") {
    val testWords = Array[String]("up", "down", "left", "right", "river", "Lake", "Square", "mountain", "desert", "north", "bridge", "south", "side", "Red", "Yellow", "Blue", "Green", "Black", "Samurai", "Ninja", "Rockstar", "Wall", "Street")
    val separators = Array[String]("-", "_")

    def makeTestCase(): (String, String, String) = { //input, expected, clue
      val separator = separators(Random.nextInt(separators.length))
      val words = List.fill(Random.between(2, 12))(testWords(Random.nextInt(testWords.length)))
      val input = words.mkString(separator)
      val expected = words.head ++ words.tail.map(_.capitalize).mkString
      val clue = if (input.head.isUpper) {
        s"\nInput delimited with $separator and leading uppercase should return UpperCamelCase"
      } else {
        s"\nInput delimited with $separator and leading lowercase should return lowerCamelCase"
      }
      (input, expected, clue)
    }

    val testCases = LazyList.continually(makeTestCase()).distinct.take(30)

    testCases.foreach {
      case (input, expected, clue) =>
        val obtained = toCamelCase(input)
        assertEquals(expected, obtained, clue)
    }
  }



}