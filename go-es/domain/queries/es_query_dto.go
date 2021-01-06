package queries

type EsQuery struct {
	Equals    []FieldValue `json:"equals"`
	NotEquals []FieldValue `json:"not_equals"`
	Like      []FieldValue `json:"like"`
	NotLike   []FieldValue `json:"not_like"`
}

type FieldValue struct {
	Field string      `json:"field"`
	Value interface{} `json:"value"`
}
