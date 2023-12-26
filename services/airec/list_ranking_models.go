package airec

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

// ListRankingModels invokes the airec.ListRankingModels API synchronously
func (client *Client) ListRankingModels(request *ListRankingModelsRequest) (response *ListRankingModelsResponse, err error) {
	response = CreateListRankingModelsResponse()
	err = client.DoAction(request, response)
	return
}

// ListRankingModelsWithChan invokes the airec.ListRankingModels API asynchronously
func (client *Client) ListRankingModelsWithChan(request *ListRankingModelsRequest) (<-chan *ListRankingModelsResponse, <-chan error) {
	responseChan := make(chan *ListRankingModelsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRankingModels(request)
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

// ListRankingModelsWithCallback invokes the airec.ListRankingModels API asynchronously
func (client *Client) ListRankingModelsWithCallback(request *ListRankingModelsRequest, callback func(response *ListRankingModelsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRankingModelsResponse
		var err error
		defer close(result)
		response, err = client.ListRankingModels(request)
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

// ListRankingModelsRequest is the request struct for api ListRankingModels
type ListRankingModelsRequest struct {
	*requests.RoaRequest
	InstanceId     string           `position:"Path" name:"instanceId"`
	Size           requests.Integer `position:"Query" name:"size"`
	RankingModelId string           `position:"Query" name:"rankingModelId"`
	Page           requests.Integer `position:"Query" name:"page"`
}

// ListRankingModelsResponse is the response struct for api ListRankingModels
type ListRankingModelsResponse struct {
	*responses.BaseResponse
	Code      string       `json:"code" xml:"code"`
	Message   string       `json:"message" xml:"message"`
	RequestId string       `json:"requestId" xml:"requestId"`
	Result    []ResultItem `json:"result" xml:"result"`
}

// CreateListRankingModelsRequest creates a request to invoke ListRankingModels API
func CreateListRankingModelsRequest() (request *ListRankingModelsRequest) {
	request = &ListRankingModelsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "ListRankingModels", "/v2/openapi/instances/[instanceId]/ranking-models", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListRankingModelsResponse creates a response to parse from ListRankingModels response
func CreateListRankingModelsResponse() (response *ListRankingModelsResponse) {
	response = &ListRankingModelsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}