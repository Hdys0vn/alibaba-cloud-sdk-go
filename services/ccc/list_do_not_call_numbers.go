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

// ListDoNotCallNumbers invokes the ccc.ListDoNotCallNumbers API synchronously
func (client *Client) ListDoNotCallNumbers(request *ListDoNotCallNumbersRequest) (response *ListDoNotCallNumbersResponse, err error) {
	response = CreateListDoNotCallNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// ListDoNotCallNumbersWithChan invokes the ccc.ListDoNotCallNumbers API asynchronously
func (client *Client) ListDoNotCallNumbersWithChan(request *ListDoNotCallNumbersRequest) (<-chan *ListDoNotCallNumbersResponse, <-chan error) {
	responseChan := make(chan *ListDoNotCallNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDoNotCallNumbers(request)
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

// ListDoNotCallNumbersWithCallback invokes the ccc.ListDoNotCallNumbers API asynchronously
func (client *Client) ListDoNotCallNumbersWithCallback(request *ListDoNotCallNumbersRequest, callback func(response *ListDoNotCallNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDoNotCallNumbersResponse
		var err error
		defer close(result)
		response, err = client.ListDoNotCallNumbers(request)
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

// ListDoNotCallNumbersRequest is the request struct for api ListDoNotCallNumbers
type ListDoNotCallNumbersRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Query" name:"PageNumber"`
	SearchPattern string           `position:"Query" name:"SearchPattern"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	Scope         string           `position:"Query" name:"Scope"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// ListDoNotCallNumbersResponse is the response struct for api ListDoNotCallNumbers
type ListDoNotCallNumbersResponse struct {
	*responses.BaseResponse
	Code           string                     `json:"Code" xml:"Code"`
	HttpStatusCode int                        `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string                     `json:"Message" xml:"Message"`
	RequestId      string                     `json:"RequestId" xml:"RequestId"`
	Data           DataInListDoNotCallNumbers `json:"Data" xml:"Data"`
}

// CreateListDoNotCallNumbersRequest creates a request to invoke ListDoNotCallNumbers API
func CreateListDoNotCallNumbersRequest() (request *ListDoNotCallNumbersRequest) {
	request = &ListDoNotCallNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListDoNotCallNumbers", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDoNotCallNumbersResponse creates a response to parse from ListDoNotCallNumbers response
func CreateListDoNotCallNumbersResponse() (response *ListDoNotCallNumbersResponse) {
	response = &ListDoNotCallNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}