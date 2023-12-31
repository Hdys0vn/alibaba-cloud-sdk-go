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

// ListLogstashLog invokes the elasticsearch.ListLogstashLog API synchronously
func (client *Client) ListLogstashLog(request *ListLogstashLogRequest) (response *ListLogstashLogResponse, err error) {
	response = CreateListLogstashLogResponse()
	err = client.DoAction(request, response)
	return
}

// ListLogstashLogWithChan invokes the elasticsearch.ListLogstashLog API asynchronously
func (client *Client) ListLogstashLogWithChan(request *ListLogstashLogRequest) (<-chan *ListLogstashLogResponse, <-chan error) {
	responseChan := make(chan *ListLogstashLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLogstashLog(request)
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

// ListLogstashLogWithCallback invokes the elasticsearch.ListLogstashLog API asynchronously
func (client *Client) ListLogstashLogWithCallback(request *ListLogstashLogRequest, callback func(response *ListLogstashLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLogstashLogResponse
		var err error
		defer close(result)
		response, err = client.ListLogstashLog(request)
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

// ListLogstashLogRequest is the request struct for api ListLogstashLog
type ListLogstashLogRequest struct {
	*requests.RoaRequest
	InstanceId string           `position:"Path" name:"InstanceId"`
	Size       requests.Integer `position:"Query" name:"size"`
	Query      string           `position:"Query" name:"query"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	BeginTime  requests.Integer `position:"Query" name:"beginTime"`
	Page       requests.Integer `position:"Query" name:"page"`
	Type       string           `position:"Query" name:"type"`
}

// ListLogstashLogResponse is the response struct for api ListLogstashLog
type ListLogstashLogResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"RequestId" xml:"RequestId"`
	Result    []ResultItem `json:"Result" xml:"Result"`
}

// CreateListLogstashLogRequest creates a request to invoke ListLogstashLog API
func CreateListLogstashLogRequest() (request *ListLogstashLogRequest) {
	request = &ListLogstashLogRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("elasticsearch", "2017-06-13", "ListLogstashLog", "/openapi/logstashes/[InstanceId]/search-log", "elasticsearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListLogstashLogResponse creates a response to parse from ListLogstashLog response
func CreateListLogstashLogResponse() (response *ListLogstashLogResponse) {
	response = &ListLogstashLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
