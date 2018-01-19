package main

import "fmt"

func main() {
	wallet := NewWallet()

	fmt.Println("0 - Having a private ECDSA key")
	fmt.Println(byteString(wallet.PrivateKey))
	fmt.Println("=======================")
	// fmt.Println("private wallet import format")
	// fmt.Println("private wallet import format", ToWIF(wallet.PrivateKey))
	// fmt.Println("=======================")
	fmt.Println("1 - Take the corresponding public key generated with it (65 bytes, 1 byte 0x04, 32 bytes corresponding to X coordinate, 32 bytes corresponding to Y coordinate)")
	fmt.Println("raw public key", byteString(wallet.PublicKey))
	fmt.Println("=======================")
	wallet.GetAddress()
}
