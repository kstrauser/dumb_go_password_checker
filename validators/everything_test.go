package validators

import "testing"

func TestIsAllPrintable(t *testing.T) {
	r := NewRulesWithDefaults()

	r.OnlyAscii = true
	if r.IsAllPrintable("ಠ_ಠ") {
		t.Error("ಠ_ಠ isn't all ASCII")
	}

	r.OnlyAscii = false
	for _, value := range []string{
		"foo",
		"ಠ_ಠ",
	} {
		if !r.IsAllPrintable(value) {
			t.Errorf("%#v is all printable", value)
		}
	}

	for _, value := range []string{
		"one\ntwo",
	} {
		if r.IsAllPrintable(value) {
			t.Errorf("%#v isn't all printable", value)
		}
	}
}

func TestIsLongEnough(t *testing.T) {
	r := NewRulesWithDefaults()

	for _, value := range []string{
		"foofoofo",
		"ಠ_ಠಠ_ಠಠ_",
	} {
		if !r.IsLongEnough(value) {
			t.Errorf("%#v is long enough", value)
		}
	}

	for _, value := range []string{
		"foofoof",
		"ಠ_ಠಠ_ಠಠ",
	} {
		if r.IsLongEnough(value) {
			t.Errorf("%#v isn't long enough", value)
		}
	}
}

func TestIsShortEnough(t *testing.T) {
	r := NewRulesWithDefaults()
	r.MaxLength = 7

	for _, value := range []string{
		"foofoofo",
		"ಠ_ಠಠ_ಠಠ_",
	} {
		if r.IsShortEnough(value) {
			t.Errorf("%#v isn't short enough", value)
		}
	}

	for _, value := range []string{
		"foofoof",
		"ಠ_ಠಠ_ಠಠ",
	} {
		if !r.IsShortEnough(value) {
			t.Errorf("%#v is short enough", value)
		}
	}
}

func TestHasEnoughPunctuation(t *testing.T) {
	r := NewRulesWithDefaults()

	r.MinPunctuation = 0
	if !r.HasEnoughPunctuation("") {
		t.Error("Empty strings are OK for 0 punctuation")
	}

	r.MinPunctuation = 1
	for _, value := range []string{
		"hi!",
		"what, me worry?",
	} {
		if !r.HasEnoughPunctuation(value) {
			t.Errorf("%#v has enough punctuation", value)
		}
	}

	for _, value := range []string{
		"hi",
		"what me worry",
	} {
		if r.HasEnoughPunctuation(value) {
			t.Errorf("%#v doesn't have enough punctuation", value)
		}
	}
}
