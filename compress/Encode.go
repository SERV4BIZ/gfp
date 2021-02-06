package compress

import (
	"bytes"
	"compress/zlib"
)

// Encode is encode compress data from buffer bytes
func Encode(buffer []byte) ([]byte, error) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	_, err1 := w.Write(buffer)
	if err1 != nil {
		w.Close()
		b.Reset()
		return nil, err1
	}

	err2 := w.Close()
	if err2 != nil {
		b.Reset()
		return nil, err2
	}
	buff := b.Bytes()
	b.Reset()
	return buff, nil
}
