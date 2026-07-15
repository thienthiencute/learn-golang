package schedule

import (
	"fmt"
	"os"
	"os/signal"
	"gocas/cmd/schedule/internal"

	"github.com/spf13/cobra"
	"github.com/teoit/gosctx"
	"github.com/teoit/gosctx/component/fiberapp"
	"github.com/teoit/gosctx/component/gormc"
	"github.com/teoit/gosctx/component/redisc"
	"github.com/teoit/gosctx/configs"
)

var (
	serviceName = "schedule-service"
	version     = "1.0.0"
)

func newScheduleServiceCtx() gosctx.ServiceContext {
	return gosctx.NewServiceContext(
		gosctx.WithName(serviceName),
		gosctx.WithComponent(fiberapp.NewFiber(configs.KeyCompFIBER)),
		gosctx.WithComponent(gormc.NewGormDB(configs.KeyCompGorm, "")),
		gosctx.WithComponent(redisc.NewRedisc(configs.KeyCompRedis)),
		gosctx.WithComponent(gosctx.NewAppLoggerDaily(configs.KeyLoggerDaily)),
	)
}

var ScheduleCmd = &cobra.Command{
	Use:     "schedule",
	Version: version,
	Short:   "schedule go clean architecture service context",
	Long:    `schedule Go clean architecture service context`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---------------------- schedule --------------")
		serviceCtx := newScheduleServiceCtx()
		
		logSrv := serviceCtx.MustGet(configs.KeyLoggerDaily).(gosctx.AppLoggerDaily).GetLogger("schedule")
		logSrv.Info("------------>> start schedule service <<------------------")

		if err := serviceCtx.Load(); err != nil {
			logSrv.Fatal(err)
		}

		internal.RoutesSchedule(serviceCtx)

		// Block main goroutine until OS interrupt
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
	},
}
