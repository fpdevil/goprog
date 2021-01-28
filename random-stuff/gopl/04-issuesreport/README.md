# Chapter 4. Section 4.6 Text And HTML Templates

Here the book describes about basics of templates and walks through an example related to gthub issue report in `textual template` format.

## Running

Here is the sample output from running the code

```bash
go run main.go repo:golang/go is:open json decoder

INFO[0000] now calling the api https://api.github.com/search/issues?q=repo%3Agolang%2Fgo+is%3Aopen+json+decoder
51 issues:
	----------------------------------------
	Number: 40443
	User: KaiserKarel
	Title: proposal: add *json.Decoder to return type of DisallowUnknownFie
	Age: 4 days
	----------------------------------------
	Number: 33416
	User: bserdar
	Title: encoding/json: This CL adds Decoder.InternKeys
	Age: 366 days
	----------------------------------------
	Number: 40128
	User: rogpeppe
	Title: proposal: encoding/json: garbage-free reading of tokens
	Age: 23 days
```