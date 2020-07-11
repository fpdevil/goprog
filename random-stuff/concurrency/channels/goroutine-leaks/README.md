# Leaks in GoRoutines and ways to avoid them

One of `go's` greatest strength is it's ability to quickly create, schedule and run `goroutines` quite efficiently.

Since `goroutines` are cheap and easy to create, it's one of the things which makes `go` such a productive language. It's similar to how `erlang` can be used to spawn multiple processes for concurrency.

The `go` runtime handles all the necessary multiplexing for spawning any number of operating system threads so that we don't have to worry about that level of abstraction. But still they do cost resources and `goroutines` are **not garbage collected** by the runtime. So, inspite of the small memory footprint of an individual `goroutine` we cannot leave a bunch of them lying forever. There should be some way of an effective cleanup.

## Paths of cleanup

The following are some of the paths which leads to cleanup of `goroutines`.

- Once the work allocated to `goroutine` has been completed

- If the `goroutine` cannot continue its allocated task/work due to an unrecoverable error.

- When it's explicitly advised or instructed to stop working. (_aka cancellation of goroutine_)


## Cancellation of goroutines

We can establish a signal between the paent `goroutine` and it's child `goroutines` which allow the parent to signal cancellation to it's children.

The usual convention is to have a `read-only channel` named as `done` for this purpose. The _parent_ `goroutine` will pass this channel to the _child_ `goroutine` and then closes the channel once it wants to cancel the _child_ `goroutine`.

_Note_: If a `goroutine` is responsible for creating another `goroutine`, it is also the responsibility of this `goroutine` to stop the one it created.
