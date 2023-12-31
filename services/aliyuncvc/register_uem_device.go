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

// RegisterUemDevice invokes the aliyuncvc.RegisterUemDevice API synchronously
func (client *Client) RegisterUemDevice(request *RegisterUemDeviceRequest) (response *RegisterUemDeviceResponse, err error) {
	response = CreateRegisterUemDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// RegisterUemDeviceWithChan invokes the aliyuncvc.RegisterUemDevice API asynchronously
func (client *Client) RegisterUemDeviceWithChan(request *RegisterUemDeviceRequest) (<-chan *RegisterUemDeviceResponse, <-chan error) {
	responseChan := make(chan *RegisterUemDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegisterUemDevice(request)
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

// RegisterUemDeviceWithCallback invokes the aliyuncvc.RegisterUemDevice API asynchronously
func (client *Client) RegisterUemDeviceWithCallback(request *RegisterUemDeviceRequest, callback func(response *RegisterUemDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegisterUemDeviceResponse
		var err error
		defer close(result)
		response, err = client.RegisterUemDevice(request)
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

// RegisterUemDeviceRequest is the request struct for api RegisterUemDevice
type RegisterUemDeviceRequest struct {
	*requests.RpcRequest
	IP            string `position:"Query" name:"IP"`
	GroupId       string `position:"Query" name:"GroupId"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	DeviceId      string `position:"Query" name:"DeviceId"`
	GroupName     string `position:"Query" name:"GroupName"`
	Mac           string `position:"Query" name:"Mac"`
	DeviceVersion string `position:"Query" name:"DeviceVersion"`
}

// RegisterUemDeviceResponse is the response struct for api RegisterUemDevice
type RegisterUemDeviceResponse struct {
	*responses.BaseResponse
	ErrorCode  int        `json:"ErrorCode" xml:"ErrorCode"`
	Success    bool       `json:"Success" xml:"Success"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Message    string     `json:"Message" xml:"Message"`
	DeviceInfo DeviceInfo `json:"DeviceInfo" xml:"DeviceInfo"`
}

// CreateRegisterUemDeviceRequest creates a request to invoke RegisterUemDevice API
func CreateRegisterUemDeviceRequest() (request *RegisterUemDeviceRequest) {
	request = &RegisterUemDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "RegisterUemDevice", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRegisterUemDeviceResponse creates a response to parse from RegisterUemDevice response
func CreateRegisterUemDeviceResponse() (response *RegisterUemDeviceResponse) {
	response = &RegisterUemDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
