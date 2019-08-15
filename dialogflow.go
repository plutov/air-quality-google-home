package app

// DialogFlow V2

// DialogFlowRequest struct
type DialogFlowRequest struct {
	QueryResult struct {
		Action string `json:"action"`
	} `json:"queryResult"`
	OriginalRequest DialogFlowOriginalRequest `json:"originalDetectIntentRequest"`
}

// DialogFlowOriginalRequest struct
type DialogFlowOriginalRequest struct {
	Data DialogFlowOriginalRequestData `json:"payload"`
}

// DialogFlowOriginalRequestData struct
type DialogFlowOriginalRequestData struct {
	Device DialogFlowOriginalRequestDevice `json:"device"`
	User   DialogFlowOriginalRequestUser   `json:"user"`
}

// DialogFlowOriginalRequestDevice struct
type DialogFlowOriginalRequestDevice struct {
	Location DialogFlowOriginalRequestLocation `json:"location"`
}

// DialogFlowOriginalRequestUser struct
type DialogFlowOriginalRequestUser struct {
	Permissions []string `json:"permissions"`
}

// DialogFlowOriginalRequestLocation struct
type DialogFlowOriginalRequestLocation struct {
	Coordinates DialogFlowOriginalRequestCoordinates `json:"coordinates"`
}

// DialogFlowOriginalRequestCoordinates struct
type DialogFlowOriginalRequestCoordinates struct {
	Lat  float32 `json:"latitude"`
	Long float32 `json:"longitude"`
}

// DialogFlowResponse struct
type DialogFlowResponse struct {
	Data DialogFlowResponseData `json:"payload"`
}

// DialogFlowResponseData struct
type DialogFlowResponseData struct {
	Google DialogFlowResponseGoogle `json:"google"`
}

// DialogFlowResponseGoogle struct
type DialogFlowResponseGoogle struct {
	ExpectUserResponse bool                            `json:"expectUserResponse"`
	RichResponse       DialogFlowRichResponse          `json:"richResponse"`
	SystemIntent       *DialogFlowResponseSystemIntent `json:"systemIntent,omitempty"`
}

// DialogFlowRichResponse struct
type DialogFlowRichResponse struct {
	Items []DialogFlowItem `json:"items"`
}

// DialogFlowItem struct
type DialogFlowItem struct {
	SimpleResponse DialogFlowSimpleResponse `json:"simpleResponse"`
}

// DialogFlowSimpleResponse struct
type DialogFlowSimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
}

// DialogFlowResponseSystemIntent struct
type DialogFlowResponseSystemIntent struct {
	Intent string                             `json:"intent"`
	Data   DialogFlowResponseSystemIntentData `json:"data"`
}

// DialogFlowResponseSystemIntentData struct
type DialogFlowResponseSystemIntentData struct {
	Type        string   `json:"@type"`
	OptContext  string   `json:"optContext"`
	Permissions []string `json:"permissions"`
}
