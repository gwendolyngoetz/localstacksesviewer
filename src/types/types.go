package types

type SesResponse struct {
	Messages []Messages `json:"messages,omitempty"`
}

type Destination struct {
	ToAddresses []string `json:"ToAddresses,omitempty"`
}

type Messages struct {
	ID          string      `json:"Id,omitempty"`
	Region      string      `json:"Region,omitempty"`
	Destination Destination `json:"Destination,omitempty"`
	Source      string      `json:"Source,omitempty"`
	Subject     string      `json:"Subject,omitempty"`
	Body        Body        `json:"Body,omitempty"`
	Timestamp   string      `json:"Timestamp,omitempty"`
}
type Body struct {
	TextPart string `json:"text_part,omitempty"`
	HTMLPart any    `json:"html_part,omitempty"`
}
