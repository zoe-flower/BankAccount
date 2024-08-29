// See https://aka.ms/new-console-template for more information
Console.WriteLine("Hello, World!");

//1
bool myBool = true; // turn to string
//2
string myString = "Hello World!"; // turn to list of chars
//3
double myDouble = -21434454.1423484; // turn to int

List<char> charList = new List<char>();

//1
string boolToString = myBool.ToString();

//2
for (int i = 0; i < myString.Length; i++)
{
 charList.Add(myString[i]); 
}

//3
int doubleToInt = (int)myDouble;

System.Console.WriteLine(boolToString.GetType());
System.Console.WriteLine(charList.GetType());
System.Console.WriteLine(doubleToInt.GetType());


//Abs, Pow, Sqrt, Floor, Ceiling, Min, Max
System.Console.WriteLine(doubleToInt);
System.Console.WriteLine(myDouble);
System.Console.WriteLine(Math.Abs(myDouble));
System.Console.WriteLine(Math.Pow(myDouble, 2));
System.Console.WriteLine(Math.Sqrt(-myDouble));
System.Console.WriteLine(Math.Floor(myDouble));
System.Console.WriteLine(Math.Ceiling(myDouble));
System.Console.WriteLine(Math.Min(myDouble, doubleToInt));
System.Console.WriteLine(Math.Max(myDouble, doubleToInt));

string interpolString = $"Hello {boolToString}";
System.Console.WriteLine("Hello {0}", boolToString);
