# The OR Channel

Some times, we might have to combine one or more `done` channels together into a single `done` channel that closes if any of it's component channels close. The number of `done` channels are usually not known upfront during such cases.

During such cases, an **or-channel** pattern is quite helpful in combining the channels together.

```sh
⇒  go run --race main.go

OR Channel
done after 1.000582325s
```
