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

// StopOmsOpenAPIProject invokes the oceanbasepro.StopOmsOpenAPIProject API synchronously
func (client *Client) StopOmsOpenAPIProject(request *StopOmsOpenAPIProjectRequest) (response *StopOmsOpenAPIProjectResponse, err error) {
	response = CreateStopOmsOpenAPIProjectResponse()
	err = client.DoAction(request, response)
	return
}

// StopOmsOpenAPIProjectWithChan invokes the oceanbasepro.StopOmsOpenAPIProject API asynchronously
func (client *Client) StopOmsOpenAPIProjectWithChan(request *StopOmsOpenAPIProjectRequest) (<-chan *StopOmsOpenAPIProjectResponse, <-chan error) {
	responseChan := make(chan *StopOmsOpenAPIProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopOmsOpenAPIProject(request)
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

// StopOmsOpenAPIProjectWithCallback invokes the oceanbasepro.StopOmsOpenAPIProject API asynchronously
func (client *Client) StopOmsOpenAPIProjectWithCallback(request *StopOmsOpenAPIProjectRequest, callback func(response *StopOmsOpenAPIProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopOmsOpenAPIProjectResponse
		var err error
		defer close(result)
		response, err = client.StopOmsOpenAPIProject(request)
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

// StopOmsOpenAPIProjectRequest is the request struct for api StopOmsOpenAPIProject
type StopOmsOpenAPIProjectRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	WorkerGradeId string           `position:"Body" name:"WorkerGradeId"`
	ProjectId     string           `position:"Body" name:"ProjectId"`
}

// StopOmsOpenAPIProjectResponse is the response struct for api StopOmsOpenAPIProject
type StopOmsOpenAPIProjectResponse struct {
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

// CreateStopOmsOpenAPIProjectRequest creates a request to invoke StopOmsOpenAPIProject API
func CreateStopOmsOpenAPIProjectRequest() (request *StopOmsOpenAPIProjectRequest) {
	request = &StopOmsOpenAPIProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "StopOmsOpenAPIProject", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopOmsOpenAPIProjectResponse creates a response to parse from StopOmsOpenAPIProject response
func CreateStopOmsOpenAPIProjectResponse() (response *StopOmsOpenAPIProjectResponse) {
	response = &StopOmsOpenAPIProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
