package AllChallenges

import scala.annotation.tailrec

object CamelCase {


  def toCamelCase2(str: String): String =
    val splitArray = str.split("[_\\-]").toList
    (splitArray: List[String]) match
      case Nil => ""
      case x :: tailL => innerCamelCase(tailL, x)

  @tailrec
  private def innerCamelCase(xs: List[String], existing: String): String =
    (xs: List[String]) match
      case Nil => existing
      case x :: Nil => existing + x.capitalize
      case x :: tailL => innerCamelCase(tailL, existing + x.capitalize)

  def toCamelCase(str: String): String = {
    val new_str = str.split("[_-]")
    val head = new_str.head
    val tail = new_str.tail.map(a => a.capitalize).mkString("")
    head + tail
  }



}
