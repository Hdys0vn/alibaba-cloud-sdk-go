package uis

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

// CreateUis invokes the uis.CreateUis API synchronously
// api document: https://help.aliyun.com/api/uis/createuis.html
func (client *Client) CreateUis(request *CreateUisRequest) (response *CreateUisResponse, err error) {
	response = CreateCreateUisResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUisWithChan invokes the uis.CreateUis API asynchronously
// api document: https://help.aliyun.com/api/uis/createuis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUisWithChan(request *CreateUisRequest) (<-chan *CreateUisResponse, <-chan error) {
	responseChan := make(chan *CreateUisResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUis(request)
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

// CreateUisWithCallback invokes the uis.CreateUis API asynchronously
// api document: https://help.aliyun.com/api/uis/createuis.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUisWithCallback(request *CreateUisRequest, callback func(response *CreateUisResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUisResponse
		var err error
		defer close(result)
		response, err = client.CreateUis(request)
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

// CreateUisRequest is the request struct for api CreateUis
type CreateUisRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	Name                 string           `position:"Query" name:"Name"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateUisResponse is the response struct for api CreateUis
type CreateUisResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	UisId     string `json:"UisId" xml:"UisId"`
}

// CreateCreateUisRequest creates a request to invoke CreateUis API
func CreateCreateUisRequest() (request *CreateUisRequest) {
	request = &CreateUisRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "CreateUis", "uis", "openAPI")
	return
}

// CreateCreateUisResponse creates a response to parse from CreateUis response
func CreateCreateUisResponse() (response *CreateUisResponse) {
	response = &CreateUisResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
