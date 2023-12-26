package rdc

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

// CreateWorkitem invokes the rdc.CreateWorkitem API synchronously
// api document: https://help.aliyun.com/api/rdc/createworkitem.html
func (client *Client) CreateWorkitem(request *CreateWorkitemRequest) (response *CreateWorkitemResponse, err error) {
	response = CreateCreateWorkitemResponse()
	err = client.DoAction(request, response)
	return
}

// CreateWorkitemWithChan invokes the rdc.CreateWorkitem API asynchronously
// api document: https://help.aliyun.com/api/rdc/createworkitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateWorkitemWithChan(request *CreateWorkitemRequest) (<-chan *CreateWorkitemResponse, <-chan error) {
	responseChan := make(chan *CreateWorkitemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateWorkitem(request)
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

// CreateWorkitemWithCallback invokes the rdc.CreateWorkitem API asynchronously
// api document: https://help.aliyun.com/api/rdc/createworkitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateWorkitemWithCallback(request *CreateWorkitemRequest, callback func(response *CreateWorkitemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateWorkitemResponse
		var err error
		defer close(result)
		response, err = client.CreateWorkitem(request)
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

// CreateWorkitemRequest is the request struct for api CreateWorkitem
type CreateWorkitemRequest struct {
	*requests.RpcRequest
	Author         string           `position:"Body" name:"Author"`
	Subject        string           `position:"Body" name:"Subject"`
	Description    string           `position:"Body" name:"Description"`
	Stamp          string           `position:"Body" name:"Stamp"`
	AKProjectId    requests.Integer `position:"Body" name:"AKProjectId"`
	TemplateId     requests.Integer `position:"Body" name:"TemplateId"`
	AssignedTo     string           `position:"Body" name:"AssignedTo"`
	PriorityId     requests.Integer `position:"Body" name:"PriorityId"`
	SeriousLevelId requests.Integer `position:"Body" name:"SeriousLevelId"`
	ModuleIds      string           `position:"Body" name:"ModuleIds"`
	CorpIdentifier string           `position:"Query" name:"CorpIdentifier"`
	WatcherUsers   string           `position:"Body" name:"WatcherUsers"`
	Verifier       string           `position:"Body" name:"Verifier"`
	CfList         string           `position:"Body" name:"CfList"`
}

// CreateWorkitemResponse is the response struct for api CreateWorkitem
type CreateWorkitemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Data      int    `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateCreateWorkitemRequest creates a request to invoke CreateWorkitem API
func CreateCreateWorkitemRequest() (request *CreateWorkitemRequest) {
	request = &CreateWorkitemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rdc", "2018-08-21", "CreateWorkitem", "rdc", "openAPI")
	return
}

// CreateCreateWorkitemResponse creates a response to parse from CreateWorkitem response
func CreateCreateWorkitemResponse() (response *CreateWorkitemResponse) {
	response = &CreateWorkitemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
