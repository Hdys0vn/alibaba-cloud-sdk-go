package ltl

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

// AuthorizeDevice invokes the ltl.AuthorizeDevice API synchronously
func (client *Client) AuthorizeDevice(request *AuthorizeDeviceRequest) (response *AuthorizeDeviceResponse, err error) {
	response = CreateAuthorizeDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// AuthorizeDeviceWithChan invokes the ltl.AuthorizeDevice API asynchronously
func (client *Client) AuthorizeDeviceWithChan(request *AuthorizeDeviceRequest) (<-chan *AuthorizeDeviceResponse, <-chan error) {
	responseChan := make(chan *AuthorizeDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AuthorizeDevice(request)
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

// AuthorizeDeviceWithCallback invokes the ltl.AuthorizeDevice API asynchronously
func (client *Client) AuthorizeDeviceWithCallback(request *AuthorizeDeviceRequest, callback func(response *AuthorizeDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AuthorizeDeviceResponse
		var err error
		defer close(result)
		response, err = client.AuthorizeDevice(request)
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

// AuthorizeDeviceRequest is the request struct for api AuthorizeDevice
type AuthorizeDeviceRequest struct {
	*requests.RpcRequest
	ApiVersion    string `position:"Query" name:"ApiVersion"`
	DeviceId      string `position:"Query" name:"DeviceId"`
	DeviceGroupId string `position:"Query" name:"DeviceGroupId"`
	BizChainId    string `position:"Query" name:"BizChainId"`
}

// AuthorizeDeviceResponse is the response struct for api AuthorizeDevice
type AuthorizeDeviceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateAuthorizeDeviceRequest creates a request to invoke AuthorizeDevice API
func CreateAuthorizeDeviceRequest() (request *AuthorizeDeviceRequest) {
	request = &AuthorizeDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ltl", "2019-05-10", "AuthorizeDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateAuthorizeDeviceResponse creates a response to parse from AuthorizeDevice response
func CreateAuthorizeDeviceResponse() (response *AuthorizeDeviceResponse) {
	response = &AuthorizeDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
