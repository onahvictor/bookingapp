package models

//Holds data sent from handler to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMpa  map[string]float32
	Data      map[string]interface{}
	CSRFToken string // security token cross site fugery token
	Flash     string
	Warning   string
	Error     string
}

//This will only exist to be imported by packages other than itself
