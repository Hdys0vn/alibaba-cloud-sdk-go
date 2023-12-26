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

// ListLhTaskFlowAndScenario invokes the dms_enterprise.ListLhTaskFlowAndScenario API synchronously
func (client *Client) ListLhTaskFlowAndScenario(request *ListLhTaskFlowAndScenarioRequest) (response *ListLhTaskFlowAndScenarioResponse, err error) {
	response = CreateListLhTaskFlowAndScenarioResponse()
	err = client.DoAction(request, response)
	return
}

// ListLhTaskFlowAndScenarioWithChan invokes the dms_enterprise.ListLhTaskFlowAndScenario API asynchronously
func (client *Client) ListLhTaskFlowAndScenarioWithChan(request *ListLhTaskFlowAndScenarioRequest) (<-chan *ListLhTaskFlowAndScenarioResponse, <-chan error) {
	responseChan := make(chan *ListLhTaskFlowAndScenarioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLhTaskFlowAndScenario(request)
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

// ListLhTaskFlowAndScenarioWithCallback invokes the dms_enterprise.ListLhTaskFlowAndScenario API asynchronously
func (client *Client) ListLhTaskFlowAndScenarioWithCallback(request *ListLhTaskFlowAndScenarioRequest, callback func(response *ListLhTaskFlowAndScenarioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLhTaskFlowAndScenarioResponse
		var err error
		defer close(result)
		response, err = client.ListLhTaskFlowAndScenario(request)
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

// ListLhTaskFlowAndScenarioRequest is the request struct for api ListLhTaskFlowAndScenario
type ListLhTaskFlowAndScenarioRequest struct {
	*requests.RpcRequest
	UserId  requests.Integer `position:"Query" name:"UserId"`
	Tid     requests.Integer `position:"Query" name:"Tid"`
	SpaceId requests.Integer `position:"Query" name:"SpaceId"`
}

// ListLhTaskFlowAndScenarioResponse is the response struct for api ListLhTaskFlowAndScenario
type ListLhTaskFlowAndScenarioResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	ErrorCode       string          `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage    string          `json:"ErrorMessage" xml:"ErrorMessage"`
	Success         bool            `json:"Success" xml:"Success"`
	RawDAGList      RawDAGList      `json:"RawDAGList" xml:"RawDAGList"`
	ScenarioDAGList ScenarioDAGList `json:"ScenarioDAGList" xml:"ScenarioDAGList"`
}

// CreateListLhTaskFlowAndScenarioRequest creates a request to invoke ListLhTaskFlowAndScenario API
func CreateListLhTaskFlowAndScenarioRequest() (request *ListLhTaskFlowAndScenarioRequest) {
	request = &ListLhTaskFlowAndScenarioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListLhTaskFlowAndScenario", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListLhTaskFlowAndScenarioResponse creates a response to parse from ListLhTaskFlowAndScenario response
func CreateListLhTaskFlowAndScenarioResponse() (response *ListLhTaskFlowAndScenarioResponse) {
	response = &ListLhTaskFlowAndScenarioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
