package av

// #cgo pkg-config: libavformat
// #include <libavformat/avformat.h>
import "C"

type (
	Codec        C.struct_AVCodec
	CodecContext C.struct_AVCodecContext
	Stream       C.struct_AVStream
	Packet       C.struct_AVPacket
	Frame        C.struct_AVFrame
)
