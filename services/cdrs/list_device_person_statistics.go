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

// ListDevicePersonStatistics invokes the cdrs.ListDevicePersonStatistics API synchronously
func (client *Client) ListDevicePersonStatistics(request *ListDevicePersonStatisticsRequest) (response *ListDevicePersonStatisticsResponse, err error) {
	response = CreateListDevicePersonStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDevicePersonStatisticsWithChan invokes the cdrs.ListDevicePersonStatistics API asynchronously
func (client *Client) ListDevicePersonStatisticsWithChan(request *ListDevicePersonStatisticsRequest) (<-chan *ListDevicePersonStatisticsResponse, <-chan error) {
	responseChan := make(chan *ListDevicePersonStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDevicePersonStatistics(request)
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

// ListDevicePersonStatisticsWithCallback invokes the cdrs.ListDevicePersonStatistics API asynchronously
func (client *Client) ListDevicePersonStatisticsWithCallback(request *ListDevicePersonStatisticsRequest, callback func(response *ListDevicePersonStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDevicePersonStatisticsResponse
		var err error
		defer close(result)
		response, err = client.ListDevicePersonStatistics(request)
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

// ListDevicePersonStatisticsRequest is the request struct for api ListDevicePersonStatistics
type ListDevicePersonStatisticsRequest struct {
	*requests.RpcRequest
	StatisticsType string `position:"Body" name:"StatisticsType"`
	CorpId         string `position:"Body" name:"CorpId"`
	EndTime        string `position:"Body" name:"EndTime"`
	StartTime      string `position:"Body" name:"StartTime"`
	DataSourceId   string `position:"Body" name:"DataSourceId"`
}

// ListDevicePersonStatisticsResponse is the response struct for api ListDevicePersonStatistics
type ListDevicePersonStatisticsResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListDevicePersonStatisticsRequest creates a request to invoke ListDevicePersonStatistics API
func CreateListDevicePersonStatisticsRequest() (request *ListDevicePersonStatisticsRequest) {
	request = &ListDevicePersonStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListDevicePersonStatistics", "", "")
	request.Method = requests.POST
	return
}

// CreateListDevicePersonStatisticsResponse creates a response to parse from ListDevicePersonStatistics response
func CreateListDevicePersonStatisticsResponse() (response *ListDevicePersonStatisticsResponse) {
	response = &ListDevicePersonStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
