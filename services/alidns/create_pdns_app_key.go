package alidns

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

// CreatePdnsAppKey invokes the alidns.CreatePdnsAppKey API synchronously
func (client *Client) CreatePdnsAppKey(request *CreatePdnsAppKeyRequest) (response *CreatePdnsAppKeyResponse, err error) {
	response = CreateCreatePdnsAppKeyResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePdnsAppKeyWithChan invokes the alidns.CreatePdnsAppKey API asynchronously
func (client *Client) CreatePdnsAppKeyWithChan(request *CreatePdnsAppKeyRequest) (<-chan *CreatePdnsAppKeyResponse, <-chan error) {
	responseChan := make(chan *CreatePdnsAppKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePdnsAppKey(request)
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

// CreatePdnsAppKeyWithCallback invokes the alidns.CreatePdnsAppKey API asynchronously
func (client *Client) CreatePdnsAppKeyWithCallback(request *CreatePdnsAppKeyRequest, callback func(response *CreatePdnsAppKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePdnsAppKeyResponse
		var err error
		defer close(result)
		response, err = client.CreatePdnsAppKey(request)
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

// CreatePdnsAppKeyRequest is the request struct for api CreatePdnsAppKey
type CreatePdnsAppKeyRequest struct {
	*requests.RpcRequest
	Lang string `position:"Query" name:"Lang"`
}

// CreatePdnsAppKeyResponse is the response struct for api CreatePdnsAppKey
type CreatePdnsAppKeyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreatePdnsAppKeyRequest creates a request to invoke CreatePdnsAppKey API
func CreateCreatePdnsAppKeyRequest() (request *CreatePdnsAppKeyRequest) {
	request = &CreatePdnsAppKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Alidns", "2015-01-09", "CreatePdnsAppKey", "alidns", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePdnsAppKeyResponse creates a response to parse from CreatePdnsAppKey response
func CreateCreatePdnsAppKeyResponse() (response *CreatePdnsAppKeyResponse) {
	response = &CreatePdnsAppKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
