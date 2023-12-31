package oam

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

// UpdateGroupTagById invokes the oam.UpdateGroupTagById API synchronously
// api document: https://help.aliyun.com/api/oam/updategrouptagbyid.html
func (client *Client) UpdateGroupTagById(request *UpdateGroupTagByIdRequest) (response *UpdateGroupTagByIdResponse, err error) {
	response = CreateUpdateGroupTagByIdResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateGroupTagByIdWithChan invokes the oam.UpdateGroupTagById API asynchronously
// api document: https://help.aliyun.com/api/oam/updategrouptagbyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateGroupTagByIdWithChan(request *UpdateGroupTagByIdRequest) (<-chan *UpdateGroupTagByIdResponse, <-chan error) {
	responseChan := make(chan *UpdateGroupTagByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateGroupTagById(request)
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

// UpdateGroupTagByIdWithCallback invokes the oam.UpdateGroupTagById API asynchronously
// api document: https://help.aliyun.com/api/oam/updategrouptagbyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateGroupTagByIdWithCallback(request *UpdateGroupTagByIdRequest, callback func(response *UpdateGroupTagByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateGroupTagByIdResponse
		var err error
		defer close(result)
		response, err = client.UpdateGroupTagById(request)
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

// UpdateGroupTagByIdRequest is the request struct for api UpdateGroupTagById
type UpdateGroupTagByIdRequest struct {
	*requests.RpcRequest
	TagId       *[]UpdateGroupTagByIdTagId `position:"Query" name:"TagId"  type:"Repeated"`
	GroupTagId  string                     `position:"Query" name:"GroupTagId"`
	Name        string                     `position:"Query" name:"Name"`
	Description string                     `position:"Query" name:"Description"`
	UserId      string                     `position:"Query" name:"UserId"`
}

// UpdateGroupTagByIdTagId is a repeated param struct in UpdateGroupTagByIdRequest
type UpdateGroupTagByIdTagId struct {
	NewId string `name:"NewId"`
	OldId string `name:"OldId"`
}

// UpdateGroupTagByIdResponse is the response struct for api UpdateGroupTagById
type UpdateGroupTagByIdResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateGroupTagByIdRequest creates a request to invoke UpdateGroupTagById API
func CreateUpdateGroupTagByIdRequest() (request *UpdateGroupTagByIdRequest) {
	request = &UpdateGroupTagByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "UpdateGroupTagById", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateGroupTagByIdResponse creates a response to parse from UpdateGroupTagById response
func CreateUpdateGroupTagByIdResponse() (response *UpdateGroupTagByIdResponse) {
	response = &UpdateGroupTagByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
