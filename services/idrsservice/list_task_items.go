package idrsservice

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

// ListTaskItems invokes the idrsservice.ListTaskItems API synchronously
func (client *Client) ListTaskItems(request *ListTaskItemsRequest) (response *ListTaskItemsResponse, err error) {
	response = CreateListTaskItemsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTaskItemsWithChan invokes the idrsservice.ListTaskItems API asynchronously
func (client *Client) ListTaskItemsWithChan(request *ListTaskItemsRequest) (<-chan *ListTaskItemsResponse, <-chan error) {
	responseChan := make(chan *ListTaskItemsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTaskItems(request)
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

// ListTaskItemsWithCallback invokes the idrsservice.ListTaskItems API asynchronously
func (client *Client) ListTaskItemsWithCallback(request *ListTaskItemsRequest, callback func(response *ListTaskItemsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTaskItemsResponse
		var err error
		defer close(result)
		response, err = client.ListTaskItems(request)
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

// ListTaskItemsRequest is the request struct for api ListTaskItems
type ListTaskItemsRequest struct {
	*requests.RpcRequest
	TaskId string `position:"Query" name:"TaskId"`
}

// ListTaskItemsResponse is the response struct for api ListTaskItems
type ListTaskItemsResponse struct {
	*responses.BaseResponse
	Code      string     `json:"Code" xml:"Code"`
	Message   string     `json:"Message" xml:"Message"`
	RequestId string     `json:"RequestId" xml:"RequestId"`
	Data      []DataItem `json:"Data" xml:"Data"`
}

// CreateListTaskItemsRequest creates a request to invoke ListTaskItems API
func CreateListTaskItemsRequest() (request *ListTaskItemsRequest) {
	request = &ListTaskItemsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("idrsservice", "2020-06-30", "ListTaskItems", "idrsservice", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListTaskItemsResponse creates a response to parse from ListTaskItems response
func CreateListTaskItemsResponse() (response *ListTaskItemsResponse) {
	response = &ListTaskItemsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
