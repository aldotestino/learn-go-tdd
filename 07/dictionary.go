package main

// To initialize an empty map:
// var m map[string]string ❌
// var m = map[string]string{} ✅
// var m = make(map[string]string) ✅
// m := map[string]string{} ✅
// m := make(map[string]string) ✅

// DictionaryErr custom error type
type DictionaryErr string

// Error error interface implemented by DictionaryErr
func (e DictionaryErr) Error() string {
	return string(e)
}

// ErrNotFound throwed by Search
// ErrWordExists throwed by Add
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// Dictionary is a map
type Dictionary map[string]string

// Adding methods to map doesn't need `*` on the receiver type

// Search a method of the Dictionary type
func (d Dictionary) Search(word string) (string, error) {
	// ok is a boolena that indicate if exists a value for a
	// given key
	defintion, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return defintion, nil
}

// Add adds a word definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
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

// Update updates a word definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete deletes a given word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
