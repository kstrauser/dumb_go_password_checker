package validators

import "unicode"

type Rules struct {
	OnlyAscii      bool
	MinLength      int
	MaxLength      int
	MinPunctuation int
}

type ValidatorPair struct {
	validator func(string) bool
	message   string
}

func NewRulesWithDefaults() *Rules {
	r := Rules{
		MinLength:      8,
		MaxLength:      128,
		OnlyAscii:      true,
		MinPunctuation: 1,
	}
	return &r
}

func (r Rules) IsAllPrintable(value string) bool {
	for _, char := range value {
		if !unicode.IsPrint(char) {
			return false
		}
		if r.OnlyAscii && int(char) >= 0x80 {
			return false
		}
	}
	return true
}

func (r Rules) IsLongEnough(value string) bool {
	return len([]rune(value)) >= r.MinLength
}

func (r Rules) IsShortEnough(value string) bool {
	return len([]rune(value)) <= r.MaxLength
}

func (r Rules) HasEnoughPunctuation(value string) bool {
	if r.MinPunctuation == 0 {
		return true
	}
	count := 0
	for _, char := range value {
		if unicode.IsPunct(char) {
			count += 1
		}
	}
	return count >= r.MinPunctuation
}

func (r Rules) Validate(value string) []string {
	errors := []string{}
	validators := []ValidatorPair{
		{r.IsAllPrintable, "The password contains invalid characters."},
		{r.IsLongEnough, "The password isn't long enough."},
		{r.IsShortEnough, "The password is too long."},
		{r.HasEnoughPunctuation, "The password doesn't have enough punctuation."},
	}

	for _, item := range validators {
		if !item.validator(value) {
			errors = append(errors, item.message)
		}
	}

	return errors
}
