package ecd

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

// DetachCen invokes the ecd.DetachCen API synchronously
func (client *Client) DetachCen(request *DetachCenRequest) (response *DetachCenResponse, err error) {
	response = CreateDetachCenResponse()
	err = client.DoAction(request, response)
	return
}

// DetachCenWithChan invokes the ecd.DetachCen API asynchronously
func (client *Client) DetachCenWithChan(request *DetachCenRequest) (<-chan *DetachCenResponse, <-chan error) {
	responseChan := make(chan *DetachCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DetachCen(request)
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

// DetachCenWithCallback invokes the ecd.DetachCen API asynchronously
func (client *Client) DetachCenWithCallback(request *DetachCenRequest, callback func(response *DetachCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DetachCenResponse
		var err error
		defer close(result)
		response, err = client.DetachCen(request)
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

// DetachCenRequest is the request struct for api DetachCen
type DetachCenRequest struct {
	*requests.RpcRequest
	OfficeSiteId string `position:"Query" name:"OfficeSiteId"`
}

// DetachCenResponse is the response struct for api DetachCen
type DetachCenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDetachCenRequest creates a request to invoke DetachCen API
func CreateDetachCenRequest() (request *DetachCenRequest) {
	request = &DetachCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "DetachCen", "", "")
	request.Method = requests.POST
	return
}

// CreateDetachCenResponse creates a response to parse from DetachCen response
func CreateDetachCenResponse() (response *DetachCenResponse) {
	response = &DetachCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
