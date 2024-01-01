## Installation
``` 
go get github.com/cursusdb/cursusdb-go
```

OR

``` 
go mod download github.com/cursusdb/cursusdb-go
```

## Usage
``` 
package main

import (
	"fmt"
	cursusdbgo "github.com/cursusdb/cursusdb-go"
)

func main() {
	var client *cursusdbgo.Client

	client = &cursusdbgo.Client{
		TLS:         false,
		ClusterHost: "0.0.0.0",
		ClusterPort: 7681,
		Username:    "someuser",
		Password:    "somepassword",
		ClusterReadTimeout: time.Now().Add(time.Second * 10)
	}

	err := client.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer client.Close()

	res, err := client.Query(`ping;`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(res)

}

```