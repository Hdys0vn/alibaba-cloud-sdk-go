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

// ModifyBackend invokes the cloudapi.ModifyBackend API synchronously
func (client *Client) ModifyBackend(request *ModifyBackendRequest) (response *ModifyBackendResponse, err error) {
	response = CreateModifyBackendResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBackendWithChan invokes the cloudapi.ModifyBackend API asynchronously
func (client *Client) ModifyBackendWithChan(request *ModifyBackendRequest) (<-chan *ModifyBackendResponse, <-chan error) {
	responseChan := make(chan *ModifyBackendResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBackend(request)
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

// ModifyBackendWithCallback invokes the cloudapi.ModifyBackend API asynchronously
func (client *Client) ModifyBackendWithCallback(request *ModifyBackendRequest, callback func(response *ModifyBackendResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBackendResponse
		var err error
		defer close(result)
		response, err = client.ModifyBackend(request)
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

// ModifyBackendRequest is the request struct for api ModifyBackend
type ModifyBackendRequest struct {
	*requests.RpcRequest
	BackendId     string `position:"Query" name:"BackendId"`
	Description   string `position:"Query" name:"Description"`
	BackendName   string `position:"Query" name:"BackendName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	BackendType   string `position:"Query" name:"BackendType"`
}

// ModifyBackendResponse is the response struct for api ModifyBackend
type ModifyBackendResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyBackendRequest creates a request to invoke ModifyBackend API
func CreateModifyBackendRequest() (request *ModifyBackendRequest) {
	request = &ModifyBackendRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "ModifyBackend", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyBackendResponse creates a response to parse from ModifyBackend response
func CreateModifyBackendResponse() (response *ModifyBackendResponse) {
	response = &ModifyBackendResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
