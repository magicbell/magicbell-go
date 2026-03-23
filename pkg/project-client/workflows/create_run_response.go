package workflows

import "encoding/json"

type CreateRunResponse struct {
	// Identifier of the workflow run that was created.
	Id *string `json:"id,omitempty" required:"true"`
}

func (c *CreateRunResponse) GetId() *string {
	if c == nil {
		return nil
	}
	return c.Id
}

func (c *CreateRunResponse) SetId(id string) {
	c.Id = &id
}

func (c CreateRunResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateRunResponse to string"
	}
	return string(jsonData)
}
