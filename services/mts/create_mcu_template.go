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

// CreateMcuTemplate invokes the mts.CreateMcuTemplate API synchronously
func (client *Client) CreateMcuTemplate(request *CreateMcuTemplateRequest) (response *CreateMcuTemplateResponse, err error) {
	response = CreateCreateMcuTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMcuTemplateWithChan invokes the mts.CreateMcuTemplate API asynchronously
func (client *Client) CreateMcuTemplateWithChan(request *CreateMcuTemplateRequest) (<-chan *CreateMcuTemplateResponse, <-chan error) {
	responseChan := make(chan *CreateMcuTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMcuTemplate(request)
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

// CreateMcuTemplateWithCallback invokes the mts.CreateMcuTemplate API asynchronously
func (client *Client) CreateMcuTemplateWithCallback(request *CreateMcuTemplateRequest, callback func(response *CreateMcuTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMcuTemplateResponse
		var err error
		defer close(result)
		response, err = client.CreateMcuTemplate(request)
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

// CreateMcuTemplateRequest is the request struct for api CreateMcuTemplate
type CreateMcuTemplateRequest struct {
	*requests.RpcRequest
	Template             string           `position:"Query" name:"Template"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateMcuTemplateResponse is the response struct for api CreateMcuTemplate
type CreateMcuTemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

// CreateCreateMcuTemplateRequest creates a request to invoke CreateMcuTemplate API
func CreateCreateMcuTemplateRequest() (request *CreateMcuTemplateRequest) {
	request = &CreateMcuTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "CreateMcuTemplate", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMcuTemplateResponse creates a response to parse from CreateMcuTemplate response
func CreateCreateMcuTemplateResponse() (response *CreateMcuTemplateResponse) {
	response = &CreateMcuTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
