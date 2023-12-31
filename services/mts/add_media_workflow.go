package mts

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

// AddMediaWorkflow invokes the mts.AddMediaWorkflow API synchronously
func (client *Client) AddMediaWorkflow(request *AddMediaWorkflowRequest) (response *AddMediaWorkflowResponse, err error) {
	response = CreateAddMediaWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// AddMediaWorkflowWithChan invokes the mts.AddMediaWorkflow API asynchronously
func (client *Client) AddMediaWorkflowWithChan(request *AddMediaWorkflowRequest) (<-chan *AddMediaWorkflowResponse, <-chan error) {
	responseChan := make(chan *AddMediaWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddMediaWorkflow(request)
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

// AddMediaWorkflowWithCallback invokes the mts.AddMediaWorkflow API asynchronously
func (client *Client) AddMediaWorkflowWithCallback(request *AddMediaWorkflowRequest, callback func(response *AddMediaWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddMediaWorkflowResponse
		var err error
		defer close(result)
		response, err = client.AddMediaWorkflow(request)
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

// AddMediaWorkflowRequest is the request struct for api AddMediaWorkflow
type AddMediaWorkflowRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	Topology             string           `position:"Query" name:"Topology"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TriggerMode          string           `position:"Query" name:"TriggerMode"`
	Name                 string           `position:"Query" name:"Name"`
}

// AddMediaWorkflowResponse is the response struct for api AddMediaWorkflow
type AddMediaWorkflowResponse struct {
	*responses.BaseResponse
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	MediaWorkflow MediaWorkflow `json:"MediaWorkflow" xml:"MediaWorkflow"`
}

// CreateAddMediaWorkflowRequest creates a request to invoke AddMediaWorkflow API
func CreateAddMediaWorkflowRequest() (request *AddMediaWorkflowRequest) {
	request = &AddMediaWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddMediaWorkflow", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddMediaWorkflowResponse creates a response to parse from AddMediaWorkflow response
func CreateAddMediaWorkflowResponse() (response *AddMediaWorkflowResponse) {
	response = &AddMediaWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
