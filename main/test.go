package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "fmt"
	_ "github.com/go-sql-driver/mysql"
	hdwallet "go-ethereum-hdwallet"
	"log"
	_ "log"
	"strings"
	"time"
)

func main() {

	for i:=0;i<5000;i++ {
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

		fmt.Println("账户", account.Address.Hex()) // 0xC49926C4124cEe1cbA0Ea94Ea31a6c12318df947

		prikey, _ := wallet.PrivateKeyHex(account)

		fmt.Println(prikey)

		insertData(account.Address.Hex(),prikey,mnemonic)


	}
	


}

func insertData(address string, prikey string, word string) int64 {
	db, err := sql.Open("mysql", "root:Wu123456789!@#@tcp(159.138.4.111:3306)/xrphelp?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}

	db.SetConnMaxLifetime(100*time.Second)  //最大连接周期，超过时间的连接就close
	db.SetMaxOpenConns(100)//设置最大连接数
	db.SetMaxIdleConns(16) //设置闲置连接数
	defer db.Close()

	stmt, err := db.Prepare(`INSERT wallets (address,prikey,word) values (?,?,?)`)
	if err == nil {

	}
	res, err := stmt.Exec(address, prikey, word)
	if err == nil {

	}
	// checkErr(err)
	id, err := res.LastInsertId()
	if err == nil {

	}
	fmt.Println("成功插入:",id)
	time.Sleep(time.Second*)
    return id
}
