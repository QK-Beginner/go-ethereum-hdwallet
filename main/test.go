package main

import (
	"database/sql"
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


	prikey ,_:=wallet.PrivateKeyHex(account)

	fmt.Println(prikey)


//	insert(account.Address.Hex(),prikey,words[0])

	db,err := sql.Open("mysql","root:Wu123456789!@#@tcp(159.138.4.111:3306)/xrphelp?charset=utf8")
	if err != nil{
		fmt.Println(err)
	}
	ret, _ := db.Exec("insert into wallets(address,prikey,word) values(address, prikey,word)");
	//获取插入ID
	ins_id, _ := ret.LastInsertId();
	fmt.Println(ins_id);


	//fmt.Println(account.Address.Hex()) // 0x8230645aC28A4EdD1b0B53E7Cd8019744E9dD559*/
}



