package dataworks_public

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

// ListProjectIds invokes the dataworks_public.ListProjectIds API synchronously
func (client *Client) ListProjectIds(request *ListProjectIdsRequest) (response *ListProjectIdsResponse, err error) {
	response = CreateListProjectIdsResponse()
	err = client.DoAction(request, response)
	return
}

// ListProjectIdsWithChan invokes the dataworks_public.ListProjectIds API asynchronously
func (client *Client) ListProjectIdsWithChan(request *ListProjectIdsRequest) (<-chan *ListProjectIdsResponse, <-chan error) {
	responseChan := make(chan *ListProjectIdsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProjectIds(request)
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

// ListProjectIdsWithCallback invokes the dataworks_public.ListProjectIds API asynchronously
func (client *Client) ListProjectIdsWithCallback(request *ListProjectIdsRequest, callback func(response *ListProjectIdsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProjectIdsResponse
		var err error
		defer close(result)
		response, err = client.ListProjectIds(request)
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

// ListProjectIdsRequest is the request struct for api ListProjectIds
type ListProjectIdsRequest struct {
	*requests.RpcRequest
	UserId string `position:"Query" name:"UserId"`
}

// ListProjectIdsResponse is the response struct for api ListProjectIds
type ListProjectIdsResponse struct {
	*responses.BaseResponse
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	ProjectIds []int64 `json:"ProjectIds" xml:"ProjectIds"`
}

// CreateListProjectIdsRequest creates a request to invoke ListProjectIds API
func CreateListProjectIdsRequest() (request *ListProjectIdsRequest) {
	request = &ListProjectIdsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListProjectIds", "", "")
	request.Method = requests.POST
	return
}

// CreateListProjectIdsResponse creates a response to parse from ListProjectIds response
func CreateListProjectIdsResponse() (response *ListProjectIdsResponse) {
	response = &ListProjectIdsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
