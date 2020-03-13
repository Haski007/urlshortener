package config

import (
	"time"
)

// Port on which server will be listened.
var Port = ":8080"

// HostName on which server will be listened.
var HostName = "localhost"

// Expiration - expiration time (in UnixNano)
// when cache item will be removed.
var Expiration = time.Minute * 2


///////////////////////////// DataBase \\\\\\\\\\\\\\\\\\\\\\\\\\\

// DataBaseHost - hostname on which mongoDB is hosting.
var DataBaseHost = "192.168.99.101"

// DataBasePort - port on which mongoDB is hosting.
var DataBasePort = ":27017"