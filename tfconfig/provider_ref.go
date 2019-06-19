package tfconfig

// ProviderRef is a reference to a provider configuration within a module.
// It represents the contents of a "provider" argument in a resource, or
// a value in the "providers" map for a module call.
type ProviderRef struct {
	Name  string `json:"name"`
	Alias string `json:"alias,omitempty"` // Empty if the default provider configuration is referenced
}

type providersSortedByName []*ProviderRef

func (a providersSortedByName) Len() int {
	return len(a)
}

func (a providersSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a providersSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name || (a[i].Name == a[j].Name && a[i].Alias < a[j].Alias)
}