package integrations

import (
	"encoding/json"
	"github.com/magicbell/magicbell-go/pkg/user-client/internal/unmarshal"
	"github.com/magicbell/magicbell-go/pkg/user-client/util"
)

type InboxConfigPayload struct {
	// Image overrides for assets used in the inbox UI.
	Images *util.Nullable[Images] `json:"images,omitempty" required:"true"`
	// Locale code (ISO language tag) used to localize built-in strings.
	Locale *util.Nullable[string] `json:"locale,omitempty" required:"true" minLength:"2"`
	// Visual customization options for the hosted inbox widget.
	Theme *util.Nullable[Theme] `json:"theme,omitempty" required:"true"`
}

func (i *InboxConfigPayload) GetImages() *util.Nullable[Images] {
	if i == nil {
		return nil
	}
	return i.Images
}

func (i *InboxConfigPayload) SetImages(images util.Nullable[Images]) {
	i.Images = &images
}

func (i *InboxConfigPayload) SetImagesNull() {
	i.Images = &util.Nullable[Images]{IsNull: true}
}

func (i *InboxConfigPayload) GetLocale() *util.Nullable[string] {
	if i == nil {
		return nil
	}
	return i.Locale
}

func (i *InboxConfigPayload) SetLocale(locale util.Nullable[string]) {
	i.Locale = &locale
}

func (i *InboxConfigPayload) SetLocaleNull() {
	i.Locale = &util.Nullable[string]{IsNull: true}
}

func (i *InboxConfigPayload) GetTheme() *util.Nullable[Theme] {
	if i == nil {
		return nil
	}
	return i.Theme
}

func (i *InboxConfigPayload) SetTheme(theme util.Nullable[Theme]) {
	i.Theme = &theme
}

func (i *InboxConfigPayload) SetThemeNull() {
	i.Theme = &util.Nullable[Theme]{IsNull: true}
}

func (i InboxConfigPayload) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InboxConfigPayload to string"
	}
	return string(jsonData)
}

func (i *InboxConfigPayload) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, i)
}
