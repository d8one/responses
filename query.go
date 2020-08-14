package responses

import (
	"net/http"
	"net/url"
)

// Gets a string of the first value of
// the query in a URL (/?key=value)
func GetQueryString(r *http.Request, key string) string {
	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return ""
	}
	return m[key][0]
}

// GetQuery Gets the query part of the url in a http.Request
// it only returns the first value of every query.
func GetQuerys(r *http.Request) map[string]string {
	result := make(map[string]string)
	m, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		return result
	}
	for key, values := range m {
		result[key] = values[0]

	}
	return result
}
