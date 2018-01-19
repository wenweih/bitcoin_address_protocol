package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

func byteString(b []byte) (s string) {
	s = ""
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%02X", b[i])
	}
	return s
}

// b58encode encodes a byte slice b into a base-58 encoded string.
func b58encode(b []byte) (s string) {
	/* See https://en.bitcoin.it/wiki/Base58Check_encoding */

	const BITCOIN_BASE58_TABLE = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* Convert big endian bytes to big int */
	x := new(big.Int).SetBytes(b)

	/* Initialize */
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	/* Convert big int to string */
	for x.Cmp(zero) > 0 {
		/* x, r = (x / 58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(BITCOIN_BASE58_TABLE[r.Int64()]) + s
	}

	return s
}

// b58checkencode encodes version ver and byte slice b into a base-58 check encoded string.
func b58checkencode(ver uint8, b []byte) (s string) {
	/* Prepend version */
	fmt.Println("4 - Add version byte in front of RIPEMD-160 hash (0x00 for Main Network)")
	bcpy := append([]byte{ver}, b...)
	fmt.Println(byteString(bcpy))
	fmt.Println("=======================")

	/* Create a new SHA256 context */
	sha256H := sha256.New()

	/* SHA256 Hash #1 */
	fmt.Println("5 - Perform SHA-256 hash on the extended RIPEMD-160 result")
	sha256H.Reset()
	sha256H.Write(bcpy)
	hash1 := sha256H.Sum(nil)
	fmt.Println(byteString(hash1))
	fmt.Println("=======================")

	/* SHA256 Hash #2 */
	fmt.Println("6 - Perform SHA-256 hash on the result of the previous SHA-256 hash")
	sha256H.Reset()
	sha256H.Write(hash1)
	hash2 := sha256H.Sum(nil)
	fmt.Println(byteString(hash2))
	fmt.Println("=======================")

	/* Append first four bytes of hash */
	fmt.Println("7 - Take the first 4 bytes of the second SHA-256 hash. This is the address checksum")
	fmt.Println(byteString(hash2[0:4]))
	fmt.Println("=======================")

	fmt.Println("8 - Add the 4 checksum bytes from stage 7 at the end of extended RIPEMD-160 hash from stage 4. This is the 25-byte binary Bitcoin Address.")
	bcpy = append(bcpy, hash2[0:4]...)
	fmt.Println(byteString(bcpy))
	fmt.Println("=======================")

	/* Encode base58 string */
	s = b58encode(bcpy)

	/* For number of leading 0's in bytes, prepend 1 */
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}
	fmt.Println("9 - Convert the result from a byte string into a base58 string using Base58Check encoding. This is the most commonly used Bitcoin Address format")
	fmt.Println(s)
	fmt.Println("=======================")

	return s
}

// paddedAppend appends the src byte slice to dst, returning the new slice.
// If the length of the source is smaller than the passed size, leading zero
// bytes are appended to the dst slice before appending src.
func paddedAppend(size uint, dst, src []byte) []byte {
	for i := 0; i < int(size)-len(src); i++ {
		dst = append(dst, 0)
	}
	return append(dst, src...)
}
