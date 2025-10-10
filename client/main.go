package client

import (
	"artifacts/state"
	"artifacts/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type ArtifactsMMOClient struct {
	client   *http.Client
	baseUrl  string
	apiToken string
}

func NewArtifactsMMOClient(baseUrl string, apiToken string) *ArtifactsMMOClient {
	transport := &authTransport{
		apiToken: apiToken,
		base:     http.DefaultTransport,
	}
	return &ArtifactsMMOClient{
		client: &http.Client{
			Transport: transport,
		},
		baseUrl:  baseUrl,
		apiToken: apiToken,
	}
}

func (ac *ArtifactsMMOClient) GetAccountDetails() (*AccountDetails, error) {
	response, err := ac.client.Do(ac.CreateRequest("GET", "/my/details", nil))
	if err != nil {
		return nil, err
	}
	defer ac.CloseBody(response)
	return utils.DecodeJSONBody[AccountDetails](response)
}

func (ac *ArtifactsMMOClient) GetServerStatus() (*ServerStatusData, error) {
	response, err := ac.client.Do(ac.CreateRequest("GET", "", nil))
	if err != nil {
		return nil, err
	}
	defer ac.CloseBody(response)
	return utils.DecodeJSONBody[ServerStatusData](response)
}

func (ac *ArtifactsMMOClient) GetMyCharacters() (*CharactersData, error) {
	response, err := ac.client.Do(ac.CreateRequest("GET", "my/characters", nil))
	if err != nil {
		return nil, err
	}
	defer ac.CloseBody(response)
	return utils.DecodeJSONBody[CharactersData](response)
}

func (ac *ArtifactsMMOClient) ActionMove(character *state.Character, toMap *state.Map) (*CharacterMovementData, error) {
	data := map[string]interface{}{
		"map_id": toMap.MapID,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	response, err := ac.client.Do(ac.CreateRequest(
		"POST",
		"my/"+character.Name+"/action/move",
		bytes.NewBuffer(jsonData)),
	)
	if err != nil {
		return nil, err
	}
	defer ac.CloseBody(response)
	return utils.DecodeJSONBody[CharacterMovementData](response)
}

func (ac *ArtifactsMMOClient) ActionGathering(character *state.Character) (*GatheringData, error) {
	response, err := ac.client.Do(
		ac.CreateRequest(
			"POST",
			"my/"+character.Name+"/action/gathering",
			nil,
		),
	)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf(response.Status)
	}
	defer ac.CloseBody(response)
	return utils.DecodeJSONBody[GatheringData](response)
}

func (ac *ArtifactsMMOClient) ActionDepositItem(character *state.Character, items []state.Item) (*DepositItemData, error) {
	jsonData, err := json.Marshal(items)
	if err != nil {
		return nil, err
	}
	jsonPretty, err := json.MarshalIndent(items, "", "  ")
	fmt.Println(string(jsonPretty))
	response, err := ac.client.Do(
		ac.CreateRequest(
			"POST",
			"my/"+character.Name+"/action/bank/deposit/item",
			bytes.NewBuffer(jsonData),
		),
	)
	if err != nil {
		return nil, err
	}
	defer ac.CloseBody(response)
	return utils.DecodeJSONBody[DepositItemData](response)
}
