package chapter

var registry []*Chapter

// Register adds a chapter to the global registry.
// Called via init() in each chapter's package during startup.
func Register(c *Chapter) {
	registry = append(registry, c)
}

// All returns all registered chapters.
func All() []*Chapter {
	return registry
}

// Get returns the chapter with the given number, or nil if not found.
func Get(number int) *Chapter {
	for _, c := range registry {
		if c.Number == number {
			return c
		}
	}
	return nil
}
