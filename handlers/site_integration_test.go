package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLanding(t *testing.T) {
	rr, err := genericGETTest("GET", "/", home)
	if err != nil {
		log.Fatal(err)
	}
	if status := rr.Code; status != http.StatusOK {
		log.Fatal("It must return 200 Code")
	}

	if resp := rr.Body.String(); resp != "Hello World" {
		log.Fatal("The Expected Response Value has changed")
	}
}

func TestStatus(t *testing.T) {
	rr, err := genericGETTest("GET", "/status", status)
	if err != nil {
		log.Fatal(err)
	}

	if code := rr.Code; code != http.StatusOK {
		log.Fatal("Expected Code is 200; Seen is: ", code)
	}

	resp := rr.Body.String()
	kk := make(map[string]interface{})
	err = json.Unmarshal([]byte(resp), &kk)
	if err != nil {
		log.Fatal(err)
	}

	println(kk["success"].(bool))
	println(kk["message"].(string))
	data := kk["data"].(map[string]interface{})
	println(data["time"].(string))
}

func genericGETTest(method, path string, handler func(w http.ResponseWriter, r *http.Request)) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)
	return rr, nil
}
