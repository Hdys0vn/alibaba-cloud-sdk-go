package swas_open

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

// DescribeDatabaseErrorLogs invokes the swas_open.DescribeDatabaseErrorLogs API synchronously
func (client *Client) DescribeDatabaseErrorLogs(request *DescribeDatabaseErrorLogsRequest) (response *DescribeDatabaseErrorLogsResponse, err error) {
	response = CreateDescribeDatabaseErrorLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDatabaseErrorLogsWithChan invokes the swas_open.DescribeDatabaseErrorLogs API asynchronously
func (client *Client) DescribeDatabaseErrorLogsWithChan(request *DescribeDatabaseErrorLogsRequest) (<-chan *DescribeDatabaseErrorLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeDatabaseErrorLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDatabaseErrorLogs(request)
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

// DescribeDatabaseErrorLogsWithCallback invokes the swas_open.DescribeDatabaseErrorLogs API asynchronously
func (client *Client) DescribeDatabaseErrorLogsWithCallback(request *DescribeDatabaseErrorLogsRequest, callback func(response *DescribeDatabaseErrorLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDatabaseErrorLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDatabaseErrorLogs(request)
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

// DescribeDatabaseErrorLogsRequest is the request struct for api DescribeDatabaseErrorLogs
type DescribeDatabaseErrorLogsRequest struct {
	*requests.RpcRequest
	DatabaseInstanceId string           `position:"Query" name:"DatabaseInstanceId"`
	EndTime            string           `position:"Query" name:"EndTime"`
	StartTime          string           `position:"Query" name:"StartTime"`
	PageNumber         requests.Integer `position:"Query" name:"PageNumber"`
	PageSize           requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeDatabaseErrorLogsResponse is the response struct for api DescribeDatabaseErrorLogs
type DescribeDatabaseErrorLogsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	ErrorLogs  []ErrorLog `json:"ErrorLogs" xml:"ErrorLogs"`
}

// CreateDescribeDatabaseErrorLogsRequest creates a request to invoke DescribeDatabaseErrorLogs API
func CreateDescribeDatabaseErrorLogsRequest() (request *DescribeDatabaseErrorLogsRequest) {
	request = &DescribeDatabaseErrorLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("SWAS-OPEN", "2020-06-01", "DescribeDatabaseErrorLogs", "SWAS-OPEN", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDatabaseErrorLogsResponse creates a response to parse from DescribeDatabaseErrorLogs response
func CreateDescribeDatabaseErrorLogsResponse() (response *DescribeDatabaseErrorLogsResponse) {
	response = &DescribeDatabaseErrorLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
