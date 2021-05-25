package object

import (
	"fmt"
	"hash/fnv"

	"github.com/lczm/as/ast"
)

// Types
const (
	BOOL     = "BOOL"
	INTEGER  = "INTEGER"
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	STRING   = "STRING"
	BUILTIN  = "BULITIN" // builtin functions from the host language
	LIST     = "LIST"
	HASHMAP  = "HASHMAP"
)

// All types implement this interface
type Object interface {
	RawType() string
	Type() string
	String() string
}

// All the call-able objects will implement this interface
// i.e. functions
type Callable interface {
	Call() Object
}

// TODO : Type can either be an `object` or a `string`
// But for now I think string will work just fine?
// Objects will implement RawType() which can then
// be used to supply this value.
type HashValue struct {
	Type  string
	Value int
}

// All hashable objects will implement this interface
// i.e. anything that can be used as a hashmap key
type Hashable interface {
	Hash() HashValue
}

// Boolean type
type Bool struct {
	Value bool
}

func (b *Bool) RawType() string {
	return BOOL
}

func (b *Bool) Type() string {
	return fmt.Sprintf("<type: %s>", BOOL)
}

func (b *Bool) String() string {
	if b.Value == true {
		return "true"
	}
	return "false"
}

func (b *Bool) Hash() HashValue {
	if b.Value == true {
		return HashValue{Type: b.RawType(), Value: 1}
	} else {
		return HashValue{Type: b.RawType(), Value: 0}
	}
}

// Integer type
type Integer struct {
	Value int64
}

func (i *Integer) RawType() string {
	return INTEGER
}

func (i *Integer) Type() string {
	return fmt.Sprintf("<type: %s>", INTEGER)
}

func (i *Integer) String() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Hash() HashValue {
	return HashValue{Type: i.RawType(), Value: int(i.Value)}
}

// String type
type String struct {
	Value string
}

func (s *String) RawType() string {
	return STRING
}

func (s *String) Type() string {
	return fmt.Sprintf("<type: %s>", STRING)
}

func (s *String) String() string {
	return s.Value
}

func (s *String) Hash() HashValue {
	// FNV-1a hash
	hash := fnv.New64a()
	hash.Write([]byte(s.Value))
	return HashValue{Type: s.RawType(), Value: int(hash.Sum64())}
}

// Function type, it is an Object as well as a Callable
type Function struct {
	FunctionStatement ast.FunctionStatement
}

func (f *Function) RawType() string {
	return FUNCTION
}

func (f *Function) Type() string {
	return fmt.Sprintf("<type: %s>", FUNCTION)
}

func (f *Function) String() string {
	return fmt.Sprintf("Function : <%s>", f.FunctionStatement.Name.Literal)
}

type BuiltinFunction struct {
	Name string
	Fn   func(args ...Object) Object
}

func (bf *BuiltinFunction) RawType() string {
	return BUILTIN
}

func (bf *BuiltinFunction) Type() string {
	return fmt.Sprintf("<type: %s>", BUILTIN)
}

func (bf *BuiltinFunction) String() string {
	return fmt.Sprintf("BulitinFunction: <%s>", bf.Name)
}

// The call functions should return an object
// in the case of something like
// var x = function(...)
func (f *Function) Call() {}

// Return type, this is only for the interpreter and is not for use normally.
type Return struct {
	Value Object
}

func (r *Return) RawType() string {
	return RETURN
}

func (r *Return) Type() string {
	return fmt.Sprintf("Return: <%s>", RETURN)
}

func (r *Return) String() string {
	return r.Value.String()
}

// Container types - Lists/Hashmaps
// List container type
type List struct {
	Value []Object
}

func (l *List) RawType() string {
	return LIST
}

func (l *List) Type() string {
	return fmt.Sprintf("List: <%s>", LIST)
}

func (l *List) String() string {
	var valueStrings []string
	for i := 0; i < len(l.Value); i++ {
		valueStrings = append(valueStrings, l.Value[i].String())
	}
	// Sprintf can automatically convert an array of strings into
	// a string for the output.
	return fmt.Sprintf("%s\n", valueStrings)
}

// Hashmap container type
// Note : Hashmaps maps object.Object to object.Object
type HashMap struct {
	Value map[HashValue]Object
}

func (hm *HashMap) RawType() string {
	return HASHMAP
}

func (hm *HashMap) Type() string {
	return fmt.Sprintf("HashMap: <%s>", HASHMAP)
}

func (hm *HashMap) String() string {
	var valueStrings []string
	// for key, value := range hm.Value {
	// 	valueStrings = append(valueStrings, key.Value, ": ", value.String())
	// }
	// Sprintf can automatically convert an array of strings into
	// a string for the output.
	return fmt.Sprintf("%s\n", valueStrings)
}
