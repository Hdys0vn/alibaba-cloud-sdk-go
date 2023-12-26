package sas

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

// CreateSuspEventNote invokes the sas.CreateSuspEventNote API synchronously
func (client *Client) CreateSuspEventNote(request *CreateSuspEventNoteRequest) (response *CreateSuspEventNoteResponse, err error) {
	response = CreateCreateSuspEventNoteResponse()
	err = client.DoAction(request, response)
	return
}

// CreateSuspEventNoteWithChan invokes the sas.CreateSuspEventNote API asynchronously
func (client *Client) CreateSuspEventNoteWithChan(request *CreateSuspEventNoteRequest) (<-chan *CreateSuspEventNoteResponse, <-chan error) {
	responseChan := make(chan *CreateSuspEventNoteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateSuspEventNote(request)
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

// CreateSuspEventNoteWithCallback invokes the sas.CreateSuspEventNote API asynchronously
func (client *Client) CreateSuspEventNoteWithCallback(request *CreateSuspEventNoteRequest, callback func(response *CreateSuspEventNoteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateSuspEventNoteResponse
		var err error
		defer close(result)
		response, err = client.CreateSuspEventNote(request)
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

// CreateSuspEventNoteRequest is the request struct for api CreateSuspEventNote
type CreateSuspEventNoteRequest struct {
	*requests.RpcRequest
	EventId requests.Integer `position:"Query" name:"EventId"`
	Note    string           `position:"Query" name:"Note"`
}

// CreateSuspEventNoteResponse is the response struct for api CreateSuspEventNote
type CreateSuspEventNoteResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreateSuspEventNoteRequest creates a request to invoke CreateSuspEventNote API
func CreateCreateSuspEventNoteRequest() (request *CreateSuspEventNoteRequest) {
	request = &CreateSuspEventNoteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "CreateSuspEventNote", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateSuspEventNoteResponse creates a response to parse from CreateSuspEventNote response
func CreateCreateSuspEventNoteResponse() (response *CreateSuspEventNoteResponse) {
	response = &CreateSuspEventNoteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}