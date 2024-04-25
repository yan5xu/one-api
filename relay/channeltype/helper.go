package channeltype

import "github.com/songquanpeng/one-api/relay/apitype"

func ToAPIType(channelType int) int {
	apiType := apitype.OpenAI
	switch channelType {
	case Anthropic:
		apiType = apitype.Anthropic
	case Baidu:
		apiType = apitype.Baidu
	case PaLM:
		apiType = apitype.PaLM
	case Zhipu:
		apiType = apitype.Zhipu
	case Ali:
		apiType = apitype.Ali
	case Xunfei:
		apiType = apitype.Xunfei
	case AIProxyLibrary:
		apiType = apitype.AIProxyLibrary
	case Tencent:
		apiType = apitype.Tencent
	case Gemini:
		apiType = apitype.Gemini
	case Ollama:
		apiType = apitype.Ollama
	case AwsClaude:
		apiType = apitype.AwsClaude
	case Coze:
		apiType = apitype.Coze
	case Cohere:
		apiType = apitype.Cohere
	case Vertex:
		apiType = apitype.Vertex
	}

	return apiType
}
