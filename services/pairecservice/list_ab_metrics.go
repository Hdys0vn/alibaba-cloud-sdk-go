package pairecservice

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

// ListABMetrics invokes the pairecservice.ListABMetrics API synchronously
func (client *Client) ListABMetrics(request *ListABMetricsRequest) (response *ListABMetricsResponse, err error) {
	response = CreateListABMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// ListABMetricsWithChan invokes the pairecservice.ListABMetrics API asynchronously
func (client *Client) ListABMetricsWithChan(request *ListABMetricsRequest) (<-chan *ListABMetricsResponse, <-chan error) {
	responseChan := make(chan *ListABMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListABMetrics(request)
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

// ListABMetricsWithCallback invokes the pairecservice.ListABMetrics API asynchronously
func (client *Client) ListABMetricsWithCallback(request *ListABMetricsRequest, callback func(response *ListABMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListABMetricsResponse
		var err error
		defer close(result)
		response, err = client.ListABMetrics(request)
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

// ListABMetricsRequest is the request struct for api ListABMetrics
type ListABMetricsRequest struct {
	*requests.RoaRequest
	Realtime    requests.Boolean `position:"Query" name:"Realtime"`
	InstanceId  string           `position:"Query" name:"InstanceId"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	Name        string           `position:"Query" name:"Name"`
	SceneId     string           `position:"Query" name:"SceneId"`
	Type        string           `position:"Query" name:"Type"`
	PageNumber  requests.Integer `position:"Query" name:"PageNumber"`
	TableMetaId string           `position:"Query" name:"TableMetaId"`
}

// ListABMetricsResponse is the response struct for api ListABMetrics
type ListABMetricsResponse struct {
	*responses.BaseResponse
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	TotalCount int64           `json:"TotalCount" xml:"TotalCount"`
	ABMetrics  []ABMetricsItem `json:"ABMetrics" xml:"ABMetrics"`
}

// CreateListABMetricsRequest creates a request to invoke ListABMetrics API
func CreateListABMetricsRequest() (request *ListABMetricsRequest) {
	request = &ListABMetricsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "ListABMetrics", "/api/v1/abmetrics", "", "")
	request.Method = requests.GET
	return
}

// CreateListABMetricsResponse creates a response to parse from ListABMetrics response
func CreateListABMetricsResponse() (response *ListABMetricsResponse) {
	response = &ListABMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}