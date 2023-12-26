package vcs

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

// ListBodyAlgorithmResults invokes the vcs.ListBodyAlgorithmResults API synchronously
func (client *Client) ListBodyAlgorithmResults(request *ListBodyAlgorithmResultsRequest) (response *ListBodyAlgorithmResultsResponse, err error) {
	response = CreateListBodyAlgorithmResultsResponse()
	err = client.DoAction(request, response)
	return
}

// ListBodyAlgorithmResultsWithChan invokes the vcs.ListBodyAlgorithmResults API asynchronously
func (client *Client) ListBodyAlgorithmResultsWithChan(request *ListBodyAlgorithmResultsRequest) (<-chan *ListBodyAlgorithmResultsResponse, <-chan error) {
	responseChan := make(chan *ListBodyAlgorithmResultsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListBodyAlgorithmResults(request)
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

// ListBodyAlgorithmResultsWithCallback invokes the vcs.ListBodyAlgorithmResults API asynchronously
func (client *Client) ListBodyAlgorithmResultsWithCallback(request *ListBodyAlgorithmResultsRequest, callback func(response *ListBodyAlgorithmResultsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListBodyAlgorithmResultsResponse
		var err error
		defer close(result)
		response, err = client.ListBodyAlgorithmResults(request)
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

// ListBodyAlgorithmResultsRequest is the request struct for api ListBodyAlgorithmResults
type ListBodyAlgorithmResultsRequest struct {
	*requests.RpcRequest
	AlgorithmType string `position:"Body" name:"AlgorithmType"`
	CorpId        string `position:"Body" name:"CorpId"`
	CapStyle      string `position:"Body" name:"CapStyle"`
	EndTime       string `position:"Body" name:"EndTime"`
	StartTime     string `position:"Body" name:"StartTime"`
	PageNumber    string `position:"Body" name:"PageNumber"`
	DataSourceId  string `position:"Body" name:"DataSourceId"`
	PageSize      string `position:"Body" name:"PageSize"`
}

// ListBodyAlgorithmResultsResponse is the response struct for api ListBodyAlgorithmResults
type ListBodyAlgorithmResultsResponse struct {
	*responses.BaseResponse
	Code      string                         `json:"Code" xml:"Code"`
	Message   string                         `json:"Message" xml:"Message"`
	RequestId string                         `json:"RequestId" xml:"RequestId"`
	Data      DataInListBodyAlgorithmResults `json:"Data" xml:"Data"`
}

// CreateListBodyAlgorithmResultsRequest creates a request to invoke ListBodyAlgorithmResults API
func CreateListBodyAlgorithmResultsRequest() (request *ListBodyAlgorithmResultsRequest) {
	request = &ListBodyAlgorithmResultsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListBodyAlgorithmResults", "", "")
	request.Method = requests.POST
	return
}

// CreateListBodyAlgorithmResultsResponse creates a response to parse from ListBodyAlgorithmResults response
func CreateListBodyAlgorithmResultsResponse() (response *ListBodyAlgorithmResultsResponse) {
	response = &ListBodyAlgorithmResultsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
