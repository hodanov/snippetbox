package forms

type errors map[string][]string

// An Add() is the function to add error messages for a given field to the map.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// A Get() is the method to retrieve the first error messages for a given field from the map.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
