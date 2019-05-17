package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	"log"
)

func createKs() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "secret"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3
}

func importKs() {

	dir := "./tmp"
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := file.Name()
		jsonByte, err := ioutil.ReadFile("./tmp/"+filename);
		if err != nil{
			log.Fatal(err)
		}
		ks,err := keystore.DecryptKey(jsonByte, "secret");
		fmt.Println(ks)
	}
}

func main() {
	//createKs()
	importKs()
}
