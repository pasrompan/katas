package AllChallenges

import scala.annotation.tailrec

object CountBits {

  def countBits2(num: Int): Int = {
    @tailrec
    def countBits2(num: Int, acc: Int): Int =
      if num == 0 then acc else
        countBits2(num >> 1, acc + (num & 1))

    countBits2(num, 0)
  }


    def countBits(n: Int): Int = {
      @tailrec
      def loop(count: Int, num: Int): Int = num match {
        case 0 => count
        case _ => loop(count + (num & 1), num >>> 1)
      }

      loop(0, n)
    }

}
