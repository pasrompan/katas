import scala.util.Random

// Description:
// In a new system you need a list of ids of your user in many places.
// Please enrich List API and provide shortcut functions for getting this list.

trait OwnId {
  def id: Int
}

final case class User(id: Int, name: String) extends OwnId

// sample data
val l = Range.apply(1,5).map(i => User(i, Random.alphanumeric.take(2*i).toString())).toList

// test
l.ids == List(1,2,3,4)



// on of possible solutions

implicit final class ImprovedList[A <: OwnId](l: List[A]) {
  def ids: List[Int] = l.map(_.id)
}