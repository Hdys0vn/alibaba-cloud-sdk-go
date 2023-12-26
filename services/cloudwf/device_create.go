package cloudwf

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

// DeviceCreate invokes the cloudwf.DeviceCreate API synchronously
// api document: https://help.aliyun.com/api/cloudwf/devicecreate.html
func (client *Client) DeviceCreate(request *DeviceCreateRequest) (response *DeviceCreateResponse, err error) {
	response = CreateDeviceCreateResponse()
	err = client.DoAction(request, response)
	return
}

// DeviceCreateWithChan invokes the cloudwf.DeviceCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/devicecreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeviceCreateWithChan(request *DeviceCreateRequest) (<-chan *DeviceCreateResponse, <-chan error) {
	responseChan := make(chan *DeviceCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeviceCreate(request)
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

// DeviceCreateWithCallback invokes the cloudwf.DeviceCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/devicecreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeviceCreateWithCallback(request *DeviceCreateRequest, callback func(response *DeviceCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeviceCreateResponse
		var err error
		defer close(result)
		response, err = client.DeviceCreate(request)
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

// DeviceCreateRequest is the request struct for api DeviceCreate
type DeviceCreateRequest struct {
	*requests.RpcRequest
	DeviceNum      string           `position:"Query" name:"DeviceNum"`
	DevicePosition string           `position:"Query" name:"DevicePosition"`
	DeviceName     string           `position:"Query" name:"DeviceName"`
	DeviceType     requests.Integer `position:"Query" name:"DeviceType"`
	Sid            requests.Integer `position:"Query" name:"Sid"`
}

// DeviceCreateResponse is the response struct for api DeviceCreate
type DeviceCreateResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateDeviceCreateRequest creates a request to invoke DeviceCreate API
func CreateDeviceCreateRequest() (request *DeviceCreateRequest) {
	request = &DeviceCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "DeviceCreate", "cloudwf", "openAPI")
	return
}

// CreateDeviceCreateResponse creates a response to parse from DeviceCreate response
func CreateDeviceCreateResponse() (response *DeviceCreateResponse) {
	response = &DeviceCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}