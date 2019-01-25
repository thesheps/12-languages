---
title: Month 1 - C#. Object-Oriented Design.
---

# Objects Vs. Classes
All of our snippets thus far have involved `Classes`. Our _Program_ class with a Main method which executes a simple Console application. Our `User` class with its cheeky little GetUsername method. Soooo what's an Object?

1. A `Class` is the _definition_ of an `Object`
   ```csharp
   public class User 
   {
   }
   ```

2. An `Object` is an _instance_ of a `Class`
   ```csharp
   var steve = new User();
   ```

# Inheritance
Inheritance is the relationship between "parent" and "child" Classes. When a class "inherits from" a parent, it also inherits any public methods and properties. It also must implement any `interface` that its parent implements (more on this shortly).

Consider this example:

```csharp
public class User
{
    public DateTime DateOfBirth { get; set; }
}

public class Student : User
{
}

public class Teacher : User
{
}
```

Both `Student` and `Teacher` are both derivatives/children of `User`. They both have access to a DateTime property called DateOfBirth.

# Types & Reflection
C# is a strongly-typed language. Every class in C# has "metadata" which describes what the class is named, where it lives and its mother's maiden name is. `Type` is a class, too. Which is kind of confusing.

```csharp
public class User
{
}

public static class Program
{
    public static void Main()
    {
        var steve = new User();
        Console.WriteLine(steve.GetType());

        // This prints the word 'User'.
    }
}
```