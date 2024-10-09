package xutils

import (
	"fmt"
	"net/http"
)

func RequestFullURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	if r.URL.Fragment == "" {
		return fmt.Sprintf("%s://%s%s?%s", scheme, r.Host, r.URL.Path, r.URL.RawQuery)
	}
	return fmt.Sprintf("%s://%s%s?%s#%s", scheme, r.Host, r.URL.Path, r.URL.RawQuery, r.URL.Fragment)
}

func SliceToInterface[T any](items []T) []interface{} {
	list := make([]interface{}, 0, len(items))

	for i := range items {
		list = append(list, items[i])
	}

	return list
}
