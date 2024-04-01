package users

import (
	"fmt"
	"time"

	"github.com/jorojas/gofromscratch/models"
)

func CreateUser() {
	u := new(models.User)
	u.AddUser(10, "JOROJAS", time.Now(), true)

	fmt.Println(u)
}
