package GoStack

import (
	"encoding/json"
	"errors"
)

//StackError gives data members for marshalling to check if you need to backoff
type StackError struct {
	ErrorID      int    `json:"error_id"`
	ErrorMessage string `json:"error_message"`
	ErrorName    string `json:"error_name"`
}

func (s *StackError) Error() string {
	return s.ErrorMessage
}

func checkError(data []byte) error {

	var err StackError
	json.Unmarshal(data, &err)

	if err.ErrorID == 502 {
		return errors.New(err.ErrorMessage)
	}
	return nil

}
