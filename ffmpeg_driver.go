package ffmpeg

import (
	"github.com/hybridgroup/gobot"
)

type FfmpegDriver struct {
	gobot.Driver
}

type FfmpegInterface interface {
}

func NewFfmpegDriver(a *FfmpegAdaptor, name string) *FfmpegDriver {
	return &FfmpegDriver{
		Driver: *gobot.NewDriver(
			name,
			"ffmpeg.FfmpegDriver",
			a,
		),
	}
}

func (f *FfmpegDriver) adaptor() *FfmpegAdaptor {
	return f.Driver.Adaptor().(*FfmpegAdaptor)
}

func (f *FfmpegDriver) Start() bool { return true }
func (f *FfmpegDriver) Halt() bool  { return true }

func (s *FfmpegDriver) StartServer(args []string) (err error) {
	return startServer(args)
}

func (s *FfmpegDriver) StartStreamer(args []string) (err error) {
	return startStreamer(args)
}

func (s *FfmpegDriver) StopServer() (err error) {
	return stopServer()
}

func (s *FfmpegDriver) StopStreamer() (err error) {
	return stopStreamer()
}
