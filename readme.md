## Installation
``` 
go get github.com/cursusdb/cursusdb-go
```

## Usage
``` 
package main

import (
	"fmt"
	cursusdbgo "github.com/cursusdb/cursusdb-go"
)

func main() {
	var cursusdb *cursusdbgo.CursusDB

	cursusdb = &cursusdbgo.CursusDB{
		TLS:         false,
		ClusterHost: "0.0.0.0",
		ClusterPort: 7682,
		Username:    "someuser",
		Password:    "somepassword",
	}

	err := cursusdb.NewClient()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer cursusdb.Close()

	res, err := cursusdb.Query(`select * from users;`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)

}

```