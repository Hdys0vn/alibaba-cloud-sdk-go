package devops_rdc

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

// GetTaskListFilter invokes the devops_rdc.GetTaskListFilter API synchronously
func (client *Client) GetTaskListFilter(request *GetTaskListFilterRequest) (response *GetTaskListFilterResponse, err error) {
	response = CreateGetTaskListFilterResponse()
	err = client.DoAction(request, response)
	return
}

// GetTaskListFilterWithChan invokes the devops_rdc.GetTaskListFilter API asynchronously
func (client *Client) GetTaskListFilterWithChan(request *GetTaskListFilterRequest) (<-chan *GetTaskListFilterResponse, <-chan error) {
	responseChan := make(chan *GetTaskListFilterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTaskListFilter(request)
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

// GetTaskListFilterWithCallback invokes the devops_rdc.GetTaskListFilter API asynchronously
func (client *Client) GetTaskListFilterWithCallback(request *GetTaskListFilterRequest, callback func(response *GetTaskListFilterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTaskListFilterResponse
		var err error
		defer close(result)
		response, err = client.GetTaskListFilter(request)
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

// GetTaskListFilterRequest is the request struct for api GetTaskListFilter
type GetTaskListFilterRequest struct {
	*requests.RpcRequest
	InvolveMembers        string           `position:"Body" name:"InvolveMembers"`
	ExecutorId            string           `position:"Body" name:"ExecutorId"`
	OrderCondition        string           `position:"Body" name:"OrderCondition"`
	SprintId              string           `position:"Body" name:"SprintId"`
	Extra                 string           `position:"Body" name:"Extra"`
	PageSize              requests.Integer `position:"Body" name:"PageSize"`
	ScenarioFieldConfigId string           `position:"Body" name:"ScenarioFieldConfigId"`
	IsDone                requests.Boolean `position:"Body" name:"IsDone"`
	ObjectType            string           `position:"Body" name:"ObjectType"`
	ProjectId             string           `position:"Body" name:"ProjectId"`
	PageToken             string           `position:"Body" name:"PageToken"`
	Order                 string           `position:"Body" name:"Order"`
	TagId                 string           `position:"Body" name:"TagId"`
	TaskFlowStatusId      string           `position:"Body" name:"TaskFlowStatusId"`
	DueDateStart          string           `position:"Body" name:"DueDateStart"`
	CreatorId             string           `position:"Body" name:"CreatorId"`
	Priority              string           `position:"Body" name:"Priority"`
	DueDateEnd            string           `position:"Body" name:"DueDateEnd"`
	OrgId                 string           `position:"Body" name:"OrgId"`
	Name                  string           `position:"Body" name:"Name"`
}

// GetTaskListFilterResponse is the response struct for api GetTaskListFilter
type GetTaskListFilterResponse struct {
	*responses.BaseResponse
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	Successful string `json:"Successful" xml:"Successful"`
	Object     Object `json:"Object" xml:"Object"`
}

// CreateGetTaskListFilterRequest creates a request to invoke GetTaskListFilter API
func CreateGetTaskListFilterRequest() (request *GetTaskListFilterRequest) {
	request = &GetTaskListFilterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetTaskListFilter", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTaskListFilterResponse creates a response to parse from GetTaskListFilter response
func CreateGetTaskListFilterResponse() (response *GetTaskListFilterResponse) {
	response = &GetTaskListFilterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
