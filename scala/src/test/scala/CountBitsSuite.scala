import AllChallenges.CountBits.*

import java.math.*

// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class CountBitsSuite extends munit.FunSuite {


  test("example test that succeeds") {
    val obtained = countBits(10)
    val expected = 2
    assertEquals(obtained, expected)
  }

}