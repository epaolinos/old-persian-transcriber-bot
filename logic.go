package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func transcribe(input string) string {
	var builder strings.Builder

	runes := []rune(strings.ToLower(input))

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		if i+1 < len(runes) {
			next := runes[i+1]
			syll := string(char) + string(next)

			if val, ok := contextualSigns[syll]; ok {
				builder.WriteString(val)
				i++
				continue
			}
		}

		if val, ok := baseSigns[char]; ok {
			builder.WriteString(val)

			if i+1 < len(runes) {
				vocals := []rune{'a', 'ā', 'e', 'i', 'o', 'u'}
				isCons := !slices.Contains(vocals, char)

				next := runes[i+1]
				if isCons && (next == 'a' || next == 'e') {
					i++
				}
			}

			continue
		}

		if unicode.IsDigit(char) {
			var number strings.Builder

			for i < len(runes) && unicode.IsDigit(runes[i]) {
				number.WriteString(string(runes[i]))
				i++
			}

			builder.WriteString(numbersProcessing(number.String()))
			i--
			continue
		}

		builder.WriteString(string(char))
	}

	builder.WriteString(findLogograms(input))

	return builder.String()
}

func numbersProcessing(number string) string {
	num, _ := strconv.ParseInt(number, 10, 64)
	var result strings.Builder

	hunds := num / 100
	result.WriteString(strings.Repeat("𐏕", int(hunds)))
	num %= 100

	twents := num / 20
	result.WriteString(strings.Repeat("𐏔", int(twents)))
	num %= 20

	if num >= 10 {
		result.WriteString("𐏓")
	}
	num %= 10

	twos := num / 2
	result.WriteString(strings.Repeat("𐏒", int(twos)))
	num %= 2

	if num == 1 {
		result.WriteString("𐏑")
	}

	return result.String()
}

func findLogograms(input string) string {
	cleanInput := strings.ToLower(input)
	words := strings.Fields(cleanInput)

	var found []string
	seen := make(map[string]bool)

	for _, word := range words {
		normalized := strings.ReplaceAll(word, "ā", "a")

		if logo, ok := logograms[normalized]; ok && !seen[normalized] {
			found = append(found, fmt.Sprintf("• %s: %s", word, logo))
			seen[normalized] = true
		}
	}

	if len(found) == 0 {
		return ""
	}

	return "\n\nPossible logograms:\n" + strings.Join(found, "\n")
}
