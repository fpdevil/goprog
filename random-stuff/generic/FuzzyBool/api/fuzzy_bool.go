package api

import "fmt"

// FuzzyBool is the type for denoting boolean values
type FuzzyBool struct{ value float32 }

// New creates a new instance of the FuzzyBool for clients
// it takes an argument which complies to an empty interface
// which indicates that any data type can be used, which ensures
// that the application will not crash with a wrong input
func New(value interface{}) (*FuzzyBool, error) {
	amount, err := float32ForValue(value)
	return &FuzzyBool{amount}, err
}

// float32ForValue is used for type casting a value in the range of
// [0.0, 1.0] into a float32 value
func float32ForValue(value interface{}) (fuzzy float32, err error) {
	// perform type switching for type cast
	switch value := value.(type) {
	case float32:
		fuzzy = value
	case float64:
		fuzzy = float32(value)
	case int:
		fuzzy = float32(value)
	case bool:
		fuzzy = 0
		if value {
			fuzzy = 1
		}
	default:
		return 0, fmt.Errorf("float32ForValue(): %v is not a valid number or boolean", value)
	}
	if fuzzy < 0 {
		fuzzy = 0
	} else if fuzzy > 0 {
		fuzzy = 1
	}
	return fuzzy, nil
}

// String function for satisfying the Stringer interface for rendering the
// fuzzy booleans.
func (fuzzy *FuzzyBool) String() string {
	return fmt.Sprintf("%.0f%%", 100*fuzzy.value)
}

// Set function sets a specific value for the fuzzy boolean making the
// fuzzy booleans mutable.
func (fuzzy *FuzzyBool) Set(value interface{}) (err error) {
	fuzzy.value, err = float32ForValue(value)
	return err
}

// Copy function is for the custom types which are passed around as pointers.
func (fuzzy *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool{fuzzy.value}
}

// Not function serves the functionality of boolean logical NOT
func (fuzzy *FuzzyBool) Not() *FuzzyBool {
	return &FuzzyBool{1 - fuzzy.value}
}

// And function serves the functionality of boolean logical AND over the supplied
// arguments. It needs atleast one argument complying to *FuzzyBool to work with
// the first argument. It can accept 0 or more arguments. Essentially, it just
// returns the minimum of the given fuzzy values.
func (fuzzy *FuzzyBool) And(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	minimum := fuzzy.value
	rest = append(rest, first)
	for _, other := range rest {
		if minimum > other.value {
			minimum = other.value
		}
	}
	return &FuzzyBool{minimum}
}

// Or function serves the functionality of boolean logical OR.
// Its logic is much similar to the And function defined earlier, with the
// only difference being Or uses maximum variable instead of minimum.
func (fuzzy *FuzzyBool) Or(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	maximum := fuzzy.value
	rest = append(rest, first)
	for _, other := range rest {
		if maximum < other.value {
			maximum = other.value
		}
	}
	return &FuzzyBool{maximum}
}

//!+
// Below methods are for comparing the fuzzy booleans interms of float32s

// Less function checks if one value os less than the other
func (fuzzy *FuzzyBool) Less(other *FuzzyBool) bool {
	return fuzzy.value < other.value
}

// Equal function checks if one value is equal to the other
func (fuzzy *FuzzyBool) Equal(other *FuzzyBool) bool {
	return fuzzy.value == other.value
}

// Bool function returns the boolean form of the given float32 based
// fuzzy boolean value.
func (fuzzy *FuzzyBool) Bool() bool {
	return fuzzy.value >= .5
}

// Float function returns the 64-bit floating point form of the given float32
// base fuzzy boolean value.
func (fuzzy *FuzzyBool) Float() float64 {
	return float64(fuzzy.value)
}

//!-
