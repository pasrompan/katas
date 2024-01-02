public static class TimeFormat
{
   public static string GetReadableTime(int seconds)
  {
    int hours = seconds / 3600;
    int minutes = (seconds % 3600) / 60;
    seconds = seconds % 60;

    return string.Format("{0:D2}:{1:D2}:{2:D2}", hours, minutes, seconds);
  }
}