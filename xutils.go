package xutils

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode"

	"github.com/go-faster/errors"
)

func RequestFullURL(r *http.Request) string {
	scheme := "http"
	
	if forwardedProto := r.Header.Get("X-Forwarded-Proto"); forwardedProto == "https" {
		scheme = "https"
	} else if r.TLS != nil {
		scheme = "https"
	}

	if r.URL.Fragment == "" {
		return fmt.Sprintf("%s://%s%s?%s", scheme, r.Host, r.URL.Path, r.URL.RawQuery)
	}
	return fmt.Sprintf("%s://%s%s?%s#%s", scheme, r.Host, r.URL.Path, r.URL.RawQuery, r.URL.Fragment)
}

func RequestCtxFullURL(ctx context.Context) string {
	r, ok := ctx.Value(CtxKeyRequest).(*http.Request)
	if !ok {
		return ""
	}
	return RequestFullURL(r)
}

func RequestToCtx(ctx context.Context, r *http.Request) context.Context {
	return context.WithValue(ctx, CtxKeyRequest, r)
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

func IsImageFile(file multipart.File) (bool, error) {
    // Read the first 512 bytes (common size for file type detection)
    buffer := make([]byte, 512)
    n, err := file.Read(buffer)
    if err != nil {
        return false, fmt.Errorf("error reading file: %v", err)
    }

    // Reset the file pointer to the beginning after reading
    _, err = file.Seek(0, 0)
    if err != nil {
        return false, fmt.Errorf("error resetting file pointer: %v", err)
    }

    // Trim buffer to actual bytes read
    buffer = buffer[:n]

    // Detect content type based on file header
    contentType := http.DetectContentType(buffer)

    // Check if the content type is an image
    switch contentType {
    case "image/jpeg", "image/png", "image/gif", "image/bmp", "image/webp":
        return true, nil
    default:
        return false, nil
    }
}
