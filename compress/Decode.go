package compress

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"

	"github.com/SERV4BIZ/gfp/handler"
)

// Decode is decode compress data from buffer bytes
func Decode(buffer []byte) ([]byte, error) {
	var b bytes.Buffer
	b.Write(buffer)
	r, err1 := zlib.NewReader(&b)
	if handler.Error(err1) {
		r.Close()
		b.Reset()
		return nil, err1
	}

	defer r.Close()
	defer b.Reset()
	p, err2 := ioutil.ReadAll(r)
	if handler.Error(err2) {
		return nil, err2
	}
	return p, nil
}
