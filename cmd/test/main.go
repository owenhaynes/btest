package main

import (
	"time"

	"google.golang.org/protobuf/types/known/anypb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {

	t := timestamppb.New(time.Now())
	_, err := anypb.New(t)
	if err != nil {
		panic(err)
	}
}
