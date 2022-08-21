package AllKatas

import scala.annotation.tailrec

object Katas {
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
}
