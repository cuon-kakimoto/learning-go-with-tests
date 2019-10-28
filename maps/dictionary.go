package main

var (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// https://dave.cheney.net/2016/04/07/constant-errors
type DictionaryErr string

func (e DictionaryErr) Error() string {
    return string(e)
}

type Dictionary map[string]string
func (d Dictionary) Search(word string) (string, error){
	definition, ok := d[word]

	if !ok{
		return "", ErrNotFound
	}

	return definition, nil
}

// HACK
// An interesting property of maps is that you can modify them without passing them as a pointer. This is because map is a reference type. 
// mapは参照だからpointerで渡さなくていいらしい
// https://en.wikipedia.org/wiki/Hash_table
func (d Dictionary) Add(word, definition string) error{

	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error{
	_, err := d.Search(word)

	switch err{
	case ErrNotFound:
		return ErrWordDoesNotExist
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