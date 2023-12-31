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

// CreateDeviceAlarm invokes the vs.CreateDeviceAlarm API synchronously
func (client *Client) CreateDeviceAlarm(request *CreateDeviceAlarmRequest) (response *CreateDeviceAlarmResponse, err error) {
	response = CreateCreateDeviceAlarmResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDeviceAlarmWithChan invokes the vs.CreateDeviceAlarm API asynchronously
func (client *Client) CreateDeviceAlarmWithChan(request *CreateDeviceAlarmRequest) (<-chan *CreateDeviceAlarmResponse, <-chan error) {
	responseChan := make(chan *CreateDeviceAlarmResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDeviceAlarm(request)
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

// CreateDeviceAlarmWithCallback invokes the vs.CreateDeviceAlarm API asynchronously
func (client *Client) CreateDeviceAlarmWithCallback(request *CreateDeviceAlarmRequest, callback func(response *CreateDeviceAlarmResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDeviceAlarmResponse
		var err error
		defer close(result)
		response, err = client.CreateDeviceAlarm(request)
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

// CreateDeviceAlarmRequest is the request struct for api CreateDeviceAlarm
type CreateDeviceAlarmRequest struct {
	*requests.RpcRequest
	StartTime  requests.Integer `position:"Query" name:"StartTime"`
	Alarm      requests.Integer `position:"Query" name:"Alarm"`
	Id         string           `position:"Query" name:"Id"`
	ObjectType requests.Integer `position:"Query" name:"ObjectType"`
	ShowLog    string           `position:"Query" name:"ShowLog"`
	EndTime    requests.Integer `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	SubAlarm   requests.Integer `position:"Query" name:"SubAlarm"`
	Expire     requests.Integer `position:"Query" name:"Expire"`
	ChannelId  requests.Integer `position:"Query" name:"ChannelId"`
}

// CreateDeviceAlarmResponse is the response struct for api CreateDeviceAlarm
type CreateDeviceAlarmResponse struct {
	*responses.BaseResponse
	Url        string `json:"Url" xml:"Url"`
	AlarmId    string `json:"AlarmId" xml:"AlarmId"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Expire     int64  `json:"Expire" xml:"Expire"`
	AlarmDelay int64  `json:"AlarmDelay" xml:"AlarmDelay"`
}

// CreateCreateDeviceAlarmRequest creates a request to invoke CreateDeviceAlarm API
func CreateCreateDeviceAlarmRequest() (request *CreateDeviceAlarmRequest) {
	request = &CreateDeviceAlarmRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "CreateDeviceAlarm", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDeviceAlarmResponse creates a response to parse from CreateDeviceAlarm response
func CreateCreateDeviceAlarmResponse() (response *CreateDeviceAlarmResponse) {
	response = &CreateDeviceAlarmResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
