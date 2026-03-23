package integrations

import "encoding/json"

type Enterprise struct {
	// Enterprise grid identifier.
	Id *string `json:"id,omitempty" required:"true"`
	// Enterprise grid name.
	Name *string `json:"name,omitempty" required:"true"`
}

func (e *Enterprise) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *Enterprise) SetId(id string) {
	e.Id = &id
}

func (e *Enterprise) GetName() *string {
	if e == nil {
		return nil
	}
	return e.Name
}

func (e *Enterprise) SetName(name string) {
	e.Name = &name
}

func (e Enterprise) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: Enterprise to string"
	}
	return string(jsonData)
}
