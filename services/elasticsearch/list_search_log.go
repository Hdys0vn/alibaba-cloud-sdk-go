package elasticsearch

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

// ListSearchLog invokes the elasticsearch.ListSearchLog API synchronously
func (client *Client) ListSearchLog(request *ListSearchLogRequest) (response *ListSearchLogResponse, err error) {
	response = CreateListSearchLogResponse()
	err = client.DoAction(request, response)
	return
}

// ListSearchLogWithChan invokes the elasticsearch.ListSearchLog API asynchronously
func (client *Client) ListSearchLogWithChan(request *ListSearchLogRequest) (<-chan *ListSearchLogResponse, <-chan error) {
	responseChan := make(chan *ListSearchLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSearchLog(request)
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

// ListSearchLogWithCallback invokes the elasticsearch.ListSearchLog API asynchronously
func (client *Client) ListSearchLogWithCallback(request *ListSearchLogRequest, callback func(response *ListSearchLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSearchLogResponse
		var err error
		defer close(result)
		response, err = client.ListSearchLog(request)
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

// ListSearchLogRequest is the request struct for api ListSearchLog
type ListSearchLogRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	Size       requests.Integer `position:"Query" name:"size"`
	Query      string           `position:"Query" name:"query"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	BeginTime  requests.Integer `position:"Query" name:"beginTime"`
	Page       requests.Integer `position:"Query" name:"page"`
	Type       string           `position:"Query" name:"type"`
}

// ListSearchLogResponse is the response struct for api ListSearchLog
type ListSearchLogResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Headers   Headers      `json:"Headers" xml:"Headers"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateListSearchLogRequest creates a request to invoke ListSearchLog API
func CreateListSearchLogRequest() (request *ListSearchLogRequest) {
	request = &ListSearchLogRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListSearchLog", "/openapi/instances/[InstanceId]/search-log", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListSearchLogResponse creates a response to parse from ListSearchLog response
func CreateListSearchLogResponse() (response *ListSearchLogResponse) {
	response = &ListSearchLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
