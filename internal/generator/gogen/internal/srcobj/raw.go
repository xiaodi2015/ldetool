package srcobj

import (
	"fmt"
	"io"
)

// Raw just puts text directly into the generated code
type Raw string

// Dump ...
func (r Raw) Dump(w io.Writer) error {
	_, err := io.WriteString(w, string(r))
	return err
}

// Literal generates literal representation of object fed
func Literal(data interface{}) Source {
	return Raw(fmt.Sprintf("%v", data))
}

// HexU64 generates hexadecimal representation of uint64 fed
func HexU64(data uint64) Source {
	return Raw(fmt.Sprintf("0x%x", data))
}
