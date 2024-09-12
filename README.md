# kukumail.go
Golang Library for "m.kuku.lu" | Temp Mail 
## example

### Create Email

```go
package main

import (
	"fmt"
	"github.com/evex-dev/kukumailgo"
)

func main() {
	s := kukumailgo.New("yourSessionHash", "yourCesrToken")
	err := s.CreateDesignateEmail("fjsaofjoa2342", "tatsu.uk")
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
  mail, err := s.CreateRandomEmail()
	if err != nil {
		fmt.Println(err)
		return
	}
	*/
}
```

### GetReceivedEmail

```go
package main

import (
	"fmt"
	"github.com/evex-dev/kukumailgo"
)

func main() {
	s := kukumailgo.New("yourSessionHash", "yourCesrToken")
	mails, err := s.GetAllReceivedEmail()
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
  mails, err := s.SearchReceivedEmail("discord")
	if err != nil {
		fmt.Println(err)
		return
	}
	*/
}
```

### GetEmailContent

```go
package main

import (
	"fmt"
	"github.com/evex-dev/kukumailgo"
)

func main() {
	s := kukumailgo.New("yourSessionHash", "yourCesrToken")
	
	content, err := s.GetEmailContent(mail)
	if err != nil {
		fmt.Println(err)
		return
	}
}

```
