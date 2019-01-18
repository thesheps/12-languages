# Month 1 - C#
To really understand the history of C#, we need to take a step back in time to the year 1990...

It was a year that saw cinema screens _rocked_ by the smouldering hair cuts of Demi Moore and Patrick Swayze. A year that the UK's Prime Minister Margaret Thatcher stepped down shortly after having introduced the _wildly_ popular Poll Tax. It was the year that the Hubble Space Telescope was launched from the space shuttle Discovery - propelling $1.5 billion of NASA technology to a modest altitude of some 340 miles above Earth. 

But even more impressive, even more _profoundly important_ - In a small town in California, a crack group of nerds had given birth to a C-based Form Generator known as __Ruby__ (no not that Ruby).

<p align="center">
<img src="./img/ruby_team.gif" alt="Lock up your daughters">
</p>

Immediately snaffled up by Bill Gates as a "cool" extension to one of their ongoing in-house projects (the outrageously-named __Project Thunder__) - __Visual Basic__ was introduced to baying hordes of listbox enthusiasts at the Atlanta World Trade Show in the spring of 1991.

<p align="center">
<img src="./img/bill.gif" alt="Needs more jpeg">
</p>

Whilst there were numerous languages preceding Visual Basic that enabled devs to create GUIs; this was the first time that a "drag 'n' drop" simplicity had been extended to an enterprise-y development environment. Not only could you now piece together multiple screens based off discrete components, but with its introduction of _Events_; you could easily create callback methods which would be executed on the click of a button/scroll of a listbox:


```vb
' mmmm listboxes

Private Sub List1_Click()
    Print "Clicked - " & Me.List1.List(Me.List1.ListIndex)
End Sub
```

Visual Basic would go on to have a shelf-life of some 18 more years before it was finally put out to pasture but, with its origins so deeply entrenched in the Windows API, its spirit continues to live on today.

## The .NET Framework
A somewhat rambling introduction to what I was aiming to be quite a _techy_ publication; but I thought a bit of history might be jolly before jumping straight into the language proper.

In 2002 Microsoft released the first candidate of the .NET Framework for Windows XP (Service Pack 1). Along with the Visual Studio .NET IDE which contained a bunch of redesigned __PROJECT THUNDER__ components (let's not forget those listboxes); .NET Framework 1.0 provided a CLR for their new flagship languages:

- VB .NET
- C# .NET*

What's a CLR I hear you ask? From the horse's mouth

> "[the] common language runtime, which runs the code and provides services that make the development process easier."

Put simply, the common language runtime is the virtual machine component of the .NET Framework and is responsible for actually executing your application. Code written in either VB .NET or C# .NET are both compiled to IL (intermediate language, or "byte code"), and executed by the CLR.

As for C#? This had been in development since 1999 by one of Borland's chief architects __Anders Hejlsberg__, as a C-Like Object Oriented Language. I literally just found out this was to be called __PROJECT COOL (C-Like Object Oriented Language)__ before pesky copyright legislation got in the way. To think I may otherwise have been able to advertise myself as a Cool Developer to any prospective employers. For shame.

In any case, C# .NET provided a similarly-native "Windowsy" experience to that of VB6/.NET; but with much more of a C-Like feel to proceedings:

```VB
' VB .NET
If nCnt <= nMax Then 
   nTotal += nCnt  
   nCnt += 1       
Else
   nTotal += nCnt
   nCnt -= 1       
End If
```

```C#
// C# .NET
if (nCnt <= nMax)
{
   nTotal += nCnt;
   nCnt++;
}
else
{
   nTotal +=nCnt;
   nCnt--;
}
```

Just look at those semicolons. Delicious.

