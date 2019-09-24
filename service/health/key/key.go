package key

var Providers = [...]string{"aws", "azure", "kvm"}

func IsKnownProvider(value string) bool {
	for i := range Providers {
		if Providers[i] == value {
			return true
		}
	}
	return false
}
