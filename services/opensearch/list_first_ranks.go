package opensearch

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/hdys0vn/alibaba-cloud-sdk-go/sdk/responses"
)

// ListFirstRanks invokes the opensearch.ListFirstRanks API synchronously
func (client *Client) ListFirstRanks(request *ListFirstRanksRequest) (response *ListFirstRanksResponse, err error) {
	response = CreateListFirstRanksResponse()
	err = client.DoAction(request, response)
	return
}

// ListFirstRanksWithChan invokes the opensearch.ListFirstRanks API asynchronously
func (client *Client) ListFirstRanksWithChan(request *ListFirstRanksRequest) (<-chan *ListFirstRanksResponse, <-chan error) {
	responseChan := make(chan *ListFirstRanksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFirstRanks(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListFirstRanksWithCallback invokes the opensearch.ListFirstRanks API asynchronously
func (client *Client) ListFirstRanksWithCallback(request *ListFirstRanksRequest, callback func(response *ListFirstRanksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFirstRanksResponse
		var err error
		defer close(result)
		response, err = client.ListFirstRanks(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListFirstRanksRequest is the request struct for api ListFirstRanks
type ListFirstRanksRequest struct {
	*requests.RoaRequest
	AppId            requests.Integer `position:"Path" name:"appId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// ListFirstRanksResponse is the response struct for api ListFirstRanks
type ListFirstRanksResponse struct {
	*responses.BaseResponse
	RequestId string          `json:"requestId" xml:"requestId"`
	Result    []FirstRankItem `json:"result" xml:"result"`
}

// CreateListFirstRanksRequest creates a request to invoke ListFirstRanks API
func CreateListFirstRanksRequest() (request *ListFirstRanksRequest) {
	request = &ListFirstRanksRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListFirstRanks", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/first-ranks", "", "")
	request.Method = requests.GET
	return
}

// CreateListFirstRanksResponse creates a response to parse from ListFirstRanks response
func CreateListFirstRanksResponse() (response *ListFirstRanksResponse) {
	response = &ListFirstRanksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
