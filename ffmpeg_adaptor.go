package ffmpeg

import (
  "github.com/hybridgroup/gobot"
)

type FfmpegAdaptor struct {
  gobot.Adaptor
}

func NewFfmpegAdaptor(name string) *FfmpegAdaptor {
  return &FfmpegAdaptor{
    Adaptor: *gobot.NewAdaptor(
      name,
      "ffmpeg.FfmpegAdaptor",
    ),
  }
}

func (f *FfmpegAdaptor) Connect() bool {
  return true
}

func (f *FfmpegAdaptor) Finalize() bool {
  return true
}
