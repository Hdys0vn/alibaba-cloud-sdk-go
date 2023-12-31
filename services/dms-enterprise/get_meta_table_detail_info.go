package dms_enterprise

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

// GetMetaTableDetailInfo invokes the dms_enterprise.GetMetaTableDetailInfo API synchronously
func (client *Client) GetMetaTableDetailInfo(request *GetMetaTableDetailInfoRequest) (response *GetMetaTableDetailInfoResponse, err error) {
	response = CreateGetMetaTableDetailInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetMetaTableDetailInfoWithChan invokes the dms_enterprise.GetMetaTableDetailInfo API asynchronously
func (client *Client) GetMetaTableDetailInfoWithChan(request *GetMetaTableDetailInfoRequest) (<-chan *GetMetaTableDetailInfoResponse, <-chan error) {
	responseChan := make(chan *GetMetaTableDetailInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMetaTableDetailInfo(request)
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

// GetMetaTableDetailInfoWithCallback invokes the dms_enterprise.GetMetaTableDetailInfo API asynchronously
func (client *Client) GetMetaTableDetailInfoWithCallback(request *GetMetaTableDetailInfoRequest, callback func(response *GetMetaTableDetailInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMetaTableDetailInfoResponse
		var err error
		defer close(result)
		response, err = client.GetMetaTableDetailInfo(request)
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

// GetMetaTableDetailInfoRequest is the request struct for api GetMetaTableDetailInfo
type GetMetaTableDetailInfoRequest struct {
	*requests.RpcRequest
	Tid       requests.Integer `position:"Query" name:"Tid"`
	TableGuid string           `position:"Query" name:"TableGuid"`
}

// GetMetaTableDetailInfoResponse is the response struct for api GetMetaTableDetailInfo
type GetMetaTableDetailInfoResponse struct {
	*responses.BaseResponse
	RequestId    string     `json:"RequestId" xml:"RequestId"`
	ErrorCode    string     `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string     `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool       `json:"Success" xml:"Success"`
	DetailInfo   DetailInfo `json:"DetailInfo" xml:"DetailInfo"`
}

// CreateGetMetaTableDetailInfoRequest creates a request to invoke GetMetaTableDetailInfo API
func CreateGetMetaTableDetailInfoRequest() (request *GetMetaTableDetailInfoRequest) {
	request = &GetMetaTableDetailInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetMetaTableDetailInfo", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMetaTableDetailInfoResponse creates a response to parse from GetMetaTableDetailInfo response
func CreateGetMetaTableDetailInfoResponse() (response *GetMetaTableDetailInfoResponse) {
	response = &GetMetaTableDetailInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
