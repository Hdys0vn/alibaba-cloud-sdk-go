package eas

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

// DescribeServiceLog invokes the eas.DescribeServiceLog API synchronously
func (client *Client) DescribeServiceLog(request *DescribeServiceLogRequest) (response *DescribeServiceLogResponse, err error) {
	response = CreateDescribeServiceLogResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServiceLogWithChan invokes the eas.DescribeServiceLog API asynchronously
func (client *Client) DescribeServiceLogWithChan(request *DescribeServiceLogRequest) (<-chan *DescribeServiceLogResponse, <-chan error) {
	responseChan := make(chan *DescribeServiceLogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeServiceLog(request)
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

// DescribeServiceLogWithCallback invokes the eas.DescribeServiceLog API asynchronously
func (client *Client) DescribeServiceLogWithCallback(request *DescribeServiceLogRequest, callback func(response *DescribeServiceLogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServiceLogResponse
		var err error
		defer close(result)
		response, err = client.DescribeServiceLog(request)
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

// DescribeServiceLogRequest is the request struct for api DescribeServiceLog
type DescribeServiceLogRequest struct {
	*requests.RoaRequest
	Ip          string           `position:"Query" name:"Ip"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	EndTime     string           `position:"Query" name:"EndTime"`
	ServiceName string           `position:"Path" name:"ServiceName"`
	StartTime   string           `position:"Query" name:"StartTime"`
	ClusterId   string           `position:"Path" name:"ClusterId"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	Keyword     string           `position:"Query" name:"Keyword"`
}

// DescribeServiceLogResponse is the response struct for api DescribeServiceLog
type DescribeServiceLogResponse struct {
	*responses.BaseResponse
	RequestId    string   `json:"RequestId" xml:"RequestId"`
	PageNum      int64    `json:"PageNum" xml:"PageNum"`
	TotalCount   int64    `json:"TotalCount" xml:"TotalCount"`
	TotalPageNum int64    `json:"TotalPageNum" xml:"TotalPageNum"`
	Logs         []string `json:"Logs" xml:"Logs"`
}

// CreateDescribeServiceLogRequest creates a request to invoke DescribeServiceLog API
func CreateDescribeServiceLogRequest() (request *DescribeServiceLogRequest) {
	request = &DescribeServiceLogRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("eas", "2021-07-01", "DescribeServiceLog", "/api/v2/services/[ClusterId]/[ServiceName]/logs", "eas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeServiceLogResponse creates a response to parse from DescribeServiceLog response
func CreateDescribeServiceLogResponse() (response *DescribeServiceLogResponse) {
	response = &DescribeServiceLogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
