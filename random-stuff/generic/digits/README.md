# Print the digits

An exercise **1.8** from the excellent book _Programming in Go: Creating Applications for the 21st Century_

```bash
⇒  go build -o digits

⇒  ./digits

	usage: digits [-b | --bar] <whole-number>
	-b --bar draw an underbar and an overbar

⇒  ./digits --bar 369
***************************
3333333  6666666  9999999
      3  6        9     9
      3  6        9     9
3333333  6666666  9999999
      3  6     6        9
      3  6     6        9
3333333  6666666        9
***************************

```
