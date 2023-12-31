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

// CreateProjectTask invokes the teambition_aliyun.CreateProjectTask API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/createprojecttask.html
func (client *Client) CreateProjectTask(request *CreateProjectTaskRequest) (response *CreateProjectTaskResponse, err error) {
	response = CreateCreateProjectTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProjectTaskWithChan invokes the teambition_aliyun.CreateProjectTask API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/createprojecttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProjectTaskWithChan(request *CreateProjectTaskRequest) (<-chan *CreateProjectTaskResponse, <-chan error) {
	responseChan := make(chan *CreateProjectTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProjectTask(request)
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

// CreateProjectTaskWithCallback invokes the teambition_aliyun.CreateProjectTask API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/createprojecttask.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateProjectTaskWithCallback(request *CreateProjectTaskRequest, callback func(response *CreateProjectTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProjectTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateProjectTask(request)
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

// CreateProjectTaskRequest is the request struct for api CreateProjectTask
type CreateProjectTaskRequest struct {
	*requests.RpcRequest
	Note                  string           `position:"Body" name:"Note"`
	Visible               string           `position:"Body" name:"Visible"`
	ExecutorId            string           `position:"Body" name:"ExecutorId"`
	TaskFlowStatusId      string           `position:"Body" name:"TaskFlowStatusId"`
	StartDate             string           `position:"Body" name:"StartDate"`
	Priority              requests.Integer `position:"Body" name:"Priority"`
	ParentTaskId          string           `position:"Body" name:"ParentTaskId"`
	OrgId                 string           `position:"Body" name:"OrgId"`
	Content               string           `position:"Body" name:"Content"`
	SprintId              string           `position:"Body" name:"SprintId"`
	DueDate               string           `position:"Body" name:"DueDate"`
	ScenarioFieldConfigId string           `position:"Body" name:"ScenarioFieldConfigId"`
	ProjectId             string           `position:"Body" name:"ProjectId"`
	TaskListId            string           `position:"Body" name:"TaskListId"`
}

// CreateProjectTaskResponse is the response struct for api CreateProjectTask
type CreateProjectTaskResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     Object `json:"Object" xml:"Object"`
}

// CreateCreateProjectTaskRequest creates a request to invoke CreateProjectTask API
func CreateCreateProjectTaskRequest() (request *CreateProjectTaskRequest) {
	request = &CreateProjectTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "CreateProjectTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateProjectTaskResponse creates a response to parse from CreateProjectTask response
func CreateCreateProjectTaskResponse() (response *CreateProjectTaskResponse) {
	response = &CreateProjectTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
