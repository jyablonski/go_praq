package main

var (
	ErrorNotFound         = DictionaryErr("could not find the work you're looking for")
	ErrorWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrorWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

type Dictionary map[string]string

// add error handling to return an error when a word is not found
func (d Dictionary) Search(word string) (string, error) {
	defintion, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return defintion, nil

}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrorNotFound:
		return ErrorWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
