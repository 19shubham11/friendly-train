package maps

import (
	"testing"
)


func assertStrings(t *testing.T, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}


func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    if definition != got {
        t.Errorf("got %q want %q", got, definition)
    }
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
    if got == nil {
        if want == nil {
            return
        }
        t.Fatal("expected to get an error.")
    }
}


func TestSearch(t *testing.T) {
	t.Run("Search a word that exists", func(t *testing.T){
		dictionary := Dictionary{"test": "this is just a test"}
	
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Search a word that does not exists", func(t *testing.T){
		dictionary := Dictionary{"test": "this is just a test"}
	
		_, err := dictionary.Search("test1")
		want := ErrorDoesNotExist.Error()

		if err == nil {
            t.Fatal("expected to get an error.")
        }

		assertStrings(t, err.Error(), want)
	})
}


func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T){
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"
	
		err := dictionary.Add(word, definition)
		
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing word", func(t *testing.T){
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		
		err := dictionary.Add(word, definition)
		
		assertError(t, err, ErrorWordExists)
	})
}


func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "key"
		definition := "key key"
	
		definitionWanted := "key key updated!"
	
		dictionary := Dictionary{word: definition}
	
	
		err := dictionary.Update(word, definitionWanted)
	
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definitionWanted)
	})

	t.Run("An unknown word", func(t *testing.T) {
		word := "key"
		definition := "key key"
	
		definitionWanted := "key key updated!"
	
		dictionary := Dictionary{word: definition}
	
	
		err := dictionary.Update("hey", definitionWanted)
		
		assertError(t, err, ErrorCannotUpdate)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "heyy"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if (err != ErrorDoesNotExist) {
		t.Errorf("Expected %q to be deleted", word)
	}

}