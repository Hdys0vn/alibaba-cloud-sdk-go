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

// UpdateDeviceHeartBeat invokes the aliyuncvc.UpdateDeviceHeartBeat API synchronously
func (client *Client) UpdateDeviceHeartBeat(request *UpdateDeviceHeartBeatRequest) (response *UpdateDeviceHeartBeatResponse, err error) {
	response = CreateUpdateDeviceHeartBeatResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDeviceHeartBeatWithChan invokes the aliyuncvc.UpdateDeviceHeartBeat API asynchronously
func (client *Client) UpdateDeviceHeartBeatWithChan(request *UpdateDeviceHeartBeatRequest) (<-chan *UpdateDeviceHeartBeatResponse, <-chan error) {
	responseChan := make(chan *UpdateDeviceHeartBeatResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDeviceHeartBeat(request)
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

// UpdateDeviceHeartBeatWithCallback invokes the aliyuncvc.UpdateDeviceHeartBeat API asynchronously
func (client *Client) UpdateDeviceHeartBeatWithCallback(request *UpdateDeviceHeartBeatRequest, callback func(response *UpdateDeviceHeartBeatResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDeviceHeartBeatResponse
		var err error
		defer close(result)
		response, err = client.UpdateDeviceHeartBeat(request)
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

// UpdateDeviceHeartBeatRequest is the request struct for api UpdateDeviceHeartBeat
type UpdateDeviceHeartBeatRequest struct {
	*requests.RpcRequest
	Message string `position:"Query" name:"Message"`
}

// UpdateDeviceHeartBeatResponse is the response struct for api UpdateDeviceHeartBeat
type UpdateDeviceHeartBeatResponse struct {
	*responses.BaseResponse
	ErrorCode  int        `json:"ErrorCode" xml:"ErrorCode"`
	Message    string     `json:"Message" xml:"Message"`
	Success    bool       `json:"Success" xml:"Success"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	DeviceInfo DeviceInfo `json:"DeviceInfo" xml:"DeviceInfo"`
}

// CreateUpdateDeviceHeartBeatRequest creates a request to invoke UpdateDeviceHeartBeat API
func CreateUpdateDeviceHeartBeatRequest() (request *UpdateDeviceHeartBeatRequest) {
	request = &UpdateDeviceHeartBeatRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aliyuncvc", "2019-10-30", "UpdateDeviceHeartBeat", "aliyuncvc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDeviceHeartBeatResponse creates a response to parse from UpdateDeviceHeartBeat response
func CreateUpdateDeviceHeartBeatResponse() (response *UpdateDeviceHeartBeatResponse) {
	response = &UpdateDeviceHeartBeatResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
