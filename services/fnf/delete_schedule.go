package fnf

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

// DeleteSchedule invokes the fnf.DeleteSchedule API synchronously
func (client *Client) DeleteSchedule(request *DeleteScheduleRequest) (response *DeleteScheduleResponse, err error) {
	response = CreateDeleteScheduleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteScheduleWithChan invokes the fnf.DeleteSchedule API asynchronously
func (client *Client) DeleteScheduleWithChan(request *DeleteScheduleRequest) (<-chan *DeleteScheduleResponse, <-chan error) {
	responseChan := make(chan *DeleteScheduleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSchedule(request)
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

// DeleteScheduleWithCallback invokes the fnf.DeleteSchedule API asynchronously
func (client *Client) DeleteScheduleWithCallback(request *DeleteScheduleRequest, callback func(response *DeleteScheduleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteScheduleResponse
		var err error
		defer close(result)
		response, err = client.DeleteSchedule(request)
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

// DeleteScheduleRequest is the request struct for api DeleteSchedule
type DeleteScheduleRequest struct {
	*requests.RpcRequest
	ScheduleName string `position:"Query" name:"ScheduleName"`
	RequestId    string `position:"Query" name:"RequestId"`
	FlowName     string `position:"Query" name:"FlowName"`
}

// DeleteScheduleResponse is the response struct for api DeleteSchedule
type DeleteScheduleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteScheduleRequest creates a request to invoke DeleteSchedule API
func CreateDeleteScheduleRequest() (request *DeleteScheduleRequest) {
	request = &DeleteScheduleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "DeleteSchedule", "fnf", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDeleteScheduleResponse creates a response to parse from DeleteSchedule response
func CreateDeleteScheduleResponse() (response *DeleteScheduleResponse) {
	response = &DeleteScheduleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
