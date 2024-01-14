using System;
using System.Linq;

public class StripCommentsSolution
{
    public static string StripComments(string text, string[] commentSymbols)
    {
        return string.Join("\n", text.Split('\n')
            .Select(line => commentSymbols
                .Aggregate(line, (current, symbol) => current.Split(new[] { symbol }, StringSplitOptions.None)[0].TrimEnd())));
    }
}