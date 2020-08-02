# Exercise 4.13 from The Go Programming Language

The `JSON` based web service of the `Open Movie Database` lets you search https://omdbapi.com/ for a movie by name and download its poster image. Write a tool poster that downloads the poster image for the movie named on the command line.

## OMDB - The Open Movie Database

`The Open Movie Database API` is a RESTful web service to obtain movie information from `IMDB`. All it needs is an `apikey` which can be obtained free by registering with their site. `OMDB Api` have generously made the service available for free.

Once we obtain the `apikey`, looking for the movie information is as simple as the below call.

```html
http://www.omdbapi.com/?apikey=<apikey_value>&t=<movie_name>
```

## poster.go

`poster.go` takes the required parameters `apikey` and the `movie name` as per the contract and fetches the movie information, parses the `json` data for `poster` link and fetches the same from it's link and stores it locally as per the exercise. In case if the movie title contains spaces, we will convert the spaces to `_` and write to file system.

### Examples

Here are a couple of examples:

1. Search for *Interstellar* movie.

```bash
⇒  go run poster.go 1a678916 "Interstellar"
INFO[0000] calling movie db @ https://omdbapi.com/?apikey=1a678916&t=Interstellar
INFO[0000] fetching the poster from https://m.media-amazon.com/images/M/MV5BZjdkOTU3MDktN2IxOS00OGEyLWFmMjktY2FiMmZkNWIyODZiXkEyXkFqcGdeQXVyMTMxODk2OTU@._V1_SX300.jpg
```

The above fetches and stores the below movie poster.

![Raiders Of The Lost Ark](Interstellar.jpg)

2. Search for *Raiders of the Lost Ark* movie.

```bash
⇒  go run poster.go 1a678916 "Raiders Of The Lost Ark"
INFO[0000] calling movie db @ https://omdbapi.com/?apikey=1a678916&t=Raiders+Of+The+Lost+Ark
INFO[0000] fetching the poster from https://m.media-amazon.com/images/M/MV5BMjA0ODEzMTc1Nl5BMl5BanBnXkFtZTcwODM2MjAxNA@@._V1_SX300.jpg
```

The above fetches and stores the below movie poster.

![Raiders Of The Lost Ark](Raiders_of_the_Lost_Ark.jpg)

3. Search for *Non existent* movie.

We will end up with an error as below and nothing is written to the image file.

```bash
⇒  go run poster.go 1a678916 "omdbapi"
INFO[0000] calling movie db @ https://omdbapi.com/?apikey=1a678916&t=omdbapi
INFO[0000] fetching the poster from
error calling : Get "": unsupported protocol scheme ""ERRO[0000] error writing 'Get "": unsupported protocol scheme ""'
```