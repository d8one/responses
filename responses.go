package responses

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

// Write writes a response
func Write(w http.ResponseWriter, data interface{}) error {
	switch data.(type) {
	case string:
		_, err := w.Write([]byte(data.(string)))
		return err

	case []byte:
		_, err := w.Write(data.([]byte))
		return err

	case int:
		s := strconv.Itoa(data.(int))
		_, err := w.Write([]byte(s))
		return err

	case float64:
		s := fmt.Sprintf("%f", data.(float64))
		_, err := w.Write([]byte(s))
		return err

	}
	return fmt.Errorf("data type of %v not supported", reflect.TypeOf(data))

}

func BadStatusOnError(err error, errText string) {
	if err != nil {
		return
	}
	if errText == "" {
		http.Error()
	}
}
