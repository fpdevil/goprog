# Word Frequencies

Perform textual analysis over the supplied input files in order to get a count of frequencies of words in the files.

The frequency data is presented in two different and equally sensible formats:

- as an alphabetical list of words with their freqencies and
- as an ordered list of freqency values and the words that have the corresponding frequencies.

## Output

```shell
⇒  go run main.go README.md
{"level":"info","main":"word-frequencies","msg":"starting Word Frequency service...","time":"2021-02-14T02:21:42-06:00"}
Word  Frequency
bss         1
buildinfo   1
data        8
frequencies 1
go          1
gopclntab   1
gosymtab    1
itablink    1
linkedit    1
nl          1
noptrbss    1
noptrdata   1
pagezero    1
ptr         1
rodata      1
stub        1
symbol      2
text        9
typelink    1
word        1
Frequency → Words
1 bss, buildinfo, frequencies, go, gopclntab, gosymtab, itablink, linkedit, nl, noptrbss, noptrdata, pagezero, ptr, rodata, stub, typelink, word
2 symbol
8 data
9 text
```
