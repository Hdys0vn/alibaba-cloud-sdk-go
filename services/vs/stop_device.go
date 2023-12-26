package vs

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

// StopDevice invokes the vs.StopDevice API synchronously
func (client *Client) StopDevice(request *StopDeviceRequest) (response *StopDeviceResponse, err error) {
	response = CreateStopDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// StopDeviceWithChan invokes the vs.StopDevice API asynchronously
func (client *Client) StopDeviceWithChan(request *StopDeviceRequest) (<-chan *StopDeviceResponse, <-chan error) {
	responseChan := make(chan *StopDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDevice(request)
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

// StopDeviceWithCallback invokes the vs.StopDevice API asynchronously
func (client *Client) StopDeviceWithCallback(request *StopDeviceRequest, callback func(response *StopDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDeviceResponse
		var err error
		defer close(result)
		response, err = client.StopDevice(request)
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

// StopDeviceRequest is the request struct for api StopDevice
type StopDeviceRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Id        string           `position:"Query" name:"Id"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// StopDeviceResponse is the response struct for api StopDevice
type StopDeviceResponse struct {
	*responses.BaseResponse
	Id        string `json:"Id" xml:"Id"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopDeviceRequest creates a request to invoke StopDevice API
func CreateStopDeviceRequest() (request *StopDeviceRequest) {
	request = &StopDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "StopDevice", "", "")
	request.Method = requests.POST
	return
}

// CreateStopDeviceResponse creates a response to parse from StopDevice response
func CreateStopDeviceResponse() (response *StopDeviceResponse) {
	response = &StopDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
