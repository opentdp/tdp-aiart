package chatbot

type ChatModel struct {
	Name      string
	CreatedAt int64
}

type ReqeustParam struct {
	Model    string
	Messages []ChatMessage
}

type ChatResponse struct {
	Index   int
	Reason  string
	Message ChatMessage
}

type ChatMessage struct {
	Role    string
	Content string
}

func Models() ([]*ChatModel, error) {

	return OpenaiModels()

}

func Create(rq *ReqeustParam) (*ChatResponse, error) {

	return OpenaiChat(rq)

}

func CreateStream(rq *ReqeustParam) (chan ChatResponse, error) {

	return OpenaiChatStream(rq)

}
