package types_test

import (
	"bytes"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/alexanderbez/protorobo/types"
	"github.com/ethereum/go-ethereum/rlp"
	proto "github.com/gogo/protobuf/proto"
)

const (
	KB = 1024
	MB = KB * 1024
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func buildTinyMessage() *types.MyMessage {
	addr := make(types.Bytes, 32)
	hash := make(types.Hash, 256)
	code := make(types.Bytes, MB/2)

	rand.Read(addr)
	rand.Read(hash)
	rand.Read(code)

	return &types.MyMessage{
		Address: addr,
		Hash:    hash,
		Code:    code,
	}
}

func buildLargeMessage() *types.MyMessage {
	addr := make(types.Bytes, 32)
	hash := make(types.Hash, 256)
	code := make(types.Bytes, 2*MB)

	rand.Read(addr)
	rand.Read(hash)
	rand.Read(code)

	return &types.MyMessage{
		Address: addr,
		Hash:    hash,
		Code:    code,
	}
}

func BenchmarkGogoProtoTinyEncode(b *testing.B) {
	msg := buildTinyMessage()

	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkGogoProtoLargeEncode(b *testing.B) {
	msg := buildLargeMessage()

	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkGogoProtoTinyDecode(b *testing.B) {
	msg := buildTinyMessage()

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		decMsg := &types.MyMessage{}

		err := proto.Unmarshal(data, decMsg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkGogoProtoLargeDecode(b *testing.B) {
	msg := buildLargeMessage()

	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		decMsg := &types.MyMessage{}

		err := proto.Unmarshal(data, decMsg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkETHRLPTinyEncode(b *testing.B) {
	msg := buildTinyMessage()

	for i := 0; i < b.N; i++ {
		msgBuff := new(bytes.Buffer)

		err := rlp.Encode(msgBuff, msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkETHRLPLargeEncode(b *testing.B) {
	msg := buildLargeMessage()

	for i := 0; i < b.N; i++ {
		msgBuff := new(bytes.Buffer)

		err := rlp.Encode(msgBuff, msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkETHRLPTinyDecode(b *testing.B) {
	msg := buildTinyMessage()
	msgBuff := new(bytes.Buffer)

	err := rlp.Encode(msgBuff, msg)
	if err != nil {
		log.Fatal(err)
	}

	msgBytes := msgBuff.Bytes()

	for i := 0; i < b.N; i++ {
		decMsg := &types.MyMessage{}

		err := rlp.DecodeBytes(msgBytes, decMsg)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkETHRLPLargeDecode(b *testing.B) {
	msg := buildLargeMessage()
	msgBuff := new(bytes.Buffer)

	err := rlp.Encode(msgBuff, msg)
	if err != nil {
		log.Fatal(err)
	}

	msgBytes := msgBuff.Bytes()

	for i := 0; i < b.N; i++ {
		decMsg := &types.MyMessage{}

		err := rlp.DecodeBytes(msgBytes, decMsg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
