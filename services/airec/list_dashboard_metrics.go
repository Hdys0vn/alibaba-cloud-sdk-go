package airec

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

// ListDashboardMetrics invokes the airec.ListDashboardMetrics API synchronously
func (client *Client) ListDashboardMetrics(request *ListDashboardMetricsRequest) (response *ListDashboardMetricsResponse, err error) {
	response = CreateListDashboardMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDashboardMetricsWithChan invokes the airec.ListDashboardMetrics API asynchronously
func (client *Client) ListDashboardMetricsWithChan(request *ListDashboardMetricsRequest) (<-chan *ListDashboardMetricsResponse, <-chan error) {
	responseChan := make(chan *ListDashboardMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDashboardMetrics(request)
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

// ListDashboardMetricsWithCallback invokes the airec.ListDashboardMetrics API asynchronously
func (client *Client) ListDashboardMetricsWithCallback(request *ListDashboardMetricsRequest, callback func(response *ListDashboardMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDashboardMetricsResponse
		var err error
		defer close(result)
		response, err = client.ListDashboardMetrics(request)
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

// ListDashboardMetricsRequest is the request struct for api ListDashboardMetrics
type ListDashboardMetricsRequest struct {
	*requests.RoaRequest
	MetricType string           `position:"Query" name:"metricType"`
	InstanceId string           `position:"Path" name:"instanceId"`
	EndTime    requests.Integer `position:"Query" name:"endTime"`
	StartTime  requests.Integer `position:"Query" name:"startTime"`
}

// ListDashboardMetricsResponse is the response struct for api ListDashboardMetrics
type ListDashboardMetricsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"requestId" xml:"requestId"`
	Code      string       `json:"code" xml:"code"`
	Message   string       `json:"message" xml:"message"`
	Result    []ResultItem `json:"result" xml:"result"`
}

// CreateListDashboardMetricsRequest creates a request to invoke ListDashboardMetrics API
func CreateListDashboardMetricsRequest() (request *ListDashboardMetricsRequest) {
	request = &ListDashboardMetricsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "ListDashboardMetrics", "/v2/openapi/instances/[instanceId]/dashboard/metrics", "airec", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListDashboardMetricsResponse creates a response to parse from ListDashboardMetrics response
func CreateListDashboardMetricsResponse() (response *ListDashboardMetricsResponse) {
	response = &ListDashboardMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
