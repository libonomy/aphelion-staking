package main

import (
	"fmt"
	"os"

	amino "github.com/evdatsion/go-amino"
	cryptoAmino "github.com/evdatsion/tendermint/crypto/encoding/amino"
)

func main() {
	cdc := amino.NewCodec()
	cryptoAmino.RegisterAmino(cdc)
	cdc.PrintTypes(os.Stdout)
	fmt.Println("")
}
