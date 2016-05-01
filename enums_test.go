package osuapi

import (
	"testing"
)

// these tests are really just here to bring the coverage up a bit

func TestGenre(t *testing.T) {
	if GenreElectronic.String() != "electronic" {
		t.Fatal("expected GenreElectronic.String() to return 'electronic', got", GenreElectronic.String(), "instead")
	}
}
func TestGenreOOB(t *testing.T) {
	if Genre(14910).String() != "14910" {
		t.Fatal("expected Genre(14910).String() to return '14910', got", Genre(14910).String(), "instead")
	}
}

func TestApprovedStatus(t *testing.T) {
	if StatusGraveyard.String() != "graveyard" {
		t.Fatal("expected StatusGraveyard.String() to return 'graveyard', got", StatusGraveyard.String(), "instead")
	}
}
func TestApprovedStatusOOB(t *testing.T) {
	if ApprovedStatus(-41).String() != "-41" {
		t.Fatal("expected ApprovedStatus(-41).String() to return '-41', got", ApprovedStatus(-41).String(), "instead")
	}
}

func TestLanguage(t *testing.T) {
	if LanguageChinese.String() != "Chinese" {
		t.Fatal("expected LanguageChinese.String() to return 'Chinese', got", LanguageChinese.String(), "instead")
	}
}
func TestLanguageOOB(t *testing.T) {
	if Language(1337).String() != "1337" {
		t.Fatal("expected Language(1337).String() to return '1337', got", Language(1337).String(), "instead")
	}
}

func TestMode(t *testing.T) {
	if ModeOsuMania.String() != "osu!mania" {
		t.Fatal("expected ModeOsuMania.String() to return 'osu!mania', got", ModeOsuMania.String(), "instead")
	}
}
func TestModeOOB(t *testing.T) {
	if Mode(-414141).String() != "-414141" {
		t.Fatal("expected Mode(-414141).String() to return '-414141', got", Mode(1337).String(), "instead")
	}
}
