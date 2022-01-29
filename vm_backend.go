package web

/* Web function representation in VM */
type Function struct {
	Args map[string]interface{}
}

/* Web class represantation in VM */
type Class struct {
	Attributes map[string]interface{}
	Methods    map[string]Function
}

/* Main VM struct */
type VM struct {
	classes map[string]Class
}
