---
title: Month 1 - C#. Hello, World.
---

### Hello, World ###

```csharp
using System;

namespace Examples
{
   // Program class - Entry-point for application
   public class Program
   {
      public static void Main(string[] args)
      {
         Console.WriteLine("Hello World!");
      }
   }
}
```

Here we see the lesser-spotted Hello World in its natural habitat. Let's deep-dive into this code snippet and explain what everything does!

- **Line 1**

  Here we import the System _namespace_ such that we can use the Console _class_ on line 5.

- **Line 3**

  We've defined a namespace here called "Examples". Note the curlies which open and close the Namespace block. Everything in this code file is contained in this Namespace block.

- **Line 5**
  This is a simple comment. Single-line comments are specified using:

  ```csharp
  // Double Slashes
  ```

  and you can surround a few different lines (a code-block) with:

  ```csharp
  /*
  Slash-Asterisk Pairs
  */
  ```

- **Line 6**

  We've defined a Class here called Program. More on these shortly!

- **Line 8**

  This is the declaration for a _method_ called Main. It takes a _string array_ called args. This will be defined as the _entry point_ of the application.

- **Line 10**

  Here we execute the WriteLine _method_ of the Console _class_, and pass our "Hello World" string. Unlike the other lines in this snippet, We end this instruction with a semicolon. All statements in C# require terminating with a semicolon.

Now we've explored some of the _anatomy_ of a C# Hello World, [Click here](./deep-dive.md) as we move from dipping our toes to a full on snorkelling expedition into the clear, cool waters of C#!