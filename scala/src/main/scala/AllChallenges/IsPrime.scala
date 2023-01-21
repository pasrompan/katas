package AllChallenges

import scala.annotation.tailrec

object IsPrime {

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
}
