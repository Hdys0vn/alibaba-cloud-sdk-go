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

// DeleteSignature invokes the cloudapi.DeleteSignature API synchronously
func (client *Client) DeleteSignature(request *DeleteSignatureRequest) (response *DeleteSignatureResponse, err error) {
	response = CreateDeleteSignatureResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSignatureWithChan invokes the cloudapi.DeleteSignature API asynchronously
func (client *Client) DeleteSignatureWithChan(request *DeleteSignatureRequest) (<-chan *DeleteSignatureResponse, <-chan error) {
	responseChan := make(chan *DeleteSignatureResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSignature(request)
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

// DeleteSignatureWithCallback invokes the cloudapi.DeleteSignature API asynchronously
func (client *Client) DeleteSignatureWithCallback(request *DeleteSignatureRequest, callback func(response *DeleteSignatureResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSignatureResponse
		var err error
		defer close(result)
		response, err = client.DeleteSignature(request)
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

// DeleteSignatureRequest is the request struct for api DeleteSignature
type DeleteSignatureRequest struct {
	*requests.RpcRequest
	SignatureId   string `position:"Query" name:"SignatureId"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
}

// DeleteSignatureResponse is the response struct for api DeleteSignature
type DeleteSignatureResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteSignatureRequest creates a request to invoke DeleteSignature API
func CreateDeleteSignatureRequest() (request *DeleteSignatureRequest) {
	request = &DeleteSignatureRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudAPI", "2016-07-14", "DeleteSignature", "apigateway", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSignatureResponse creates a response to parse from DeleteSignature response
func CreateDeleteSignatureResponse() (response *DeleteSignatureResponse) {
	response = &DeleteSignatureResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
