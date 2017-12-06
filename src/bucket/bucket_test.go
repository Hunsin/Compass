package bucket

import (
	"flag"
	"os"
	"testing"

	_ "github.com/lib/pq"
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
	flag.StringVar(&name, "db", "bucket", "bucket test database name")
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
	_, err = Open(host, port, name, usr, pwd, ssl)
	if err != nil {
		t.Errorf("Open exits with error %v", err)
	}
}

func TestInitTables(t *testing.T) {
	err := bk.InitTables()
	if err != nil {
		t.Errorf("InitTables exits with error %v", err)
	}

	for _, n := range []string{"averages", "daily", "securities"} {
		_, err = bk.db.Exec("DROP TABLE " + n)
		if err != nil {
			t.Errorf("Table %s not created, exits with error %v", n, err)
		}
	}
}
