package main

import (
	"fmt"
)

func main() {

	crypted_passwd := "cdiiddwpgswtgt"

	for decal := 1; decal < 27; decal++ {
		fmt.Printf("decalage (%d) --> ", decal)
		fmt.Println(decrypt(crypted_passwd, decal))
	}
}

func decrypt(crypted_passwd string, decal int) (string) {

	decrypt_passwd := ""

	for i := 0; i < len(crypted_passwd); i++ {
		decrypt_passwd += string(change_letter(crypted_passwd[i], byte(decal)))
	}

	return (decrypt_passwd)
}

func change_letter(c byte, decal byte) (byte) {

	if c + decal <= 122 {
		return (c + decal)
	}

	new := (c - 122 + decal) + 96
	return (new)
}