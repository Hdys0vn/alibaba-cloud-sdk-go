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

// ModifyDeviceBackground invokes the aliyuncvc.ModifyDeviceBackground API synchronously
func (client *Client) ModifyDeviceBackground(request *ModifyDeviceBackgroundRequest) (response *ModifyDeviceBackgroundResponse, err error) {
	response = CreateModifyDeviceBackgroundResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDeviceBackgroundWithChan invokes the aliyuncvc.ModifyDeviceBackground API asynchronously
func (client *Client) ModifyDeviceBackgroundWithChan(request *ModifyDeviceBackgroundRequest) (<-chan *ModifyDeviceBackgroundResponse, <-chan error) {
	responseChan := make(chan *ModifyDeviceBackgroundResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDeviceBackground(request)
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

// ModifyDeviceBackgroundWithCallback invokes the aliyuncvc.ModifyDeviceBackground API asynchronously
func (client *Client) ModifyDeviceBackgroundWithCallback(request *ModifyDeviceBackgroundRequest, callback func(response *ModifyDeviceBackgroundResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDeviceBackgroundResponse
		var err error
		defer close(result)
		response, err = client.ModifyDeviceBackground(request)
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

// ModifyDeviceBackgroundRequest is the request struct for api ModifyDeviceBackground
type ModifyDeviceBackgroundRequest struct {
	*requests.RpcRequest
	SerialNumber string `position:"Body" name:"SerialNumber"`
	Picture      string `position:"Body" name:"Picture"`
}

// ModifyDeviceBackgroundResponse is the response struct for api ModifyDeviceBackground
type ModifyDeviceBackgroundResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDeviceBackgroundRequest creates a request to invoke ModifyDeviceBackground API
func CreateModifyDeviceBackgroundRequest() (request *ModifyDeviceBackgroundRequest) {
	request = &ModifyDeviceBackgroundRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "ModifyDeviceBackground", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDeviceBackgroundResponse creates a response to parse from ModifyDeviceBackground response
func CreateModifyDeviceBackgroundResponse() (response *ModifyDeviceBackgroundResponse) {
	response = &ModifyDeviceBackgroundResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
