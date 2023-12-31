package vcs

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

// DeleteRecords invokes the vcs.DeleteRecords API synchronously
func (client *Client) DeleteRecords(request *DeleteRecordsRequest) (response *DeleteRecordsResponse, err error) {
	response = CreateDeleteRecordsResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteRecordsWithChan invokes the vcs.DeleteRecords API asynchronously
func (client *Client) DeleteRecordsWithChan(request *DeleteRecordsRequest) (<-chan *DeleteRecordsResponse, <-chan error) {
	responseChan := make(chan *DeleteRecordsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteRecords(request)
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

// DeleteRecordsWithCallback invokes the vcs.DeleteRecords API asynchronously
func (client *Client) DeleteRecordsWithCallback(request *DeleteRecordsRequest, callback func(response *DeleteRecordsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteRecordsResponse
		var err error
		defer close(result)
		response, err = client.DeleteRecords(request)
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

// DeleteRecordsRequest is the request struct for api DeleteRecords
type DeleteRecordsRequest struct {
	*requests.RpcRequest
	AlgorithmType string `position:"Body" name:"AlgorithmType"`
	CorpId        string `position:"Body" name:"CorpId"`
	AttributeName string `position:"Body" name:"AttributeName"`
	OperatorType  string `position:"Body" name:"OperatorType"`
	Value         string `position:"Body" name:"Value"`
}

// DeleteRecordsResponse is the response struct for api DeleteRecords
type DeleteRecordsResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateDeleteRecordsRequest creates a request to invoke DeleteRecords API
func CreateDeleteRecordsRequest() (request *DeleteRecordsRequest) {
	request = &DeleteRecordsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "DeleteRecords", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteRecordsResponse creates a response to parse from DeleteRecords response
func CreateDeleteRecordsResponse() (response *DeleteRecordsResponse) {
	response = &DeleteRecordsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
