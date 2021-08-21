package sproc

import "log"

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

func NewStream(reader Reader, store Store) Streamer {
	return &stream{reader: reader, store: store}
}

type stream struct {
	reader    Reader
	store     Store
	transform func(val interface{}) (k interface{}, v interface{})
	reduced   []interface{}
}

func (s *stream) Transform(call func(val interface{}) (k interface{}, v interface{})) Streamer {
	s.transform = call
	return s
}

func (s *stream) ToStore() Streamer {
	s.reader.Read(func(v interface{}) {
		k, v := s.transform(v)
		err := s.store.Set(k, v)
		if err != nil {
			log.Fatal(err)
		}
	})
	return s
}

func (s *stream) Sort(sorter func([]interface{}) []interface{}) Streamer {
	if s.reduced == nil {
		aa, _ := s.store.GetAll()
		s.reduced = aa
	}
	s.reduced = sorter(s.reduced)
	return s
}

func (s *stream) Limit(limit int) Streamer {
	if s.reduced == nil {
		aa, _ := s.store.GetAll()
		s.reduced = aa
	}
	if limit > len(s.reduced) {
		limit = len(s.reduced)
	}
	s.reduced = s.reduced[0:limit]
	return s
}

func (s *stream) Print(print func(val interface{})) Streamer {
	if s.reduced == nil {
		aa, _ := s.store.GetAll()
		s.reduced = aa
	}
	for _, a := range s.reduced {
		print(a)
	}
	return s
}
