package strategies

import (
	"sort"
)

// Factory constructs a new Strategy instance
type Factory func() Strategy

// registry holds all registered strategies
var registry []Factory

// Register adds a strategy factory to the registry. Call from init() in each strategy file.
func Register(factory Factory) {
	if factory == nil {
		return
	}
	registry = append(registry, factory)
}

// DiscoverFactories returns all registered factories sorted by strategy Name().
func DiscoverFactories() []Factory {
	// Build name/factory pairs to sort by name
	type pair struct {
		name    string
		factory Factory
	}
	pairs := make([]pair, 0, len(registry))
	for _, f := range registry {
		inst := f()
		pairs = append(pairs, pair{name: inst.Name(), factory: f})
	}

	sort.Slice(pairs, func(i, j int) bool { return pairs[i].name < pairs[j].name })

	result := make([]Factory, len(pairs))
	for i, p := range pairs {
		result[i] = p.factory
	}
	return result
}

// Discover returns fresh instances of all registered strategies, sorted by Name().
func Discover() []Strategy {
	factories := DiscoverFactories()
	instances := make([]Strategy, 0, len(factories))
	for _, f := range factories {
		instances = append(instances, f())
	}
	return instances
}

// NewByIndex constructs a fresh Strategy instance by sorted index.
func NewByIndex(idx int) Strategy {
	factories := DiscoverFactories()
	if idx < 0 || idx >= len(factories) {
		return nil
	}
	return factories[idx]()
}

// Count returns the number of registered strategies.
func Count() int { return len(registry) }
