import AllKatas.Katas.toCamelCase

// For more information on writing tests, see
// https://scalameta.org/munit/docs/getting-started.html
class ToCamelCaseTests extends munit.FunSuite {
  test("example test that succeeds") {
    val obtained = toCamelCase("the_Stealth_Warrior")
    val expected = "theStealthWarrior"
    assertEquals(obtained, expected)
  }
}