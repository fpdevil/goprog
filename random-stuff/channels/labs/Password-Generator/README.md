# Password Generator

This is the exercise from `section 07` of the [Golang For Tourists] series from `verrol`

The target for this exercise is to write a complete `golang` application to `generate unique passwords`.

## Specifications & Requirements

Below are all the required specifications to be enforced.

1. If using a *Western keyboard*, use the following character classes in your password:
   * *letters*: `a-z` and `A-Z`
   * *numbers*: `0-9`
   * *special characters*: `-=~!@#$%^&*()_+[]{}|;':",./<>?`
2. User may specify password length using `'-l'` option, with the minimum length of `8` and a maximum length of `128`. If no value is provided, a default value of `16` will be considered.
   * _TIP_: Use the `standard flag pacakge`.
3. **MUST** use `select statement` to pick character from `character class`
4. Prefer more *letters* to *numbers*, and more `numbers` to `special characters`.
   * **HINT**: Use `case statements` to increase probability of one character set over another.

## Result

*Sample Run*

```bash
Sample Run
⇒  go run main.go -m 10 -np 3
/// Multiple Producers Running Sequentially ///

Producer #001
 	Number of Elements: 5
	Sub-total: 28
Producer #002
 	Number of Elements: 7
	Sub-total: 94
Producer #003
 	Number of Elements: 5
	Sub-total: 62
Total count: 17
Total sum: 184
```


[Golang For Tourists]: https://github.com/striversity/glft/tree/master/sec07