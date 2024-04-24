package kana

const (
	KATAKANA_FILE_PATH = "katakana.json"
)

var (
	Katakana []Kana
)

func init() {
	UnmarshalKana(KATAKANA_FILE_PATH, &Katakana)
}

func RandomKatakana() Kana {
	return RandomKana(&Katakana)
}
