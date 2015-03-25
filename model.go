package inforce

import (
	"appengine"
	"appengine/datastore"
)

type Counter struct {
	Count int64
}

// Gets a key for a specified counter by name
func counterKey(c appengine.Context, counterName string) *datastore.Key {
	return datastore.NewKey(c, "Counter", counterName, 0, nil)
}

// Gets the number of entities
func NumberOf(c appengine.Context, counterName string) (int64, error) {
	k := counterKey(c, counterName)
	var counter Counter
	if err := datastore.Get(c, k, &counter); err != nil {
		return 0, err
	}
	return counter.Count, nil
}

func updateCounter(c appengine.Context, counterName string) error {
	key := counterKey(c, counterName)
	count := new(Counter)
	err := datastore.RunInTransaction(c, func(c appengine.Context) error {
		err := datastore.Get(c, key, count)
		if err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}
		count.Count++
		_, err = datastore.Put(c, key, count)
		return err
	}, nil)

	if err != nil {
		c.Errorf("Transaction failed: %v", err)
	}
	return err
}
