package models

type Request struct {
	Query string `json:"query"`
}

type Response struct {
	Data struct {
		Get struct {
			QuestionTest []struct {
				Additional struct {
					Generate struct {
						Error        *string `json:"error"`
						SingleResult string  `json:"singleResult"`
					} `json:"generate"`
				} `json:"_additional"`
				Answer   string `json:"answer"`
				Category string `json:"category"`
				Question string `json:"question"`
			} `json:"QuestionTest"`
		} `json:"Get"`
	} `json:"data"`
}

type RAGRequest struct {
	Prompt string `json:"prompt"`
}

type DataRequest struct {
	Class      string                 `json:"class"`
	Properties map[string]interface{} `json:"properties"`
}
