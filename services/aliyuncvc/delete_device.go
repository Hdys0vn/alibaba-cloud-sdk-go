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

// DeleteDevice invokes the aliyuncvc.DeleteDevice API synchronously
func (client *Client) DeleteDevice(request *DeleteDeviceRequest) (response *DeleteDeviceResponse, err error) {
	response = CreateDeleteDeviceResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDeviceWithChan invokes the aliyuncvc.DeleteDevice API asynchronously
func (client *Client) DeleteDeviceWithChan(request *DeleteDeviceRequest) (<-chan *DeleteDeviceResponse, <-chan error) {
	responseChan := make(chan *DeleteDeviceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDevice(request)
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

// DeleteDeviceWithCallback invokes the aliyuncvc.DeleteDevice API asynchronously
func (client *Client) DeleteDeviceWithCallback(request *DeleteDeviceRequest, callback func(response *DeleteDeviceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDeviceResponse
		var err error
		defer close(result)
		response, err = client.DeleteDevice(request)
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

// DeleteDeviceRequest is the request struct for api DeleteDevice
type DeleteDeviceRequest struct {
	*requests.RpcRequest
	GroupId string `position:"Body" name:"GroupId"`
	SN      string `position:"Body" name:"SN"`
}

// DeleteDeviceResponse is the response struct for api DeleteDevice
type DeleteDeviceResponse struct {
	*responses.BaseResponse
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDeviceRequest creates a request to invoke DeleteDevice API
func CreateDeleteDeviceRequest() (request *DeleteDeviceRequest) {
	request = &DeleteDeviceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "DeleteDevice", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDeviceResponse creates a response to parse from DeleteDevice response
func CreateDeleteDeviceResponse() (response *DeleteDeviceResponse) {
	response = &DeleteDeviceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
