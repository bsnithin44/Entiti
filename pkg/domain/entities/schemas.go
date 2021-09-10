package entities

import "time"

// EntityType
type EntityTypeSystem struct {
	Id        string    `json:"id" cql:"id"`
	CreatedAt time.Time `json:"createdAt" cql:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" cql:"updated_at"`
}
type EntityTypeReadOnly struct {
	PluralName   *string `json:"pluralName" cql:"plura_name"`
	SingularName *string `json:"singularName" cql:"singular_name"`
	UniqueCode   *string `json:"unique_code" cql:"unique_code"`
}

type EntityTypeBase struct {
	IsActive *bool `json:"isActive" cql:"is_active"`
}
type CreateEntityTypeStruct struct {
	EntityTypeReadOnly
	EntityTypeBase
}
type EntityTypeInDb struct {
	EntityTypeSystem
	EntityTypeReadOnly
	EntityTypeBase
}
