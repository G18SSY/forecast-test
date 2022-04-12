# Forecast example  

This application is a small example for reading the current days assignments in [Go](https://go.dev/).

The script fetches from the unofficial [Forecast API](https://help.getharvest.com/forecast/faqs/faq-list/api/) to get todays assignments.

## Running
To get it running locally:

1. Install the latest version of [Go](https://go.dev/doc/install)
2. Create a config file in the root directory called `local.settings.json` with your access token and account ID. These can be created/found [here](https://id.getharvest.com/oauth2/access_tokens/new). Make sure you enter your Forecast ID and not your Harvest one (they're different).
``` JSON
{
    "accessToken": "",
    "accountId": ""
}
```
3. Run the script
``` bash
go run .
```

The output for today will be printed to your terminal window, if nothing shows then make sure that you have got work assigned today in Forecast.