using System.Collections.Generic;
using System.Linq;
using System;
using System.Diagnostics;
using System.Text.RegularExpressions;
using System.Text;

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

    internal static int CountBits(long v)
    {
        int sumValue = 0;
        if (v < 0)
        {
          return 0;
        }

        var a = v;
        while (a>0){
          var b = 0b_0001;
          var c = a & b;
          if ((a & b) == 1)
          {
            sumValue++;
          }
          a = a >> 1;          
        }

        return sumValue;        
    }


    internal static bool IsPrime(int num)
    {
      if (num <=1)
        return false;

      if (num == 2)
        return true;
      
      var iterationStep = 1;
      for (int i = 2; i*i <= num; i = i+iterationStep)
      {
        var modu = num % i;
        if (modu == 0)
        {
          return false;
        }
        if (i==3)
        {
          iterationStep = 2;
        }
      }

      return true;
    }

    internal static string BreakCamelCase(string str)
    {
      var charArray = str.ToCharArray();
      StringBuilder sb = new StringBuilder("", 50);
      foreach (var c in charArray)
      {
        if (Char.IsUpper(c))
        {
          sb.Append($" ");
        }
        sb.Append(c);

      }
      return sb.ToString();


    }
}

//https://www.codewars.com/kata/526571aae218b8ee490006f4/train/csharp