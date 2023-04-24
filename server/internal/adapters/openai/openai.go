package openai

type OpenaiService struct {
	TextGenerator TextGenerator
}

type TextGenerator interface {
	NewMessages(message string) (string, error)
}

func NewOpenaiService(apiKey string) *OpenaiService {
	return &OpenaiService{
		TextGenerator: NewTextGenerator(apiKey, "https://api.openai.com/v1/chat/completions"),
	}
}
