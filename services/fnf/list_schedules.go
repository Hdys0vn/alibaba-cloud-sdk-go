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

// ListSchedules invokes the fnf.ListSchedules API synchronously
func (client *Client) ListSchedules(request *ListSchedulesRequest) (response *ListSchedulesResponse, err error) {
	response = CreateListSchedulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListSchedulesWithChan invokes the fnf.ListSchedules API asynchronously
func (client *Client) ListSchedulesWithChan(request *ListSchedulesRequest) (<-chan *ListSchedulesResponse, <-chan error) {
	responseChan := make(chan *ListSchedulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSchedules(request)
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

// ListSchedulesWithCallback invokes the fnf.ListSchedules API asynchronously
func (client *Client) ListSchedulesWithCallback(request *ListSchedulesRequest, callback func(response *ListSchedulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSchedulesResponse
		var err error
		defer close(result)
		response, err = client.ListSchedules(request)
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

// ListSchedulesRequest is the request struct for api ListSchedules
type ListSchedulesRequest struct {
	*requests.RpcRequest
	NextToken string           `position:"Query" name:"NextToken"`
	RequestId string           `position:"Query" name:"RequestId"`
	Limit     requests.Integer `position:"Query" name:"Limit"`
	FlowName  string           `position:"Query" name:"FlowName"`
}

// ListSchedulesResponse is the response struct for api ListSchedules
type ListSchedulesResponse struct {
	*responses.BaseResponse
	NextToken string          `json:"NextToken" xml:"NextToken"`
	RequestId string          `json:"RequestId" xml:"RequestId"`
	Schedules []SchedulesItem `json:"Schedules" xml:"Schedules"`
}

// CreateListSchedulesRequest creates a request to invoke ListSchedules API
func CreateListSchedulesRequest() (request *ListSchedulesRequest) {
	request = &ListSchedulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("fnf", "2019-03-15", "ListSchedules", "fnf", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListSchedulesResponse creates a response to parse from ListSchedules response
func CreateListSchedulesResponse() (response *ListSchedulesResponse) {
	response = &ListSchedulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
