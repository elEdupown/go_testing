package users

import (
	"fmt"
	"time"

	"github.com/elEdupown/go_testing/modelos"
)

func MiUsuario() {
	u := new(modelos.User)
	u.AddUser(1, "Eduardo", time.Now(), true)
	fmt.Println(u)
}
