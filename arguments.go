package cwl

// Argument represents ""
type Argument struct {
	Value   string
	Binding *Binding
}

// New constructs an "Argument" struct from any interface.
func (_ Argument) New(i interface{}) Argument {
	dest := Argument{}
	switch x := i.(type) {
	case string:
		dest.Value = x
	case map[string]interface{}:
		dest.Binding = Binding{}.New(x)
	}
	return dest
}

// Arguments ...
type Arguments []Argument

// New constructs "Arguments" struct.
func (_ Arguments) New(i interface{}) Arguments {
	dest := Arguments{}
	switch x := i.(type) {
	case []interface{}:
		for _, v := range x {
			dest = append(dest, Argument{}.New(v))
		}
	default:
		dest = append(dest, Argument{}.New(x))
	}
	return dest
}
