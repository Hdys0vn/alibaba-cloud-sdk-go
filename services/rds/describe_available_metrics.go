package rds

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

// DescribeAvailableMetrics invokes the rds.DescribeAvailableMetrics API synchronously
func (client *Client) DescribeAvailableMetrics(request *DescribeAvailableMetricsRequest) (response *DescribeAvailableMetricsResponse, err error) {
	response = CreateDescribeAvailableMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAvailableMetricsWithChan invokes the rds.DescribeAvailableMetrics API asynchronously
func (client *Client) DescribeAvailableMetricsWithChan(request *DescribeAvailableMetricsRequest) (<-chan *DescribeAvailableMetricsResponse, <-chan error) {
	responseChan := make(chan *DescribeAvailableMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAvailableMetrics(request)
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

// DescribeAvailableMetricsWithCallback invokes the rds.DescribeAvailableMetrics API asynchronously
func (client *Client) DescribeAvailableMetricsWithCallback(request *DescribeAvailableMetricsRequest, callback func(response *DescribeAvailableMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAvailableMetricsResponse
		var err error
		defer close(result)
		response, err = client.DescribeAvailableMetrics(request)
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

// DescribeAvailableMetricsRequest is the request struct for api DescribeAvailableMetrics
type DescribeAvailableMetricsRequest struct {
	*requests.RpcRequest
	DBInstanceName       string           `position:"Query" name:"DBInstanceName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
}

// DescribeAvailableMetricsResponse is the response struct for api DescribeAvailableMetrics
type DescribeAvailableMetricsResponse struct {
	*responses.BaseResponse
	TotalRecordCount int       `json:"TotalRecordCount" xml:"TotalRecordCount"`
	RequestId        string    `json:"RequestId" xml:"RequestId"`
	DBInstanceName   string    `json:"DBInstanceName" xml:"DBInstanceName"`
	Items            []Metrics `json:"Items" xml:"Items"`
}

// CreateDescribeAvailableMetricsRequest creates a request to invoke DescribeAvailableMetrics API
func CreateDescribeAvailableMetricsRequest() (request *DescribeAvailableMetricsRequest) {
	request = &DescribeAvailableMetricsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeAvailableMetrics", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeAvailableMetricsResponse creates a response to parse from DescribeAvailableMetrics response
func CreateDescribeAvailableMetricsResponse() (response *DescribeAvailableMetricsResponse) {
	response = &DescribeAvailableMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
