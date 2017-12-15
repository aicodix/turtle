Who doesn't remember Turtle graphics?

Hacked this quickly to make a nice background image of the Hilbert curve using the Lindenmayer system:

```
// Hilbert curve
rules['A'] = "-BF+AFA+FB-"
rules['B'] = "+AF-BFB-FA+"
axiom := "A"
level := 6
```
The code above produces a nice image like this:

![hilbert.png](https://raw.githubusercontent.com/aicodix/turtle/master/hilbert.png)

Here some more examples:

```
// Koch curve
rules['F'] = "F+F-F-F+F"
axiom := "F"
level := 4
```
![koch.png](https://raw.githubusercontent.com/aicodix/turtle/master/koch.png)

```
// Dragon curve
rules['X'] = "X+YF+"
rules['Y'] = "-FX-Y"
axiom := "FX"
level := 12
```
![dragon.png](https://raw.githubusercontent.com/aicodix/turtle/master/dragon.png)
