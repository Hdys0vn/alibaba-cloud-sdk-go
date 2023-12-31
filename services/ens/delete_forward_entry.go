package ens

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

// DeleteForwardEntry invokes the ens.DeleteForwardEntry API synchronously
func (client *Client) DeleteForwardEntry(request *DeleteForwardEntryRequest) (response *DeleteForwardEntryResponse, err error) {
	response = CreateDeleteForwardEntryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteForwardEntryWithChan invokes the ens.DeleteForwardEntry API asynchronously
func (client *Client) DeleteForwardEntryWithChan(request *DeleteForwardEntryRequest) (<-chan *DeleteForwardEntryResponse, <-chan error) {
	responseChan := make(chan *DeleteForwardEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteForwardEntry(request)
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

// DeleteForwardEntryWithCallback invokes the ens.DeleteForwardEntry API asynchronously
func (client *Client) DeleteForwardEntryWithCallback(request *DeleteForwardEntryRequest, callback func(response *DeleteForwardEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteForwardEntryResponse
		var err error
		defer close(result)
		response, err = client.DeleteForwardEntry(request)
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

// DeleteForwardEntryRequest is the request struct for api DeleteForwardEntry
type DeleteForwardEntryRequest struct {
	*requests.RpcRequest
	ForwardEntryId string `position:"Query" name:"ForwardEntryId"`
}

// DeleteForwardEntryResponse is the response struct for api DeleteForwardEntry
type DeleteForwardEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteForwardEntryRequest creates a request to invoke DeleteForwardEntry API
func CreateDeleteForwardEntryRequest() (request *DeleteForwardEntryRequest) {
	request = &DeleteForwardEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DeleteForwardEntry", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteForwardEntryResponse creates a response to parse from DeleteForwardEntry response
func CreateDeleteForwardEntryResponse() (response *DeleteForwardEntryResponse) {
	response = &DeleteForwardEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
