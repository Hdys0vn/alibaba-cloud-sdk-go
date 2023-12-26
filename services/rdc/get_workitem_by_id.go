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

// GetWorkitemById invokes the rdc.GetWorkitemById API synchronously
// api document: https://help.aliyun.com/api/rdc/getworkitembyid.html
func (client *Client) GetWorkitemById(request *GetWorkitemByIdRequest) (response *GetWorkitemByIdResponse, err error) {
	response = CreateGetWorkitemByIdResponse()
	err = client.DoAction(request, response)
	return
}

// GetWorkitemByIdWithChan invokes the rdc.GetWorkitemById API asynchronously
// api document: https://help.aliyun.com/api/rdc/getworkitembyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetWorkitemByIdWithChan(request *GetWorkitemByIdRequest) (<-chan *GetWorkitemByIdResponse, <-chan error) {
	responseChan := make(chan *GetWorkitemByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetWorkitemById(request)
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

// GetWorkitemByIdWithCallback invokes the rdc.GetWorkitemById API asynchronously
// api document: https://help.aliyun.com/api/rdc/getworkitembyid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetWorkitemByIdWithCallback(request *GetWorkitemByIdRequest, callback func(response *GetWorkitemByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetWorkitemByIdResponse
		var err error
		defer close(result)
		response, err = client.GetWorkitemById(request)
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

// GetWorkitemByIdRequest is the request struct for api GetWorkitemById
type GetWorkitemByIdRequest struct {
	*requests.RpcRequest
	CorpIdentifier string           `position:"Query" name:"CorpIdentifier"`
	Id             requests.Integer `position:"Body" name:"Id"`
}

// GetWorkitemByIdResponse is the response struct for api GetWorkitemById
type GetWorkitemByIdResponse struct {
	*responses.BaseResponse
	Code      int                   `json:"Code" xml:"Code"`
	Success   string                `json:"Success" xml:"Success"`
	RequestId string                `json:"RequestId" xml:"RequestId"`
	Data      DataInGetWorkitemById `json:"Data" xml:"Data"`
}

// CreateGetWorkitemByIdRequest creates a request to invoke GetWorkitemById API
func CreateGetWorkitemByIdRequest() (request *GetWorkitemByIdRequest) {
	request = &GetWorkitemByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rdc", "2018-08-21", "GetWorkitemById", "rdc", "openAPI")
	return
}

// CreateGetWorkitemByIdResponse creates a response to parse from GetWorkitemById response
func CreateGetWorkitemByIdResponse() (response *GetWorkitemByIdResponse) {
	response = &GetWorkitemByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
