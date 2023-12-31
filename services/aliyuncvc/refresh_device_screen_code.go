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

// RefreshDeviceScreenCode invokes the aliyuncvc.RefreshDeviceScreenCode API synchronously
func (client *Client) RefreshDeviceScreenCode(request *RefreshDeviceScreenCodeRequest) (response *RefreshDeviceScreenCodeResponse, err error) {
	response = CreateRefreshDeviceScreenCodeResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshDeviceScreenCodeWithChan invokes the aliyuncvc.RefreshDeviceScreenCode API asynchronously
func (client *Client) RefreshDeviceScreenCodeWithChan(request *RefreshDeviceScreenCodeRequest) (<-chan *RefreshDeviceScreenCodeResponse, <-chan error) {
	responseChan := make(chan *RefreshDeviceScreenCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshDeviceScreenCode(request)
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

// RefreshDeviceScreenCodeWithCallback invokes the aliyuncvc.RefreshDeviceScreenCode API asynchronously
func (client *Client) RefreshDeviceScreenCodeWithCallback(request *RefreshDeviceScreenCodeRequest, callback func(response *RefreshDeviceScreenCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshDeviceScreenCodeResponse
		var err error
		defer close(result)
		response, err = client.RefreshDeviceScreenCode(request)
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

// RefreshDeviceScreenCodeRequest is the request struct for api RefreshDeviceScreenCode
type RefreshDeviceScreenCodeRequest struct {
	*requests.RpcRequest
	SerialNumber string `position:"Body" name:"SerialNumber"`
}

// RefreshDeviceScreenCodeResponse is the response struct for api RefreshDeviceScreenCode
type RefreshDeviceScreenCodeResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRefreshDeviceScreenCodeRequest creates a request to invoke RefreshDeviceScreenCode API
func CreateRefreshDeviceScreenCodeRequest() (request *RefreshDeviceScreenCodeRequest) {
	request = &RefreshDeviceScreenCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "RefreshDeviceScreenCode", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRefreshDeviceScreenCodeResponse creates a response to parse from RefreshDeviceScreenCode response
func CreateRefreshDeviceScreenCodeResponse() (response *RefreshDeviceScreenCodeResponse) {
	response = &RefreshDeviceScreenCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
