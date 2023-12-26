package avatar

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

// ClientAuth invokes the avatar.ClientAuth API synchronously
func (client *Client) ClientAuth(request *ClientAuthRequest) (response *ClientAuthResponse, err error) {
	response = CreateClientAuthResponse()
	err = client.DoAction(request, response)
	return
}

// ClientAuthWithChan invokes the avatar.ClientAuth API asynchronously
func (client *Client) ClientAuthWithChan(request *ClientAuthRequest) (<-chan *ClientAuthResponse, <-chan error) {
	responseChan := make(chan *ClientAuthResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClientAuth(request)
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

// ClientAuthWithCallback invokes the avatar.ClientAuth API asynchronously
func (client *Client) ClientAuthWithCallback(request *ClientAuthRequest, callback func(response *ClientAuthResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClientAuthResponse
		var err error
		defer close(result)
		response, err = client.ClientAuth(request)
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

// ClientAuthRequest is the request struct for api ClientAuth
type ClientAuthRequest struct {
	*requests.RpcRequest
	DeviceId   string           `position:"Query" name:"DeviceId"`
	DeviceType string           `position:"Query" name:"DeviceType"`
	License    string           `position:"Query" name:"License"`
	AppId      string           `position:"Query" name:"AppId"`
	TenantId   requests.Integer `position:"Query" name:"TenantId"`
	DeviceInfo string           `position:"Query" name:"DeviceInfo"`
}

// ClientAuthResponse is the response struct for api ClientAuth
type ClientAuthResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateClientAuthRequest creates a request to invoke ClientAuth API
func CreateClientAuthRequest() (request *ClientAuthRequest) {
	request = &ClientAuthRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "ClientAuth", "", "")
	request.Method = requests.POST
	return
}

// CreateClientAuthResponse creates a response to parse from ClientAuth response
func CreateClientAuthResponse() (response *ClientAuthResponse) {
	response = &ClientAuthResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
