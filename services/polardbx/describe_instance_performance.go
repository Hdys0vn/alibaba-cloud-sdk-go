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

// DescribeInstancePerformance invokes the polardbx.DescribeInstancePerformance API synchronously
func (client *Client) DescribeInstancePerformance(request *DescribeInstancePerformanceRequest) (response *DescribeInstancePerformanceResponse, err error) {
	response = CreateDescribeInstancePerformanceResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeInstancePerformanceWithChan invokes the polardbx.DescribeInstancePerformance API asynchronously
func (client *Client) DescribeInstancePerformanceWithChan(request *DescribeInstancePerformanceRequest) (<-chan *DescribeInstancePerformanceResponse, <-chan error) {
	responseChan := make(chan *DescribeInstancePerformanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstancePerformance(request)
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

// DescribeInstancePerformanceWithCallback invokes the polardbx.DescribeInstancePerformance API asynchronously
func (client *Client) DescribeInstancePerformanceWithCallback(request *DescribeInstancePerformanceRequest, callback func(response *DescribeInstancePerformanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstancePerformanceResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstancePerformance(request)
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

// DescribeInstancePerformanceRequest is the request struct for api DescribeInstancePerformance
type DescribeInstancePerformanceRequest struct {
	*requests.RpcRequest
	DbInstanceName string `position:"Query" name:"DbInstanceName"`
	Keys           string `position:"Query" name:"Keys"`
	EndTime        string `position:"Query" name:"EndTime"`
	StartTime      string `position:"Query" name:"StartTime"`
	NodeId         string `position:"Query" name:"NodeId"`
}

// DescribeInstancePerformanceResponse is the response struct for api DescribeInstancePerformance
type DescribeInstancePerformanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeInstancePerformanceRequest creates a request to invoke DescribeInstancePerformance API
func CreateDescribeInstancePerformanceRequest() (request *DescribeInstancePerformanceRequest) {
	request = &DescribeInstancePerformanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardbx", "2020-02-02", "DescribeInstancePerformance", "polardbx", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeInstancePerformanceResponse creates a response to parse from DescribeInstancePerformance response
func CreateDescribeInstancePerformanceResponse() (response *DescribeInstancePerformanceResponse) {
	response = &DescribeInstancePerformanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
