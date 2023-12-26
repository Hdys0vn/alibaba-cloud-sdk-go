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

// QueryMdsUpgradeTaskDetail invokes the mpaas.QueryMdsUpgradeTaskDetail API synchronously
func (client *Client) QueryMdsUpgradeTaskDetail(request *QueryMdsUpgradeTaskDetailRequest) (response *QueryMdsUpgradeTaskDetailResponse, err error) {
	response = CreateQueryMdsUpgradeTaskDetailResponse()
	err = client.DoAction(request, response)
	return
}

// QueryMdsUpgradeTaskDetailWithChan invokes the mpaas.QueryMdsUpgradeTaskDetail API asynchronously
func (client *Client) QueryMdsUpgradeTaskDetailWithChan(request *QueryMdsUpgradeTaskDetailRequest) (<-chan *QueryMdsUpgradeTaskDetailResponse, <-chan error) {
	responseChan := make(chan *QueryMdsUpgradeTaskDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMdsUpgradeTaskDetail(request)
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

// QueryMdsUpgradeTaskDetailWithCallback invokes the mpaas.QueryMdsUpgradeTaskDetail API asynchronously
func (client *Client) QueryMdsUpgradeTaskDetailWithCallback(request *QueryMdsUpgradeTaskDetailRequest, callback func(response *QueryMdsUpgradeTaskDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMdsUpgradeTaskDetailResponse
		var err error
		defer close(result)
		response, err = client.QueryMdsUpgradeTaskDetail(request)
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

// QueryMdsUpgradeTaskDetailRequest is the request struct for api QueryMdsUpgradeTaskDetail
type QueryMdsUpgradeTaskDetailRequest struct {
	*requests.RpcRequest
	TenantId    string           `position:"Body" name:"TenantId"`
	TaskId      requests.Integer `position:"Body" name:"TaskId"`
	AppId       string           `position:"Body" name:"AppId"`
	WorkspaceId string           `position:"Body" name:"WorkspaceId"`
}

// QueryMdsUpgradeTaskDetailResponse is the response struct for api QueryMdsUpgradeTaskDetail
type QueryMdsUpgradeTaskDetailResponse struct {
	*responses.BaseResponse
	ResultMessage string        `json:"ResultMessage" xml:"ResultMessage"`
	ResultCode    string        `json:"ResultCode" xml:"ResultCode"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	ResultContent ResultContent `json:"ResultContent" xml:"ResultContent"`
}

// CreateQueryMdsUpgradeTaskDetailRequest creates a request to invoke QueryMdsUpgradeTaskDetail API
func CreateQueryMdsUpgradeTaskDetailRequest() (request *QueryMdsUpgradeTaskDetailRequest) {
	request = &QueryMdsUpgradeTaskDetailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mPaaS", "2020-10-28", "QueryMdsUpgradeTaskDetail", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryMdsUpgradeTaskDetailResponse creates a response to parse from QueryMdsUpgradeTaskDetail response
func CreateQueryMdsUpgradeTaskDetailResponse() (response *QueryMdsUpgradeTaskDetailResponse) {
	response = &QueryMdsUpgradeTaskDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
