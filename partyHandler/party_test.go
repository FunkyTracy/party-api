package partyHandler

import (
	"reflect"
	"testing"
)

var (
	expectedParties = []party{gnomeParty, vampParty}
)

func TestGetParties_HappyPath_Title(t *testing.T) {
	title := "Vampire"
	req := partyRequest{
		Title: &title,
	}

	resp, err := GetParties(req)

	if err != nil {
		t.Errorf("Expected no errors, got %s", err.Error())
	}

	if !reflect.DeepEqual(resp, expectedParties) {
		t.Errorf("Expected response to be %v got %v", expectedParties, resp)
	}
}

func TestGetParties_HappyPath_Theme(t *testing.T) {
	theme := "gothic"
	req := partyRequest{
		Theme: &theme,
	}

	resp, err := GetParties(req)

	if err != nil {
		t.Errorf("Expected no errors, got %s", err.Error())
	}

	if !reflect.DeepEqual(resp, expectedParties) {
		t.Errorf("Expected response to be %v got %v", expectedParties, resp)
	}
}
