package shareinterfaces

import share_types "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"

type Workflow[T any] = func(share_types.Command[T]) ([]Event, error)

type Event interface {
	IsEvent()
}
