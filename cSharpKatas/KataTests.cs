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
  private Number? num;
  
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

[Test]
    public void StripCommentsTest()
    {
        var input = "apples, pears # and bananas\ngrapes\nbananas !apples";
        var expected = "apples, pears\ngrapes\nbananas";
        var symbols = new string[] { "#", "!" };
        Assert.AreEqual(expected, StripCommentsSolution.StripComments(input, symbols));

        input = "a #b\nc\nd $e f g";
        expected = "a\nc\nd";
        symbols = new string[] { "#", "$" };
        Assert.AreEqual(expected, StripCommentsSolution.StripComments(input, symbols));
    }
}

public class ExampleTests
{
    [Test]
    public void TestBasic()
    {
        var expectations = new Dictionary<string, string[]>{
        { "8", new[] { "5", "7", "8", "9", "0" } },
        {"11",  new[]{"11", "22", "44", "12", "21", "14", "41", "24", "42" } },
        {"369", new[] { "339","366","399","658","636","258","268","669","668","266","369","398","256","296","259","368","638","396","238","356","659","639","666","359","336","299","338","696","269","358","656","698","699","298","236","239" } }
  };

        foreach (var pin in expectations)
        {
            CollectionAssert.AreEquivalent(pin.Value, GetPin.GetPINs(pin.Key), "PIN: " + pin);
        }
    }
}