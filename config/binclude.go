// Code generated by binclude; DO NOT EDIT.

package config

import (
	"github.com/lu4p/binclude"
	"time"
)

var (
	_binclude0	= []byte("mongo:\r\n  host: 192.168.1.129\r\n  port: 27017\r\n  username: admin\r\n  password: pw123456\r\n\r\n# PROD为正式\r\nenv: DEV\r\njwt-key: WC7fkchY6FL6eS")
	BinFS		= binclude.FileSystem{"../assets": {Filename: "assets", Mode: 2147484159, ModTime: time.Unix(1593222293, 339024600), Compression: binclude.None, Content: nil}, "../assets/config.yml": {Filename: "config.yml", Mode: 438, ModTime: time.Unix(1593222293, 339000000), Compression: binclude.None, Content: _binclude0}}
)