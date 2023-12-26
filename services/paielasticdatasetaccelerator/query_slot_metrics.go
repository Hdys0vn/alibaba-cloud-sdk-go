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

// QuerySlotMetrics invokes the paielasticdatasetaccelerator.QuerySlotMetrics API synchronously
func (client *Client) QuerySlotMetrics(request *QuerySlotMetricsRequest) (response *QuerySlotMetricsResponse, err error) {
	response = CreateQuerySlotMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// QuerySlotMetricsWithChan invokes the paielasticdatasetaccelerator.QuerySlotMetrics API asynchronously
func (client *Client) QuerySlotMetricsWithChan(request *QuerySlotMetricsRequest) (<-chan *QuerySlotMetricsResponse, <-chan error) {
	responseChan := make(chan *QuerySlotMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QuerySlotMetrics(request)
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

// QuerySlotMetricsWithCallback invokes the paielasticdatasetaccelerator.QuerySlotMetrics API asynchronously
func (client *Client) QuerySlotMetricsWithCallback(request *QuerySlotMetricsRequest, callback func(response *QuerySlotMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QuerySlotMetricsResponse
		var err error
		defer close(result)
		response, err = client.QuerySlotMetrics(request)
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

// QuerySlotMetricsRequest is the request struct for api QuerySlotMetrics
type QuerySlotMetricsRequest struct {
	*requests.RoaRequest
	MetricType string `position:"Query" name:"MetricType"`
	TimeStep   string `position:"Query" name:"TimeStep"`
	EndTime    string `position:"Query" name:"EndTime"`
	SlotId     string `position:"Path" name:"SlotId"`
	StartTime  string `position:"Query" name:"StartTime"`
	Dimensions string `position:"Query" name:"Dimensions"`
}

// QuerySlotMetricsResponse is the response struct for api QuerySlotMetrics
type QuerySlotMetricsResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Period    string        `json:"Period" xml:"Period"`
	Metrics   []MetricsItem `json:"Metrics" xml:"Metrics"`
}

// CreateQuerySlotMetricsRequest creates a request to invoke QuerySlotMetrics API
func CreateQuerySlotMetricsRequest() (request *QuerySlotMetricsRequest) {
	request = &QuerySlotMetricsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PAIElasticDatasetAccelerator", "2022-08-01", "QuerySlotMetrics", "/api/v1/slots/[SlotId]/metrics/action/query", "datasetacc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQuerySlotMetricsResponse creates a response to parse from QuerySlotMetrics response
func CreateQuerySlotMetricsResponse() (response *QuerySlotMetricsResponse) {
	response = &QuerySlotMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
