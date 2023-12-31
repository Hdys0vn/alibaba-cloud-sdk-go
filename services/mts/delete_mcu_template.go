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

// DeleteMcuTemplate invokes the mts.DeleteMcuTemplate API synchronously
func (client *Client) DeleteMcuTemplate(request *DeleteMcuTemplateRequest) (response *DeleteMcuTemplateResponse, err error) {
	response = CreateDeleteMcuTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMcuTemplateWithChan invokes the mts.DeleteMcuTemplate API asynchronously
func (client *Client) DeleteMcuTemplateWithChan(request *DeleteMcuTemplateRequest) (<-chan *DeleteMcuTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteMcuTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMcuTemplate(request)
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

// DeleteMcuTemplateWithCallback invokes the mts.DeleteMcuTemplate API asynchronously
func (client *Client) DeleteMcuTemplateWithCallback(request *DeleteMcuTemplateRequest, callback func(response *DeleteMcuTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMcuTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteMcuTemplate(request)
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

// DeleteMcuTemplateRequest is the request struct for api DeleteMcuTemplate
type DeleteMcuTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId           string           `position:"Query" name:"TemplateId"`
}

// DeleteMcuTemplateResponse is the response struct for api DeleteMcuTemplate
type DeleteMcuTemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

// CreateDeleteMcuTemplateRequest creates a request to invoke DeleteMcuTemplate API
func CreateDeleteMcuTemplateRequest() (request *DeleteMcuTemplateRequest) {
	request = &DeleteMcuTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "DeleteMcuTemplate", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMcuTemplateResponse creates a response to parse from DeleteMcuTemplate response
func CreateDeleteMcuTemplateResponse() (response *DeleteMcuTemplateResponse) {
	response = &DeleteMcuTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
