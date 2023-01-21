import AllChallenges.IsPrime.*

import java.math.*

// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class IsPrimeSuite extends munit.FunSuite {

  test("primeChecks"){
    val obtained = isPrime(2)
    val expected = true
    assertEquals(obtained, expected)
  }

  test("primeChecks") {
    val obtained = isPrime(13)
    val expected = true
    assertEquals(obtained, expected)
  }

  test("primeChecks") {
    val obtained = isPrime(9)
    val expected = false
    assertEquals(obtained, expected)
  }

}