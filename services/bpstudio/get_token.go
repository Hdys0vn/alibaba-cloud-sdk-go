package bpstudio

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

// GetToken invokes the bpstudio.GetToken API synchronously
func (client *Client) GetToken(request *GetTokenRequest) (response *GetTokenResponse, err error) {
	response = CreateGetTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetTokenWithChan invokes the bpstudio.GetToken API asynchronously
func (client *Client) GetTokenWithChan(request *GetTokenRequest) (<-chan *GetTokenResponse, <-chan error) {
	responseChan := make(chan *GetTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetToken(request)
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

// GetTokenWithCallback invokes the bpstudio.GetToken API asynchronously
func (client *Client) GetTokenWithCallback(request *GetTokenRequest, callback func(response *GetTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTokenResponse
		var err error
		defer close(result)
		response, err = client.GetToken(request)
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

// GetTokenRequest is the request struct for api GetToken
type GetTokenRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Body" name:"ResourceGroupId"`
}

// GetTokenResponse is the response struct for api GetToken
type GetTokenResponse struct {
	*responses.BaseResponse
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetTokenRequest creates a request to invoke GetToken API
func CreateGetTokenRequest() (request *GetTokenRequest) {
	request = &GetTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BPStudio", "2021-09-31", "GetToken", "bpstudio", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetTokenResponse creates a response to parse from GetToken response
func CreateGetTokenResponse() (response *GetTokenResponse) {
	response = &GetTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
