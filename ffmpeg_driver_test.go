package ffmpeg

import (
  "github.com/hybridgroup/gobot"
  "testing"
)

func initTestFfmpegDriver() *FfmpegDriver {
  return NewFfmpegDriver(NewFfmpegAdaptor("myAdaptor"), "myDriver")
}

func TestFfmpegDriverStart(t *testing.T) {
  d := initTestFfmpegDriver()
  gobot.Expect(t, d.Start(), true)
}

func TestFfmpegDriverHalt(t *testing.T) {
  d := initTestFfmpegDriver()
  gobot.Expect(t, d.Halt(), true)
}
