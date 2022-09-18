import AllKatas.Kata.*

import java.math.*

// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class KataSuite extends munit.FunSuite {
  test("example test that succeeds") {
    val obtained = toCamelCase("the_Stealth_Warrior")
    val expected = "theStealthWarrior"
    assertEquals(obtained, expected)
  }


  test("example test that succeeds") {
    val obtained = countBits(10)
    val expected = 2
    assertEquals(obtained, expected)
  }

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