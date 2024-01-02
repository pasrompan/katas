using System;

public class User1
{
    private static readonly int[] Ranks = new int[]
    {
        -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8
    };

    public int rank { get; set; }
    public int progress { get; set; }

    public User1()
    {
        rank = Ranks[0]; // Start at -8
        progress = 0;
    }

    private int RankIndex(int rank1)
    {
        if (rank1 == 0 || rank1 < -8 || rank1 > 8)
            throw new ArgumentException("Invalid rank");

        return Array.IndexOf(Ranks, rank1);
    }

    public void incProgress(int activityRank)
    {
        if (activityRank == 0 || activityRank < -8 || activityRank > 8)
            throw new ArgumentException("Invalid activity rank");

        if (rank == Ranks[^1]) // Max rank reached
            return;

        int currentRankIndex = RankIndex(rank);
        int activityRankIndex = RankIndex(activityRank);
        int rankDifference = activityRankIndex - currentRankIndex;

        if (rankDifference == 0)
        {
            progress += 3;
        }
        else if (rankDifference == -1)
        {
            progress += 1;
        }
        else if (rankDifference > 0)
        {
            progress += 10 * rankDifference * rankDifference;
        }

        UpdateRank();
    }

    private void UpdateRank()
    {
        while (progress >= 100 && rank != Ranks[^1])
        {
            progress -= 100;
            rank = Ranks[RankIndex(rank) + 1];
        }

        if (rank == Ranks[^1]) // If max rank reached, reset progress to 0
            progress = 0;
    }
}