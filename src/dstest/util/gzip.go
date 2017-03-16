package util

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"io"
	"strings"
)

func GzipString(str string) []byte {
	buf := new(bytes.Buffer)
	w := gzip.NewWriter(buf)
	defer w.Close()
	if _, err := io.WriteString(w, str); err != nil {
		panic(err)
	}
	if err := w.Flush(); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func UnGzipBytes(b []byte) (string, error) {
	br := bytes.NewReader(b)
	r, err := gzip.NewReader(br)
	if err != nil {
		return "", err
	}

	var str string
	for {
		p := make([]byte, 4*1024)
		n, err := r.Read(p)
		if err == io.EOF || n <= 0 {
			break
		} else if err != nil {
			return "", err
		}
		str += string(p[:n])
	}

	return str, nil
}

func IsGzipCompressed(s string) bool {
	gzipID, _ := hex.DecodeString("1f8b")
	return strings.HasPrefix(s, string(gzipID))
}
