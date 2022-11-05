package internal

import (
	"github.com/crayonwow/base_service/internal/user"
	"github.com/crayonwow/base_service/pkg/di"
)

func Module() di.Module {
	return di.NewModule().
		Append(user.Module())
}
