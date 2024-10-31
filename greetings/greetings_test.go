package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"
	want := regexp.MustCompile(`\b` + name + `\b`) //busca una coincidencia exacta con el nombre en expresion regular
	msg, err := Hello("Juan")
	if !want.MatchString(msg) || err != nil{
		t.Fatalf(`Hello("%s") = %q, %v, quiere coincidencia para %#q, nil`, name, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T){
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}