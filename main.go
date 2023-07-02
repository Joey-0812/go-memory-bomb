package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	pstruct "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/protobuf/proto"
)

func EncodeMetadata(i int64) string {
	pb := &pstruct.Struct{
		Fields: map[string]*pstruct.Value{
			"WORKLOAD_NAME": &pstruct.Value{
				Kind: &pstruct.Value_StringValue{
					StringValue: strings.Repeat("a", 128) + strconv.FormatInt(i, 10),
				},
			},
		},
	}
	bytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal(err)
	}
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}
	url := "http://127.0.0.1:8080"
	var i int64
	for i < 10_000 {
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("x-envoy-peer-metadata-id", strconv.FormatInt(i, 10))
		req.Header.Add("x-envoy-peer-metadata", EncodeMetadata(i))
		_, _ = client.Do(req)
		i++
		log.Printf("clent request: %+v", i)
	}
}
