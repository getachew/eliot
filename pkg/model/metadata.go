package model

var (
	nameKey      = "name"
	namespaceKey = "namespace"
)

// Metadata is metadata that all resources must have
type Metadata struct {
	Name      string `validate:"required,gt=0,alphanumOrDash"  json:"name"          yaml:"name"`
	Namespace string `validate:"omitempty,gt=0,alphanumOrDash"     json:"namespace"     yaml:"namespace"`
}

// NewMetadata creates new metadata with name and metadata fields
func NewMetadata(name, namespace string) Metadata {
	return Metadata{
		Name:      name,
		Namespace: namespace,
	}
}
