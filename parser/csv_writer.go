package parser

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func SaveCsv(dstFileName string, transactions []Transaction) error {
	_ = os.Remove(dstFileName)
	_ = os.MkdirAll(filepath.Dir(dstFileName), 0777)
	f, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_RDWR, 0755)
	defer f.Close()

	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	// Write headers
	err = w.Write([]string{"DATUM", "TYP", "BESCHREIBUNG", "ZAHLUNGSEINGANG", "ZAHLUNGSAUSGANG", "SALDO"})
	if err != nil {
		return fmt.Errorf("failed to write csv headers: %w", err)
	}

	for _, t := range transactions {
		if err := w.Write(t.ToRecords()); err != nil {
			return fmt.Errorf("failed to write csv row: %w", err)
		}
	}

	return nil

}
