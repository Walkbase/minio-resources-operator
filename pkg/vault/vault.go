package vault

import (
	"fmt"

	"github.com/hashicorp/vault/api"
)

var client api.Client

func init() {
	client, err := api.NewClient(nil)
	if err != nil {
		panic("Vault is fucked")
	}
}

func HasKeys(user string) (bool, error) {
	l := client.Logical()

	secret, err := l.Read("miniokeys/" + user)

	fmt.Println("sekret %v", secret)
	fmt.Println("error" + err.Error())

	return true, nil
}
