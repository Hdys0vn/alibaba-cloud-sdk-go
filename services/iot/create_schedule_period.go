package iot

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

// CreateSchedulePeriod invokes the iot.CreateSchedulePeriod API synchronously
func (client *Client) CreateSchedulePeriod(request *CreateSchedulePeriodRequest) (response *CreateSchedulePeriodResponse, err error) {
	response = CreateCreateSchedulePeriodResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSchedulePeriodWithChan invokes the iot.CreateSchedulePeriod API asynchronously
func (client *Client) CreateSchedulePeriodWithChan(request *CreateSchedulePeriodRequest) (<-chan *CreateSchedulePeriodResponse, <-chan error) {
	responseChan := make(chan *CreateSchedulePeriodResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSchedulePeriod(request)
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

// CreateSchedulePeriodWithCallback invokes the iot.CreateSchedulePeriod API asynchronously
func (client *Client) CreateSchedulePeriodWithCallback(request *CreateSchedulePeriodRequest, callback func(response *CreateSchedulePeriodResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSchedulePeriodResponse
		var err error
		defer close(result)
		response, err = client.CreateSchedulePeriod(request)
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

// CreateSchedulePeriodRequest is the request struct for api CreateSchedulePeriod
type CreateSchedulePeriodRequest struct {
	*requests.RpcRequest
	ScheduleCode     string `position:"Body" name:"ScheduleCode"`
	Description      string `position:"Body" name:"Description"`
	StartTime        string `position:"Body" name:"StartTime"`
	IotInstanceId    string `position:"Body" name:"IotInstanceId"`
	EndTime          string `position:"Body" name:"EndTime"`
	SoundCodeContent string `position:"Body" name:"SoundCodeContent"`
	ApiProduct       string `position:"Body" name:"ApiProduct"`
	ApiRevision      string `position:"Body" name:"ApiRevision"`
}

// CreateSchedulePeriodResponse is the response struct for api CreateSchedulePeriod
type CreateSchedulePeriodResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         string `json:"Data" xml:"Data"`
}

// CreateCreateSchedulePeriodRequest creates a request to invoke CreateSchedulePeriod API
func CreateCreateSchedulePeriodRequest() (request *CreateSchedulePeriodRequest) {
	request = &CreateSchedulePeriodRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "CreateSchedulePeriod", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateSchedulePeriodResponse creates a response to parse from CreateSchedulePeriod response
func CreateCreateSchedulePeriodResponse() (response *CreateSchedulePeriodResponse) {
	response = &CreateSchedulePeriodResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
