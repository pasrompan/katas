package AllKatas

import scala.annotation.tailrec

object Kata {
  def toCamelCase(str: String): String =
    val splitArray = str.split("[_\\-]").toList
    (splitArray:List[String]) match
      case Nil => ""
      case x::tailL => innerCamelCase(tailL,x)

  @tailrec
  private def innerCamelCase(xs: List[String], existing: String):String =
    (xs:List[String]) match
      case Nil => existing
      case x::Nil => existing + x.capitalize
      case x::tailL => innerCamelCase(tailL, existing + x.capitalize)



  def countBits(num: Int): Int =
    @tailrec
    def countBits(num: Int, acc: Int): Int =
      if num == 0 then acc else
          countBits(num>>1, acc+(num & 1))
    countBits(num, 0)

  def f(a: String)(b: Int)(c: Boolean): String =
  "(" + a + ", " + b + ", " + c + ")"

  def isPrime(a: Int) : Boolean =

    @tailrec
    def iter(n: Int, acc: Int) :Boolean =
      n match
        case x if acc*acc>x => true
        case x if x % acc ==0 => false
        case _ => iter(n, acc+1)
    a match
      case x if x<=1 => false
      case x if x==2 => true
      case x if x>2 => iter(x, 2)

  val partialApplication1 = f("Scala")

  val partialApplication2 = partialApplication1(42)
}
