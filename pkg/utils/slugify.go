package utils

import (
	"regexp"
	"strings"
)

func Slugify(input string) string {
	// Türkçe karakterleri İngilizce karakterlere çevir
	turkishToEnglish := strings.NewReplacer(
		"ı", "i",
		"ğ", "g",
		"ü", "u",
		"ş", "s",
		"ç", "c",
		"ö", "o",
		"İ", "i",
		"Ğ", "g",
		"Ü", "u",
		"Ş", "s",
		"Ç", "c",
		"Ö", "o",
	)

	// Metni küçük harfe çevir
	input = strings.ToLower(input)

	// Türkçe karakterleri İngilizce karakterlere çevir
	input = turkishToEnglish.Replace(input)

	// Harf dışındaki karakterleri "-" ile değiştir
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		panic(err)
	}
	input = reg.ReplaceAllString(input, "-")

	// Başta ve sonda "-" karakterlerini kaldır
	input = strings.Trim(input, "-")

	return input
}
