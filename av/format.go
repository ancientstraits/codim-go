package av

// #cgo pkg-config: libavformat
// #include <libavformat/avformat.h>
import "C"

type (
	FormatContext C.struct_AVFormatContext
	InputFormat   C.struct_AVInputFormat
	OutputFormat  C.struct_AVOutputFormat
)

// Creates a new FormatContext.
func newFormatContext() FormatContext {
	return FormatContext(*C.avformat_alloc_context())
}

// Creates a new FormatContext for output.
func newOutputFormatContext(out_fmt *OutputFormat, fmt_name, filename string) (FormatContext, error) {
	fmt_ctx := &C.struct_AVFormatContext{}
	C.avformat_alloc_output_context2(&fmt_ctx, (*C.struct_AVOutputFormat)(out_fmt), C.CString(fmt_name), C.CString(filename))
}

func (fmt_ctx *FormatContext) AllocOutputContext2(of *OutputFormat, fmtname, filename string) {}
