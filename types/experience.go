package types

type ExperienceType string

const (
	ExperienceTypeCustom   ExperienceType = "custom"
	ExperienceTypeDiscord  ExperienceType = "discord"
	ExperienceTypeFile     ExperienceType = "file"
	ExperienceTypeLink     ExperienceType = "link"
	ExperienceTypeContent  ExperienceType = "content"
	ExperienceTypeSoftware ExperienceType = "software"
)

type ExperienceDiscordCancelActionType string

const (
	ExperienceDiscordCancelActionTypeNoAction       ExperienceDiscordCancelActionType = "noaction"
	ExperienceDiscordCancelActionTypeRemoveRole     ExperienceDiscordCancelActionType = "removerole"
	ExperienceDiscordCancelActionTypeKickUser       ExperienceDiscordCancelActionType = "kickuser"
	ExperienceDiscordCancelActionTypeRemoveAllRoles ExperienceDiscordCancelActionType = "removeallroles"
)

type custom any

type discord struct {
	RoleID       string                            `json:"role_id"`
	CancelAction ExperienceDiscordCancelActionType `json:"cancel_action"`
}

type file any

type link struct {
	URL   string `json:"url"`
	Oauth bool   `json:"oauth"`
}

type content struct {
	DocumentIds []string `json:"document_ids"`
}

type software struct {
	Links []struct {
		URL      string  `json:"url"`
		Version  float64 `json:"version"`
		Platform string  `json:"platform"`
	} `json:"links"`
	Category         string   `json:"category"`
	SupportedSites   []string `json:"supported_sites"`
	SupportedRegions []string `json:"supported_regions"`
}

type properties struct {
	Custom   *custom
	FIle     *file
	Discord  *discord
	Link     *link
	Content  *content
	Software *software
}

type Experience struct {
	ID             string         `json:"id"`
	ExperienceType ExperienceType `json:"experience_type"`
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Properties     properties     `json:"properties"`
	Products       []string       `json:"products"`
	AccessPasses   []string       `json:"access_passes"`
}
