package main
import (
  "testing"
)
func TestGoRocks(t *testing.T){
	got := greeting("Code.education Rocks!!!")
	want:= "Code.education Rocks!!!"

	if got != want{
		t.Errorf("Retornou a mensagem errada.!!")
	}

}
