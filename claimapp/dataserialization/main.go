package main

import (
	"fmt"
	"time"

	v1 "github.com/eswaribala/claimapp/dataserialization/protobufsrc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	claim := &v1.Claim{
		ClaimId:     123,
		ClaimType:   "Alice",
		CreatedAt:   timestamppb.Now(),
		Status:      "Pending",
		Description: "Test claim",
		ClaimAmount: 1000,
	}

	// Encode (marshal) to []byte
	b, err := proto.Marshal(claim)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Serialized bytes (%d): %v\n", len(b), b[:min(16, len(b))])

	// Decode (unmarshal) to struct
	var out v1.Claim
	if err := proto.Unmarshal(b, &out); err != nil {
		panic(err)
	}
	fmt.Printf("Decoded: id=%d type=%s created_at=%T\n", out.ClaimId, out.ClaimType, out.CreatedAt)

	// Presence checks (proto3 with well-known types or oneofs)
	if out.CreatedAt != nil && out.CreatedAt.AsTime().After(time.Time{}) {
		fmt.Println("created_at present:", out.CreatedAt.AsTime())
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
