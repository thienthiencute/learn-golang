package pubsub

import (
	"fmt"
	"os"
	"os/signal"
	"gocas/cmd/pubsub/internal"

	"github.com/spf13/cobra"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/fiberapp"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/component/redisc"
	"github.com/teoit/gosctx/configs"
)

var (
	serviceName = "pubsub-service"
	version     = "1.0.0"
)

func newPubSubServiceCtx() gosctx.ServiceContext {
	return gosctx.NewServiceContext(
		gosctx.WithName(serviceName),
		gosctx.WithComponent(fiberapp.NewFiber(configs.KeyCompFIBER)),
		gosctx.WithComponent(gormc.NewGormDB(configs.KeyCompGorm, "")),
		gosctx.WithComponent(redisc.NewRedisc(configs.KeyCompRedis)),
		gosctx.WithComponent(gosctx.NewAppLoggerDaily(configs.KeyLoggerDaily)),
	)
}

var PubSubCmd = &cobra.Command{
	Use:     "pubsub",
	Version: version,
	Short:   " pubsub go clean architecture service context",
	Long:    `pubsub Go clean architecture service context tranning youtube`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---------------------- pubsub --------------")
		serviceCtx := newPubSubServiceCtx()
		
		logSrv := serviceCtx.MustGet(configs.KeyLoggerDaily).(gosctx.AppLoggerDaily).GetLogger("pubsub")
		logSrv.Info("------------>> start pubsub service <<------------------")

		if err := serviceCtx.Load(); err != nil {
			logSrv.Fatal(err)
		}

		internal.RoutesPubsub(serviceCtx)

		// Block main goroutine until OS interrupt
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
	},
}
