package mpaas

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

// GetMcubeUpgradeTaskInfo invokes the mpaas.GetMcubeUpgradeTaskInfo API synchronously
func (client *Client) GetMcubeUpgradeTaskInfo(request *GetMcubeUpgradeTaskInfoRequest) (response *GetMcubeUpgradeTaskInfoResponse, err error) {
	response = CreateGetMcubeUpgradeTaskInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetMcubeUpgradeTaskInfoWithChan invokes the mpaas.GetMcubeUpgradeTaskInfo API asynchronously
func (client *Client) GetMcubeUpgradeTaskInfoWithChan(request *GetMcubeUpgradeTaskInfoRequest) (<-chan *GetMcubeUpgradeTaskInfoResponse, <-chan error) {
	responseChan := make(chan *GetMcubeUpgradeTaskInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMcubeUpgradeTaskInfo(request)
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

// GetMcubeUpgradeTaskInfoWithCallback invokes the mpaas.GetMcubeUpgradeTaskInfo API asynchronously
func (client *Client) GetMcubeUpgradeTaskInfoWithCallback(request *GetMcubeUpgradeTaskInfoRequest, callback func(response *GetMcubeUpgradeTaskInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMcubeUpgradeTaskInfoResponse
		var err error
		defer close(result)
		response, err = client.GetMcubeUpgradeTaskInfo(request)
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

// GetMcubeUpgradeTaskInfoRequest is the request struct for api GetMcubeUpgradeTaskInfo
type GetMcubeUpgradeTaskInfoRequest struct {
	*requests.RpcRequest
	TenantId    string           `position:"Body" name:"TenantId"`
	TaskId      requests.Integer `position:"Body" name:"TaskId"`
	AppId       string           `position:"Body" name:"AppId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// GetMcubeUpgradeTaskInfoResponse is the response struct for api GetMcubeUpgradeTaskInfo
type GetMcubeUpgradeTaskInfoResponse struct {
	*responses.BaseResponse
	ResultMessage string        `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string        `json:"ResultCode" xml:"ResultCode"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	GetTaskResult GetTaskResult `json:"GetTaskResult" xml:"GetTaskResult"`
}

// CreateGetMcubeUpgradeTaskInfoRequest creates a request to invoke GetMcubeUpgradeTaskInfo API
func CreateGetMcubeUpgradeTaskInfoRequest() (request *GetMcubeUpgradeTaskInfoRequest) {
	request = &GetMcubeUpgradeTaskInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "GetMcubeUpgradeTaskInfo", "", "")
	request.Method = requests.POST
	return
}

// CreateGetMcubeUpgradeTaskInfoResponse creates a response to parse from GetMcubeUpgradeTaskInfo response
func CreateGetMcubeUpgradeTaskInfoResponse() (response *GetMcubeUpgradeTaskInfoResponse) {
	response = &GetMcubeUpgradeTaskInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
