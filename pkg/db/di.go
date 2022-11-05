package db

import (
	"github.com/crayonwow/base_service/pkg/db/mysql"
	"github.com/crayonwow/base_service/pkg/db/redis"
	"github.com/crayonwow/base_service/pkg/di"
)

func Module() di.Module {
	return di.NewModule().
		Append(mysql.Module()).
		Append(redis.Module())
}
