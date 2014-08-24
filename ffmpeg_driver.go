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
func (f *FfmpegDriver) Halt() bool { return true }
