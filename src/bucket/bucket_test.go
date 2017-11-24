package bucket

import (
	"flag"
	"os"
	"testing"
)

var (
	bk   *Bucket
	host string
	port int
	name string
	usr  string
	pwd  string
	ssl  bool
)

func init() {
	flag.StringVar(&host, "host", "localhost", "bucket test database host")
	flag.StringVar(&name, "name", "bucket", "bucket test database name")
	flag.StringVar(&usr, "user", os.Getenv("USER"), "bucket test database user")
	flag.StringVar(&pwd, "pwd", "pwd", "bucket test database password")
	flag.BoolVar(&ssl, "ssl", false, "bucket test database ssl mode")
	flag.IntVar(&port, "port", 5432, "bucket test database port")
}

func TestMain(m *testing.M) {
	flag.Parse()
	bk, _ = Open(host, port, name, usr, pwd, ssl)
	c := m.Run()
	os.Exit(c)
}

func TestOpen(t *testing.T) {
	var err error
	bk, err = Open(host, port, name, usr, pwd, ssl)
	if err != nil {
		t.Errorf("Open exits with error %v", err)
	}
}

func TestCreateTables(t *testing.T) {
	err := bk.CreateTables()
	if err != nil {
		t.Errorf("CreateTables exits with error %v", err)
	}
}
