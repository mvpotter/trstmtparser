package main

import (
	"flag"
	"github.com/mvpotter/trstmtparser/parser"
	"log/slog"
)

func main() {
	src := flag.String("src", "Statement.pdf", "Trade Republic PDF statement")
	dst := flag.String("dst", "Statement.csv", "Converted CSV statement")
	flag.Parse()

	transactions, err := parser.Parse(*src)
	if err != nil {
		slog.Error("unable to parse PDF", "err", err)
		return
	}

	err = parser.SaveCsv(*dst, transactions)
	if err != nil {
		slog.Error("unable to write CSV", "err", err)
		return
	}

	for _, transaction := range transactions {
		slog.Info("transaction", "data", transaction.ToRecords())
	}
}
