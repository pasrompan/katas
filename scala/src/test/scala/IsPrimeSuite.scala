import AllChallenges.IsPrime.*

import java.math.*

// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class IsPrimeSuite extends munit.FunSuite {

  test("test cases for IsPrime"){
    test("test cases for count bits") {

      val testCases = List(
        (0, false),
        (1, false),
        (2, true),
        (73, true),
        (75, true),
        (-1, false),
        (100003, true),
        (100004, false)
      )

      testCases.foreach {
        case (input, expected) =>
          val obtained = isPrime(input)
          assertEquals(obtained, expected)
      }
    }
  }

}