package edas

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

// RollbackApplication invokes the edas.RollbackApplication API synchronously
func (client *Client) RollbackApplication(request *RollbackApplicationRequest) (response *RollbackApplicationResponse, err error) {
	response = CreateRollbackApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// RollbackApplicationWithChan invokes the edas.RollbackApplication API asynchronously
func (client *Client) RollbackApplicationWithChan(request *RollbackApplicationRequest) (<-chan *RollbackApplicationResponse, <-chan error) {
	responseChan := make(chan *RollbackApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RollbackApplication(request)
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

// RollbackApplicationWithCallback invokes the edas.RollbackApplication API asynchronously
func (client *Client) RollbackApplicationWithCallback(request *RollbackApplicationRequest, callback func(response *RollbackApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RollbackApplicationResponse
		var err error
		defer close(result)
		response, err = client.RollbackApplication(request)
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

// RollbackApplicationRequest is the request struct for api RollbackApplication
type RollbackApplicationRequest struct {
	*requests.RoaRequest
	AppId          string           `position:"Query" name:"AppId"`
	GroupId        string           `position:"Query" name:"GroupId"`
	BatchWaitTime  requests.Integer `position:"Query" name:"BatchWaitTime"`
	Batch          requests.Integer `position:"Query" name:"Batch"`
	HistoryVersion string           `position:"Query" name:"HistoryVersion"`
}

// RollbackApplicationResponse is the response struct for api RollbackApplication
type RollbackApplicationResponse struct {
	*responses.BaseResponse
	ChangeOrderId string `json:"ChangeOrderId" xml:"ChangeOrderId"`
	Code          int    `json:"Code" xml:"Code"`
	Message       string `json:"Message" xml:"Message"`
	RequestId     string `json:"RequestId" xml:"RequestId"`
}

// CreateRollbackApplicationRequest creates a request to invoke RollbackApplication API
func CreateRollbackApplicationRequest() (request *RollbackApplicationRequest) {
	request = &RollbackApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "RollbackApplication", "/pop/v5/changeorder/co_rollback", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRollbackApplicationResponse creates a response to parse from RollbackApplication response
func CreateRollbackApplicationResponse() (response *RollbackApplicationResponse) {
	response = &RollbackApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}