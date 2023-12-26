package cloudesl

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

// UnbindEslDevice invokes the cloudesl.UnbindEslDevice API synchronously
func (client *Client) UnbindEslDevice(request *UnbindEslDeviceRequest) (response *UnbindEslDeviceResponse, err error) {
	response = CreateUnbindEslDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindEslDeviceWithChan invokes the cloudesl.UnbindEslDevice API asynchronously
func (client *Client) UnbindEslDeviceWithChan(request *UnbindEslDeviceRequest) (<-chan *UnbindEslDeviceResponse, <-chan error) {
	responseChan := make(chan *UnbindEslDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindEslDevice(request)
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

// UnbindEslDeviceWithCallback invokes the cloudesl.UnbindEslDevice API asynchronously
func (client *Client) UnbindEslDeviceWithCallback(request *UnbindEslDeviceRequest, callback func(response *UnbindEslDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindEslDeviceResponse
		var err error
		defer close(result)
		response, err = client.UnbindEslDevice(request)
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

// UnbindEslDeviceRequest is the request struct for api UnbindEslDevice
type UnbindEslDeviceRequest struct {
	*requests.RpcRequest
	ExtraParams   string           `position:"Body" name:"ExtraParams"`
	ContainerName string           `position:"Body" name:"ContainerName"`
	StoreId       string           `position:"Body" name:"StoreId"`
	Layer         requests.Integer `position:"Body" name:"Layer"`
	EslBarCode    string           `position:"Body" name:"EslBarCode"`
	ItemBarCode   string           `position:"Body" name:"ItemBarCode"`
	Column        string           `position:"Body" name:"Column"`
	Shelf         string           `position:"Body" name:"Shelf"`
}

// UnbindEslDeviceResponse is the response struct for api UnbindEslDevice
type UnbindEslDeviceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
}

// CreateUnbindEslDeviceRequest creates a request to invoke UnbindEslDevice API
func CreateUnbindEslDeviceRequest() (request *UnbindEslDeviceRequest) {
	request = &UnbindEslDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "UnbindEslDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateUnbindEslDeviceResponse creates a response to parse from UnbindEslDevice response
func CreateUnbindEslDeviceResponse() (response *UnbindEslDeviceResponse) {
	response = &UnbindEslDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
