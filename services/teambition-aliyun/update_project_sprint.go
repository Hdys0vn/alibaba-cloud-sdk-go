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

// UpdateProjectSprint invokes the teambition_aliyun.UpdateProjectSprint API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/updateprojectsprint.html
func (client *Client) UpdateProjectSprint(request *UpdateProjectSprintRequest) (response *UpdateProjectSprintResponse, err error) {
	response = CreateUpdateProjectSprintResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateProjectSprintWithChan invokes the teambition_aliyun.UpdateProjectSprint API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/updateprojectsprint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateProjectSprintWithChan(request *UpdateProjectSprintRequest) (<-chan *UpdateProjectSprintResponse, <-chan error) {
	responseChan := make(chan *UpdateProjectSprintResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateProjectSprint(request)
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

// UpdateProjectSprintWithCallback invokes the teambition_aliyun.UpdateProjectSprint API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/updateprojectsprint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateProjectSprintWithCallback(request *UpdateProjectSprintRequest, callback func(response *UpdateProjectSprintResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateProjectSprintResponse
		var err error
		defer close(result)
		response, err = client.UpdateProjectSprint(request)
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

// UpdateProjectSprintRequest is the request struct for api UpdateProjectSprint
type UpdateProjectSprintRequest struct {
	*requests.RpcRequest
	ExecutorId  string `position:"Body" name:"ExecutorId"`
	Description string `position:"Body" name:"Description"`
	StartDate   string `position:"Body" name:"StartDate"`
	OrgId       string `position:"Body" name:"OrgId"`
	SprintId    string `position:"Body" name:"SprintId"`
	DueDate     string `position:"Body" name:"DueDate"`
	Name        string `position:"Body" name:"Name"`
	ProjectId   string `position:"Body" name:"ProjectId"`
}

// UpdateProjectSprintResponse is the response struct for api UpdateProjectSprint
type UpdateProjectSprintResponse struct {
	*responses.BaseResponse
	Successful bool   `json:"Successful" xml:"Successful"`
	ErrorCode  string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg   string `json:"ErrorMsg" xml:"ErrorMsg"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Object     bool   `json:"Object" xml:"Object"`
}

// CreateUpdateProjectSprintRequest creates a request to invoke UpdateProjectSprint API
func CreateUpdateProjectSprintRequest() (request *UpdateProjectSprintRequest) {
	request = &UpdateProjectSprintRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "UpdateProjectSprint", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateProjectSprintResponse creates a response to parse from UpdateProjectSprint response
func CreateUpdateProjectSprintResponse() (response *UpdateProjectSprintResponse) {
	response = &UpdateProjectSprintResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
