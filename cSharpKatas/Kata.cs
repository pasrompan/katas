using System.Collections.Generic;
using System.Linq;
using System;
using System.Diagnostics;

public class Kata
{

    internal static string ToCamelCase(string v)
    {
       var output = "";

       string[] totalSubs = v.Split("-_".ToCharArray());

       bool isFirstWord = true;

       if (totalSubs.Count() <= 1)
       {
        return v;
       }

       foreach (var subitem in totalSubs)
       {
        string s1=subitem; 

        if (s1.Length <1){
          output+=s1;
          continue;
        }
        string s=s1.Substring(0,1);

        if (isFirstWord)
        {
          isFirstWord = false;
        }
        else
        {
          s = s.ToUpper();
        }
        string s2 = s1.Substring(1,subitem.Length - 1);
        output += s + s2;
       }

        return output;
    }

    internal static bool IsPrime(int n)
    {
        if (n < 2)
        {
          return false;
        }
        if (n<=3)
        {
          return true;
        }
        int half = (int) Math.Round(Math.Sqrt(n));
        int iterateStep = 1;
        for (var i=2; i<=half; i=i+iterateStep)
        {
          var modResult = n%i;
          if (modResult == 0 )
          {
            return false;
          }
          if (i==3)
          {
            iterateStep=2;
          }
        }
        return true;
    }
}