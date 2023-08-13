# YNAB & Coinbase Sync

Currently, there is no official way to sync a [Coinbase](https://www.coinbase.com/) account to YNAB ([You Need A Budget](https://www.youneedabudget.com/)) for an accurate view of your Net Worth.

This Go cron uses the YNAB and Coinbase APIs to automatically update an investment or tracking account in YNAB. A [python version](https://github.com/Rjbaird/YNAB-Coinbase-Sync) is also available.

## Tech Stack

- [Go](https://go.dev/)
- [YNAB API](https://api.youneedabudget.com/v1)
- [Coinbase Wallet API](https://developers.coinbase.com/docs/wallet/api-key-authentication)
- [CoinAPI.io](https://www.coinapi.io/)
  
## Environment Variables

To run this project, you will need to add the following environment variables. The YNAB and Coinbase API keys are free and the CoinAPI is free up to 100 calls per day.

**YNAB**

`YNAB_TOKEN` 

`BUDGET_ID` 

`ACCOUNT_ID`

**Coinbase**

`COINBASE_KEY` 

`COINBASE_SECRET`

**CoinAPI**

`EXCHANGE_KEY`

**Optional**

`PORT`

## Feedback & Support

If you have any feedback or need support, feel free to reach out at rjbaird09@gmail.com