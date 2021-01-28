package compress

import (
	"bytes"
	"compress/zlib"

	"github.com/SERV4BIZ/handler"
)

// Encode is encode compress data from buffer bytes
func Encode(buffer []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err1 := w.Write(buffer)
	if handler.Error(err1) {
		w.Close()
		b.Reset()
		return nil, err1
	}

	err2 := w.Close()
	if handler.Error(err2) {
		b.Reset()
		return nil, err2
	}
	buff := b.Bytes()
	b.Reset()
	return buff, nil
}
