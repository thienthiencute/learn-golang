package server

import (
	"fmt"
	"gocas/cmd/server/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/fiberapp"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/component/redisc"
	"github.com/teoit/gosctx/configs"
)

var (
	serviceName = "server-service"
	version     = "1.0.0"
)

func newServerServiceCtx() gosctx.ServiceContext {
	return gosctx.NewServiceContext(
		gosctx.WithName(serviceName),
		gosctx.WithComponent(fiberapp.NewFiber(configs.KeyCompFIBER)),
		gosctx.WithComponent(gormc.NewGormDB(configs.KeyCompGorm, "")),
		gosctx.WithComponent(redisc.NewRedisc(configs.KeyCompRedis)),
		gosctx.WithComponent(gosctx.NewAppLoggerDaily(configs.KeyLoggerDaily)),
	)
}

var ServerCmd = &cobra.Command{
	Use:     "server",
	Version: version,
	Short:   "server go clean architecture service context",
	Long:    `server go clean architecture service context trainning`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("----------- Server ----------")
		serviceCtx := newServerServiceCtx()

		logSrv := serviceCtx.MustGet(configs.KeyLoggerDaily).(gosctx.AppLoggerDaily).GetLogger("server")

		logSrv.Info("------------ start server --------------")

		if err := serviceCtx.Load(); err != nil {
			logSrv.Fatal(err)
		}

		fiberComp := serviceCtx.MustGet(configs.KeyCompFIBER).(fiberapp.FiberComponent)

		appFiber := fiberComp.GetApp()

		appFiber.Get("/", func(c *fiber.Ctx) error {
			return c.SendString("Hello server")
		})

		internal.RoutesServer(appFiber, serviceCtx)

		if err := appFiber.Listen(fmt.Sprintf(":%d", fiberComp.GetPort())); err != nil {
			logSrv.Fatal(err)
		}

	},
}
