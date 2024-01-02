public class User
{
    public int rank { get; set; }
    public int progress { get; set; }

    public User()
    {
        rank = -8;
        progress = 0;
    }

    public void incProgress(int activityRank)
    {
        if (activityRank < -8 || activityRank > 8 || activityRank == 0)
            throw new ArgumentException("Invalid activity rank");

        if (rank == 8)
            return;

        int rankDifference = activityRank - (rank < 0 && activityRank > 0 ? rank + 1 : rank);

        if (rankDifference == 0)
            progress += 3;
        else if (rankDifference == -1)
            progress += 1;
        else if (rankDifference > 0)
            progress += 10 * rankDifference * rankDifference;
        
        //10 x (r)^2 +c = 121

        while (progress >= 100 && rank != 8)
        {
            progress -= 100;
            rank = rank == -1 ? 1 : rank + 1;

            if (rank == 8)
            {
                progress = 0;
                break;
            }
        }
    }
}