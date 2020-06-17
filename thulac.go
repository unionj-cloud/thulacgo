package thulacgo

/*
#cgo CXXFLAGS: -I./cppthulac -DLOGGING_LEVEL=LL_WARNING -O3 -Wall
#include <stdlib.h>
#include "thulac.h"
*/
import "C"
import "unsafe"

type Thulacgo struct {
	lac C.Thulac
}

func NewThulacgo(modelpath string, userpath string, justseg bool, t2s bool, ufilter bool, separator byte) (*Thulacgo) {
	mpath, upath := C.CString(modelpath), C.CString(userpath)
	defer C.free(unsafe.Pointer(mpath))
	defer C.free(unsafe.Pointer(upath))

	_t2s := C.int(0)
	if t2s {
		_t2s = C.int(1)
	}
	_justseg := C.int(0)
	if justseg {
		_justseg = C.int(1)
	}
	_ufilter := C.int(0)
	if ufilter {
		_ufilter = C.int(1)
	}
	sep := C.char(separator)
	lac := C.NewThulac(mpath, upath, _justseg, _t2s, _ufilter, sep)
	return &Thulacgo{
		lac,
	}
}

func (self *Thulacgo) Deinit() {
	C.Deinit(self.lac)
}
func (self *Thulacgo) Seg(text string) string {
	return C.GoString(C.Seg(self.lac, C.CString(text)))
}
