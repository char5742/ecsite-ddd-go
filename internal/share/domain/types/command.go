package sharetypes

import (
	"context"
	"time"

	identitytypes "github.com/char5742/ecsite-ddd-go/internal/identity/domain/types"
)

type Command[T any] struct {
	Context    context.Context
	Data       T
	Timestamp  time.Time
	IdentityID identitytypes.IdentityID
}
