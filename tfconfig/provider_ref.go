package tfconfig

// ProviderRef is a reference to a provider configuration within a module.
// It represents the contents of a "provider" argument in a resource, or
// a value in the "providers" map for a module call.
type ProviderRef struct {
	Name  string `json:"name"`
	Alias string `json:"alias,omitempty"` // Empty if the default provider configuration is referenced
}

type providerRefs []*ProviderRef

func (a providerRefs) Len() int {
	return len(a)
}

func (a providerRefs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a providerRefs) Less(i, j int) bool {
	return a[i].Name < a[j].Name || (a[i].Name == a[j].Name && a[i].Alias < a[j].Alias)
}

func (a providerRefs) contains(provider ProviderRef) bool {
	for _, existing := range a {
		if *existing == provider {
			return true
		}
	}
	return false
}