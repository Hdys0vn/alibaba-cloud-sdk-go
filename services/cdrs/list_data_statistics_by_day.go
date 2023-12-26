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

// ListDataStatisticsByDay invokes the cdrs.ListDataStatisticsByDay API synchronously
func (client *Client) ListDataStatisticsByDay(request *ListDataStatisticsByDayRequest) (response *ListDataStatisticsByDayResponse, err error) {
	response = CreateListDataStatisticsByDayResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataStatisticsByDayWithChan invokes the cdrs.ListDataStatisticsByDay API asynchronously
func (client *Client) ListDataStatisticsByDayWithChan(request *ListDataStatisticsByDayRequest) (<-chan *ListDataStatisticsByDayResponse, <-chan error) {
	responseChan := make(chan *ListDataStatisticsByDayResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataStatisticsByDay(request)
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

// ListDataStatisticsByDayWithCallback invokes the cdrs.ListDataStatisticsByDay API asynchronously
func (client *Client) ListDataStatisticsByDayWithCallback(request *ListDataStatisticsByDayRequest, callback func(response *ListDataStatisticsByDayResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataStatisticsByDayResponse
		var err error
		defer close(result)
		response, err = client.ListDataStatisticsByDay(request)
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

// ListDataStatisticsByDayRequest is the request struct for api ListDataStatisticsByDay
type ListDataStatisticsByDayRequest struct {
	*requests.RpcRequest
	CorpId    string `position:"Body" name:"CorpId"`
	EndTime   string `position:"Body" name:"EndTime"`
	StartTime string `position:"Body" name:"StartTime"`
}

// ListDataStatisticsByDayResponse is the response struct for api ListDataStatisticsByDay
type ListDataStatisticsByDayResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	TotalCount int64   `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListDataStatisticsByDayRequest creates a request to invoke ListDataStatisticsByDay API
func CreateListDataStatisticsByDayRequest() (request *ListDataStatisticsByDayRequest) {
	request = &ListDataStatisticsByDayRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "ListDataStatisticsByDay", "", "")
	request.Method = requests.POST
	return
}

// CreateListDataStatisticsByDayResponse creates a response to parse from ListDataStatisticsByDay response
func CreateListDataStatisticsByDayResponse() (response *ListDataStatisticsByDayResponse) {
	response = &ListDataStatisticsByDayResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
