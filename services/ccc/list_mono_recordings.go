package ccc

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

// ListMonoRecordings invokes the ccc.ListMonoRecordings API synchronously
func (client *Client) ListMonoRecordings(request *ListMonoRecordingsRequest) (response *ListMonoRecordingsResponse, err error) {
	response = CreateListMonoRecordingsResponse()
	err = client.DoAction(request, response)
	return
}

// ListMonoRecordingsWithChan invokes the ccc.ListMonoRecordings API asynchronously
func (client *Client) ListMonoRecordingsWithChan(request *ListMonoRecordingsRequest) (<-chan *ListMonoRecordingsResponse, <-chan error) {
	responseChan := make(chan *ListMonoRecordingsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListMonoRecordings(request)
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

// ListMonoRecordingsWithCallback invokes the ccc.ListMonoRecordings API asynchronously
func (client *Client) ListMonoRecordingsWithCallback(request *ListMonoRecordingsRequest, callback func(response *ListMonoRecordingsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListMonoRecordingsResponse
		var err error
		defer close(result)
		response, err = client.ListMonoRecordings(request)
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

// ListMonoRecordingsRequest is the request struct for api ListMonoRecordings
type ListMonoRecordingsRequest struct {
	*requests.RpcRequest
	ContactId  string `position:"Query" name:"ContactId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// ListMonoRecordingsResponse is the response struct for api ListMonoRecordings
type ListMonoRecordingsResponse struct {
	*responses.BaseResponse
	Code           string         `json:"Code" xml:"Code"`
	HttpStatusCode int            `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string         `json:"Message" xml:"Message"`
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	Data           []RecordingDTO `json:"Data" xml:"Data"`
}

// CreateListMonoRecordingsRequest creates a request to invoke ListMonoRecordings API
func CreateListMonoRecordingsRequest() (request *ListMonoRecordingsRequest) {
	request = &ListMonoRecordingsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "ListMonoRecordings", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListMonoRecordingsResponse creates a response to parse from ListMonoRecordings response
func CreateListMonoRecordingsResponse() (response *ListMonoRecordingsResponse) {
	response = &ListMonoRecordingsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
