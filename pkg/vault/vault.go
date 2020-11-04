package vault

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

func init() {
	c, err := api.NewClient(nil)

	if err != nil {
		panic("poop")
	}

	s, err := c.Logical().List("secret/")
	fmt.Println(s)

}
