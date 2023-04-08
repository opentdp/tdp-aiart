package artman

type ReqeustParams struct {
	Action  string
	Payload any
}

type ImagePayload struct {
	Prompt         string
	NegativePrompt string
	InputImage     string
	Styles         []string
	Strength       float64
	ResultConfig   any
	LogoAdd        int64
}

type TextPayload struct {
	Prompt         string
	NegativePrompt string
	Styles         []string
	ResultConfig   any
	LogoAdd        int64
}
