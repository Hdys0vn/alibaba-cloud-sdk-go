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

// CreateClientKey invokes the kms.CreateClientKey API synchronously
func (client *Client) CreateClientKey(request *CreateClientKeyRequest) (response *CreateClientKeyResponse, err error) {
	response = CreateCreateClientKeyResponse()
	err = client.DoAction(request, response)
	return
}

// CreateClientKeyWithChan invokes the kms.CreateClientKey API asynchronously
func (client *Client) CreateClientKeyWithChan(request *CreateClientKeyRequest) (<-chan *CreateClientKeyResponse, <-chan error) {
	responseChan := make(chan *CreateClientKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateClientKey(request)
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

// CreateClientKeyWithCallback invokes the kms.CreateClientKey API asynchronously
func (client *Client) CreateClientKeyWithCallback(request *CreateClientKeyRequest, callback func(response *CreateClientKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateClientKeyResponse
		var err error
		defer close(result)
		response, err = client.CreateClientKey(request)
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

// CreateClientKeyRequest is the request struct for api CreateClientKey
type CreateClientKeyRequest struct {
	*requests.RpcRequest
	NotBefore string `position:"Query" name:"NotBefore"`
	NotAfter  string `position:"Query" name:"NotAfter"`
	Password  string `position:"Query" name:"Password"`
	AapName   string `position:"Query" name:"AapName"`
}

// CreateClientKeyResponse is the response struct for api CreateClientKey
type CreateClientKeyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ClientKeyId    string `json:"ClientKeyId" xml:"ClientKeyId"`
	KeyAlgorithm   string `json:"KeyAlgorithm" xml:"KeyAlgorithm"`
	PrivateKeyData string `json:"PrivateKeyData" xml:"PrivateKeyData"`
	NotBefore      string `json:"NotBefore" xml:"NotBefore"`
	NotAfter       string `json:"NotAfter" xml:"NotAfter"`
}

// CreateCreateClientKeyRequest creates a request to invoke CreateClientKey API
func CreateCreateClientKeyRequest() (request *CreateClientKeyRequest) {
	request = &CreateClientKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "CreateClientKey", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateClientKeyResponse creates a response to parse from CreateClientKey response
func CreateCreateClientKeyResponse() (response *CreateClientKeyResponse) {
	response = &CreateClientKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
