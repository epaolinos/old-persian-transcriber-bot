package main

var baseSigns = map[rune]string{
	'a': "𐎠",
	'ā': "𐎠",
	'e': "𐎠",
	'o': "𐎠𐎢",
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
	'r': "𐎼",
	's': "𐎿",
	'š': "𐏁",
	't': "𐎫",
	'θ': "𐎰",
	'v': "𐎺",
	'x': "𐎧",
	'y': "𐎹",
	'z': "𐏀",

	' ': "𐏐",
}

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
}

var logograms = map[string]string{
	"ahuramazda": "...",
}

// TODO: замены по диакритикам
var diacritics = map[string]string{
	"a'": "ā",
	"s'": "š",
	"t'": "θ",
	"c'": "ç",
}
