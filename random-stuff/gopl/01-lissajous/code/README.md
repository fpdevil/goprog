# Lissajous Curve plotting

`Lissajous`  figure or  `Bowditch`  curve, is  the graph  of  a system  of parametric equations which describe complex harmonic motion.

These figures are created through a superposition of two perpendicular *sinusoidal waves*. One  mirror oscillates in the **X**  direction while the other one oscillates in the **Y**  direction.

The laser reflected from the mirrors traces  patterns which depend  on the relative  frequencies of the sounds. The below _2_ equations represent Lissajous figures.

```js
x = Asin(pt)
y = Bsin(qt + φ)
```

> where  the constants `A` & `B` represent  the amplitudes  along  `X` & `Y`
> directions The values `A` and `B`  determine the scale of the figure, with
> the entire graph being contained with in a box of dimension *2A X 2B*.

```text
p & q represents the angular frequencies along X & Y directions.
	1 ≤ p ≤ q
	if n = q/p (where n is irrational, the equations may be written as)

	x = Asin(t)
	y = Bsin(nt + φ) ; 0 ≤ φ ≤ π/2p

t is a parameter representing time passed in seconds
	0 ≤ t ≤ 2π

φ is the phase  shift which is inserted to account  for the change in
phases of the two vibrations.
	0 ≤ φ ≤ π/2
```
## Running the code
```shell
# open in browser running via a web server
⇒  go run main.go --options=web
# dump to a local file
⇒  go run main.go --options=file
# get the usage
⇒  go run main.go --help
Lissajous Figure Generation
------------------------------------
usage main -option=<web|file>
 - if option=web go to browser and check http://localhost:8000
 - if option=file open the output file lissajous.gif in a browser
```

Refer [mathcurve] for checking the curve details

[mathcurve]: https://mathcurve.com/courbes2d.gb/lissajous/lissajous.shtml
