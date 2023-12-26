package ram

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

// CreateVirtualMFADevice invokes the ram.CreateVirtualMFADevice API synchronously
func (client *Client) CreateVirtualMFADevice(request *CreateVirtualMFADeviceRequest) (response *CreateVirtualMFADeviceResponse, err error) {
	response = CreateCreateVirtualMFADeviceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateVirtualMFADeviceWithChan invokes the ram.CreateVirtualMFADevice API asynchronously
func (client *Client) CreateVirtualMFADeviceWithChan(request *CreateVirtualMFADeviceRequest) (<-chan *CreateVirtualMFADeviceResponse, <-chan error) {
	responseChan := make(chan *CreateVirtualMFADeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateVirtualMFADevice(request)
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

// CreateVirtualMFADeviceWithCallback invokes the ram.CreateVirtualMFADevice API asynchronously
func (client *Client) CreateVirtualMFADeviceWithCallback(request *CreateVirtualMFADeviceRequest, callback func(response *CreateVirtualMFADeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateVirtualMFADeviceResponse
		var err error
		defer close(result)
		response, err = client.CreateVirtualMFADevice(request)
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

// CreateVirtualMFADeviceRequest is the request struct for api CreateVirtualMFADevice
type CreateVirtualMFADeviceRequest struct {
	*requests.RpcRequest
	VirtualMFADeviceName string `position:"Query" name:"VirtualMFADeviceName"`
}

// CreateVirtualMFADeviceResponse is the response struct for api CreateVirtualMFADevice
type CreateVirtualMFADeviceResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	VirtualMFADevice VirtualMFADevice `json:"VirtualMFADevice" xml:"VirtualMFADevice"`
}

// CreateCreateVirtualMFADeviceRequest creates a request to invoke CreateVirtualMFADevice API
func CreateCreateVirtualMFADeviceRequest() (request *CreateVirtualMFADeviceRequest) {
	request = &CreateVirtualMFADeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ram", "2015-05-01", "CreateVirtualMFADevice", "Ram", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateVirtualMFADeviceResponse creates a response to parse from CreateVirtualMFADevice response
func CreateCreateVirtualMFADeviceResponse() (response *CreateVirtualMFADeviceResponse) {
	response = &CreateVirtualMFADeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
