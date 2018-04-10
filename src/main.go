package main

import (
  "fmt"
  "net/http"
  "github.com/shamoto-donuts/go/codec"
  "bytes"
)

type TestStruct struct 
{
  TestInt int
  TestInt16 int16
  TestInt32 int32
  TestInt64 int64
  TestUint uint
  TestUint16 uint16
  TestUint32 uint32
  TestUint64 uint64
}

func msgPackHandler(w http.ResponseWriter, r *http.Request) {
  bufbody := new(bytes.Buffer)
  bufbody.ReadFrom(r.Body)
  mh := &codec.MsgpackHandle{}
  dec := codec.NewDecoderBytes(bufbody.Bytes(), mh)
  out := TestStruct{}
  dec.Decode(&out)
  fmt.Println(out)
  w.Write([]byte("OK"))
}

func main() {
  //server
  println("Server Start")
  http.HandleFunc("/test", msgPackHandler)
  http.ListenAndServe(":1212", nil)
}