package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)

	return &Environment{store: s}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}

	return obj, ok
}

// For assigning
func (e *Environment) Set(name string, val Object) Object {
	_, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Set(name, val)
	}

	e.store[name] = val

	return val
}

// For local scope variables
func (e *Environment) Let(name string, val Object) Object {
	e.store[name] = val

	return val
}
