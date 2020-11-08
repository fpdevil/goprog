package apachereport

// A tread-safe safe map having string keys and interface{} values (which
// essentially can take any tupe of values) and that can safely be shared
// among as many goroutines as we would like without locking.

// SafeMap interface provides handlers for Insert, Delete, Find , Length,
// Update and Close operations to its satisfiers
type SafeMap interface {
	Insert(string, interface{})
	Delete(string)
	Find(string) (interface{}, bool)
	Len() int
	Update(string, UpdateFunc)
	Close() map[string]interface{}
}

// UpdateFunc type provides a convenient way of specifying the signature
// of an update function.
type UpdateFunc func(interface{}, bool) interface{}

// safeMap type depends on a channel which can send and receive values of
// the custom type commandData declared next
type safeMap chan commandData

type commandAction int

const (
	remove commandAction = iota
	end
	find
	insert
	length
	update
)

// commandData struct has values that specify action that is to be taken
// and also the data necessary to perform the action.
type commandData struct {
	action  commandAction
	key     string
	value   interface{}
	result  chan<- interface{}
	data    chan<- map[string]interface{}
	updater UpdateFunc
}

// findResult struct will be what returned to the caller
// when a find is invoked over safemap
type findResult struct {
	value interface{}
	found bool
}

func (sm safeMap) Insert(key string, value interface{}) {
	sm <- commandData{action: insert, key: key, value: value}
}

func (sm safeMap) Delete(key string) {
	sm <- commandData{action: remove, key: key}
}

func (sm safeMap) Find(key string) (value interface{}, found bool) {
	reply := make(chan interface{})
	sm <- commandData{action: find, key: key, result: reply}
	result := (<-reply).(findResult)
	return result.value, result.found
}

func (sm safeMap) Len() int {
	reply := make(chan interface{})
	sm <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

func (sm safeMap) Update(key string, updater UpdateFunc) {
	sm <- commandData{action: update, key: key, updater: updater}
}

func (sm safeMap) Close() map[string]interface{} {
	reply := make(chan map[string]interface{})
	sm <- commandData{action: end, data: reply}
	return <-reply
}

// New is an exported function that creates an instance of
// the safeMap and returns the same outside of the package
func New() SafeMap {
	sm := make(safeMap)
	go sm.run()
	return sm
}

func (sm safeMap) run() {
	store := make(map[string]interface{})
	for command := range sm {
		switch command.action {
		case insert:
			store[command.key] = command.value
		case remove:
			delete(store, command.key)
		case find:
			value, found := store[command.key]
			command.result <- findResult{value, found}
		case length:
			command.result <- len(store)
		case update:
			value, found := store[command.key]
			store[command.key] = command.updater(value, found)
		case end:
			close(sm)
			command.data <- store
		}
	}
}
