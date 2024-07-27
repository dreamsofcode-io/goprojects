# Project 02 - Backend API

The goal of this project is to create an http+json API for a calculator service.

## Overview

An example version of this API can be found at https://calculator.dreamsofcode.io, which you can use whilst it's available. This calculator is stateless, meaning that there is no data stored in a database or in memory.

## Requirements

The API should conform to the given OpenAPI spec found in this directory, which can also be viewed at this URL.

### Production Ready

In order to make this API more production ready, there's a few other requiements you'll need to consider

#### Input Validation

You should never trust input from the user. This means you'll need to be sure to validate and sanitize any inputs. Division by zero is a common cause
panics when it comes to applications, so you'll want to make sure you're handling it.

#### Error feedback
Additionally, it's a good idea to let the user know when they've made a mistake with input, so they can fix it if the mistake was innocent.

#### Logging

In order to be able to debug issues that occur, you're going to want to log out each request as it comes in, as well as any associated data such as the status code, ip address, and what the request path was.

For logging in Go, I recommend using the `log/slog` package, which provides structured logging in either JSON, or Text format. You can create a logger using the following snippet:

##### Text
```go
logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
```

##### JSON
```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
```

#### CORS

You may want to add cors in if you intend to hit this directly from a website. The [following package](github.com/rs/cors) makes it really easy to set up cors

## Recommended Packages

Some of the packages that I used for my implementation include:

- `net/http` - This is the http package of the standard library which I use for routing and setting up an http server. If you want to know how to use this package for advanced routing, [I have a video you can check out](https://youtu.be/H7tbjKFSg58).
- `encoding/json` - This package is used for encoding and decoding JSON from the request and to the response bodies.
- `log/slog` - For structured logging
- `github.com/rs/cors` - For cors, if you need it.

## Resources

You can access my implementation of this API at [https://calculator.dreamsofcode.io](https://calculator-api.dreamsofcode.io)

## Additional Tasks

- Add in rate limiter to prevent misuse of the API
- Add in token authentication to prevent anyone unauthorized from using the API
- Add in a database to keep track of all of the calculations that have taken place
- Add in support for floating point numbers as well.
- Create an associated http client that can work with the calculator API.
- Create a frontend that makes use of your API.
- Add in a middleware that adds a request ID to the http.Request object.
