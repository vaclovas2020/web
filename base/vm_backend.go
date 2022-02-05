package base

/* Web function representation in VM */
type Function struct {
	Args    map[string]*interface{}
	Handler *FunctionHandler
}

/* Web class represantation in VM */
type Class struct {
	Attributes map[string]*interface{}
	Methods    map[string]Function
}

/* User defined function handler */
type FunctionHandler func(args map[string]*interface{}) (*interface{}, error)
