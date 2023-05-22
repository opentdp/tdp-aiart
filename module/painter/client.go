package painter

type ReqeustParam struct {
	Action         string
	Subject        string
	Prompt         string
	NegativePrompt string
	InputImage     string
	Strength       float64
	Styles         []string
	ResultConfig   map[string]any
	LogoAdd        int64
	LogoParam      map[string]any
	Status         string
}

type ResponseData struct {
	InputFile  string
	OutputFile string
}

func Create(rq *ReqeustParam) (*ResponseData, error) {

	outputImage, err := TencentAiart(rq)

	if err != nil {
		return nil, err
	}

	return saveObject(rq, outputImage)

}
