package aliyuncvc

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

// GetDeviceToken invokes the aliyuncvc.GetDeviceToken API synchronously
func (client *Client) GetDeviceToken(request *GetDeviceTokenRequest) (response *GetDeviceTokenResponse, err error) {
	response = CreateGetDeviceTokenResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceTokenWithChan invokes the aliyuncvc.GetDeviceToken API asynchronously
func (client *Client) GetDeviceTokenWithChan(request *GetDeviceTokenRequest) (<-chan *GetDeviceTokenResponse, <-chan error) {
	responseChan := make(chan *GetDeviceTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceToken(request)
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

// GetDeviceTokenWithCallback invokes the aliyuncvc.GetDeviceToken API asynchronously
func (client *Client) GetDeviceTokenWithCallback(request *GetDeviceTokenRequest, callback func(response *GetDeviceTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceTokenResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceToken(request)
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

// GetDeviceTokenRequest is the request struct for api GetDeviceToken
type GetDeviceTokenRequest struct {
	*requests.RpcRequest
	Token string `position:"Query" name:"Token"`
	SN    string `position:"Query" name:"SN"`
}

// GetDeviceTokenResponse is the response struct for api GetDeviceToken
type GetDeviceTokenResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Token     string `json:"Token" xml:"Token"`
}

// CreateGetDeviceTokenRequest creates a request to invoke GetDeviceToken API
func CreateGetDeviceTokenRequest() (request *GetDeviceTokenRequest) {
	request = &GetDeviceTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "GetDeviceToken", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDeviceTokenResponse creates a response to parse from GetDeviceToken response
func CreateGetDeviceTokenResponse() (response *GetDeviceTokenResponse) {
	response = &GetDeviceTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
