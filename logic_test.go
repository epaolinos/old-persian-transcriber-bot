package main

import "testing"

// TestTranscribe checks if the transcription logic correctly handles
// base signs, contextual syllables, and inherent vowels.
func TestTranscribe(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Base characters",
			input:    "adam",
			expected: "𐎠𐎭𐎶", // a-d-m (assuming inherent 'a' is handled)
		},
		{
			name:     "Contextual syllable 'di'",
			input:    "didā",
			expected: "𐎮𐎡𐎭𐎠", // d(i)-i-d-ā
		},
		{
			name:     "Apostrophe notation",
			input:    "s'a",
			expected: "𐏁", // ša
		},
		{
			name:     "Inherent vowel skipping",
			input:    "ba",
			expected: "𐎲", // b (a is skipped as it's inherent)
		},
		{
			name:     "Short example",
			input:    "θātiy Dārayavauš",
			expected: "𐎰𐎠𐎫𐎡𐎹𐏐𐎭𐎠𐎼𐎹𐎺𐎢𐏁",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := transcribe(tt.input)
			// Note: We might need to trim logograms for exact comparison
			// if transcribe always appends them.
			if result != tt.expected {
				t.Errorf("transcribe(%s) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

// TestNumbersProcessing verifies the additive cuneiform numeric system.
func TestNumbersProcessing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1", "𐏑"},
		{"2", "𐏒"},
		{"10", "𐏓"},
		{"20", "𐏔"},
		{"33", "𐏔𐏓𐏒𐏑"}, // 20 + 10 + 2 + 1
		{"100", "𐏕"},
		{"259", "𐏕𐏕𐏔𐏔𐏓𐏒𐏒𐏒𐏒𐏑"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := numbersProcessing(tt.input)
			if result != tt.expected {
				t.Errorf("numbersProcessing(%s) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}
