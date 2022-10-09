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

  [Test]
  public void CountBits()
  {
    Assert.AreEqual(0, Kata.CountBits(0));
    Assert.AreEqual(1, Kata.CountBits(4));
    Assert.AreEqual(3, Kata.CountBits(7));
    Assert.AreEqual(2, Kata.CountBits(9));
    Assert.AreEqual(2, Kata.CountBits(10));
    Assert.AreEqual(16, Kata.CountBits(6015773801));
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


     private static IEnumerable<TestCaseData> testCases
    {
      get
      {
        yield return new TestCaseData("camelCasing").Returns("camel Casing");
        yield return new TestCaseData("camelCasingTest").Returns("camel Casing Test");
      }
    }
  
    [Test, TestCaseSource("testCases")]
    public string Test(string str) => Kata.BreakCamelCase(str);


[TestFixture]
public class NumberTest
{
  private Number num;
  
  [SetUp]
  public void SetUp() 
  {
    num = new Number();
  }

  [TearDown]
  public void TearDown()
  {
    num = null;
  }

  [Test]
  public void Tests()
  {
    Assert.NotNull(num);
    Assert.AreEqual(7, num.DigitalRoot(16));       
    Assert.AreEqual(6, num.DigitalRoot(456)); 
    Assert.AreEqual(6, num.DigitalRoot(132189));
    Assert.AreEqual(9, num.DigitalRoot(9));
    Assert.AreEqual(9, num.DigitalRoot(2000000000000000007));
  }
}
  }
}