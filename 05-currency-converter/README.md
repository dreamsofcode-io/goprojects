# 05 - Currency Converter

The goal of this project is to create a currency converter TUI application.

## Requirements

- Convert between a number of base currencies. Recommended to start with USD,GBP,EUR & JPY.
- Use the huh package in order to create the terminal form
- You'll need to make use of a third party API in order to obtain currency conversion data.

## Packages Used

- net/http for http requests to the currency exchange api
- github.com/charmbracelet/huh for the TUI interface form
- encoding/json in order to marshal the data for the api

## Technical Considerations

### Currency API

There are a number of different apis out there in order to obtain exchange data. In my implementation
I went with [https://openexchangerates.org](https://openexchangerates.org/) which provides a free
API, but does require signing up.

### Conversion

The free api only provides conversions with USD as the base currency, however you can still
use it to perform conversions between other currencies.

The code I used for this was as follows:

```go
rateFrom := data.Rates[currencyFrom]
rateTo := data.Rates[currencyTo]

converted := amount * (rateTo / rateFrom)
```

### Supported Currencies

For the initial implementation, it may make sense to support a few currencies at first.

In my case, I supported only

- USD
- EUR
- GBP
- JPY

## Improvements

### More currency support

The API you're using should provide a list of currencies they support. One improvement would be to use this list in order to define the supported currencies of your app. You'll want to consider how to make this a good UX however.
