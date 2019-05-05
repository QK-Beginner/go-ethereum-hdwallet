package main

import (
	"fmt"
	_ "fmt"
	hdwallet "go-ethereum-hdwallet"
	"log"
	_ "log"
	"strings"
)

func main() {
	//生成助记词
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		log.Fatal(err)
	}
	words := strings.Split(mnemonic, " ")
	if len(words) != 12 {
		log.Fatal(err)
	}
	fmt.Printf("助记词: %s\n", words)

	////根据助记词生成
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("账户",account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947




	//prikey ,_:=wallet.PrivateKey(account)

	//fmt.Println(prikey)


	//path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	//account, err = wallet.Derive(path, false)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559*/
}
