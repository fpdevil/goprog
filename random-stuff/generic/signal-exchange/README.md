# Cond

As per the official documentation, `Cond` is defined as below:

> Cond implements a condition variable, a rendezvous point
> for goroutines waiting for or announcing the occurrence
> of an event.

An `event` can be ay arbitrary signal between two or more goroutines, which carry no information apart from he fact that it has occurred. Very often, while working with execution of goroutines we might have to wait for one such signals

Using `Cond` provides a better and efficient approach for such case.

Let's see an example where there are a couple of goroutines with one waiting for a signal and the other sending signals. For instance, if we have a `queue` of fixed length `2`, and there are `10` items which needs to be pushed into the `queue`. We would like to `enqueue` the items as soon as there is enough room, so we want to be notified as soon as if there is such a soom available in the `queue`.
