package parser

import (
	"fmt"
	"github.com/goodsign/monday"
	"strconv"
	"strings"
	"time"
	_ "time/tzdata"
)

func parseDEDate(val string) (time.Time, error) {
	loc, _ := time.LoadLocation("Europe/Berlin")
	t, err := monday.ParseInLocation("2 January 2006", val, loc, monday.LocaleDeDE)
	if err != nil {
		t, err = monday.ParseInLocation("2 Jan. 2006", val, loc, monday.LocaleDeDE)
		if err != nil {
			return time.Time{}, fmt.Errorf("unable to parse date = %v, %w", val, err)
		}
	}

	return t, nil
}

func parseAmount(value string) (float64, error) {
	valueWithoutDots := strings.Replace(value, ".", "", 1)
	valueWithDotSeparator := strings.Replace(valueWithoutDots, ",", ".", 1)
	parsedFloat, err := strconv.ParseFloat(valueWithDotSeparator, 64)
	if err != nil {
		return 0, err
	}

	return parsedFloat, nil
}

func isInt(val string) bool {
	_, err := strconv.Atoi(val)
	return err == nil
}

func formatFloat(val float64) string {
	if val == 0 {
		return ""
	}
	return strconv.FormatFloat(val, 'f', -1, 64)
}
