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

// AddRamMember invokes the rdc.AddRamMember API synchronously
// api document: https://help.aliyun.com/api/rdc/addrammember.html
func (client *Client) AddRamMember(request *AddRamMemberRequest) (response *AddRamMemberResponse, err error) {
	response = CreateAddRamMemberResponse()
	err = client.DoAction(request, response)
	return
}

// AddRamMemberWithChan invokes the rdc.AddRamMember API asynchronously
// api document: https://help.aliyun.com/api/rdc/addrammember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddRamMemberWithChan(request *AddRamMemberRequest) (<-chan *AddRamMemberResponse, <-chan error) {
	responseChan := make(chan *AddRamMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddRamMember(request)
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

// AddRamMemberWithCallback invokes the rdc.AddRamMember API asynchronously
// api document: https://help.aliyun.com/api/rdc/addrammember.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddRamMemberWithCallback(request *AddRamMemberRequest, callback func(response *AddRamMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddRamMemberResponse
		var err error
		defer close(result)
		response, err = client.AddRamMember(request)
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

// AddRamMemberRequest is the request struct for api AddRamMember
type AddRamMemberRequest struct {
	*requests.RpcRequest
	CorpIdentifier  string `position:"Body" name:"CorpIdentifier"`
	StaffIdentifier string `position:"Body" name:"StaffIdentifier"`
}

// AddRamMemberResponse is the response struct for api AddRamMember
type AddRamMemberResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Success   bool     `json:"Success" xml:"Success"`
	Code      int      `json:"Code" xml:"Code"`
	Message   string   `json:"Message" xml:"Message"`
	Data      []string `json:"Data" xml:"Data"`
}

// CreateAddRamMemberRequest creates a request to invoke AddRamMember API
func CreateAddRamMemberRequest() (request *AddRamMemberRequest) {
	request = &AddRamMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rdc", "2018-08-21", "AddRamMember", "rdc", "openAPI")
	return
}

// CreateAddRamMemberResponse creates a response to parse from AddRamMember response
func CreateAddRamMemberResponse() (response *AddRamMemberResponse) {
	response = &AddRamMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
