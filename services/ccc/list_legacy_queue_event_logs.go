package ccc

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

// ListLegacyQueueEventLogs invokes the ccc.ListLegacyQueueEventLogs API synchronously
func (client *Client) ListLegacyQueueEventLogs(request *ListLegacyQueueEventLogsRequest) (response *ListLegacyQueueEventLogsResponse, err error) {
	response = CreateListLegacyQueueEventLogsResponse()
	err = client.DoAction(request, response)
	return
}

// ListLegacyQueueEventLogsWithChan invokes the ccc.ListLegacyQueueEventLogs API asynchronously
func (client *Client) ListLegacyQueueEventLogsWithChan(request *ListLegacyQueueEventLogsRequest) (<-chan *ListLegacyQueueEventLogsResponse, <-chan error) {
	responseChan := make(chan *ListLegacyQueueEventLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLegacyQueueEventLogs(request)
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

// ListLegacyQueueEventLogsWithCallback invokes the ccc.ListLegacyQueueEventLogs API asynchronously
func (client *Client) ListLegacyQueueEventLogsWithCallback(request *ListLegacyQueueEventLogsRequest, callback func(response *ListLegacyQueueEventLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLegacyQueueEventLogsResponse
		var err error
		defer close(result)
		response, err = client.ListLegacyQueueEventLogs(request)
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

// ListLegacyQueueEventLogsRequest is the request struct for api ListLegacyQueueEventLogs
type ListLegacyQueueEventLogsRequest struct {
	*requests.RpcRequest
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	StartTime    requests.Integer `position:"Query" name:"StartTime"`
	PageNumber   requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	SkillGroupId string           `position:"Query" name:"SkillGroupId"`
	PageSize     requests.Integer `position:"Query" name:"PageSize"`
}

// ListLegacyQueueEventLogsResponse is the response struct for api ListLegacyQueueEventLogs
type ListLegacyQueueEventLogsResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string                         `json:"Code" xml:"Code"`
	Message        string                         `json:"Message" xml:"Message"`
	RequestId      string                         `json:"RequestId" xml:"RequestId"`
	Success        bool                           `json:"Success" xml:"Success"`
	Data           DataInListLegacyQueueEventLogs `json:"Data" xml:"Data"`
}

// CreateListLegacyQueueEventLogsRequest creates a request to invoke ListLegacyQueueEventLogs API
func CreateListLegacyQueueEventLogsRequest() (request *ListLegacyQueueEventLogsRequest) {
	request = &ListLegacyQueueEventLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListLegacyQueueEventLogs", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLegacyQueueEventLogsResponse creates a response to parse from ListLegacyQueueEventLogs response
func CreateListLegacyQueueEventLogsResponse() (response *ListLegacyQueueEventLogsResponse) {
	response = &ListLegacyQueueEventLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
