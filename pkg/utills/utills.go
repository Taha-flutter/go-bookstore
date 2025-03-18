package utills

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the body of the request and unmarshals it into the provided interface.
func ParseBody(r *http.Request, x any) error {
	// Read the entire body from the request
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// Return an error if reading the body fails
		return err
	}

	// Unmarshal the JSON body into the provided interface
	if err := json.Unmarshal(body, x); err != nil {
		// Return an error if unmarshalling the body fails
		return err
	}

	// If everything is successful, return nil
	return nil
}

func ErrorResponse(w http.ResponseWriter, statuscode int, message string) {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statuscode)

	errorResponse := map[string]string{"error": message}

	res, _ := json.Marshal(errorResponse)

	w.Write(res)

}
