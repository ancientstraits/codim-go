package codim_go

import (
	"github.com/ancientstraits/codim_go/av"
)

type VideoContext struct {
	duration uint64
	pts      uint64
	fmt_ctx  *av.FormatContext
	cod      *av.Codec
	cod_ctx  *av.CodecContext
	stream   *av.Stream
	pkt      *av.Packet
	frame    *av.Frame
}

// Creates a new VideoContext.
func newVideoContext(video_path string, width uint, height uint, fps uint) (VideoContext, error) {}
// Closes a VideoContext.
func (vc *VideoContext) Close()                                                                  {}
// Writes the VideoContext's frame to its video.
func (vc *VideoContext) WriteFrame() error                                                       {}
