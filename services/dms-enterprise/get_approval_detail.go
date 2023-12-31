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

// GetApprovalDetail invokes the dms_enterprise.GetApprovalDetail API synchronously
func (client *Client) GetApprovalDetail(request *GetApprovalDetailRequest) (response *GetApprovalDetailResponse, err error) {
	response = CreateGetApprovalDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetApprovalDetailWithChan invokes the dms_enterprise.GetApprovalDetail API asynchronously
func (client *Client) GetApprovalDetailWithChan(request *GetApprovalDetailRequest) (<-chan *GetApprovalDetailResponse, <-chan error) {
	responseChan := make(chan *GetApprovalDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetApprovalDetail(request)
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

// GetApprovalDetailWithCallback invokes the dms_enterprise.GetApprovalDetail API asynchronously
func (client *Client) GetApprovalDetailWithCallback(request *GetApprovalDetailRequest, callback func(response *GetApprovalDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetApprovalDetailResponse
		var err error
		defer close(result)
		response, err = client.GetApprovalDetail(request)
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

// GetApprovalDetailRequest is the request struct for api GetApprovalDetail
type GetApprovalDetailRequest struct {
	*requests.RpcRequest
	Tid                requests.Integer `position:"Query" name:"Tid"`
	WorkflowInstanceId requests.Integer `position:"Query" name:"WorkflowInstanceId"`
}

// GetApprovalDetailResponse is the response struct for api GetApprovalDetail
type GetApprovalDetailResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	ErrorCode      string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool           `json:"Success" xml:"Success"`
	ApprovalDetail ApprovalDetail `json:"ApprovalDetail" xml:"ApprovalDetail"`
}

// CreateGetApprovalDetailRequest creates a request to invoke GetApprovalDetail API
func CreateGetApprovalDetailRequest() (request *GetApprovalDetailRequest) {
	request = &GetApprovalDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetApprovalDetail", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetApprovalDetailResponse creates a response to parse from GetApprovalDetail response
func CreateGetApprovalDetailResponse() (response *GetApprovalDetailResponse) {
	response = &GetApprovalDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
