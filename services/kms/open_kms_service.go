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

// OpenKmsService invokes the kms.OpenKmsService API synchronously
func (client *Client) OpenKmsService(request *OpenKmsServiceRequest) (response *OpenKmsServiceResponse, err error) {
	response = CreateOpenKmsServiceResponse()
	err = client.DoAction(request, response)
	return
}

// OpenKmsServiceWithChan invokes the kms.OpenKmsService API asynchronously
func (client *Client) OpenKmsServiceWithChan(request *OpenKmsServiceRequest) (<-chan *OpenKmsServiceResponse, <-chan error) {
	responseChan := make(chan *OpenKmsServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenKmsService(request)
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

// OpenKmsServiceWithCallback invokes the kms.OpenKmsService API asynchronously
func (client *Client) OpenKmsServiceWithCallback(request *OpenKmsServiceRequest, callback func(response *OpenKmsServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenKmsServiceResponse
		var err error
		defer close(result)
		response, err = client.OpenKmsService(request)
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

// OpenKmsServiceRequest is the request struct for api OpenKmsService
type OpenKmsServiceRequest struct {
	*requests.RpcRequest
}

// OpenKmsServiceResponse is the response struct for api OpenKmsService
type OpenKmsServiceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOpenKmsServiceRequest creates a request to invoke OpenKmsService API
func CreateOpenKmsServiceRequest() (request *OpenKmsServiceRequest) {
	request = &OpenKmsServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "OpenKmsService", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOpenKmsServiceResponse creates a response to parse from OpenKmsService response
func CreateOpenKmsServiceResponse() (response *OpenKmsServiceResponse) {
	response = &OpenKmsServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
