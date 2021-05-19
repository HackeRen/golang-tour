package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	"tour/internal/timer"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式转换",
	Long:  "时间格式转换",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		now := timer.GetNowTime()
		log.Printf("输出结果         : %s, %d", now.Format("2006-01-02 15:04:05"), now.Unix())
		log.Printf("输出结果(RFC3339): %s, %d", now.Format(time.RFC3339), now.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		location, _ := time.LoadLocation("Asia/Shanghai")
		if calculateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			if !strings.Contains(calculateTime, " ") {
				layout = "2006-01-02"
			}

			currentTimer, err = time.ParseInLocation(layout, calculateTime, location)
			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		calculateTime, err := timer.GetCalculateTime(currentTimer, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}
		location, _ = time.LoadLocation("America/New_York")
		//log.Printf("输出结果: %s, %d", calculateTime.Format("2006-01-02 15:04:05"), calculateTime.Unix())
		log.Printf("输出结果: %s, %d", time.Unix(calculateTime.Unix(), 0).In(location).Format("2006-01-02 15:04:05"), calculateTime.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或已格式化")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效时间单位为"ns", "us" (or "μs"), "ms", "s", "m", "h"`)
}
