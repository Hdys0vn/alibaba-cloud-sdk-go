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

// ModifyDeviceAlarm invokes the vs.ModifyDeviceAlarm API synchronously
func (client *Client) ModifyDeviceAlarm(request *ModifyDeviceAlarmRequest) (response *ModifyDeviceAlarmResponse, err error) {
	response = CreateModifyDeviceAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDeviceAlarmWithChan invokes the vs.ModifyDeviceAlarm API asynchronously
func (client *Client) ModifyDeviceAlarmWithChan(request *ModifyDeviceAlarmRequest) (<-chan *ModifyDeviceAlarmResponse, <-chan error) {
	responseChan := make(chan *ModifyDeviceAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDeviceAlarm(request)
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

// ModifyDeviceAlarmWithCallback invokes the vs.ModifyDeviceAlarm API asynchronously
func (client *Client) ModifyDeviceAlarmWithCallback(request *ModifyDeviceAlarmRequest, callback func(response *ModifyDeviceAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDeviceAlarmResponse
		var err error
		defer close(result)
		response, err = client.ModifyDeviceAlarm(request)
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

// ModifyDeviceAlarmRequest is the request struct for api ModifyDeviceAlarm
type ModifyDeviceAlarmRequest struct {
	*requests.RpcRequest
	Id        string           `position:"Query" name:"Id"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AlarmId   string           `position:"Query" name:"AlarmId"`
	ChannelId requests.Integer `position:"Query" name:"ChannelId"`
	Status    requests.Integer `position:"Query" name:"Status"`
}

// ModifyDeviceAlarmResponse is the response struct for api ModifyDeviceAlarm
type ModifyDeviceAlarmResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDeviceAlarmRequest creates a request to invoke ModifyDeviceAlarm API
func CreateModifyDeviceAlarmRequest() (request *ModifyDeviceAlarmRequest) {
	request = &ModifyDeviceAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "ModifyDeviceAlarm", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyDeviceAlarmResponse creates a response to parse from ModifyDeviceAlarm response
func CreateModifyDeviceAlarmResponse() (response *ModifyDeviceAlarmResponse) {
	response = &ModifyDeviceAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
