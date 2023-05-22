package chatbot

type ReqeustParam struct {
	Model    string
	Messages []ChatMessage
}

type ChatResponse struct {
	Messages     []ChatMessage
	FinishReason string
}

type ChatMessage struct {
	Index   int
	Role    string
	Content string
}

func Create(rq *ReqeustParam) (*ChatResponse, error) {

	return OpenaiChat(rq)

}

func CreateStream(rq *ReqeustParam) (chan ChatResponse, error) {

	return OpenaiChatStream(rq)

}
