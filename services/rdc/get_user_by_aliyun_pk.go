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

// GetUserByAliyunPk invokes the rdc.GetUserByAliyunPk API synchronously
// api document: https://help.aliyun.com/api/rdc/getuserbyaliyunpk.html
func (client *Client) GetUserByAliyunPk(request *GetUserByAliyunPkRequest) (response *GetUserByAliyunPkResponse, err error) {
	response = CreateGetUserByAliyunPkResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserByAliyunPkWithChan invokes the rdc.GetUserByAliyunPk API asynchronously
// api document: https://help.aliyun.com/api/rdc/getuserbyaliyunpk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserByAliyunPkWithChan(request *GetUserByAliyunPkRequest) (<-chan *GetUserByAliyunPkResponse, <-chan error) {
	responseChan := make(chan *GetUserByAliyunPkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserByAliyunPk(request)
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

// GetUserByAliyunPkWithCallback invokes the rdc.GetUserByAliyunPk API asynchronously
// api document: https://help.aliyun.com/api/rdc/getuserbyaliyunpk.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserByAliyunPkWithCallback(request *GetUserByAliyunPkRequest, callback func(response *GetUserByAliyunPkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserByAliyunPkResponse
		var err error
		defer close(result)
		response, err = client.GetUserByAliyunPk(request)
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

// GetUserByAliyunPkRequest is the request struct for api GetUserByAliyunPk
type GetUserByAliyunPkRequest struct {
	*requests.RpcRequest
	Pk string `position:"Query" name:"Pk"`
}

// GetUserByAliyunPkResponse is the response struct for api GetUserByAliyunPk
type GetUserByAliyunPkResponse struct {
	*responses.BaseResponse
	Code      int                     `json:"Code" xml:"Code"`
	Success   bool                    `json:"Success" xml:"Success"`
	RequestId string                  `json:"RequestId" xml:"RequestId"`
	Data      DataInGetUserByAliyunPk `json:"Data" xml:"Data"`
}

// CreateGetUserByAliyunPkRequest creates a request to invoke GetUserByAliyunPk API
func CreateGetUserByAliyunPkRequest() (request *GetUserByAliyunPkRequest) {
	request = &GetUserByAliyunPkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rdc", "2018-08-21", "GetUserByAliyunPk", "rdc", "openAPI")
	return
}

// CreateGetUserByAliyunPkResponse creates a response to parse from GetUserByAliyunPk response
func CreateGetUserByAliyunPkResponse() (response *GetUserByAliyunPkResponse) {
	response = &GetUserByAliyunPkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
