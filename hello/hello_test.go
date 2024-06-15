package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to people", func(t *testing.T) {
		got:= Hello("Chris","kikuyu")
		want:= "Hello, Chris"
		assertCorrectMessage(t,got,want)
	})
	t.Run("Say 'Hello, World'when an empty string is supplied",func(t *testing.T) {
		got:=Hello("","kijerumani")
		want := "Hello, World"
		assertCorrectMessage(t,got,want)
	})
	t.Run("Greating in spanish",func(t *testing.T) {
		got:= Hello("Ediel","Spanish")
		want:= "Hola, Ediel"
		assertCorrectMessage(t,got,want)
	})
	
	t.Run("Greating in French",func(t *testing.T) {
		got:= Hello("Lukas","French")
		want:= "Bonjour, Lukas"
		assertCorrectMessage(t,got,want)
	})
}
func assertCorrectMessage(t testing.TB,got,want string ){
	t.Helper()
	if got !=want {
		t.Errorf("Found %s but required %s",got,want)
	}
}