package utils

import (
	"fmt"
	"time"
)

var bulanIndo = map[time.Month]string{
	1:  "Januari",
	2:  "Februari",
	3:  "Maret",
	4:  "April",
	5:  "Mei",
	6:  "Juni",
	7:  "Juli",
	8:  "Agustus",
	9:  "September",
	10: "Oktober",
	11: "November",
	12: "Desember",
}

func FormatTanggalIndo(dateStr string) (string, error) {
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", err
	}
	day := parsedTime.Day()
	month := bulanIndo[parsedTime.Month()]
	year := parsedTime.Year()

	return fmt.Sprintf("%d %s %d", day, month, year), nil
}