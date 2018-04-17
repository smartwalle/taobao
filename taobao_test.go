package taobao

import (
	"os"
	"testing"
)

var client *TaoBao

func TestMain(m *testing.M) {

	client = New("23201003", "1e2dfd16981d75142597fd10131b17b5", true)

	os.Exit(m.Run())
}
