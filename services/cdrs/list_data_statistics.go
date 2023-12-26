package cdrs

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

// ListDataStatistics invokes the cdrs.ListDataStatistics API synchronously
func (client *Client) ListDataStatistics(request *ListDataStatisticsRequest) (response *ListDataStatisticsResponse, err error) {
	response = CreateListDataStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataStatisticsWithChan invokes the cdrs.ListDataStatistics API asynchronously
func (client *Client) ListDataStatisticsWithChan(request *ListDataStatisticsRequest) (<-chan *ListDataStatisticsResponse, <-chan error) {
	responseChan := make(chan *ListDataStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataStatistics(request)
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

// ListDataStatisticsWithCallback invokes the cdrs.ListDataStatistics API asynchronously
func (client *Client) ListDataStatisticsWithCallback(request *ListDataStatisticsRequest, callback func(response *ListDataStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataStatisticsResponse
		var err error
		defer close(result)
		response, err = client.ListDataStatistics(request)
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

// ListDataStatisticsRequest is the request struct for api ListDataStatistics
type ListDataStatisticsRequest struct {
	*requests.RpcRequest
	Schema       string `position:"Body" name:"Schema"`
	BackCategory string `position:"Body" name:"BackCategory"`
}

// ListDataStatisticsResponse is the response struct for api ListDataStatistics
type ListDataStatisticsResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListDataStatisticsRequest creates a request to invoke ListDataStatistics API
func CreateListDataStatisticsRequest() (request *ListDataStatisticsRequest) {
	request = &ListDataStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListDataStatistics", "", "")
	request.Method = requests.POST
	return
}

// CreateListDataStatisticsResponse creates a response to parse from ListDataStatistics response
func CreateListDataStatisticsResponse() (response *ListDataStatisticsResponse) {
	response = &ListDataStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}