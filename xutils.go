package xutils

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/go-faster/errors"
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

func StringToInt64(v string) (int64, error) {

	// Remove all non-numeric characters
	cleanedStr := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, v)

	// Convert the cleaned string to int64
	num, err := strconv.ParseInt(cleanedStr, 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "cant ParseInt. %s", v)
	}

	return num, nil
}

func Bulk[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
