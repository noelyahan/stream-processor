package sproc

// Streamer is the interface to wrap high level data i/o streaming functionality
type Streamer interface {
	// Transform is a Map like function to mutate, add custom logic down to the stream
	Transform(func(val interface{}) (k interface{}, v interface{})) Streamer

	// ToStore is a Reducer like function to aggregate all final state to the state store, also final call function
	ToStore() Streamer

	// Sort is a stream operation to sort final states
	Sort(sorter func([]interface{}) []interface{}) Streamer

	// Limit is a stream operation to limit final state results
	Limit(limit int) Streamer

	// Print is a stream operation to print final state results
	Print(print func(val interface{})) Streamer
}
