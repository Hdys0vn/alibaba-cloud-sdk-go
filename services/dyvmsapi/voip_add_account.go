package dyvmsapi

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

// VoipAddAccount invokes the dyvmsapi.VoipAddAccount API synchronously
func (client *Client) VoipAddAccount(request *VoipAddAccountRequest) (response *VoipAddAccountResponse, err error) {
	response = CreateVoipAddAccountResponse()
	err = client.DoAction(request, response)
	return
}

// VoipAddAccountWithChan invokes the dyvmsapi.VoipAddAccount API asynchronously
func (client *Client) VoipAddAccountWithChan(request *VoipAddAccountRequest) (<-chan *VoipAddAccountResponse, <-chan error) {
	responseChan := make(chan *VoipAddAccountResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VoipAddAccount(request)
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

// VoipAddAccountWithCallback invokes the dyvmsapi.VoipAddAccount API asynchronously
func (client *Client) VoipAddAccountWithCallback(request *VoipAddAccountRequest, callback func(response *VoipAddAccountResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VoipAddAccountResponse
		var err error
		defer close(result)
		response, err = client.VoipAddAccount(request)
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

// VoipAddAccountRequest is the request struct for api VoipAddAccount
type VoipAddAccountRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DeviceId             string           `position:"Query" name:"DeviceId"`
}

// VoipAddAccountResponse is the response struct for api VoipAddAccount
type VoipAddAccountResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Module    string `json:"Module" xml:"Module"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateVoipAddAccountRequest creates a request to invoke VoipAddAccount API
func CreateVoipAddAccountRequest() (request *VoipAddAccountRequest) {
	request = &VoipAddAccountRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "VoipAddAccount", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVoipAddAccountResponse creates a response to parse from VoipAddAccount response
func CreateVoipAddAccountResponse() (response *VoipAddAccountResponse) {
	response = &VoipAddAccountResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
