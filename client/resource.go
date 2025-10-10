package client

import (
	"artifacts/state"
	"artifacts/utils"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

func (ac *ArtifactsMMOClient) GetResources() ([]state.Resource, error) {
	firstPage, err := ac.getResourcesPage(1, 100)
	if err != nil {
		return nil, err
	}

	if firstPage.Pages <= 1 {
		return firstPage.Data, nil
	}

	allResources := make([]state.Resource, 0, firstPage.Total)
	allResources = append(allResources, firstPage.Data...)

	var mu sync.Mutex
	var wg sync.WaitGroup
	errChan := make(chan error, firstPage.Pages-1)

	for page := 2; page <= firstPage.Pages; page++ {
		wg.Add(1)
		go func(pageNum int) {
			defer wg.Done()

			pageData, err := ac.getResourcesPage(pageNum, 100)
			if err != nil {
				errChan <- fmt.Errorf("page %d: %w", pageNum, err)
				return
			}

			mu.Lock()
			allResources = append(allResources, pageData.Data...)
			mu.Unlock()
		}(page)
	}

	wg.Wait()
	close(errChan)

	var errors []string
	for err := range errChan {
		errors = append(errors, err.Error())
	}

	if len(errors) > 0 {
		return allResources, fmt.Errorf("errors fetching resources: %s", strings.Join(errors, "; "))
	}

	return allResources, nil
}

func (ac *ArtifactsMMOClient) getResourcesPage(page, size int) (*ResourceData, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 100
	}

	url := fmt.Sprintf("%s/resources?page=%d&size=%d", ac.baseUrl, page, size)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	response, err := ac.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("execute request: %w", err)
	}
	defer ac.CloseBody(response)
	return utils.DecodeJSONBody[ResourceData](response)
}
