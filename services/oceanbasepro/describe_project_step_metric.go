package oceanbasepro

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

// DescribeProjectStepMetric invokes the oceanbasepro.DescribeProjectStepMetric API synchronously
func (client *Client) DescribeProjectStepMetric(request *DescribeProjectStepMetricRequest) (response *DescribeProjectStepMetricResponse, err error) {
	response = CreateDescribeProjectStepMetricResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProjectStepMetricWithChan invokes the oceanbasepro.DescribeProjectStepMetric API asynchronously
func (client *Client) DescribeProjectStepMetricWithChan(request *DescribeProjectStepMetricRequest) (<-chan *DescribeProjectStepMetricResponse, <-chan error) {
	responseChan := make(chan *DescribeProjectStepMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProjectStepMetric(request)
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

// DescribeProjectStepMetricWithCallback invokes the oceanbasepro.DescribeProjectStepMetric API asynchronously
func (client *Client) DescribeProjectStepMetricWithCallback(request *DescribeProjectStepMetricRequest, callback func(response *DescribeProjectStepMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProjectStepMetricResponse
		var err error
		defer close(result)
		response, err = client.DescribeProjectStepMetric(request)
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

// DescribeProjectStepMetricRequest is the request struct for api DescribeProjectStepMetric
type DescribeProjectStepMetricRequest struct {
	*requests.RpcRequest
	MetricType     string           `position:"Body" name:"MetricType"`
	StepName       string           `position:"Body" name:"StepName"`
	Aggregator     string           `position:"Body" name:"Aggregator"`
	MaxPointNum    requests.Integer `position:"Body" name:"MaxPointNum"`
	EndTimestamp   requests.Integer `position:"Body" name:"EndTimestamp"`
	BeginTimestamp requests.Integer `position:"Body" name:"BeginTimestamp"`
	ProjectId      string           `position:"Body" name:"ProjectId"`
}

// DescribeProjectStepMetricResponse is the response struct for api DescribeProjectStepMetric
type DescribeProjectStepMetricResponse struct {
	*responses.BaseResponse
}

// CreateDescribeProjectStepMetricRequest creates a request to invoke DescribeProjectStepMetric API
func CreateDescribeProjectStepMetricRequest() (request *DescribeProjectStepMetricRequest) {
	request = &DescribeProjectStepMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeProjectStepMetric", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProjectStepMetricResponse creates a response to parse from DescribeProjectStepMetric response
func CreateDescribeProjectStepMetricResponse() (response *DescribeProjectStepMetricResponse) {
	response = &DescribeProjectStepMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}