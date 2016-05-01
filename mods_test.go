package osuapi

import (
	"testing"
)

func TestParseMods(t *testing.T) {
	if ParseMods("HDHRDT").String() != "HDHRDT" {
		t.Fatal("Expected Mods.String() to return HDHRDT, returned", ParseMods("HDHRDT").String())
	}
}

func TestParseModsShouldIgnore(t *testing.T) {
	mods := ParseMods("EZHDHRPFMEMETIgnore me I'm useless")
	if mods.String() != "EZHDHRPF" {
		t.Fatal("Expected HDEZHRPF, got", mods.String())
	}
}
