package model

type Response struct {

	// message
	Message string `json:"message,omitempty"`

	// messages
	Messages any `json:"messages,omitempty"`
}

type ResponseWithID struct {

	// Id
	ID string `json:"Id,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// messages
	Messages any `json:"messages,omitempty"`
}
