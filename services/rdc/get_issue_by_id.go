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

// GetIssueById invokes the rdc.GetIssueById API synchronously
// api document: https://help.aliyun.com/api/rdc/getissuebyid.html
func (client *Client) GetIssueById(request *GetIssueByIdRequest) (response *GetIssueByIdResponse, err error) {
	response = CreateGetIssueByIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetIssueByIdWithChan invokes the rdc.GetIssueById API asynchronously
// api document: https://help.aliyun.com/api/rdc/getissuebyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetIssueByIdWithChan(request *GetIssueByIdRequest) (<-chan *GetIssueByIdResponse, <-chan error) {
	responseChan := make(chan *GetIssueByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetIssueById(request)
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

// GetIssueByIdWithCallback invokes the rdc.GetIssueById API asynchronously
// api document: https://help.aliyun.com/api/rdc/getissuebyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetIssueByIdWithCallback(request *GetIssueByIdRequest, callback func(response *GetIssueByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetIssueByIdResponse
		var err error
		defer close(result)
		response, err = client.GetIssueById(request)
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

// GetIssueByIdRequest is the request struct for api GetIssueById
type GetIssueByIdRequest struct {
	*requests.RpcRequest
	CorpIdentifier string           `position:"Query" name:"CorpIdentifier"`
	Id             requests.Integer `position:"Body" name:"Id"`
}

// GetIssueByIdResponse is the response struct for api GetIssueById
type GetIssueByIdResponse struct {
	*responses.BaseResponse
	Code      int                `json:"Code" xml:"Code"`
	Success   string             `json:"Success" xml:"Success"`
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Data      DataInGetIssueById `json:"Data" xml:"Data"`
}

// CreateGetIssueByIdRequest creates a request to invoke GetIssueById API
func CreateGetIssueByIdRequest() (request *GetIssueByIdRequest) {
	request = &GetIssueByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rdc", "2018-08-21", "GetIssueById", "rdc", "openAPI")
	return
}

// CreateGetIssueByIdResponse creates a response to parse from GetIssueById response
func CreateGetIssueByIdResponse() (response *GetIssueByIdResponse) {
	response = &GetIssueByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}