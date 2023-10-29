package service

type (
	ChannelMessage struct {
		ID             string
		Author         string `json:"author"`
		Content        string `json:"content"`
		CreatedAtMonth string `json:"created_at_month"`
		CreatedAtYear  string `json:"created_at_year"`

		SaveApproved bool `json:"approved"`
	}

	ChannelMessageResponse struct {
		Content string `json:"content"`
	}

	PhraseForEternity struct {
		Phrase string `json:"phrase"`
		Author string `json:"author"`
		Date   string `json:"date"`
	}
)
