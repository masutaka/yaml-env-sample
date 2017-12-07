package main

import (
	"fmt"
	"log"
	"os"

	"github.com/seihmd/confield"
	"gopkg.in/yaml.v2"
)

type dbSetting struct {
	Password confield.F
	User     confield.F
	Name     confield.F
	Host     confield.F
}

//go:generate go-assets-builder -s="/config" -o bindata.go config

func main() {
	var yml = make([]byte, 1024)

	n, err := Assets.Files["/settings.yml"].Read(yml)
	if err != nil {
		panic(err)
	}

	if int64(n) != Assets.Files["/settings.yml"].Size() {
		log.Fatal("size miss match")
	}

	// 	yml2 := `
	// password: $CONFIELD_DBPASS|mypass
	// user: $CONFIELD_DBUSER|root
	// name: $CONFIELD_DBNAME|db_dev
	// host: $CONFIELD_DBHOST|localhost
	// `
	//
	// 	fmt.Println(string(yml))
	//
	// 	yml = []byte(yml2)

	fmt.Println(string(yml))

	conf := dbSetting{}
	yaml.Unmarshal(yml, &conf)

	os.Setenv("CONFIELD_DBPASS", "dbpass")
	os.Setenv("CONFIELD_DBUSER", "dbuser")

	fmt.Println(conf.Password.String()) // "dbpass" from env var
	fmt.Println(conf.User.String())     // "dbuser" from env var
	fmt.Println(conf.Name.String())     // "db_dev" as default value
	fmt.Println(conf.Host.String())     // "localhost" as default value
}
