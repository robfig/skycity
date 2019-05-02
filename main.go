package main

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/jsonpb"
	_ "github.com/robfig/teamcity-skycfg/teamcity"
	"github.com/stripe/skycfg"
)

func main() {
	// config, err := skycfg.Load(context.Background(), filename, skycfg.WithProtoRegistry(&protoRegistry{}))
	ctx := context.Background()
	config, err := skycfg.Load(ctx, "testdata/pipeline.sky")
	if err != nil {
		panic(err)
	}
	messages, err := config.Main(ctx)
	if err != nil {
		panic(err)
	}

	var jsonMarshaler = &jsonpb.Marshaler{OrigName: true}
	for _, msg := range messages {
		marshaled, err := jsonMarshaler.MarshalToString(msg)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(marshaled)
	}

}
