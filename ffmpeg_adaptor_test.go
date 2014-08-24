package ffmpeg

import (
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestFfmpegAdaptor() *FfmpegAdaptor {
	return NewFfmpegAdaptor("myAdaptor")
}

func TestFfmpegAdaptorConnect(t *testing.T) {
	a := initTestFfmpegAdaptor()
	gobot.Expect(t, a.Connect(), true)
}

func TestFfmpegAdaptorFinalize(t *testing.T) {
	a := initTestFfmpegAdaptor()
	gobot.Expect(t, a.Finalize(), true)
}
