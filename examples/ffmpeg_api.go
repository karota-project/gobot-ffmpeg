package main

import (
	"fmt"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/api"
	"github.com/karota-project/gobot-ffmpeg"
)

func main() {
	master := gobot.NewGobot()
	api.NewAPI(master).Start()

	ffmpegAdaptor := ffmpeg.NewFfmpegAdaptor("ffmpeg-a01")
	ffmpegDriver := ffmpeg.NewFfmpegDriver(ffmpegAdaptor, "ffmpeg-d01")

	master.AddRobot(
		gobot.NewRobot(
			"ffmpeg",
			[]gobot.Connection{ffmpegAdaptor},
			[]gobot.Device{ffmpegDriver},
			func() {
				fmt.Println("work")
			}))

	master.Start()
}
