Who doesn't remember Turtle graphics?

Hacked this quickly to make a nice background image of the Hilbert curve using the Lindenmayer system:

```
// Hilbert curve
rules := strings.NewReplacer(
	"A", "-BF+AFA+FB-",
	"B", "+AF-BFB-FA+")
axiom := "A"
level := 6
```
The code above produces a nice image like this:

![hilbert.png](https://raw.githubusercontent.com/aicodix/turtle/master/hilbert.png)

Here some more examples:

```
// Koch curve
rules := strings.NewReplacer(
	"F", "F+F-F-F+F")
axiom := "F"
level := 4
```
![koch.png](https://raw.githubusercontent.com/aicodix/turtle/master/koch.png)

```
// Dragon curve
rules := strings.NewReplacer(
	"X", "X+YF+",
	"Y", "-FX-Y")
axiom := "FX"
level := 12
```
![dragon.png](https://raw.githubusercontent.com/aicodix/turtle/master/dragon.png)
