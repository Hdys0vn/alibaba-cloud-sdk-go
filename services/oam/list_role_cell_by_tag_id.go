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

// ListRoleCellByTagId invokes the oam.ListRoleCellByTagId API synchronously
// api document: https://help.aliyun.com/api/oam/listrolecellbytagid.html
func (client *Client) ListRoleCellByTagId(request *ListRoleCellByTagIdRequest) (response *ListRoleCellByTagIdResponse, err error) {
	response = CreateListRoleCellByTagIdResponse()
	err = client.DoAction(request, response)
	return
}

// ListRoleCellByTagIdWithChan invokes the oam.ListRoleCellByTagId API asynchronously
// api document: https://help.aliyun.com/api/oam/listrolecellbytagid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRoleCellByTagIdWithChan(request *ListRoleCellByTagIdRequest) (<-chan *ListRoleCellByTagIdResponse, <-chan error) {
	responseChan := make(chan *ListRoleCellByTagIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListRoleCellByTagId(request)
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

// ListRoleCellByTagIdWithCallback invokes the oam.ListRoleCellByTagId API asynchronously
// api document: https://help.aliyun.com/api/oam/listrolecellbytagid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListRoleCellByTagIdWithCallback(request *ListRoleCellByTagIdRequest, callback func(response *ListRoleCellByTagIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListRoleCellByTagIdResponse
		var err error
		defer close(result)
		response, err = client.ListRoleCellByTagId(request)
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

// ListRoleCellByTagIdRequest is the request struct for api ListRoleCellByTagId
type ListRoleCellByTagIdRequest struct {
	*requests.RpcRequest
	Id string `position:"Query" name:"Id"`
}

// ListRoleCellByTagIdResponse is the response struct for api ListRoleCellByTagId
type ListRoleCellByTagIdResponse struct {
	*responses.BaseResponse
	Code      string                    `json:"Code" xml:"Code"`
	Message   string                    `json:"Message" xml:"Message"`
	RequestId string                    `json:"RequestId" xml:"RequestId"`
	Data      DataInListRoleCellByTagId `json:"Data" xml:"Data"`
}

// CreateListRoleCellByTagIdRequest creates a request to invoke ListRoleCellByTagId API
func CreateListRoleCellByTagIdRequest() (request *ListRoleCellByTagIdRequest) {
	request = &ListRoleCellByTagIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Oam", "2017-01-01", "ListRoleCellByTagId", "", "")
	request.Method = requests.POST
	return
}

// CreateListRoleCellByTagIdResponse creates a response to parse from ListRoleCellByTagId response
func CreateListRoleCellByTagIdResponse() (response *ListRoleCellByTagIdResponse) {
	response = &ListRoleCellByTagIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
