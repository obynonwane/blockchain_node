package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func FormatDobStringToTime(dobString, layout string) (*time.Time, error) {

	log.Println("The Layout", layout)
	// 2. Parse the string into a time.Time object
	dobTime, err := time.Parse(layout, dobString)
	if err != nil {
		return nil, fmt.Errorf("Invalid date format. Use YYYY-MM-DD")
	}
	return &dobTime, nil
}

func StrOrEmpty(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func PrettyLog(label string, v any) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("%s: %+v", label, v)
		return
	}
	log.Printf("%s:\n%s", label, string(b))
}
