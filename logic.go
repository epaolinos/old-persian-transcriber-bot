// Package main provides the core functionality for the Old Persian cuneiform bot.
// This file specifically handles the transcription logic and character mapping.
package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

// transcribe converts a Latin transcription string into Old Persian cuneiform.
// It handles special consonant-vowel clusters, vowel length, and numerical values.
func transcribe(input string) string {
	// First, normalize the input: convert to lowercase and replace apostrophe
	// notation (e.g., a') with proper diacritics (e.g., ā).
	normalizedInput := apostropheNormalizer.Replace(strings.ToLower(input))

	var builder strings.Builder
	runes := []rune(normalizedInput)

	for i := 0; i < len(runes); i++ {
		char := runes[i]

		// Check for contextual syllables (two-character combinations like 'di' or 'mi')
		// that have unique cuneiform signs.
		if i+1 < len(runes) {
			next := runes[i+1]
			syll := string(char) + string(next)

			if val, ok := contextualSigns[syll]; ok {
				builder.WriteString(val)
				i++ // Skip the next character as it's part of the current syllable.
				continue
			}
		}

		// Process standard characters using the base signs map.
		if val, ok := baseSigns[char]; ok {
			builder.WriteString(val)

			// If a consonant is followed by a short 'a', the vowel sign is usually
			// inherent in the consonant sign and not written separately.
			if i+1 < len(runes) {
				vocals := []rune{'a', 'ā', 'e', 'i', 'o', 'u'}
				isCons := !slices.Contains(vocals, char)

				// We process 'e' as it were Persian 'a'
				next := runes[i+1]
				if isCons && (next == 'a' || next == 'e') {
					i++ // Skip the short vowel.
				}
			}

			continue
		}

		// Handle numerical values by grouping consecutive digits.
		if unicode.IsDigit(char) {
			var number strings.Builder

			for i < len(runes) && unicode.IsDigit(runes[i]) {
				number.WriteString(string(runes[i]))
				i++
			}

			// Convert the full number string into cuneiform numeric signs.
			builder.WriteString(numbersProcessing(number.String()))
			i-- // Step back to correctly handle the loop increment.
			continue
		}

		// If no special rules apply, write the character as is.
		builder.WriteString(string(char))
	}

	// Append suggested logograms at the end of the transcription result.
	builder.WriteString(findLogograms(input))

	return builder.String()
}

// numbersProcessing converts a decimal number string into Old Persian cuneiform numeric signs.
// The system is additive, using specific symbols for hundreds, twenties, tens, twos, and units.
func numbersProcessing(number string) string {
	// Parse the string into a 64-bit integer. Error is ignored as the caller
	// ensures the string contains only digits.
	num, _ := strconv.ParseInt(number, 10, 64)
	var result strings.Builder

	// Process hundreds (𐏕).
	hunds := num / 100
	result.WriteString(strings.Repeat("𐏕", int(hunds)))
	num %= 100

	// Process twenties (𐏔).
	twents := num / 20
	result.WriteString(strings.Repeat("𐏔", int(twents)))
	num %= 20

	// Process tens (𐏓).
	if num >= 10 {
		result.WriteString("𐏓")
	}
	num %= 10

	// Process twos (𐏒).
	twos := num / 2
	result.WriteString(strings.Repeat("𐏒", int(twos)))
	num %= 2

	// Process the remaining unit (𐏑), if any.
	if num == 1 {
		result.WriteString("𐏑")
	}

	return result.String()
}

// findLogograms analyzes the input text to identify words that can be represented
// by Old Persian logograms (e.g., King, God, Country). It returns a formatted
// string list of suggested logograms.
func findLogograms(input string) string {
	// Remove common punctuation to ensure words are correctly identified
	// (e.g., "baga," should be recognized as "baga").
	cleanInput := strings.NewReplacer(",", "", ".", "", ":", "", ";", "").Replace(input)
	words := strings.Fields(cleanInput)

	var found []string
	seen := make(map[string]bool) // Track found logograms to avoid duplicates in the response.

	for _, word := range words {
		// Normalize the word by removing vowel length marks (ā, ū) for comparison.
		normalizedWord := vowelNormalizer.Replace(strings.ToLower(word))

		for base, logo := range logograms {
			// Normalize the dictionary key (base) to match the input's format.
			normalizedBase := vowelNormalizer.Replace(strings.ToLower(base))

			// Check if the current word starts with a known logogram base.
			// This prefix matching helps catch inflected forms (e.g., 'Ahuramazdāha')
			if strings.HasPrefix(normalizedWord, normalizedBase) && !seen[base] {
				found = append(found, fmt.Sprintf("• %s: %s", base, logo))
				seen[base] = true
				break // Move to the next word once a match is found.
			}
		}
	}

	// Return an empty string if no logograms were detected.
	if len(found) == 0 {
		return ""
	}

	// Join all found hints into a single block with a header.
	return "\n\nPossible logograms:\n" + strings.Join(found, "\n")
}
