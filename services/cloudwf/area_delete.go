package cloudwf

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

// AreaDelete invokes the cloudwf.AreaDelete API synchronously
// api document: https://help.aliyun.com/api/cloudwf/areadelete.html
func (client *Client) AreaDelete(request *AreaDeleteRequest) (response *AreaDeleteResponse, err error) {
	response = CreateAreaDeleteResponse()
	err = client.DoAction(request, response)
	return
}

// AreaDeleteWithChan invokes the cloudwf.AreaDelete API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/areadelete.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AreaDeleteWithChan(request *AreaDeleteRequest) (<-chan *AreaDeleteResponse, <-chan error) {
	responseChan := make(chan *AreaDeleteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AreaDelete(request)
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

// AreaDeleteWithCallback invokes the cloudwf.AreaDelete API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/areadelete.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AreaDeleteWithCallback(request *AreaDeleteRequest, callback func(response *AreaDeleteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AreaDeleteResponse
		var err error
		defer close(result)
		response, err = client.AreaDelete(request)
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

// AreaDeleteRequest is the request struct for api AreaDelete
type AreaDeleteRequest struct {
	*requests.RpcRequest
	Aid requests.Integer `position:"Query" name:"Aid"`
	Sid requests.Integer `position:"Query" name:"Sid"`
}

// AreaDeleteResponse is the response struct for api AreaDelete
type AreaDeleteResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateAreaDeleteRequest creates a request to invoke AreaDelete API
func CreateAreaDeleteRequest() (request *AreaDeleteRequest) {
	request = &AreaDeleteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "AreaDelete", "cloudwf", "openAPI")
	return
}

// CreateAreaDeleteResponse creates a response to parse from AreaDelete response
func CreateAreaDeleteResponse() (response *AreaDeleteResponse) {
	response = &AreaDeleteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
