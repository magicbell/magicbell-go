package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/project-client/shared"
)

type SmtpConfigObjectCollection struct {
	Data  []SmtpConfigObject `json:"data,omitempty"`
	Links *shared.Links      `json:"links,omitempty"`
}

func (s *SmtpConfigObjectCollection) GetData() []SmtpConfigObject {
	if s == nil {
		return nil
	}
	return s.Data
}

func (s *SmtpConfigObjectCollection) SetData(data []SmtpConfigObject) {
	s.Data = data
}

func (s *SmtpConfigObjectCollection) GetLinks() *shared.Links {
	if s == nil {
		return nil
	}
	return s.Links
}

func (s *SmtpConfigObjectCollection) SetLinks(links shared.Links) {
	s.Links = &links
}

func (s SmtpConfigObjectCollection) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SmtpConfigObjectCollection to string"
	}
	return string(jsonData)
}
