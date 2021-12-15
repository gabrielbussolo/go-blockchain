package presenter

type jsonValidChain struct {
	Message string `json:"message"`
}

func ValidChain(message string) jsonValidChain {
	return jsonValidChain{Message: message}
}
