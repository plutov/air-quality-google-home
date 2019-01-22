package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handle)
}

var (
	userMsg          = "The air quality index in your city is %d right now. %s"
	errMsg           = "Sorry, I was unable to get air quality index in your place."
	actionNotFound   = "Sorry, I am lost, please try next time."
	locationNotFound = "Sorry, I was not able to find your location. Are you from outer space?"
)

func handle(w http.ResponseWriter, r *http.Request) {
	dfReq := DialogFlowRequest{}
	dfErr := json.NewDecoder(r.Body).Decode(&dfReq)

	if dfErr == nil && dfReq.QueryResult.Action == "location" {
		handleLocationPermissionAction(w, r, dfReq)
		return
	}

	if dfErr == nil && dfReq.QueryResult.Action == "get" {
		handleGetAction(w, r, dfReq)
		return
	}

	returnAPIErrorMessage(w, actionNotFound)
}

func handleLocationPermissionAction(w http.ResponseWriter, r *http.Request, dfReq DialogFlowRequest) {
	json.NewEncoder(w).Encode(DialogFlowResponse{
		Data: DialogFlowResponseData{
			Google: DialogFlowResponseGoogle{
				ExpectUserResponse: true,
				RichResponse: DialogFlowRichResponse{
					Items: []DialogFlowItem{
						DialogFlowItem{
							SimpleResponse: DialogFlowSimpleResponse{
								TextToSpeech: "PLACEHOLDER_FOR_PERMISSION",
							},
						},
					},
				},
				SystemIntent: &DialogFlowResponseSystemIntent{
					Intent: "actions.intent.PERMISSION",
					Data: DialogFlowResponseSystemIntentData{
						Type:        "type.googleapis.com/google.actions.v2.PermissionValueSpec",
						OptContext:  "To get city for air quality check",
						Permissions: []string{"DEVICE_PRECISE_LOCATION"},
					},
				},
			},
		},
	})
}

func handleGetAction(w http.ResponseWriter, r *http.Request, dfReq DialogFlowRequest) {
	lat := dfReq.OriginalRequest.Data.Device.Location.Coordinates.Lat
	long := dfReq.OriginalRequest.Data.Device.Location.Coordinates.Long
	if lat == 0 || long == 0 {
		returnAPIErrorMessage(w, locationNotFound)
		return
	}

	index, levelDescription, aqiErr := getAirQualityByCoordinates(r, lat, long)
	if aqiErr != nil {
		returnAPIErrorMessage(w, errMsg)
		return
	}

	json.NewEncoder(w).Encode(DialogFlowResponse{
		Data: DialogFlowResponseData{
			Google: DialogFlowResponseGoogle{
				ExpectUserResponse: true,
				RichResponse: DialogFlowRichResponse{
					Items: []DialogFlowItem{
						DialogFlowItem{
							SimpleResponse: DialogFlowSimpleResponse{
								TextToSpeech: fmt.Sprintf(userMsg, index, levelDescription),
							},
						},
					},
				},
			},
		},
	})
}

func returnAPIErrorMessage(w http.ResponseWriter, msg string) {
	json.NewEncoder(w).Encode(DialogFlowResponse{
		Data: DialogFlowResponseData{
			Google: DialogFlowResponseGoogle{
				ExpectUserResponse: true,
				RichResponse: DialogFlowRichResponse{
					Items: []DialogFlowItem{
						DialogFlowItem{
							SimpleResponse: DialogFlowSimpleResponse{
								TextToSpeech: msg,
							},
						},
					},
				},
			},
		},
	})
}
