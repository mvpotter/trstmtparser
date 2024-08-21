package parser

import (
	"fmt"
	"time"
)

type Transaction struct {
	Date    time.Time
	Type    string
	Descr   string
	In      float64
	Out     float64
	Balance float64
}

func (t Transaction) String() string {
	return fmt.Sprintf(
		"date = %s | type =  %s | descr = %s | in = %.2f | out = %.2f | balance = %.2f",
		t.Date,
		t.Type,
		t.Descr,
		t.In,
		t.Out,
		t.Balance,
	)
}

func (t Transaction) ToRecords() []string {
	return []string{
		t.Date.Format("02.01.2006"),
		t.Type,
		t.Descr,
		formatFloat(t.In),
		formatFloat(t.Out),
		formatFloat(t.Balance),
	}
}
