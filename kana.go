package kana

import (
	"encoding/json"
	"io"
	"math/rand"
	"os"
)

type Kana struct {
	Kana   string `json:"kana"`
	Romaji string `json:"romaji"`
	Type   string `json:"type"`
}

func UnmarshalKana(filePath string, kana *[]Kana) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(bytes, kana)
}

func RandomKana(kana *[]Kana) Kana {
	i := rand.Intn(len(*kana))
	return (*kana)[i]
}
