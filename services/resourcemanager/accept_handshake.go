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

// AcceptHandshake invokes the resourcemanager.AcceptHandshake API synchronously
func (client *Client) AcceptHandshake(request *AcceptHandshakeRequest) (response *AcceptHandshakeResponse, err error) {
	response = CreateAcceptHandshakeResponse()
	err = client.DoAction(request, response)
	return
}

// AcceptHandshakeWithChan invokes the resourcemanager.AcceptHandshake API asynchronously
func (client *Client) AcceptHandshakeWithChan(request *AcceptHandshakeRequest) (<-chan *AcceptHandshakeResponse, <-chan error) {
	responseChan := make(chan *AcceptHandshakeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AcceptHandshake(request)
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

// AcceptHandshakeWithCallback invokes the resourcemanager.AcceptHandshake API asynchronously
func (client *Client) AcceptHandshakeWithCallback(request *AcceptHandshakeRequest, callback func(response *AcceptHandshakeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AcceptHandshakeResponse
		var err error
		defer close(result)
		response, err = client.AcceptHandshake(request)
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

// AcceptHandshakeRequest is the request struct for api AcceptHandshake
type AcceptHandshakeRequest struct {
	*requests.RpcRequest
	HandshakeId string `position:"Query" name:"HandshakeId"`
}

// AcceptHandshakeResponse is the response struct for api AcceptHandshake
type AcceptHandshakeResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Handshake Handshake `json:"Handshake" xml:"Handshake"`
}

// CreateAcceptHandshakeRequest creates a request to invoke AcceptHandshake API
func CreateAcceptHandshakeRequest() (request *AcceptHandshakeRequest) {
	request = &AcceptHandshakeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "AcceptHandshake", "", "")
	request.Method = requests.POST
	return
}

// CreateAcceptHandshakeResponse creates a response to parse from AcceptHandshake response
func CreateAcceptHandshakeResponse() (response *AcceptHandshakeResponse) {
	response = &AcceptHandshakeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
