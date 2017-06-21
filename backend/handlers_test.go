package faultinfo_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/charakoba-com/fault_info/backend"
	"github.com/charakoba-com/fault_info/backend/db"
	"github.com/charakoba-com/fault_info/backend/model"
)

var s *faultinfo.Server
var ts *httptest.Server

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	if ret == 0 {
		teardown()
	}
	os.Exit(ret)
}

func setup() {
	s = faultinfo.New()
	ts = httptest.NewServer(s)
	db.Init(nil)
}

func teardown() {
	ts.Close()
}

func TestPostInfoHandler(t *testing.T) {
	body := bytes.Buffer{}
	body.WriteString(`{"info_type": "maintenance", "service": "www", "begin": "2017-02-01 00:00:00", "detail": "creation test"}`)
	rawResponse, err := http.Post(ts.URL, "application/json", &body)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if rawResponse.StatusCode != 200 {
		t.Errorf("status 200 is expected, but %d", rawResponse.StatusCode)
		return
	}
	var response model.PostInfoHandlerResponse
	if err := json.NewDecoder(rawResponse.Body).Decode(&response); err != nil {
		t.Errorf("%s", err)
		return
	}
	if response.Message != "success" {
		t.Errorf(`"%s" != "success"`, response.Message)
		return
	}
	if response.ID == 0 {
		t.Errorf("ID should be >0")
		return
	}
}

func TestGetInfoListHandler(t *testing.T) {
	rawResponse, err := http.Get(ts.URL)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if rawResponse.StatusCode != 200 {
		t.Errorf("status 200 is expected, but %d", rawResponse.StatusCode)
		return
	}
	var response model.GetInfoListHandlerResponse
	if err := json.NewDecoder(rawResponse.Body).Decode(&response); err != nil {
		t.Errorf("%s", err)
		return
	}
	expected := model.Info{
		ID:      1,
		Type:    "maintenance",
		Service: "www",
		Begin:   time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
		End:     time.Date(2017, 1, 1, 12, 0, 0, 0, time.UTC),
		Detail:  "test data",
	}
	if response.Info == nil {
		t.Errorf("response.Info is nil")
		return
	}
	if response.Info[0] != expected {
		t.Errorf("%s != %s", response.Info[0], expected)
		return
	}

}

func TestUpdateInfoHandler(t *testing.T) {
	body := bytes.Buffer{}
	body.WriteString(`{"detail": "updated"}`)
	request, err := http.NewRequest(`PUT`, ts.URL+"/1", &body)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	rawResponse, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	var response model.UpdateInfoHandlerResponse
	if err := json.NewDecoder(rawResponse.Body).Decode(&response); err != nil {
		t.Errorf("%s", err)
		return
	}
	if response.Message != "success" {
		t.Errorf(`"%s" != "success"`, response.Message)
		return
	}
	var di db.Info
	tx, err := db.BeginTx()
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if err := di.Load(tx, 1); err != nil {
		t.Errorf("%s", err)
		return
	}
	if di.Detail != "updated" {
		t.Errorf("%s != updated", di.Detail)
		return
	}
	// teardown
	body = bytes.Buffer{}
	body.WriteString(`{"detail": "test data"}`)
	request, err = http.NewRequest(`PUT`, ts.URL+"/1", &body)
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	request.Header.Set("Content-Type", "application/json")
	http.DefaultClient.Do(request)
}

func TestGetTypesHandler(t *testing.T) {
	rawResponse, err := http.Get(ts.URL + "/types")
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if rawResponse.StatusCode != 200 {
		t.Errorf("status 200 is expected, but %d", rawResponse.StatusCode)
		return
	}
	var response model.GetTypesHandlerResponse
	if err := json.NewDecoder(rawResponse.Body).Decode(&response); err != nil {
		t.Errorf("%s", err)
		return
	}
	expected := model.TypeList{
		model.Type{Type: "maintenance"},
	}
	if response.Types == nil {
		t.Errorf("response.Types is nil")
		return
	}
	if response.Types[0] != expected[0] {
		t.Errorf("%s != %s", response.Types[0], expected[0])
		return
	}
}

func TestGetServicesHandler(t *testing.T) {
	rawResponse, err := http.Get(ts.URL + "/services")
	if err != nil {
		t.Errorf("%s", err)
		return
	}
	if rawResponse.StatusCode != 200 {
		t.Errorf("status 200 is expected, but %d", rawResponse.StatusCode)
		return
	}
	var response model.GetServicesHandlerResponse
	if err := json.NewDecoder(rawResponse.Body).Decode(&response); err != nil {
		t.Errorf("%s", err)
		return
	}
	expected := model.ServiceList{
		model.Service{Name: "www"},
	}
	if response.Services == nil {
		t.Errorf("response.Services is nil")
		return
	}
	if response.Services[0] != expected[0] {
		t.Errorf("%s != %s", response.Services[0], expected[0])
		return
	}
}
