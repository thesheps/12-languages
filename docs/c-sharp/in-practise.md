---
title: Month 1. C# - In Practise
---

# Projects
Unlike the more hippy, scripty languages out there in the wild (Python, Javascript I'm looking at you); C# wants to give you a project file to work with. Depending on whatever flavour of .NET you're using, this might be an XML file or a Json file. Whatever the file looks like, it's going to contain references to each code file in your app, along with any environment variables you might want to set.

# Testing
Where's the first place you look when you clone a new repository? That's right, the tests! If you disagree with this comment, seek medical attention immediately.

So how do we unit test in C#? There are a bunch of different libraries to help with this (all accessible through the language's de facto package manager _NuGet_ - More on this in 2 shakes of a lamb's tail). Testing itself is a _massive_ subject and is such this session is only going to focus on the tooling.

## NUnit
To get going, I recommend using this testing framework. Although not Microsoft's official library for unit testing, its maturity and popularity across the industry makes this a great starting point.

So let's write a reeeeal simple test:

```csharp
import NUnit;

[TestClass]
public class GivenIAmWritingACodeExample()
{
    [TestMethod]
    public void WhenIHaveASimpleTest_ThenItPasses()
    {
        Assert.Pass();
    }
}
```

Simple, no? On line 3 we see yet-another C# _thing_ called annotations. Other languages have them too - Essentially you can _decorate_ both `classes` and `methods` with metadata contained in classes that inherit from the `Attribute` class. You can then stick them in square brackets before `class` and `method`. In this case, `[TestClass]` and `[TestMethod]` both allow the `NUnit Test Runner` to execute these tests and execute their _assertions_.

On line 10 we do the simplest-possible-thing: `Assert.Pass()` - All this does is _force_ an Nunit "Pass". Chances are you won't end up using IRL, but again this is a super-simple example.

# Package Management
C# ships with a package manager called `NuGet`. This is essentially a command-line tool which allows you to go:

```powershell
Install-Package Nunit # Powershell / .NET Framework
```

```bash
dotnet add nunit # Bash / .NET Core
```
This ends up:
- Adding a reference to your newly-installed project in whatever flavour project file you're working with.
- Creating a __local__ "cache" folder containing the downloaded resources for your package.

There are loads more docs [here](https://www.nuget.org/) if you want to take a look!

# Signing Off
I know this has been a bit of a whistle-stop tour on C#, but I hope you've enjoyed reading it as much as I've enjoyed putting it together.

To continue your journey, I recommend taking a look at the following docs:

- [The Xplat CLI](https://docs.microsoft.com/en-us/dotnet/core/tutorials/using-with-xplat-cli)
- [Gettings Started With ASP.NET Core](https://docs.microsoft.com/en-us/aspnet/core/getting-started/?view=aspnetcore-2.2&tabs=macos)
- [Doing Things With NUnit](https://github.com/nunit/docs/wiki/NUnit-Documentation)
- [The Yellow Book](http://www.csharpcourse.com/)