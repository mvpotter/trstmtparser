package parser

import (
	"strings"
)

func Parse(filePath string) ([]Transaction, error) {
	pdfRows, err := readPdf(filePath)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	var transactionRows [][]string
	var transactions []Transaction
	incomePosition := -1
	isPayment := true
	inColName := "ZAHLUNGSEINGANG"
	inColNameLength := len(inColName)
	for _, row := range pdfRows {
		inPos := strings.Index(row, inColName)
		if inPos != -1 {
			incomePosition = inPos
		}

		tokens := strings.Fields(row)
		if len(transactionRows) == 0 && len(tokens) > 0 && isInt(tokens[0]) {
			transactionRows = append(transactionRows, tokens)
		} else if len(transactionRows) == 2 {
			if len(tokens) > 0 && isInt(tokens[0]) {
				transactionRows = append(transactionRows, tokens)
				transactions = append(transactions, toTransaction(transactionRows, isPayment))
			}
			transactionRows = [][]string{}
			isPayment = true
		} else if len(transactionRows) > 0 {
			if len(transactionRows) == 1 && len(tokens) >= 3 {
				if len(strings.TrimSpace(row[incomePosition:incomePosition+inColNameLength-1])) != 0 {
					isPayment = false
				}
			}
			transactionRows = append(transactionRows, tokens)
		}
	}

	return transactions, nil
}

func toTransaction(transactionRows [][]string, isPayment bool) Transaction {
	row1 := transactionRows[0]
	row2 := transactionRows[1]
	row3 := transactionRows[2]

	dateStr := row1[0] + " " + row1[1] + " " + row3[0]
	tType := row2[0]
	note := ""
	amountStr := ""
	balanceStr := ""

	if len(row1) > 2 {
		note = strings.Join(row1[2:], " ")
	}

	if len(row2) > 1 {
		amountStr = row2[len(row2)-4]
		balanceStr = row2[len(row2)-2]
		note = note + " " + strings.Join(row2[1:len(row2)-4], " ")
	}

	if len(row3) > 1 {
		note = note + " " + strings.Join(row3[1:], " ")
	}

	date, err := parseDEDate(dateStr)
	if err != nil {
		panic(err)
	}

	amount, err := parseAmount(amountStr)
	if err != nil {
		panic(err)
	}

	balance, err := parseAmount(balanceStr)
	if err != nil {
		panic(err)
	}

	var inAmount float64
	var outAmount float64
	if isPayment {
		outAmount = amount
	} else {
		inAmount = amount
	}

	return Transaction{
		Date:    date,
		Type:    tType,
		Descr:   strings.TrimSpace(note),
		In:      inAmount,
		Out:     outAmount,
		Balance: balance,
	}
}
