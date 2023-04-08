package artman

type ReqeustParams struct {
	Action  string
	Payload any
}

type ImagePayload struct {
	TextPayload
	InputImage string
	InputUrl   string
}

type TextPayload struct {
	Prompt         string
	NegativePrompt string
	Styles         []string
	ResultConfig   any
	LogoAdd        int64
	Strength       float64
}
