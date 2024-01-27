using System;
using System.Collections.Generic;
using System.Linq;

public class GetPin
{
    private static readonly Dictionary<char, char[]> AdjacentDigits = new Dictionary<char, char[]>
    {
        { '0', new[] { '0', '8' } },
        { '1', new[] { '1', '2', '4' } },
        { '2', new[] { '1', '2', '3', '5' } },
        { '3', new[] { '2', '3', '6' } },
        { '4', new[] { '1', '4', '5', '7' } },
        { '5', new[] { '2', '4', '5', '6', '8' } },
        { '6', new[] { '3', '5', '6', '9' } },
        { '7', new[] { '4', '7', '8' } },
        { '8', new[] { '5', '7', '8', '9', '0' } },
        { '9', new[] { '6', '8', '9' } }
    };

    public static List<string> GetPINs(string observed)
    {
        return GetPINCombinations(observed).ToList();
    }

    private static IEnumerable<string> GetPINCombinations(string observed, string prefix = "")
    {
        if (string.IsNullOrEmpty(observed))
        {
            yield return prefix;
        }
        else
        {
            char digit = observed[0];
            string remaining = observed.Substring(1);

            foreach (char adjacent in AdjacentDigits[digit])
            {
                foreach (var combination in GetPINCombinations(remaining, prefix + adjacent))
                {
                    yield return combination;
                }
            }
        }
    }
}