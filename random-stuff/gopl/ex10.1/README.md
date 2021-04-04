# Exercise 10.1

Extend the `jpeg` program so that it converts any supported **input format** to any **output format**, using `image.Decode` to detect the _input format_ and a _flag_ to select the _output format_.

## Sample run

```shell
⇒  go run convert.go -format=png popeye.jpeg
2021/02/10 02:45:05 input image format is jpeg and intended format is png
```
