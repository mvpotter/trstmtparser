# Trade Republic Statement Parser

[Trade Republic Bank](https://traderepublic.com) allows to download statements in PDF format only.
It is not handy when you want to export transactions to expenses tracker. So the module allows to 
parse PDF statement and convert it to CSV table.

## Usage

### From sources

```bash
go run ./cmd/parser -src path-to-trade-republic-statement.pdf
```

### From docker

It is necessary to bind a volume, where statement PDF is located.
In the following example, statement in located in the `stmt` folder of the directory where the command is being 
launched.
CSV destination should be provided for the same folder. Otherwise, the result will be written inside the container.

```bash
docker run -it -v ./stmt:/stmt mvpotter/trstmtparser:v0.0.1 -src /stmt/Statement.pdf -dst /stmt/Statement.csv
```
