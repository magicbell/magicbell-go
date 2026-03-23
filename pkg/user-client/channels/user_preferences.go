package channels

import "encoding/json"

type UserPreferences struct {
	Categories []Categories `json:"categories,omitempty"`
}

func (u *UserPreferences) GetCategories() []Categories {
	if u == nil {
		return nil
	}
	return u.Categories
}

func (u *UserPreferences) SetCategories(categories []Categories) {
	u.Categories = categories
}

func (u UserPreferences) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserPreferences to string"
	}
	return string(jsonData)
}
