package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"unsafe"

	"github.com/altafan/go-secp256k1-zkp"
)

func exit(msg interface{}) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	ctx, _ := secp256k1.ContextCreate(secp256k1.ContextBoth)
	defer secp256k1.ContextDestroy(ctx)

	var valuebytes [8]byte
	valueLen, err := rand.Read(valuebytes[:])
	if err != nil {
		exit(err)
	}
	value := *(*uint64)(unsafe.Pointer(&valuebytes[0]))
	fmt.Printf("value=%v, valueLen=%v\n", value, valueLen)

	var blind [32]byte
	blindLen, err := rand.Read(blind[:])
	if err != nil {
		exit(err)
	}
	fmt.Printf("blind=%s, blindLen=%v\n", hex.EncodeToString(blind[:]), blindLen)

	var key [32]byte
	_, err = rand.Read(key[:])
	if err != nil {
		exit(err)
	}
	generator, err := secp256k1.GeneratorGenerate(ctx, key[:])
	if err != nil {
		exit(err)
	}
	fmt.Printf("generator=%s\n", generator.String())

	comNone, err := secp256k1.Commit(ctx, blind[:], value, generator)
	if err != nil {
		exit(err)
	}
	fmt.Printf("comNone=%v\n", comNone)
}
