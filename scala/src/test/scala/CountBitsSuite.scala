import AllChallenges.CountBits.*

import java.math.*
import scala.util.Random


// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class CountBitsSuite extends munit.FunSuite {

  test("test cases for count bits") {

    val testCases = List(
      (0, 0),
      (4, 1),
      (7, 3),
      (9, 2),
      (10, 2),
      (26, 3),
      (77231418, 14),
      (12525589, 11),
      (3811, 8),
      (392902058, 17),
      (1044, 3),
      (10030245, 10),
      (183337941, 16),
      (20478766, 14),
      (103021, 9),
      (287, 6),
      (115370965, 15),
      (31, 5),
      (417862, 7),
      (626031, 12),
      (89, 4),
      (674259, 10)
    )

    testCases.foreach {
      case (input, expected) =>
          val obtained = countBits(input)
          assertEquals(obtained, expected)
        }
    }
  }
