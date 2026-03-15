// Package main provides the core functionality for the Old Persian cuneiform bot.
// This file contains the mapping rules for characters, syllables, and logograms.
package main

import "strings"

// baseSigns maps individual Latin characters to their primary Old Persian cuneiform signs.
var baseSigns = map[rune]string{
	'a': "𐎠",
	'ā': "𐎠",
	'e': "𐎠",  // 'e' processed as 'a'
	'o': "𐎠𐎢", // 'o' processed as 'au'
	'i': "𐎡",
	'u': "𐎢",

	'b': "𐎲",
	'c': "𐎨",
	'ç': "𐏂",
	'd': "𐎭",
	'f': "𐎳",
	'g': "𐎥",
	'h': "𐏃",
	'j': "𐎩",
	'k': "𐎣",
	'l': "𐎾",
	'm': "𐎶",
	'n': "𐎴",
	'p': "𐎱",
	'q': "𐎣", // processed as 'k'
	'r': "𐎼",
	's': "𐎿",
	'š': "𐏁",
	't': "𐎫",
	'θ': "𐎰",
	'v': "𐎺",
	'w': "𐎢", // processed as 'u'
	'x': "𐎧",
	'y': "𐎹",
	'z': "𐏀",

	' ': "𐏐",
}

// contextualSigns handles specific two-character combinations (syllables)
// that have dedicated cuneiform signs, overriding the base mapping.
var contextualSigns = map[string]string{
	"di": "𐎮𐎡",
	"du": "𐎯𐎢",
	"gu": "𐎦𐎢",
	"ji": "𐎪𐎡",
	"ku": "𐎤𐎢",
	"mi": "𐎷𐎡",
	"mu": "𐎸𐎢",
	"nu": "𐎵𐎢",
	"ru": "𐎽𐎢",
	"tu": "𐎬𐎢",
	"vi": "𐎻𐎡",
	"qu": "𐎤𐎢", // processed as 'ku'
}

// logograms maps full words or stems to their corresponding ideograms.
var logograms = map[string]string{
	"Ahuramazdā": "𐏈 / 𐏉 / 𐏊 (genitive)",
	"xšāyaθiya":  "𐏋",
	"dahyāuš":    "𐏌 / 𐏍",
	"baga":       "𐏎",
	"būmiš":      "𐏏",
}

// apostropheNormalizer converts user-friendly apostrophe notation
// into proper linguistic transcription characters with diacritics.
var apostropheNormalizer = strings.NewReplacer(
	"a'", "ā",
	"i'", "ī",
	"u'", "ū",

	"s'", "š",
	"t'", "θ",
	"c'", "ç",
)

// vowelNormalizer is used to strip length marks (macrons) from vowels
// to simplify internal string comparisons and logogram lookups.
var vowelNormalizer = strings.NewReplacer(
	"ā", "a",
	"ū", "u",
	"ī", "i",
)
