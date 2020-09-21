## Connection Log

### Example

```
package main

import (
	"fmt"

	"github.com/ifrasoft/connection"
)

func main() {

	cl := connection.NewConnectionLog()
	cl.AddConnection(&connection.Connection{Name: "MySQL", Host: "www.facebook.com", Port: "3306"})
	cl.AddConnection(&connection.Connection{Name: "MySQL", Host: "www.google.com", Port: "3306"})
	cl.AddConnection(&connection.Connection{Name: "FTP", Host: "www.google.com", Port: "21"})
	cl.AddConnection(&connection.Connection{Name: "HTTPS", Host: "www.google.com", Port: "443"})
	cl.AddConnection(&connection.Connection{Name: "HTTP", Host: "www.google.com", Port: "80"})

	results := cl.Check()
	for _, result := range results {
		fmt.Println(result.Connection.Host, result.Connection.Port, result.Status)
	}

}

```
Result
```
www.google.com 80 OK
www.google.com 443 OK
www.google.com 21 FAIL
www.google.com 21 OK
www.google.com 3306 FAIL

```

