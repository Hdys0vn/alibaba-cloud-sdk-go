package paielasticdatasetaccelerator

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

// QueryInstanceMetrics invokes the paielasticdatasetaccelerator.QueryInstanceMetrics API synchronously
func (client *Client) QueryInstanceMetrics(request *QueryInstanceMetricsRequest) (response *QueryInstanceMetricsResponse, err error) {
	response = CreateQueryInstanceMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInstanceMetricsWithChan invokes the paielasticdatasetaccelerator.QueryInstanceMetrics API asynchronously
func (client *Client) QueryInstanceMetricsWithChan(request *QueryInstanceMetricsRequest) (<-chan *QueryInstanceMetricsResponse, <-chan error) {
	responseChan := make(chan *QueryInstanceMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInstanceMetrics(request)
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

// QueryInstanceMetricsWithCallback invokes the paielasticdatasetaccelerator.QueryInstanceMetrics API asynchronously
func (client *Client) QueryInstanceMetricsWithCallback(request *QueryInstanceMetricsRequest, callback func(response *QueryInstanceMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInstanceMetricsResponse
		var err error
		defer close(result)
		response, err = client.QueryInstanceMetrics(request)
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

// QueryInstanceMetricsRequest is the request struct for api QueryInstanceMetrics
type QueryInstanceMetricsRequest struct {
	*requests.RoaRequest
	MetricType string `position:"Query" name:"MetricType"`
	InstanceId string `position:"Path" name:"InstanceId"`
	TimeStep   string `position:"Query" name:"TimeStep"`
	EndTime    string `position:"Query" name:"EndTime"`
	StartTime  string `position:"Query" name:"StartTime"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

// QueryInstanceMetricsResponse is the response struct for api QueryInstanceMetrics
type QueryInstanceMetricsResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Period    string        `json:"Period" xml:"Period"`
	Metrics   []MetricsItem `json:"Metrics" xml:"Metrics"`
}

// CreateQueryInstanceMetricsRequest creates a request to invoke QueryInstanceMetrics API
func CreateQueryInstanceMetricsRequest() (request *QueryInstanceMetricsRequest) {
	request = &QueryInstanceMetricsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "QueryInstanceMetrics", "/api/v1/instances/[InstanceId]/metrics/action/query", "datasetacc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryInstanceMetricsResponse creates a response to parse from QueryInstanceMetrics response
func CreateQueryInstanceMetricsResponse() (response *QueryInstanceMetricsResponse) {
	response = &QueryInstanceMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}