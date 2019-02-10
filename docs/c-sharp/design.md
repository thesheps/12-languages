---
title: Month 1. C# - Object-Oriented Design
---

# Objects Vs. Classes
All of our snippets thus far have involved `Classes`. Our _Program_ class with a Main method which executes a simple Console application. Our `User` class with its cheeky little GetUsername method. Soooo what's an Object?

1. A `Class` is the _definition_ of an `Object`
   ```csharp
   public class User 
   {
   }
   ```

---

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

```csharp
public class Student : User
{
}

public class Teacher : User
{
    public bool CanTeach()
    {
        return true;
    }
}

public class Headteacher : Teacher
{
    public bool IsLeGrandeFromage()
    {
        return true;
    }
}

public class Program
{
    public static void Main()
    {
        var user = new Headteacher();
        Console.WriteLine(user.IsGrandeFromage());
        Console.WriteLine(user.CanTeach());
        // Prints:
        // True
        // True
    }
}
```
In _this_ example, the Headteacher class inherits from Teacher which in turn inherits from User. This is absolutely fine. What _isn't_ absolutely fine is:

```csharp
public class UberTeacher : Teacher, Headteacher
{
}
```
Es ist verboten. You cannot  inherit from multiple `classes` in different hierarchies. The chain of inheritance is through parents, grandchildren, great-grandchildren (etc).

__INTERESTING FACT:__ All classes in C# inherit from `Object` eventually. It's objects all the way down, people!

# Interfaces
Interfaces are "contracts" that classes need to abide by, or they'll be arrested by the Police!

```csharp
public interface IGun
{
    string Shoot();
}

public class WaterPistol : IGun
{
    public string Shoot()
    {
        return "Squirt!";
    }
}

public class LaserGun : IGun
{
    public string Shoot()
    {
        return "PewPew!!";
    }
}

public class Shotgun : IGun
{
    // I don't compile, I cause everything to break.
}
```

Interfaces don't have any actual implementation details. They just define a contract for the way in which a class can choose to behave. 

This is similar to `abstract` classes (yet-another-keyword); but abstract classes can include actual functionality, too!

```csharp
public abstract class Gun
{
    public abstract string Shoot();

    public int GetWeaponStrength()
    {
        return 10;
    }
}
```

# Method Overriding and Method Overloading
Say you've got a `class` called Gun. Say you want to change the behaviour of one of its methods. You can use the `virtual` method of the parent class to say "Hey! You can change what this does, if you like!"

You can then use the `override` keyword in a derived class to change the implementation.

```csharp
public class Gun
{
    public virtual string Shoot()
    {
        return "BANG!";
    }
}

public class Shotgun : Gun
{
    public override string Shoot()
    {
        return "BOOM!";
    }
}
```

But what if the new method I want to create takes a different number of parameters? This is called `overloading`:

```csharp
public class Gun
{
    public virtual string Shoot()
    {
        return "BANG!";
    }
}

public class Shotgun : Gun
{
    public string Shoot(string sound)
    {
        return sound;
    }
}
```

Note that you don't need to use the `override` modifier here. Both implementations of Shoot() are available!

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
        // This prints 'User'.

        Console.WriteLine(1.GetType());
        // This prints 'System.Int32'
    }
}
```

How does this work? C# is able to scan the contents of these Type classes using `Reflection`. It's unbelievably powerful, and allows you to do all manner of crazy shit like set property values dynamically, run methods etc. That's all I have to say on the matter right now.

[Next](./in-practise.md) we look at C# in practise. We'll give you a couple of code examples, we'll talk a little bit about unit testing and point you at a few resources for starting different flavours of applications.