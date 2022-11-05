package main

import (
	"github.com/crayonwow/base_service/internal"
	"github.com/crayonwow/base_service/pkg/config"
	"github.com/crayonwow/base_service/pkg/db"
	"github.com/crayonwow/base_service/pkg/di"
	"github.com/crayonwow/base_service/pkg/graceful"
	"github.com/crayonwow/base_service/pkg/httpserver"
	"github.com/crayonwow/base_service/pkg/logger"
	"github.com/crayonwow/base_service/pkg/messagebroker"
)

func main() {
	di.
		NewContainer().
		Provide(
			internal.Module(),
			config.Module(),
			graceful.Module(),
			db.Module(),
			messagebroker.Module(),
			httpserver.Module(),
		).
		Invoke(logger.NewLogger).
		Run()
}
