package polardbx

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

// DescribeInstanceDbPerformance invokes the polardbx.DescribeInstanceDbPerformance API synchronously
func (client *Client) DescribeInstanceDbPerformance(request *DescribeInstanceDbPerformanceRequest) (response *DescribeInstanceDbPerformanceResponse, err error) {
	response = CreateDescribeInstanceDbPerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstanceDbPerformanceWithChan invokes the polardbx.DescribeInstanceDbPerformance API asynchronously
func (client *Client) DescribeInstanceDbPerformanceWithChan(request *DescribeInstanceDbPerformanceRequest) (<-chan *DescribeInstanceDbPerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceDbPerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceDbPerformance(request)
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

// DescribeInstanceDbPerformanceWithCallback invokes the polardbx.DescribeInstanceDbPerformance API asynchronously
func (client *Client) DescribeInstanceDbPerformanceWithCallback(request *DescribeInstanceDbPerformanceRequest, callback func(response *DescribeInstanceDbPerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceDbPerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceDbPerformance(request)
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

// DescribeInstanceDbPerformanceRequest is the request struct for api DescribeInstanceDbPerformance
type DescribeInstanceDbPerformanceRequest struct {
	*requests.RpcRequest
	DbInstanceName string `position:"Query" name:"DbInstanceName"`
	Keys           string `position:"Query" name:"Keys"`
	EndTime        string `position:"Query" name:"EndTime"`
	StartTime      string `position:"Query" name:"StartTime"`
	DbName         string `position:"Query" name:"DbName"`
}

// DescribeInstanceDbPerformanceResponse is the response struct for api DescribeInstanceDbPerformance
type DescribeInstanceDbPerformanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeInstanceDbPerformanceRequest creates a request to invoke DescribeInstanceDbPerformance API
func CreateDescribeInstanceDbPerformanceRequest() (request *DescribeInstanceDbPerformanceRequest) {
	request = &DescribeInstanceDbPerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "DescribeInstanceDbPerformance", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstanceDbPerformanceResponse creates a response to parse from DescribeInstanceDbPerformance response
func CreateDescribeInstanceDbPerformanceResponse() (response *DescribeInstanceDbPerformanceResponse) {
	response = &DescribeInstanceDbPerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}