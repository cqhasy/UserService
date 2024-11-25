package data

import (
	"reflect"
	"testing"
)

func TestOpenYaml(t *testing.T) {
	got, err := NewMysqlConfig("../configs/mysql.yaml")
	if err != nil {
		t.Error(err)
	}
	want := MysqlConfig{
		Mysql: struct {
			Dsn string `yaml:"dsn"`
		}{Dsn: "root:users@tcp(127.0.0.1:3307)/users?charset=utf8mb4&parseTime=True&loc=Local"},
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v want %v", got, want)
	}
}
