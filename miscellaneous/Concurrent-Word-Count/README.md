# An Iterative and Concurrent version of Word Count Exercise

A program to process an input file to read and parse the individual words to get a final count of each of the indiviudal words.

This is another exercise from the excellent resources from [Go Language for Tourists]

## Tasks

- TODO:

    Write a golang application that count the number of each word in one or more input files.

### Requirements

The following are the specifications:

1. The program must _not_ use concurrency for this exercise.
2. The program must accept one or more filenames as input from the commandline.
3. After processing each input file, the program must print the _list of words_ found in _all_ of the files and a count of _how many times_ each word appeared.

   - Sample output:

     | **Count**    | **Word**      |
     | ------------ | ------------- |
     | 8            | Hendrikhovna  |
     | 30           | succeeded     |
     | 32           | Russia.       |
     | 2            | worry,        |
     | 1            | carpus        |

4. The program must print the _time_ taken for processing the word counting as its last output.
   - _NOTE_: This time need not include the time for printing out the results.
   - _TIP_: Use `time.Now()` at the start of the program and `time.Since()` at the end to get the elapsed time.

[Go Language for Tourists]: https://github.com/striversity/glft