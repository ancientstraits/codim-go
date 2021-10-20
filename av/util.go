package av

// #cgo pkg-config: libavutil
// #include <libavutil/avutil.h>
// char* avstrerr(int errnum) {return av_err2str(errnum);}
import "C"

func strerr(errnum int) string {
	return C.GoString(C.avstrerr(C.int(errnum))
}
