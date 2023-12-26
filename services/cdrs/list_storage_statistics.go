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

// ListStorageStatistics invokes the cdrs.ListStorageStatistics API synchronously
func (client *Client) ListStorageStatistics(request *ListStorageStatisticsRequest) (response *ListStorageStatisticsResponse, err error) {
	response = CreateListStorageStatisticsResponse()
	err = client.DoAction(request, response)
	return
}

// ListStorageStatisticsWithChan invokes the cdrs.ListStorageStatistics API asynchronously
func (client *Client) ListStorageStatisticsWithChan(request *ListStorageStatisticsRequest) (<-chan *ListStorageStatisticsResponse, <-chan error) {
	responseChan := make(chan *ListStorageStatisticsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListStorageStatistics(request)
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

// ListStorageStatisticsWithCallback invokes the cdrs.ListStorageStatistics API asynchronously
func (client *Client) ListStorageStatisticsWithCallback(request *ListStorageStatisticsRequest, callback func(response *ListStorageStatisticsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListStorageStatisticsResponse
		var err error
		defer close(result)
		response, err = client.ListStorageStatistics(request)
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

// ListStorageStatisticsRequest is the request struct for api ListStorageStatistics
type ListStorageStatisticsRequest struct {
	*requests.RpcRequest
	CorpId string `position:"Body" name:"CorpId"`
}

// ListStorageStatisticsResponse is the response struct for api ListStorageStatistics
type ListStorageStatisticsResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListStorageStatisticsRequest creates a request to invoke ListStorageStatistics API
func CreateListStorageStatisticsRequest() (request *ListStorageStatisticsRequest) {
	request = &ListStorageStatisticsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListStorageStatistics", "", "")
	request.Method = requests.POST
	return
}

// CreateListStorageStatisticsResponse creates a response to parse from ListStorageStatistics response
func CreateListStorageStatisticsResponse() (response *ListStorageStatisticsResponse) {
	response = &ListStorageStatisticsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}