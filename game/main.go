package game

import (
	"artifacts/client"
	"artifacts/state"
	"fmt"
)

const API_KEY = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbWFpbCI6InF1ZXJkeTdAZ21haWwuY29tIiwicGFzc3dvcmRfY2hhbmdlZCI6bnVsbH0.-5lNsy4G90w8flV2h-rteKSG2pyjeNuohFdbaKPRR9c"
const API_URL = "https://api.artifactsmmo.com/"

func GameLoop() {
	aMmoClient := client.NewArtifactsMMOClient(API_URL, API_KEY)
	accountDetails, err := aMmoClient.GetAccountDetails()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	state.GameStateData.Account = accountDetails.Data
	//fmt.Println(utils.Stringify(accountDetails))
	//printAccountStatus(accountDetails)
	serverStatus, err := aMmoClient.GetServerStatus()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	state.GameStateData.Server = serverStatus.Data
	//fmt.Println(utils.Stringify(serverStatus))
	//fmt.Println(utils.Stringify(serverStatus))

	charactersData, err := aMmoClient.GetMyCharacters()
	if err != nil {
		fmt.Println("Error: ", err)
	}
	state.GameStateData.Characters = charactersData.Data
	//fmt.Println(utils.Stringify(charactersData))

	maps, err := aMmoClient.GetMaps()
	state.GameStateData.Maps = maps
	//fmt.Println(utils.Stringify(maps))

	resources, err := aMmoClient.GetResources()
	state.GameStateData.Resources = resources
	//fmt.Println(utils.Stringify(resources))
	go CharactersController(aMmoClient)
}
