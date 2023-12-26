package swas_open

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

// ReleasePublicConnection invokes the swas_open.ReleasePublicConnection API synchronously
func (client *Client) ReleasePublicConnection(request *ReleasePublicConnectionRequest) (response *ReleasePublicConnectionResponse, err error) {
	response = CreateReleasePublicConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// ReleasePublicConnectionWithChan invokes the swas_open.ReleasePublicConnection API asynchronously
func (client *Client) ReleasePublicConnectionWithChan(request *ReleasePublicConnectionRequest) (<-chan *ReleasePublicConnectionResponse, <-chan error) {
	responseChan := make(chan *ReleasePublicConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReleasePublicConnection(request)
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

// ReleasePublicConnectionWithCallback invokes the swas_open.ReleasePublicConnection API asynchronously
func (client *Client) ReleasePublicConnectionWithCallback(request *ReleasePublicConnectionRequest, callback func(response *ReleasePublicConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReleasePublicConnectionResponse
		var err error
		defer close(result)
		response, err = client.ReleasePublicConnection(request)
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

// ReleasePublicConnectionRequest is the request struct for api ReleasePublicConnection
type ReleasePublicConnectionRequest struct {
	*requests.RpcRequest
	ClientToken        string `position:"Query" name:"ClientToken"`
	DatabaseInstanceId string `position:"Query" name:"DatabaseInstanceId"`
}

// ReleasePublicConnectionResponse is the response struct for api ReleasePublicConnection
type ReleasePublicConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateReleasePublicConnectionRequest creates a request to invoke ReleasePublicConnection API
func CreateReleasePublicConnectionRequest() (request *ReleasePublicConnectionRequest) {
	request = &ReleasePublicConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "ReleasePublicConnection", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateReleasePublicConnectionResponse creates a response to parse from ReleasePublicConnection response
func CreateReleasePublicConnectionResponse() (response *ReleasePublicConnectionResponse) {
	response = &ReleasePublicConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
