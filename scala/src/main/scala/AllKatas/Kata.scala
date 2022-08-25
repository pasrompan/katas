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
}
