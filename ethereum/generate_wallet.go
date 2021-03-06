package main

import (
    "crypto/ecdsa"
    "fmt"
    "log"
    "strings"

    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/crypto"
)

var attempts int

func genAddress(inputLength int) string {
  privateKey, err := crypto.GenerateKey()
  if err != nil {
    log.Fatal(err)
  }

  privateKeyBytes := crypto.FromECDSA(privateKey)
  fmt.Println("Private key:", hexutil.Encode(privateKeyBytes)[2:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

  publicKey := privateKey.Public()
  publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
  if !ok {
    log.Fatal("error casting public key to ECDSA")
  }

  address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
  fmt.Println("Public key:", address) // 0x96216849c49358B10257cb55b28eA603c874b05E

  return strings.ToLower(address[0:inputLength])
}

func runIteration (generatedPubKey, inputText string) {
  if generatedPubKey == inputText {
    fmt.Println("Successful key generated ;-)")
  } else {
    attempts++
    fmt.Println("Attempt number:", attempts)
    runIteration(genAddress(len(inputText)), inputText)
  }
}

func main() {
  fmt.Print("What would you like it say after 0x:")
  var input string
  fmt.Scanln(&input)
  input = "0x" + input
  var inputLength = len(input)

  runIteration(genAddress(inputLength), input)
}
