package settings

import (
	"os"
	"strings"
)

type Vars struct {
	Env     string
	Version string
	DB      DbInfo
}
type DbInfo struct {
	Host     string
	Port     string
	User     string
	Password string
	Table    string
}

var V Vars

func init() {

	V.Version = "0.0.0"

	V.Env = "LOCAL"
	if temp := os.Getenv("GO_ENV"); temp != "" {
		V.Env = strings.ToUpper(temp)
	}
	switch V.Env {
	case "LOCAL":
		host := "localhost"
		port := "8889"
		user := "root"
		password := "root"
		table := "example"
		V.DB = DbInfo{Host: host, Port: port, User: user, Password: password, Table: table}
	}
}
