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

// DeleteNetworkAclEntry invokes the ens.DeleteNetworkAclEntry API synchronously
func (client *Client) DeleteNetworkAclEntry(request *DeleteNetworkAclEntryRequest) (response *DeleteNetworkAclEntryResponse, err error) {
	response = CreateDeleteNetworkAclEntryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteNetworkAclEntryWithChan invokes the ens.DeleteNetworkAclEntry API asynchronously
func (client *Client) DeleteNetworkAclEntryWithChan(request *DeleteNetworkAclEntryRequest) (<-chan *DeleteNetworkAclEntryResponse, <-chan error) {
	responseChan := make(chan *DeleteNetworkAclEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteNetworkAclEntry(request)
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

// DeleteNetworkAclEntryWithCallback invokes the ens.DeleteNetworkAclEntry API asynchronously
func (client *Client) DeleteNetworkAclEntryWithCallback(request *DeleteNetworkAclEntryRequest, callback func(response *DeleteNetworkAclEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteNetworkAclEntryResponse
		var err error
		defer close(result)
		response, err = client.DeleteNetworkAclEntry(request)
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

// DeleteNetworkAclEntryRequest is the request struct for api DeleteNetworkAclEntry
type DeleteNetworkAclEntryRequest struct {
	*requests.RpcRequest
	NetworkAclEntryId string `position:"Query" name:"NetworkAclEntryId"`
}

// DeleteNetworkAclEntryResponse is the response struct for api DeleteNetworkAclEntry
type DeleteNetworkAclEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteNetworkAclEntryRequest creates a request to invoke DeleteNetworkAclEntry API
func CreateDeleteNetworkAclEntryRequest() (request *DeleteNetworkAclEntryRequest) {
	request = &DeleteNetworkAclEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DeleteNetworkAclEntry", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteNetworkAclEntryResponse creates a response to parse from DeleteNetworkAclEntry response
func CreateDeleteNetworkAclEntryResponse() (response *DeleteNetworkAclEntryResponse) {
	response = &DeleteNetworkAclEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
