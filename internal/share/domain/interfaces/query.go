package shareinterfaces

import sharetypes "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"

// ジェネリックなクエリハンドラー
type QueryHandler[T, R any] = func(sharetypes.Query[T]) (R, error)
