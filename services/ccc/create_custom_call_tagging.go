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

// CreateCustomCallTagging invokes the ccc.CreateCustomCallTagging API synchronously
func (client *Client) CreateCustomCallTagging(request *CreateCustomCallTaggingRequest) (response *CreateCustomCallTaggingResponse, err error) {
	response = CreateCreateCustomCallTaggingResponse()
	err = client.DoAction(request, response)
	return
}

// CreateCustomCallTaggingWithChan invokes the ccc.CreateCustomCallTagging API asynchronously
func (client *Client) CreateCustomCallTaggingWithChan(request *CreateCustomCallTaggingRequest) (<-chan *CreateCustomCallTaggingResponse, <-chan error) {
	responseChan := make(chan *CreateCustomCallTaggingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateCustomCallTagging(request)
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

// CreateCustomCallTaggingWithCallback invokes the ccc.CreateCustomCallTagging API asynchronously
func (client *Client) CreateCustomCallTaggingWithCallback(request *CreateCustomCallTaggingRequest, callback func(response *CreateCustomCallTaggingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateCustomCallTaggingResponse
		var err error
		defer close(result)
		response, err = client.CreateCustomCallTagging(request)
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

// CreateCustomCallTaggingRequest is the request struct for api CreateCustomCallTagging
type CreateCustomCallTaggingRequest struct {
	*requests.RpcRequest
	InstanceId       string `position:"Query" name:"InstanceId"`
	CustomNumberList string `position:"Query" name:"CustomNumberList"`
}

// CreateCustomCallTaggingResponse is the response struct for api CreateCustomCallTagging
type CreateCustomCallTaggingResponse struct {
	*responses.BaseResponse
	RequestId      string        `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string        `json:"Code" xml:"Code"`
	Message        string        `json:"Message" xml:"Message"`
	Data           []FailureItem `json:"Data" xml:"Data"`
}

// CreateCreateCustomCallTaggingRequest creates a request to invoke CreateCustomCallTagging API
func CreateCreateCustomCallTaggingRequest() (request *CreateCustomCallTaggingRequest) {
	request = &CreateCustomCallTaggingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "CreateCustomCallTagging", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateCustomCallTaggingResponse creates a response to parse from CreateCustomCallTagging response
func CreateCreateCustomCallTaggingResponse() (response *CreateCustomCallTaggingResponse) {
	response = &CreateCustomCallTaggingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}