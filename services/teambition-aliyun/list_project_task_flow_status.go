package teambition_aliyun

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

// ListProjectTaskFlowStatus invokes the teambition_aliyun.ListProjectTaskFlowStatus API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/listprojecttaskflowstatus.html
func (client *Client) ListProjectTaskFlowStatus(request *ListProjectTaskFlowStatusRequest) (response *ListProjectTaskFlowStatusResponse, err error) {
	response = CreateListProjectTaskFlowStatusResponse()
	err = client.DoAction(request, response)
	return
}

// ListProjectTaskFlowStatusWithChan invokes the teambition_aliyun.ListProjectTaskFlowStatus API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/listprojecttaskflowstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListProjectTaskFlowStatusWithChan(request *ListProjectTaskFlowStatusRequest) (<-chan *ListProjectTaskFlowStatusResponse, <-chan error) {
	responseChan := make(chan *ListProjectTaskFlowStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListProjectTaskFlowStatus(request)
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

// ListProjectTaskFlowStatusWithCallback invokes the teambition_aliyun.ListProjectTaskFlowStatus API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/listprojecttaskflowstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListProjectTaskFlowStatusWithCallback(request *ListProjectTaskFlowStatusRequest, callback func(response *ListProjectTaskFlowStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListProjectTaskFlowStatusResponse
		var err error
		defer close(result)
		response, err = client.ListProjectTaskFlowStatus(request)
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

// ListProjectTaskFlowStatusRequest is the request struct for api ListProjectTaskFlowStatus
type ListProjectTaskFlowStatusRequest struct {
	*requests.RpcRequest
	TaskFlowId string `position:"Body" name:"TaskFlowId"`
	OrgId      string `position:"Body" name:"OrgId"`
}

// ListProjectTaskFlowStatusResponse is the response struct for api ListProjectTaskFlowStatus
type ListProjectTaskFlowStatusResponse struct {
	*responses.BaseResponse
	Successful bool             `json:"Successful" xml:"Successful"`
	ErrorCode  string           `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string           `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string           `json:"RequestId" xml:"RequestId"`
	Object     []TaskflowStatus `json:"Object" xml:"Object"`
}

// CreateListProjectTaskFlowStatusRequest creates a request to invoke ListProjectTaskFlowStatus API
func CreateListProjectTaskFlowStatusRequest() (request *ListProjectTaskFlowStatusRequest) {
	request = &ListProjectTaskFlowStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "ListProjectTaskFlowStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateListProjectTaskFlowStatusResponse creates a response to parse from ListProjectTaskFlowStatus response
func CreateListProjectTaskFlowStatusResponse() (response *ListProjectTaskFlowStatusResponse) {
	response = &ListProjectTaskFlowStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
