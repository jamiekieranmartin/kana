package kana

const (
	HIRAGANA_FILE_PATH = "hiragana.json"
)

var (
	Hiragana []Kana
)

func init() {
	UnmarshalKana(HIRAGANA_FILE_PATH, &Hiragana)
}

func RandomHiragana() Kana {
	return RandomKana(&Hiragana)
}
