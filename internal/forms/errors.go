package forms

//declared a package level type that is not exported
type errors map[string][]string

//Adds error message for a giving form fild
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//Get returns the first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
