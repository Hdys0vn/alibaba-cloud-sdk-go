package aegis

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

// DescribeLoginLogs invokes the aegis.DescribeLoginLogs API synchronously
// api document: https://help.aliyun.com/api/aegis/describeloginlogs.html
func (client *Client) DescribeLoginLogs(request *DescribeLoginLogsRequest) (response *DescribeLoginLogsResponse, err error) {
	response = CreateDescribeLoginLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLoginLogsWithChan invokes the aegis.DescribeLoginLogs API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeloginlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoginLogsWithChan(request *DescribeLoginLogsRequest) (<-chan *DescribeLoginLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeLoginLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLoginLogs(request)
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

// DescribeLoginLogsWithCallback invokes the aegis.DescribeLoginLogs API asynchronously
// api document: https://help.aliyun.com/api/aegis/describeloginlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLoginLogsWithCallback(request *DescribeLoginLogsRequest, callback func(response *DescribeLoginLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLoginLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeLoginLogs(request)
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

// DescribeLoginLogsRequest is the request struct for api DescribeLoginLogs
type DescribeLoginLogsRequest struct {
	*requests.RpcRequest
	Types       string           `position:"Query" name:"Types"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Statuses    string           `position:"Query" name:"Statuses"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	Remark      string           `position:"Query" name:"Remark"`
	Tag         string           `position:"Query" name:"Tag"`
}

// DescribeLoginLogsResponse is the response struct for api DescribeLoginLogs
type DescribeLoginLogsResponse struct {
	*responses.BaseResponse
	RequestId   string        `json:"RequestId" xml:"RequestId"`
	PageSize    int           `json:"PageSize" xml:"PageSize"`
	CurrentPage int           `json:"CurrentPage" xml:"CurrentPage"`
	TotalCount  int           `json:"TotalCount" xml:"TotalCount"`
	LogList     []LogListItem `json:"LogList" xml:"LogList"`
}

// CreateDescribeLoginLogsRequest creates a request to invoke DescribeLoginLogs API
func CreateDescribeLoginLogsRequest() (request *DescribeLoginLogsRequest) {
	request = &DescribeLoginLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeLoginLogs", "vipaegis", "openAPI")
	return
}

// CreateDescribeLoginLogsResponse creates a response to parse from DescribeLoginLogs response
func CreateDescribeLoginLogsResponse() (response *DescribeLoginLogsResponse) {
	response = &DescribeLoginLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
