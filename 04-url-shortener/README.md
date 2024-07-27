# 04 - URL Shortener

The goal of this project is to create a full URL shortener web app
using Go.

To see this in action, I have my own version hosted at [https://dree.my](https://dree.my)

Please feel free to reference this app when building your own.

## Requirements

- Web page where you can submit a form with a long URL
- Upon submitting web page, the user should then be taken to a new page which has their shortened URL.
- When using the shortened URL, it should automatically redirect the user to the long url.

## Technical Considerations

### Frontend

The frontend is should be built entirely using Go. Initially, I'd recommend using the html/template package of the standard library in order to gain a better understanding of how to use templating in Go. You can find an example site in this directory which uses it and shows how to use it with the standard library.

### Redirection

Redirection is performed in Go using the `http.Redirect` function of the `net/http` package.

This function expects the URL to redirect to, and the status code to use. In this case you'll want to decide between the http.StatusFound and the http.PermanentlyMoved status codes, and which one is preferrable. You can wla

### Database

For this, you can choose any data store you want, personally I'd recommend using something such as Valkey/Redis however provided it has persistence enabled.
