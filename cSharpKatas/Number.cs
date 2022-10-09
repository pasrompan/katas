public class Number
{
  public int DigitalRoot(long n)
  {

    if (n < 10) return (int) n;
    
    var k = n % 10;
    var l = n / 10;
    var sum = k;
    while(l > 9)
    {
        k = l % 10;
        l = l/10;
        sum += k;
    }
    sum +=l;

    return DigitalRoot(sum);
  }
}