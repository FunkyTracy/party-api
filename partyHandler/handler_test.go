package partyHandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestHandleGetParties_HappyPath_Title(t *testing.T) {
	m := mockPartyAPI{}
	expectedResp := []party{gnomeParty, vampParty}

	title := "vampire night club"
	reqBody := partyRequest{
		Title: &title,
	}
	url := "http://localhost/parties"
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))

	m.On("GetParties", reqBody).Return(expectedResp, nil)

	rr := httptest.NewRecorder()

	HandleGetParties(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var got partyResponse
	_ = json.NewDecoder(rr.Body).Decode(&got)

	if !reflect.DeepEqual(got.Parties, expectedResp) {
		t.Errorf("Expected response to be %v got %v", expectedResp, got.Parties)
	}
}

func TestHandleGetParties_HappyPath_Theme(t *testing.T) {
	m := mockPartyAPI{}
	expectedResp := []party{gnomeParty, vampParty}

	theme := "dark gothic"
	reqBody := partyRequest{
		Theme: &theme,
	}
	url := "http://localhost/parties"
	jsonBody, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonBody))

	m.On("GetParties", reqBody).Return(expectedResp, nil)

	rr := httptest.NewRecorder()

	HandleGetParties(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var got partyResponse
	_ = json.NewDecoder(rr.Body).Decode(&got)

	if !reflect.DeepEqual(got.Parties, expectedResp) {
		t.Errorf("Expected response to be %v got %v", expectedResp, got.Parties)
	}
}

type mockPartyAPI struct {
	mock.Mock
}

func (m *mockPartyAPI) GetParties(req partyRequest) ([]party, error) {
	args := m.Called(req)
	return args.Get(0).([]party), args.Error(1)
}
