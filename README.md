# URL Shortener
- Receives GET query on `http://localhost:8080` and draw simple html with input field.
- Enter any url to make it shorter and with your hostname, submit.
- It will apper shortened url which will be always redirecting on your url.

### Data storage:
- mongoDB, memoryCache.

### To run it
- Enter `make deps` at your shell.
- Enter `make`, it will compile binary `Url_shortener`
- Finally `./Url_shortener`


##### To run docker mongoDB container just `make docker-build-mongo` or `make docker-start-mongo`, or use your own mongoDB server.

### Config file
- In config file `config/config.go` you can set your:
  * Database ip address and port.
  * Host ip adress and port.
  * Expiration - time (in UnixNano) when cache item will be removed.
  
```GOlang
package config

import (
	"time"
)

// Port on which server will be listened.
var Port = ":8080"

// HostName on which server will be listened.
var HostName = "localhost"

// Expiration - time (in UnixNano)
// when cache item will be removed.
var Expiration = time.Minute * 2


///////////////////////////// DataBase \\\\\\\\\\\\\\\\\\\\\\\\\\\

// DataBaseHost - hostname on which mongoDB is hosting.
var DataBaseHost = "192.168.99.101"

// DataBasePort - port on which mongoDB is hosting.
var DataBasePort = ":27017"
```
  
