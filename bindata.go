package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets36eba4f2eb6f7f8def37b062719b6f86c6d4f395 = "password: $CONFIELD_DBPASS|mypass\nuser: $CONFIELD_DBUSER|root\nname: $CONFIELD_DBNAME|db_dev\nhost: $CONFIELD_DBHOST|localhost\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"settings.yml"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1512658028, 1512658028099181315),
		Data:     nil,
	}, "/settings.yml": &assets.File{
		Path:     "/settings.yml",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1512657611, 1512657611623784028),
		Data:     []byte(_Assets36eba4f2eb6f7f8def37b062719b6f86c6d4f395),
	}}, "")
