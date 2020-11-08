package maps


type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string{
	return string(e)
}

const (
	ErrorDoesNotExist = DictionaryErr("Word does not exist")
	ErrorWordExists = DictionaryErr("word already exists in the dictionary")
	ErrorCannotUpdate = DictionaryErr("Cannot update and the given word does not exist")
)


func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrorDoesNotExist
	}
	return definition, nil
}


func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorDoesNotExist: 
		 d[word] = definition
	
	case nil: 
		return ErrorWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorDoesNotExist:
		return ErrorCannotUpdate
	case nil:
		d[word]	= newDefinition
	default:
		return err	
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}