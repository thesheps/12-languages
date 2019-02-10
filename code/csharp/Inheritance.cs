import System;

namespace Examples
{
    public class User
    {
        public virtual string Name { get; set; }
    }

    public class Student : User
    {
        public virtual override string Name { get { return "Student"; } }
    }

    public class Teacher : User
    {
        public virtual override string Name { get { return "Teacher"; } }
    }

    public class Headteacher : Teacher
    {
        public virtual override string Name { get { return "Headteacher"; } }
    }

    public class Program
    {
        public static void Main(string[] args)
        {
            Console.WriteLine(new Student().Name); // Prints "Student"
            Console.WriteLine(new Teacher().Name); // Prints "Teacher"
            Console.WriteLine(new Headteacher().Name); // Prints "Headteacher"
        }
    }
}