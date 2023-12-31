package resourcemanager

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

// CancelHandshake invokes the resourcemanager.CancelHandshake API synchronously
func (client *Client) CancelHandshake(request *CancelHandshakeRequest) (response *CancelHandshakeResponse, err error) {
	response = CreateCancelHandshakeResponse()
	err = client.DoAction(request, response)
	return
}

// CancelHandshakeWithChan invokes the resourcemanager.CancelHandshake API asynchronously
func (client *Client) CancelHandshakeWithChan(request *CancelHandshakeRequest) (<-chan *CancelHandshakeResponse, <-chan error) {
	responseChan := make(chan *CancelHandshakeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelHandshake(request)
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

// CancelHandshakeWithCallback invokes the resourcemanager.CancelHandshake API asynchronously
func (client *Client) CancelHandshakeWithCallback(request *CancelHandshakeRequest, callback func(response *CancelHandshakeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelHandshakeResponse
		var err error
		defer close(result)
		response, err = client.CancelHandshake(request)
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

// CancelHandshakeRequest is the request struct for api CancelHandshake
type CancelHandshakeRequest struct {
	*requests.RpcRequest
	HandshakeId string `position:"Query" name:"HandshakeId"`
}

// CancelHandshakeResponse is the response struct for api CancelHandshake
type CancelHandshakeResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Handshake Handshake `json:"Handshake" xml:"Handshake"`
}

// CreateCancelHandshakeRequest creates a request to invoke CancelHandshake API
func CreateCancelHandshakeRequest() (request *CancelHandshakeRequest) {
	request = &CancelHandshakeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "CancelHandshake", "", "")
	request.Method = requests.POST
	return
}

// CreateCancelHandshakeResponse creates a response to parse from CancelHandshake response
func CreateCancelHandshakeResponse() (response *CancelHandshakeResponse) {
	response = &CancelHandshakeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
