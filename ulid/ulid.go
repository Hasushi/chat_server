package ulid

import (
	"chat_server/usecase/output_port"
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ULID struct {}

func NewULID() output_port.ULID {
	return &ULID{}
}

func (u *ULID) GenerateID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.Reader, 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}