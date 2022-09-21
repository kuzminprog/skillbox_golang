package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RequestCreate struct {
	Name    string   `json:"name"`
	Age     string   `json:"age"`
	Friends []string `json:"friends"`
}

type RequestMakeFriend struct {
	SourceID string `json:"source_id"`
	TargetID string `json:"target_id"`
}

type RequestDeleteUser struct {
	TargetID string `json:"target_id"`
}

type RequestUserId struct {
	NewAge string `json:"new age"`
}

// DataRequest type to handle data from request structures
type DataRequest interface{}

// getDataFromRequest - Gets data from JSON and sends it to DataRequest type.
// In case of incorrect processing it will return an error.
func getDataFromRequest(dataRequest DataRequest, r *http.Request) error {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &dataRequest)
	if err != nil {
		return err
	}
	return nil
}
