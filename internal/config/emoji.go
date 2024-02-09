package config

var emojis = map[string]string{
	"":            "&#x1F600;", // default icon
	"poop":        "&#x1f4a9;",
	"smile":       "&#x1F600;",
	"ping":        "&#x1f4e1;",
	"backup":      "&#128190;",
	"reboot":      "&#9211;",
	"restart":     "&#9211;",
	"box":         "&#128230;",
	"ashtonished": "&#128562;",
	"clock":       "&#128338;",
	"disk":        "&#128189;",
	"logs":        "&#128269;",
}

func lookupHTMLIcon(keyToLookup string) string {
	if emoji, found := emojis[keyToLookup]; found {
		return emoji
	}

	return keyToLookup
}
