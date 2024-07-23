# Project 02 - Backend API

The goal of this project is to create an http+json API for a calculator service.

## Overview

An example version of this API can be found at https://calculator.dreamsofcode.io, which you can use whilst it's available. This calculator is stateless, meaning that there is no data stored in a database or in memory.

## Requirements

The API should conform to the given OpenAPI spec found in this directory, which can also be viewed at this URL.

As for the underlying technologies, feel free to use any that you may wish, however personally, I recommend using the standard library net/http package as much as possible, and using PostgreSQL as the backing database.

## Testing

To make testing the API specification easier, I've created an app that can be used to test the endpoints to see if they work correctly as per the specification. You can find this in the ./test directory and it can be ran using the `go test ./...` command.

# Additional Tasks

- Add in rate limiter to prevent misuse of the API
- Add in token authentication to prevent anyone unauthorized from using the API
- Add in a database to keep track of all of the calculations that have taken place
- Add in support for floating point numbers as well.
- Create an associated http client that can work with the calculator API.
