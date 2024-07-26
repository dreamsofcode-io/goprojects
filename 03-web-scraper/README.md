# 03 - Web Scraper

The goal of this project is to build a web scraper to detect dead links in a website.

A dead link is defined as one that returns a status code in the range of 4xx or 5xx.

## Overview

The webscraper should recursively check every anchor tag on the website found on every page that belongs to the same domain as the website. If a page is on a different domain, the page itself should be checked that it's valid, but none of the links on that page should be checked.

Each discovered page should only be checked once, in order to prevent any infinite loops.

A simple decision tree diagram can be found here.

## Website

To assist with building the web scraper, I've created a website with a number of different web pages and links, some of them dead, and some of them working. 

You can either run this website yourself at localhost:8080 by entering into the [./scrapeme](./scrapeme) directory and running `go run .`

or you can use the hosted version at [https://scrape-me.dreamsofcode.io](https://scrape-me.dreamsofcode.io)

## Technical Considerations

This webscraper initially should focus on non JS rendered web pages, which require the use of a web browser based scraping solution, such as puppeteer or playwright.

## Packages

- `net/http` This package should be used to build the http client for pulling data
- `golang.org/x/net/html` Used to parse the HTML data into tokens, in order to pull out anchor tags.

## Additional Tasks

### Concurrency
Use concurrency to speed up the process. Running each page processing in it's own go routine, and passing the links through a channel to be then filtered or checked.

## Edge Cases

There are a number of different edge cases in this task such as:

- Ensuring that redirects are handled properly
- Only checking a page once
- Preventing infinite loops from recursion

