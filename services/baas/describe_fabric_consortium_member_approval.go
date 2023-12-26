package baas

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

// DescribeFabricConsortiumMemberApproval invokes the baas.DescribeFabricConsortiumMemberApproval API synchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiummemberapproval.html
func (client *Client) DescribeFabricConsortiumMemberApproval(request *DescribeFabricConsortiumMemberApprovalRequest) (response *DescribeFabricConsortiumMemberApprovalResponse, err error) {
	response = CreateDescribeFabricConsortiumMemberApprovalResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFabricConsortiumMemberApprovalWithChan invokes the baas.DescribeFabricConsortiumMemberApproval API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiummemberapproval.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricConsortiumMemberApprovalWithChan(request *DescribeFabricConsortiumMemberApprovalRequest) (<-chan *DescribeFabricConsortiumMemberApprovalResponse, <-chan error) {
	responseChan := make(chan *DescribeFabricConsortiumMemberApprovalResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFabricConsortiumMemberApproval(request)
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

// DescribeFabricConsortiumMemberApprovalWithCallback invokes the baas.DescribeFabricConsortiumMemberApproval API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiummemberapproval.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricConsortiumMemberApprovalWithCallback(request *DescribeFabricConsortiumMemberApprovalRequest, callback func(response *DescribeFabricConsortiumMemberApprovalResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFabricConsortiumMemberApprovalResponse
		var err error
		defer close(result)
		response, err = client.DescribeFabricConsortiumMemberApproval(request)
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

// DescribeFabricConsortiumMemberApprovalRequest is the request struct for api DescribeFabricConsortiumMemberApproval
type DescribeFabricConsortiumMemberApprovalRequest struct {
	*requests.RpcRequest
	Location     string `position:"Body" name:"Location"`
	ConsortiumId string `position:"Query" name:"ConsortiumId"`
}

// DescribeFabricConsortiumMemberApprovalResponse is the response struct for api DescribeFabricConsortiumMemberApproval
type DescribeFabricConsortiumMemberApprovalResponse struct {
	*responses.BaseResponse
	RequestId string                     `json:"RequestId" xml:"RequestId"`
	Success   bool                       `json:"Success" xml:"Success"`
	ErrorCode int                        `json:"ErrorCode" xml:"ErrorCode"`
	Result    []ConsortiumMemberApproval `json:"Result" xml:"Result"`
}

// CreateDescribeFabricConsortiumMemberApprovalRequest creates a request to invoke DescribeFabricConsortiumMemberApproval API
func CreateDescribeFabricConsortiumMemberApprovalRequest() (request *DescribeFabricConsortiumMemberApprovalRequest) {
	request = &DescribeFabricConsortiumMemberApprovalRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeFabricConsortiumMemberApproval", "baas", "openAPI")
	return
}

// CreateDescribeFabricConsortiumMemberApprovalResponse creates a response to parse from DescribeFabricConsortiumMemberApproval response
func CreateDescribeFabricConsortiumMemberApprovalResponse() (response *DescribeFabricConsortiumMemberApprovalResponse) {
	response = &DescribeFabricConsortiumMemberApprovalResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
