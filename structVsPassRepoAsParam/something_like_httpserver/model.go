package something_like_httpserver

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Group     string
	ProfileId string
	Token     []byte
	IsManager bool
}
