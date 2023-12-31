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

// ModifyUisAttribute invokes the uis.ModifyUisAttribute API synchronously
// api document: https://help.aliyun.com/api/uis/modifyuisattribute.html
func (client *Client) ModifyUisAttribute(request *ModifyUisAttributeRequest) (response *ModifyUisAttributeResponse, err error) {
	response = CreateModifyUisAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyUisAttributeWithChan invokes the uis.ModifyUisAttribute API asynchronously
// api document: https://help.aliyun.com/api/uis/modifyuisattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyUisAttributeWithChan(request *ModifyUisAttributeRequest) (<-chan *ModifyUisAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyUisAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyUisAttribute(request)
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

// ModifyUisAttributeWithCallback invokes the uis.ModifyUisAttribute API asynchronously
// api document: https://help.aliyun.com/api/uis/modifyuisattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyUisAttributeWithCallback(request *ModifyUisAttributeRequest, callback func(response *ModifyUisAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyUisAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyUisAttribute(request)
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

// ModifyUisAttributeRequest is the request struct for api ModifyUisAttribute
type ModifyUisAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	UisId                string           `position:"Query" name:"UisId"`
	Name                 string           `position:"Query" name:"Name"`
	Description          string           `position:"Query" name:"Description"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyUisAttributeResponse is the response struct for api ModifyUisAttribute
type ModifyUisAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyUisAttributeRequest creates a request to invoke ModifyUisAttribute API
func CreateModifyUisAttributeRequest() (request *ModifyUisAttributeRequest) {
	request = &ModifyUisAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Uis", "2018-08-21", "ModifyUisAttribute", "uis", "openAPI")
	return
}

// CreateModifyUisAttributeResponse creates a response to parse from ModifyUisAttribute response
func CreateModifyUisAttributeResponse() (response *ModifyUisAttributeResponse) {
	response = &ModifyUisAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
