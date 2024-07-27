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
