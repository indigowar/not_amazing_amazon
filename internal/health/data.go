package health

import (
	"encoding/json"
	"fmt"
)

type Dependency struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	Error          error  `json:"error,omitempty"`
	ResponseTimeMs int64  `json:"response_time_ms,omitempty"`
}

type Health struct {
	Status       string       `json:"status"`
	Error        error        `json:"error,omitempty"`
	Dependencies []Dependency `json:"dependencies,omitempty"`
}

func (h Health) MarshalJSON() ([]byte, error) {
	type Alias Health
	var errorString *string
	if h.Error != nil {
		errorString = new(string)
		*errorString = fmt.Sprint(h.Error)
	}
	return json.Marshal(&struct {
		Alias
		Error *string `json:"error,omitempty"`
	}{
		Alias: Alias(h),
		Error: errorString,
	})
}

func (d Dependency) MarshalJSON() ([]byte, error) {
	type Alias Dependency
	var errorString *string
	if d.Error != nil {
		errorString = new(string)
		*errorString = fmt.Sprint(d.Error)
	}
	return json.Marshal(&struct {
		Alias
		Error *string `json:"error,omitempty"`
	}{
		Alias: Alias(d),
		Error: errorString,
	})
}
