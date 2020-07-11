# Multiple Producers Running Sequentially

This is the exercise from `section 07` of the [Golang For Tourists] series from `verrol`

The target for this exercise is to write a complete `golang` application to demonstrate *Multiple Producers writing to the same channel sequentially*. The *max messages per producers* and *the number of producers* should be specified by the program command line parameters as `'-m'` and `'-np'` respectively.

## Specifications & Requirements

### TODO 1 - Use the 'flag' standard golang package to handle and parse the command line arguments '-m' and '-np' when provided.

`Golang` provides the `flag` package for parsing the command-line options. Hence, we don't need to check the `os.Args` variable directly for any command line arguments.

If no command-line options are provided the, program should consider default values of `'m = 100'` and `'np = 3'`.

The program should show a standard *usage message* for any invalid value(s) like `-1` for example.

* **TIP**: The function `flag.Usage()` would be useful when an invalid value is provided as argument. See the `flag` package documentation for examples.

    Here is an excerpt from `godoc` for `flag.Usage()`

    ```go
    var Usage = func() {
        fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
        PrintDefaults()
    }
        Usage prints a usage message documenting all defined command-line flags to
        CommandLine's output, which by default is os.Stderr. It is called when an
        error occurs while parsing flags. The function is a variable that may be
        changed to point to a custom function. By default it prints a simple header
        and calls PrintDefaults; for details about the format of the output and how
        to control it, see the documentation for PrintDefaults. Custom usage
        functions may choose to exit the program; by default exiting happens anyway
        as the command line's error handling strategy is set to ExitOnError.
    ```

### TODO 2 - Create a channel for string messages at the package level scope.

Our `channel` must be capable of holding all of the messages that can be produced, which is `np * m`. See sections *TODO 3 & 4*.

### TODO 3 - Write a producer function

Our `producer` function should handle the following:

  1. Takes as parameters an `id`
  2. Generate a `random number of messages between 1 to 'm'` messages to the `channel` containing its `name/id` and a `random number`.

  * For example, a message from a `producer` might look like: `"Producer 1, num: 25"`

### TODO 4 - Launch 'n' producers

For each `producer`, assign a `unique id` and a `shared channel`.

### TODO 5 - Write a consumer function

The `Consumer` is responsible for handling the below:

1. Extracting the `producer's name/id` and `random number` from the message.
2. Print the `number of messages` and their `_sum_` from each `producer`.
3. Print the `total number of messages` **sent** and `total sum` across `producers`.
4. Consumer _MUST_ use the `_range_` operator for consuming `messages` from the `channel`.

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