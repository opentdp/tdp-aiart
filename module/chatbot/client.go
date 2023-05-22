package chatbot

type ReqeustParam struct {
	Model    string
	Messages []ChatMessage
}

type ChatResponse struct {
	Messages []ChatMessage
}

type ChatMessage struct {
	Role    string
	Content string
}

func Create(rq *ReqeustParam) (*ChatResponse, error) {

	return OpenaiChat(rq)

}
