package xutils

import (
	"encoding/json"
	"fmt"
)

// interface for database row.Scan and rows.Scan
type RowScanner interface {
	Scan(dest ...any) error
}

// for scan json column
// JSONScanner is a generic interface for scanning JSON into any struct
type JSONScanner[T any] struct {
	Val *T
}

// Scan implements the sql.Scanner interface
func (js *JSONScanner[T]) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	if js.Val == nil {
		js.Val = new(T)
	}
	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, js.Val)
	case string:
		return json.Unmarshal([]byte(src), js.Val)
	}
	return fmt.Errorf("cannot scan %T into JSONScanner", src)
}
