package mopen

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

// MoPenCreateDevice invokes the mopen.MoPenCreateDevice API synchronously
// api document: https://help.aliyun.com/api/mopen/mopencreatedevice.html
func (client *Client) MoPenCreateDevice(request *MoPenCreateDeviceRequest) (response *MoPenCreateDeviceResponse, err error) {
	response = CreateMoPenCreateDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// MoPenCreateDeviceWithChan invokes the mopen.MoPenCreateDevice API asynchronously
// api document: https://help.aliyun.com/api/mopen/mopencreatedevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoPenCreateDeviceWithChan(request *MoPenCreateDeviceRequest) (<-chan *MoPenCreateDeviceResponse, <-chan error) {
	responseChan := make(chan *MoPenCreateDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MoPenCreateDevice(request)
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

// MoPenCreateDeviceWithCallback invokes the mopen.MoPenCreateDevice API asynchronously
// api document: https://help.aliyun.com/api/mopen/mopencreatedevice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoPenCreateDeviceWithCallback(request *MoPenCreateDeviceRequest, callback func(response *MoPenCreateDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MoPenCreateDeviceResponse
		var err error
		defer close(result)
		response, err = client.MoPenCreateDevice(request)
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

// MoPenCreateDeviceRequest is the request struct for api MoPenCreateDevice
type MoPenCreateDeviceRequest struct {
	*requests.RpcRequest
	DeviceName string           `position:"Body" name:"DeviceName"`
	DeviceType requests.Integer `position:"Body" name:"DeviceType"`
}

// MoPenCreateDeviceResponse is the response struct for api MoPenCreateDevice
type MoPenCreateDeviceResponse struct {
	*responses.BaseResponse
	Code        bool   `json:"Code" xml:"Code"`
	Message     string `json:"Message" xml:"Message"`
	Success     bool   `json:"Success" xml:"Success"`
	Description string `json:"Description" xml:"Description"`
	Data        Data   `json:"Data" xml:"Data"`
}

// CreateMoPenCreateDeviceRequest creates a request to invoke MoPenCreateDevice API
func CreateMoPenCreateDeviceRequest() (request *MoPenCreateDeviceRequest) {
	request = &MoPenCreateDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("MoPen", "2018-02-11", "MoPenCreateDevice", "mopen", "openAPI")
	return
}

// CreateMoPenCreateDeviceResponse creates a response to parse from MoPenCreateDevice response
func CreateMoPenCreateDeviceResponse() (response *MoPenCreateDeviceResponse) {
	response = &MoPenCreateDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}