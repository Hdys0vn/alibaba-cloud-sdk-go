package dg

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

// AddDatabaseList invokes the dg.AddDatabaseList API synchronously
func (client *Client) AddDatabaseList(request *AddDatabaseListRequest) (response *AddDatabaseListResponse, err error) {
	response = CreateAddDatabaseListResponse()
	err = client.DoAction(request, response)
	return
}

// AddDatabaseListWithChan invokes the dg.AddDatabaseList API asynchronously
func (client *Client) AddDatabaseListWithChan(request *AddDatabaseListRequest) (<-chan *AddDatabaseListResponse, <-chan error) {
	responseChan := make(chan *AddDatabaseListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDatabaseList(request)
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

// AddDatabaseListWithCallback invokes the dg.AddDatabaseList API asynchronously
func (client *Client) AddDatabaseListWithCallback(request *AddDatabaseListRequest, callback func(response *AddDatabaseListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDatabaseListResponse
		var err error
		defer close(result)
		response, err = client.AddDatabaseList(request)
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

// AddDatabaseListRequest is the request struct for api AddDatabaseList
type AddDatabaseListRequest struct {
	*requests.RpcRequest
	DatabaseString string `position:"Body" name:"DatabaseString"`
	ClientToken    string `position:"Body" name:"ClientToken"`
}

// AddDatabaseListResponse is the response struct for api AddDatabaseList
type AddDatabaseListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateAddDatabaseListRequest creates a request to invoke AddDatabaseList API
func CreateAddDatabaseListRequest() (request *AddDatabaseListRequest) {
	request = &AddDatabaseListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dg", "2019-03-27", "AddDatabaseList", "dg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddDatabaseListResponse creates a response to parse from AddDatabaseList response
func CreateAddDatabaseListResponse() (response *AddDatabaseListResponse) {
	response = &AddDatabaseListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
