package oceanbasepro

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

// ResumeOmsOpenAPIProject invokes the oceanbasepro.ResumeOmsOpenAPIProject API synchronously
func (client *Client) ResumeOmsOpenAPIProject(request *ResumeOmsOpenAPIProjectRequest) (response *ResumeOmsOpenAPIProjectResponse, err error) {
	response = CreateResumeOmsOpenAPIProjectResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeOmsOpenAPIProjectWithChan invokes the oceanbasepro.ResumeOmsOpenAPIProject API asynchronously
func (client *Client) ResumeOmsOpenAPIProjectWithChan(request *ResumeOmsOpenAPIProjectRequest) (<-chan *ResumeOmsOpenAPIProjectResponse, <-chan error) {
	responseChan := make(chan *ResumeOmsOpenAPIProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeOmsOpenAPIProject(request)
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

// ResumeOmsOpenAPIProjectWithCallback invokes the oceanbasepro.ResumeOmsOpenAPIProject API asynchronously
func (client *Client) ResumeOmsOpenAPIProjectWithCallback(request *ResumeOmsOpenAPIProjectRequest, callback func(response *ResumeOmsOpenAPIProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeOmsOpenAPIProjectResponse
		var err error
		defer close(result)
		response, err = client.ResumeOmsOpenAPIProject(request)
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

// ResumeOmsOpenAPIProjectRequest is the request struct for api ResumeOmsOpenAPIProject
type ResumeOmsOpenAPIProjectRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	WorkerGradeId string           `position:"Body" name:"WorkerGradeId"`
	ProjectId     string           `position:"Body" name:"ProjectId"`
}

// ResumeOmsOpenAPIProjectResponse is the response struct for api ResumeOmsOpenAPIProject
type ResumeOmsOpenAPIProjectResponse struct {
	*responses.BaseResponse
	Success     bool        `json:"Success" xml:"Success"`
	Code        string      `json:"Code" xml:"Code"`
	Message     string      `json:"Message" xml:"Message"`
	Advice      string      `json:"Advice" xml:"Advice"`
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	TotalCount  int64       `json:"TotalCount" xml:"TotalCount"`
	Cost        string      `json:"Cost" xml:"Cost"`
	Data        bool        `json:"Data" xml:"Data"`
	ErrorDetail ErrorDetail `json:"ErrorDetail" xml:"ErrorDetail"`
}

// CreateResumeOmsOpenAPIProjectRequest creates a request to invoke ResumeOmsOpenAPIProject API
func CreateResumeOmsOpenAPIProjectRequest() (request *ResumeOmsOpenAPIProjectRequest) {
	request = &ResumeOmsOpenAPIProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "ResumeOmsOpenAPIProject", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResumeOmsOpenAPIProjectResponse creates a response to parse from ResumeOmsOpenAPIProject response
func CreateResumeOmsOpenAPIProjectResponse() (response *ResumeOmsOpenAPIProjectResponse) {
	response = &ResumeOmsOpenAPIProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}