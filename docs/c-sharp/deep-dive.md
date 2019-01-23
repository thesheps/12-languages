---
title: Month 1 - C#. Deep Dive.
---

### Deep-Dive
Ok. We've said a whole bunch about the history of the language, a few common uses and demonstrated a very simple "Hello World". Let's get our feet wet and take a look through some language specifics now!

#### Namespaces, Classes and Methods (oh my!) ###
Ok. Now we've explored a simple code snippet, let's review a few fundamentals in a bit more depth!

##### Namespaces ####
A namespace in C# is a _container_ for stuff. It defines a grouping of logical components. You don't _need_ to use namespaces, but it is highly recommended for an application of any complexity so that you can keep everything nicely organised.

If you want to reference a class within a namespace, you use the _import_ statement. Without specifying this using statement, we could refer to it by it's _fully-qualified name_:

```csharp
System.Console
```

This is a bit ugly. Don't do this.

The general naming convention for Namespaces is:

`CompanyName.TechnologyName[.Feature][.Design]`

You'll undoubtedly see loads more conventions in the wild, but if you stick to this formula you'll not go far wrong.

**Note:** It's super important not to name any _Class_ in your Namespace with the same name as your Namespace. Whilst the compiler might not shout at you, you run the risk of introducing insidious bugs and confusion into your codebase. Don't do this.

##### Classes & Methods #####
Classes are the building blocks of applications. Every piece of functionality you write in any program will be housed in a class.

Methods are blocks of codes that _optionally_ accept a bunch of parameters and _optionally_ return a value.

Consider the following code example:

```csharp
namespace Examples
{
   public class User
   {
   }
}
```

Here we've introduced an empty class called "User". It belongs to a namespace called "Examples". It doesn't do anything though. Lazy thing. Let's give it a job.

```csharp
namespace Examples
{
   public class User
   {
      public string GetUsername()
      {
         return "JoeBloggs";
      }
   }
}
```
- **Line 3** - The Class Declaration
   Here we're saying "Hey! This is a class. It is called User."

- **Line 5** - The Method Declaration

   ```public``` - This is an _access modifier_. More on these in just a tick!

   ```string``` - This specifies the _return type_ of the method. Put simply this specifies that this method _HAS_ to return a string. If you don't use the _return_ keyword (below); your code will not compile. If your method doesn't return a value (I.E. if it _does_ a thing rather than _returning_ a thing); you can use the ```void``` keyword here instead, and omit the ```return``` line entirely.

   ```GetUsername()``` - This is the name of our method. The open-close parentheses could house a few different parameters, but right now there aren't any. It is parameterless.

- **Line 7** - The Return
   
   Here we return a ```string``` with the value "JoeBloggs". _All_ strings in C# need to be enclosed in double quotes.

##### Instances #####
But how do we actually _call_ this method? In C# there are a few different options. As it currently stands, we can make an _instance_ of this class, and call our method directly:

```csharp
import System;

var user = new User();
var username = user.GetUsername();
Console.WriteLine(username);
```

- **Line 1** - Import

   We've seen this before - Let's import everything from the System namespace, so we can use its members without having to say "System.*"

- **Lines 3-5** - Newing Up
   
   My syntax highlighter is complaining that "Newing" is not a word:

   <p align="center">
      <img src="./img/newing.png" alt="I think you'll find it is.">
   </p>
   
   I promise it is, though! By using the ```new``` keyword, we _instantiate_ a new variable with the type User. We can then use its ```GetUsername``` method on line 4 in order to grab the Username string (covered above ^^^), and latterly write it to ```stdout``` using the ```WriteLine``` method of the Console class.

##### Access Modifiers #####
So far eeeeeverything we've done has been ```public``` (open to the containing Assembly and any _referencing_ assemblies). The above example _news up_ an _instance_ of the User class.

But what if I don't _want_ to open my User class up to the Universe?

But what if I don't _need_ an instance variable? Can I not just go:

```csharp
User.GetUsername();
```

In its current form? Nope. But 
> "There's an Access Modifier™ for that!"

Taken from the [Official Docs](https://docs.microsoft.com/en-us/dotnet/csharp/language-reference/keywords/access-modifiers) (sue me, I'm lazy)

> Access modifiers are keywords used to specify the declared accessibility of a member or a type.

```public```

Access is not restricted.

```protected```

Access is limited to the containing class or types derived from the containing class.

```internal```

Access is limited to the current assembly.

```protected internal```

Access is limited to the current assembly or types derived from the containing class.

```private```

Access is limited to the containing type.

```private protected```

Access is limited to the containing class or types derived from the containing class within the current assembly.

That's a lot of information, I know - but you'll find a lot of different use cases for each of these on your travels. Be safe out there.