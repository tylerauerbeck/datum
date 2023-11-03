package idx

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"

	"github.com/jaevor/go-nanoid"
)

type ID string

func GetNewID() (string, error) {
	canonicID, err := nanoid.Standard(21)
	if err != nil {
		return "", err
	}
	return canonicID(), nil
}

func MustGetNewID() string {
	v, err := GetNewID()
	if err != nil {
		panic(err)
	}
	return v
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (u *ID) UnmarshalGQL(v interface{}) error {
	return u.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (u ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(u)))
}

// Scan implements the Scanner interface.
func (u *ID) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("id: expected a string")
	}
	*u = ID(s)
	return nil
}

// Value implements the driver Valuer interface.
func (u ID) Value() (driver.Value, error) {
	return string(u), nil
}

func (u ID) IsValid() bool {
	return u != ""
}
