using NUnit.Framework;
using System;

[TestFixture]
public class Tests
{
  [Test]
  public void KataTests()
  {
    Assert.AreEqual("theStealthWarrior", Kata.ToCamelCase("the_stealth_warrior"), "Kata.ToCamelCase('the_stealth_warrior') did not return correct value");
    Assert.AreEqual("TheStealthWarrior", Kata.ToCamelCase("The-Stealth-Warrior"), "Kata.ToCamelCase('The-Stealth-Warrior') did not return correct value");
  }

  [Test]
  public void StringTestsSimple()
  {
    Assert.AreEqual("", Kata.ToCamelCase("-_"), "wrong");
  }

  [Test]
  public void StringTestsBothDelimiters()
  {
    Assert.AreEqual("theStealthWarrior", Kata.ToCamelCase("the-stealth_warrior"), "wrong");
  }

  [TestFixture]
  public class SolutionTest
  {
    private static IEnumerable<TestCaseData> sampleTestCases
    {
      get
      {
        yield return new TestCaseData(0).Returns(false);
        yield return new TestCaseData(1).Returns(false);
        yield return new TestCaseData(2).Returns(true);
        yield return new TestCaseData(3).Returns(true);
        yield return new TestCaseData(9).Returns(false);
        yield return new TestCaseData(13).Returns(true);
      }
    }
  
    [Test, TestCaseSource("sampleTestCases")]
    public bool SampleTest(int n) => Kata.IsPrime(n);
  }
}