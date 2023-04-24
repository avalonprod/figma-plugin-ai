package openai

type OpenaiService struct {
	TextGenerator TextGenerator
	ImgGenerator  ImgGenerator
}

type TextGenerator interface {
	NewMessages(message string) (string, error)
}

type ImgGenerator interface {
	NewImg(prompt string) ([]imgItem, error)
}

func NewOpenaiService(apiKey string) *OpenaiService {
	return &OpenaiService{
		TextGenerator: NewTextGenerator(apiKey, "https://api.openai.com/v1/chat/completions"),
		ImgGenerator:  NewImgGeneration(apiKey, "https://api.openai.com/v1/images/generations"),
	}
}
