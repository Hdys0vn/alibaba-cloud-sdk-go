package cloudapi

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

// CreateBackend invokes the cloudapi.CreateBackend API synchronously
func (client *Client) CreateBackend(request *CreateBackendRequest) (response *CreateBackendResponse, err error) {
	response = CreateCreateBackendResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBackendWithChan invokes the cloudapi.CreateBackend API asynchronously
func (client *Client) CreateBackendWithChan(request *CreateBackendRequest) (<-chan *CreateBackendResponse, <-chan error) {
	responseChan := make(chan *CreateBackendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBackend(request)
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

// CreateBackendWithCallback invokes the cloudapi.CreateBackend API asynchronously
func (client *Client) CreateBackendWithCallback(request *CreateBackendRequest, callback func(response *CreateBackendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBackendResponse
		var err error
		defer close(result)
		response, err = client.CreateBackend(request)
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

// CreateBackendRequest is the request struct for api CreateBackend
type CreateBackendRequest struct {
	*requests.RpcRequest
	Description                        string           `position:"Query" name:"Description"`
	BackendName                        string           `position:"Query" name:"BackendName"`
	CreateEventBridgeServiceLinkedRole requests.Boolean `position:"Query" name:"CreateEventBridgeServiceLinkedRole"`
	SecurityToken                      string           `position:"Query" name:"SecurityToken"`
	BackendType                        string           `position:"Query" name:"BackendType"`
}

// CreateBackendResponse is the response struct for api CreateBackend
type CreateBackendResponse struct {
	*responses.BaseResponse
	BackendId string `json:"BackendId" xml:"BackendId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateBackendRequest creates a request to invoke CreateBackend API
func CreateCreateBackendRequest() (request *CreateBackendRequest) {
	request = &CreateBackendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "CreateBackend", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBackendResponse creates a response to parse from CreateBackend response
func CreateCreateBackendResponse() (response *CreateBackendResponse) {
	response = &CreateBackendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}