package nas

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

// DescribeFileSystemStatistics invokes the nas.DescribeFileSystemStatistics API synchronously
func (client *Client) DescribeFileSystemStatistics(request *DescribeFileSystemStatisticsRequest) (response *DescribeFileSystemStatisticsResponse, err error) {
	response = CreateDescribeFileSystemStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFileSystemStatisticsWithChan invokes the nas.DescribeFileSystemStatistics API asynchronously
func (client *Client) DescribeFileSystemStatisticsWithChan(request *DescribeFileSystemStatisticsRequest) (<-chan *DescribeFileSystemStatisticsResponse, <-chan error) {
	responseChan := make(chan *DescribeFileSystemStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFileSystemStatistics(request)
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

// DescribeFileSystemStatisticsWithCallback invokes the nas.DescribeFileSystemStatistics API asynchronously
func (client *Client) DescribeFileSystemStatisticsWithCallback(request *DescribeFileSystemStatisticsRequest, callback func(response *DescribeFileSystemStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFileSystemStatisticsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFileSystemStatistics(request)
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

// DescribeFileSystemStatisticsRequest is the request struct for api DescribeFileSystemStatistics
type DescribeFileSystemStatisticsRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeFileSystemStatisticsResponse is the response struct for api DescribeFileSystemStatistics
type DescribeFileSystemStatisticsResponse struct {
	*responses.BaseResponse
	RequestId            string                                    `json:"RequestId" xml:"RequestId"`
	PageSize             int                                       `json:"PageSize" xml:"PageSize"`
	PageNumber           int                                       `json:"PageNumber" xml:"PageNumber"`
	TotalCount           int                                       `json:"TotalCount" xml:"TotalCount"`
	FileSystemStatistics FileSystemStatistics                      `json:"FileSystemStatistics" xml:"FileSystemStatistics"`
	FileSystems          FileSystemsInDescribeFileSystemStatistics `json:"FileSystems" xml:"FileSystems"`
}

// CreateDescribeFileSystemStatisticsRequest creates a request to invoke DescribeFileSystemStatistics API
func CreateDescribeFileSystemStatisticsRequest() (request *DescribeFileSystemStatisticsRequest) {
	request = &DescribeFileSystemStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "DescribeFileSystemStatistics", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFileSystemStatisticsResponse creates a response to parse from DescribeFileSystemStatistics response
func CreateDescribeFileSystemStatisticsResponse() (response *DescribeFileSystemStatisticsResponse) {
	response = &DescribeFileSystemStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
