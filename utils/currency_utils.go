package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Format ke mata uang Rupiah
func FormatRupiah(amount int) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("Rp %d", amount)
}