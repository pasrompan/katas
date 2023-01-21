package AllChallenges

import scala.annotation.tailrec

object CountBits {

  def countBits(num: Int): Int =
    @tailrec
    def countBits(num: Int, acc: Int): Int =
      if num == 0 then acc else
          countBits(num>>1, acc+(num & 1))
    countBits(num, 0)

}
