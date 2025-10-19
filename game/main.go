package game

import (
	"artifacts/client"
	"artifacts/state"
	"log"
	"sync"
)

const API_KEY = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbWFpbCI6InF1ZXJkeTdAZ21haWwuY29tIiwicGFzc3dvcmRfY2hhbmdlZCI6bnVsbH0.-5lNsy4G90w8flV2h-rteKSG2pyjeNuohFdbaKPRR9c"
const API_URL = "https://api.artifactsmmo.com/"

func GameLoop() {
	log.Println("Initialize")
	aMmoClient := client.NewArtifactsMMOClient(API_URL, API_KEY)

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		defer wg.Done()
		accountDetails, err := aMmoClient.GetAccountDetails()
		if err != nil {
			log.Println("Error getting account details: ", err)
			return
		}
		state.GameStateData.Account = accountDetails.Data
	}()

	go func() {
		defer wg.Done()
		serverStatus, err := aMmoClient.GetServerStatus()
		if err != nil {
			log.Println("Error getting server status: ", err)
			return
		}
		state.GameStateData.Server = serverStatus.Data
	}()

	go func() {
		defer wg.Done()
		charactersData, err := aMmoClient.GetMyCharacters()
		if err != nil {
			log.Println("Error getting my characters: ", err)
			return
		}
		state.GameStateData.Characters = charactersData.Data
	}()

	go func() {
		defer wg.Done()
		maps, err := aMmoClient.GetMaps()
		if err != nil {
			log.Println("Error getting maps: ", err)
			return
		}
		state.GameStateData.Maps = maps
	}()

	go func() {
		defer wg.Done()
		resources, err := aMmoClient.GetResources()
		if err != nil {
			log.Println("Error getting resources: ", err)
			return
		}
		state.GameStateData.Resources = resources
	}()

	wg.Wait()
	log.Println("End of initialization")
	go CharactersController(aMmoClient)
}
