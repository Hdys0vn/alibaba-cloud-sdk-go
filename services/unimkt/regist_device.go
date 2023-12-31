package unimkt

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

// RegistDevice invokes the unimkt.RegistDevice API synchronously
func (client *Client) RegistDevice(request *RegistDeviceRequest) (response *RegistDeviceResponse, err error) {
	response = CreateRegistDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// RegistDeviceWithChan invokes the unimkt.RegistDevice API asynchronously
func (client *Client) RegistDeviceWithChan(request *RegistDeviceRequest) (<-chan *RegistDeviceResponse, <-chan error) {
	responseChan := make(chan *RegistDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RegistDevice(request)
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

// RegistDeviceWithCallback invokes the unimkt.RegistDevice API asynchronously
func (client *Client) RegistDeviceWithCallback(request *RegistDeviceRequest, callback func(response *RegistDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RegistDeviceResponse
		var err error
		defer close(result)
		response, err = client.RegistDevice(request)
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

// RegistDeviceRequest is the request struct for api RegistDevice
type RegistDeviceRequest struct {
	*requests.RpcRequest
	FirstScene        string `position:"Body" name:"FirstScene"`
	DetailAddr        string `position:"Body" name:"DetailAddr"`
	City              string `position:"Body" name:"City"`
	DeviceType        string `position:"Body" name:"DeviceType"`
	LocationName      string `position:"Body" name:"LocationName"`
	Province          string `position:"Body" name:"Province"`
	District          string `position:"Body" name:"District"`
	DeviceName        string `position:"Body" name:"DeviceName"`
	DeviceModelNumber string `position:"Body" name:"DeviceModelNumber"`
	SecondScene       string `position:"Body" name:"SecondScene"`
	Floor             string `position:"Body" name:"Floor"`
	ChannelId         string `position:"Body" name:"ChannelId"`
	OuterCode         string `position:"Body" name:"OuterCode"`
}

// RegistDeviceResponse is the response struct for api RegistDevice
type RegistDeviceResponse struct {
	*responses.BaseResponse
	Status    bool   `json:"Status" xml:"Status"`
	Msg       string `json:"Msg" xml:"Msg"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRegistDeviceRequest creates a request to invoke RegistDevice API
func CreateRegistDeviceRequest() (request *RegistDeviceRequest) {
	request = &RegistDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("UniMkt", "2018-12-12", "RegistDevice", "uniMkt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRegistDeviceResponse creates a response to parse from RegistDevice response
func CreateRegistDeviceResponse() (response *RegistDeviceResponse) {
	response = &RegistDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
