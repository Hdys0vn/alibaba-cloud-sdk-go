package kms

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

// CancelKeyDeletion invokes the kms.CancelKeyDeletion API synchronously
func (client *Client) CancelKeyDeletion(request *CancelKeyDeletionRequest) (response *CancelKeyDeletionResponse, err error) {
	response = CreateCancelKeyDeletionResponse()
	err = client.DoAction(request, response)
	return
}

// CancelKeyDeletionWithChan invokes the kms.CancelKeyDeletion API asynchronously
func (client *Client) CancelKeyDeletionWithChan(request *CancelKeyDeletionRequest) (<-chan *CancelKeyDeletionResponse, <-chan error) {
	responseChan := make(chan *CancelKeyDeletionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelKeyDeletion(request)
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

// CancelKeyDeletionWithCallback invokes the kms.CancelKeyDeletion API asynchronously
func (client *Client) CancelKeyDeletionWithCallback(request *CancelKeyDeletionRequest, callback func(response *CancelKeyDeletionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelKeyDeletionResponse
		var err error
		defer close(result)
		response, err = client.CancelKeyDeletion(request)
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

// CancelKeyDeletionRequest is the request struct for api CancelKeyDeletion
type CancelKeyDeletionRequest struct {
	*requests.RpcRequest
	KeyId string `position:"Query" name:"KeyId"`
}

// CancelKeyDeletionResponse is the response struct for api CancelKeyDeletion
type CancelKeyDeletionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelKeyDeletionRequest creates a request to invoke CancelKeyDeletion API
func CreateCancelKeyDeletionRequest() (request *CancelKeyDeletionRequest) {
	request = &CancelKeyDeletionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "CancelKeyDeletion", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelKeyDeletionResponse creates a response to parse from CancelKeyDeletion response
func CreateCancelKeyDeletionResponse() (response *CancelKeyDeletionResponse) {
	response = &CancelKeyDeletionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
