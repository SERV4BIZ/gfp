package compress

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
)

// Decode is decode compress data from buffer bytes
func Decode(buffer []byte) ([]byte, error) {
	var b bytes.Buffer
	b.Write(buffer)
	r, err1 := zlib.NewReader(&b)
	if err1 != nil {
		r.Close()
		b.Reset()
		return nil, err1
	}

	defer r.Close()
	defer b.Reset()
	p, err2 := ioutil.ReadAll(r)
	if err2 != nil {
		return nil, err2
	}
	return p, nil
}
