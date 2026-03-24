package integrations

import "encoding/json"

// Image overrides for assets used in the inbox UI.
type Images struct {
	// URL for the illustration shown when the inbox is empty.
	EmptyInboxUrl *string `json:"emptyInboxUrl,omitempty" required:"true"`
}

func (i *Images) GetEmptyInboxUrl() *string {
	if i == nil {
		return nil
	}
	return i.EmptyInboxUrl
}

func (i *Images) SetEmptyInboxUrl(emptyInboxUrl string) {
	i.EmptyInboxUrl = &emptyInboxUrl
}

func (i Images) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: Images to string"
	}
	return string(jsonData)
}
