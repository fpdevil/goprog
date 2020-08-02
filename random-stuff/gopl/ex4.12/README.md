# Exercise 4.12 XKCD

**Work In Progress**

The popular web comic `xkcd` has a `json` interface. For example, a request to `https://xkcd.com/571/info.0.json` produces a detailed description of comic *571*. Download each `URL` (once!) and build  an offline index. Write a tool `xkcd` that, using this index prints the `URL` and `transcript` of each comic that matches a search term provided on the command line.

## JSON data

Here is the `json` obtained from `xkcd`.

```json
{
  "month": "4",
  "num": 571,
  "link": "",
  "year": "2009",
  "news": "",
  "safe_title": "Can't Sleep",
  "transcript": "[[Someone is in bed, presumably trying to sleep. The top of each panel is a thought bubble showing sheep leaping over a fence.]]\n1 ... 2 ...\n<<baaa>>\n[[Two sheep are jumping from left to right.]]\n\n... 1,306 ... 1,307 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow.]]\n\n... 32,767 ... -32,768 ...\n<<baaa>> <<baaa>> <<baaa>> <<baaa>> <<baaa>>\n[[A whole flock of sheep is jumping over the fence from right to left. The would-be sleeper is sitting up.]]\nSleeper: ?\n\n... -32,767 ... -32,766 ...\n<<baaa>>\n[[Two sheep are jumping from left to right. The would-be sleeper is holding his pillow over his head.]]\n\n{{Title text: If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.}}",
  "alt": "If androids someday DO dream of electric sheep, don't forget to declare sheepCount as a long int.",
  "img": "https://imgs.xkcd.com/comics/cant_sleep.png",
  "title": "Can't Sleep",
  "day": "20"
}
```

We use the excellent `json` to `go struct` converter tool [json-to-go] to convert the raw response into a struct.


[json-to-go]: https://mholt.github.io/json-to-go/
