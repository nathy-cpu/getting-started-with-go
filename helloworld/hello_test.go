package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("Testing with empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing with english", func(t *testing.T) {
		got := Hello("John", "English")
		want := "Hello, John"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing with spanish", func(t *testing.T) {
		got := Hello("Hector", "Spanish")
		want := "Hola, Hector"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Testing with french", func(t *testing.T) {
		got := Hello("Jean", "French")
		want := "Bonjour, Jean"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got: %q want: %q", got, want)
	}
}
