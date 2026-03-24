package integrations

import "encoding/json"

type SmtpConfigObject struct {
	Config *SmtpConfig `json:"config,omitempty" required:"true"`
	Id     *string     `json:"id,omitempty" required:"true"`
	Name   *string     `json:"name,omitempty" required:"true"`
}

func (s *SmtpConfigObject) GetConfig() *SmtpConfig {
	if s == nil {
		return nil
	}
	return s.Config
}

func (s *SmtpConfigObject) SetConfig(config SmtpConfig) {
	s.Config = &config
}

func (s *SmtpConfigObject) GetId() *string {
	if s == nil {
		return nil
	}
	return s.Id
}

func (s *SmtpConfigObject) SetId(id string) {
	s.Id = &id
}

func (s *SmtpConfigObject) GetName() *string {
	if s == nil {
		return nil
	}
	return s.Name
}

func (s *SmtpConfigObject) SetName(name string) {
	s.Name = &name
}

func (s SmtpConfigObject) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SmtpConfigObject to string"
	}
	return string(jsonData)
}
