package cmd

import (
	"fmt"
	"holiday/configs"
	"holiday/data"
	"holiday/model"
	"holiday/pkg/logger"
	"holiday/router"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "holiday",
	Run: func(cmd *cobra.Command, args []string) {
		if err := configs.Init(""); err != nil {
			panic(err)
		}

		model.PostgresDB.Init()
		defer model.PostgresDB.Close()
		var newRouter = router.Router{}

		g := gin.New()
		newRouter.InitRouter(g)

		g.Run(":8080")
	},
}

var holidayCmd = &cobra.Command{
	Use: "db",
	Run: func(cmd *cobra.Command, args []string) {
		model.PostgresDB.Init()
		defer model.PostgresDB.Close()

		if args[0] == "migrate" {
			handlerModelCollection()
		}

		if args[0] == "import" {
			var allData *data.HistoryHoliday
			allData = data.NewHistoryHoliday()
			for _, one := range allData.Data {
				for _, i := range one {
					fmt.Println(i)
					yearS, monthS, dayS := handlerString(i.Start)
					yearE, monthE, dayE := handlerString(i.End)
					var holiday model.Holiday
					holiday = model.Holiday{
						EnName:     i.EnName,
						ChName:     i.ChName,
						StartYear:  yearS,
						StartMonth: int(monthS),
						StartDay:   dayS,
						EndYear:    yearE,
						EndMonth:   int(monthE),
						EndDay:     dayE,
					}
					if dbError := model.PostgresDB.DB.Where("en_name = ? AND end_year = ?", i.EnName, yearE).FirstOrCreate(&model.Holiday{}, holiday).Error; dbError != nil {
						var info = logger.Info{
							Package: "cmd",
							Action:  "import holiday",
							Message: "import date into database",
						}
						logger.PanicLog(info)
						return
					}
				}

			}
		}

	},
}

// Execute ...
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(holidayCmd)
}

func handlerString(startDate string) (year int, month time.Month, day int) {
	newString := strings.TrimSpace(startDate)
	newTime, _ := time.Parse("2006/01/02", newString)
	fmt.Println(newTime.Date())
	year, month, day = newTime.Date()
	return

}

func handlerModelCollection() {
	model.PostgresDB.DB.AutoMigrate(&model.Holiday{})
}
