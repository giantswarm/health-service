package key

// Providers is an array of providers which are currently implemented for health checking
var Providers = [...]string{"aws", "azure", "kvm"}

// IsKnownProvider is a helper function returning true iff the given value is in Providers
func IsKnownProvider(value string) bool {
	for i := range Providers {
		if Providers[i] == value {
			return true
		}
	}
	return false
}
