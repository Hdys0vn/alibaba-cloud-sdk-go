package ccc

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

// ListUnassignedNumbers invokes the ccc.ListUnassignedNumbers API synchronously
func (client *Client) ListUnassignedNumbers(request *ListUnassignedNumbersRequest) (response *ListUnassignedNumbersResponse, err error) {
	response = CreateListUnassignedNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// ListUnassignedNumbersWithChan invokes the ccc.ListUnassignedNumbers API asynchronously
func (client *Client) ListUnassignedNumbersWithChan(request *ListUnassignedNumbersRequest) (<-chan *ListUnassignedNumbersResponse, <-chan error) {
	responseChan := make(chan *ListUnassignedNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListUnassignedNumbers(request)
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

// ListUnassignedNumbersWithCallback invokes the ccc.ListUnassignedNumbers API asynchronously
func (client *Client) ListUnassignedNumbersWithCallback(request *ListUnassignedNumbersRequest, callback func(response *ListUnassignedNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListUnassignedNumbersResponse
		var err error
		defer close(result)
		response, err = client.ListUnassignedNumbers(request)
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

// ListUnassignedNumbersRequest is the request struct for api ListUnassignedNumbers
type ListUnassignedNumbersRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	SearchPattern string           `position:"Query" name:"SearchPattern"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListUnassignedNumbersResponse is the response struct for api ListUnassignedNumbers
type ListUnassignedNumbersResponse struct {
	*responses.BaseResponse
	Code           string                      `json:"Code" xml:"Code"`
	HttpStatusCode int                         `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                      `json:"Message" xml:"Message"`
	RequestId      string                      `json:"RequestId" xml:"RequestId"`
	Data           DataInListUnassignedNumbers `json:"Data" xml:"Data"`
}

// CreateListUnassignedNumbersRequest creates a request to invoke ListUnassignedNumbers API
func CreateListUnassignedNumbersRequest() (request *ListUnassignedNumbersRequest) {
	request = &ListUnassignedNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListUnassignedNumbers", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListUnassignedNumbersResponse creates a response to parse from ListUnassignedNumbers response
func CreateListUnassignedNumbersResponse() (response *ListUnassignedNumbersResponse) {
	response = &ListUnassignedNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
