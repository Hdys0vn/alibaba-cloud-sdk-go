package das

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

// DescribeAutoScalingHistory invokes the das.DescribeAutoScalingHistory API synchronously
func (client *Client) DescribeAutoScalingHistory(request *DescribeAutoScalingHistoryRequest) (response *DescribeAutoScalingHistoryResponse, err error) {
	response = CreateDescribeAutoScalingHistoryResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeAutoScalingHistoryWithChan invokes the das.DescribeAutoScalingHistory API asynchronously
func (client *Client) DescribeAutoScalingHistoryWithChan(request *DescribeAutoScalingHistoryRequest) (<-chan *DescribeAutoScalingHistoryResponse, <-chan error) {
	responseChan := make(chan *DescribeAutoScalingHistoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeAutoScalingHistory(request)
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

// DescribeAutoScalingHistoryWithCallback invokes the das.DescribeAutoScalingHistory API asynchronously
func (client *Client) DescribeAutoScalingHistoryWithCallback(request *DescribeAutoScalingHistoryRequest, callback func(response *DescribeAutoScalingHistoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeAutoScalingHistoryResponse
		var err error
		defer close(result)
		response, err = client.DescribeAutoScalingHistory(request)
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

// DescribeAutoScalingHistoryRequest is the request struct for api DescribeAutoScalingHistory
type DescribeAutoScalingHistoryRequest struct {
	*requests.RpcRequest
	InstanceId          string           `position:"Query" name:"InstanceId"`
	AutoScalingTaskType string           `position:"Query" name:"AutoScalingTaskType"`
	StartTime           requests.Integer `position:"Query" name:"StartTime"`
	EndTime             requests.Integer `position:"Query" name:"EndTime"`
}

// DescribeAutoScalingHistoryResponse is the response struct for api DescribeAutoScalingHistory
type DescribeAutoScalingHistoryResponse struct {
	*responses.BaseResponse
	Code      string                           `json:"Code" xml:"Code"`
	Message   string                           `json:"Message" xml:"Message"`
	RequestId string                           `json:"RequestId" xml:"RequestId"`
	Success   string                           `json:"Success" xml:"Success"`
	Data      DataInDescribeAutoScalingHistory `json:"Data" xml:"Data"`
}

// CreateDescribeAutoScalingHistoryRequest creates a request to invoke DescribeAutoScalingHistory API
func CreateDescribeAutoScalingHistoryRequest() (request *DescribeAutoScalingHistoryRequest) {
	request = &DescribeAutoScalingHistoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "DescribeAutoScalingHistory", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeAutoScalingHistoryResponse creates a response to parse from DescribeAutoScalingHistory response
func CreateDescribeAutoScalingHistoryResponse() (response *DescribeAutoScalingHistoryResponse) {
	response = &DescribeAutoScalingHistoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
