package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/d-o-n-u-t-s/lz4msgpack"
)

type TestStruct struct {
	TestInt16   int16
	TestInt32   int32
	TestInt64   int64
	TestUint16  uint16
	TestUint32  uint32
	TestUint64  uint64
	StringArray []string
}

func msgPackHandler(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	out := TestStruct{}
	UnPack(bufbody.Bytes(), &out)

	fmt.Println(out)
	packedData, _ := Pack(out)
	w.Write(packedData)
}

// 置き換え予定
func Pack(data interface{}) ([]byte, error) {
	d, e := lz4msgpack.MarshalAsArray(data)
	if e != nil {
		return nil, e
	}
	return d, e
}

// 置き換え予定
func UnPack(data []byte, out interface{}) error {
	return lz4msgpack.Unmarshal(data, &out)
}

func main() {
	//server
	println("Server Start")
	http.HandleFunc("/test", msgPackHandler)
	http.ListenAndServe(":1212", nil)
}
