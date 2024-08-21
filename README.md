# Trade Republic Statement Parser

[Trade Republic Bank](https://traderepublic.com) allows to download statements in PDF format only.
It is not handy when you want to export transactions to expenses tracker. So the module allows to 
parse PDF statement and convert it to CSV table.

## Usage

```bash
go run ./cmd/parser -src path-to-trade-republic-statement.pdf
```
