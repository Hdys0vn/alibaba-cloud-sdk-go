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

// CreateContactFlow invokes the ccc.CreateContactFlow API synchronously
func (client *Client) CreateContactFlow(request *CreateContactFlowRequest) (response *CreateContactFlowResponse, err error) {
	response = CreateCreateContactFlowResponse()
	err = client.DoAction(request, response)
	return
}

// CreateContactFlowWithChan invokes the ccc.CreateContactFlow API asynchronously
func (client *Client) CreateContactFlowWithChan(request *CreateContactFlowRequest) (<-chan *CreateContactFlowResponse, <-chan error) {
	responseChan := make(chan *CreateContactFlowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateContactFlow(request)
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

// CreateContactFlowWithCallback invokes the ccc.CreateContactFlow API asynchronously
func (client *Client) CreateContactFlowWithCallback(request *CreateContactFlowRequest, callback func(response *CreateContactFlowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateContactFlowResponse
		var err error
		defer close(result)
		response, err = client.CreateContactFlow(request)
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

// CreateContactFlowRequest is the request struct for api CreateContactFlow
type CreateContactFlowRequest struct {
	*requests.RpcRequest
	Description string `position:"Query" name:"Description"`
	Type        string `position:"Query" name:"Type"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	Name        string `position:"Query" name:"Name"`
	Definition  string `position:"Query" name:"Definition"`
}

// CreateContactFlowResponse is the response struct for api CreateContactFlow
type CreateContactFlowResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	Data           string `json:"Data" xml:"Data"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateContactFlowRequest creates a request to invoke CreateContactFlow API
func CreateCreateContactFlowRequest() (request *CreateContactFlowRequest) {
	request = &CreateContactFlowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "CreateContactFlow", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateContactFlowResponse creates a response to parse from CreateContactFlow response
func CreateCreateContactFlowResponse() (response *CreateContactFlowResponse) {
	response = &CreateContactFlowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
