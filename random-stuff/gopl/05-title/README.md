# Handling deferred function calls

The `findLinks` code used the output of `http.Get` as the input to `html.Parse`.

This works well if the content of the requested *URL* is indeed *HTML*, but many pages contain *images, plain text, and other file formats*.

Feeding such files into an `HTML parser` could have undesirable effects.
