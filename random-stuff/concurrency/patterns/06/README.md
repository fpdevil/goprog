# Concurrency pattern 6 - Fan-Out

`Stages` in `pipeline` can be occassionally computationally expensive. The upstream stages of a pipeline may go into a blocked state if the downstream stages take more time for processing because of them being computationally expensive. During such circumstances, we can combine multiple goroutines to handle the input from a pipeline

This particular pattern in which multiple goroutines are spawned to handle input from the pipeline is called `Fan-Out`.

## Process of Fan-Out

Handling with Fan-Out is pretty easy. We just start or spawn multiple processes to handle the task or work.