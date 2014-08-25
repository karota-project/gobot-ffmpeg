package main

import (
	"fmt"
	"time"

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

				if err := ffmpegDriver.StartStreamer(nil); err != nil {
					fmt.Println(err)
					return
				}

				time.Sleep(5 * time.Second)

				if err := ffmpegDriver.StartServer(nil); err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println("Media Streaming...")

				time.Sleep(60 * time.Second)

				if err := ffmpegDriver.StopStreamer(); err != nil {
					fmt.Println(err)
					return
				}

				if err := ffmpegDriver.StopServer(); err != nil {
					fmt.Println(err)
					return
				}

				fmt.Println("Done!")
			}))

	master.Start()
}
