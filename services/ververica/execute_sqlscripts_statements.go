package ververica

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

// ExecuteSqlscriptsStatements invokes the ververica.ExecuteSqlscriptsStatements API synchronously
func (client *Client) ExecuteSqlscriptsStatements(request *ExecuteSqlscriptsStatementsRequest) (response *ExecuteSqlscriptsStatementsResponse, err error) {
	response = CreateExecuteSqlscriptsStatementsResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteSqlscriptsStatementsWithChan invokes the ververica.ExecuteSqlscriptsStatements API asynchronously
func (client *Client) ExecuteSqlscriptsStatementsWithChan(request *ExecuteSqlscriptsStatementsRequest) (<-chan *ExecuteSqlscriptsStatementsResponse, <-chan error) {
	responseChan := make(chan *ExecuteSqlscriptsStatementsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteSqlscriptsStatements(request)
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

// ExecuteSqlscriptsStatementsWithCallback invokes the ververica.ExecuteSqlscriptsStatements API asynchronously
func (client *Client) ExecuteSqlscriptsStatementsWithCallback(request *ExecuteSqlscriptsStatementsRequest, callback func(response *ExecuteSqlscriptsStatementsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteSqlscriptsStatementsResponse
		var err error
		defer close(result)
		response, err = client.ExecuteSqlscriptsStatements(request)
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

// ExecuteSqlscriptsStatementsRequest is the request struct for api ExecuteSqlscriptsStatements
type ExecuteSqlscriptsStatementsRequest struct {
	*requests.RoaRequest
	Workspace  string `position:"Path" name:"workspace"`
	ParamsJson string `position:"Body" name:"paramsJson"`
	Namespace  string `position:"Path" name:"namespace"`
}

// ExecuteSqlscriptsStatementsResponse is the response struct for api ExecuteSqlscriptsStatements
type ExecuteSqlscriptsStatementsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateExecuteSqlscriptsStatementsRequest creates a request to invoke ExecuteSqlscriptsStatements API
func CreateExecuteSqlscriptsStatementsRequest() (request *ExecuteSqlscriptsStatementsRequest) {
	request = &ExecuteSqlscriptsStatementsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "ExecuteSqlscriptsStatements", "/pop/workspaces/[workspace]/sql/v1beta1/namespaces/[namespace]/sqlscripts:execute-multi", "", "")
	request.Method = requests.POST
	return
}

// CreateExecuteSqlscriptsStatementsResponse creates a response to parse from ExecuteSqlscriptsStatements response
func CreateExecuteSqlscriptsStatementsResponse() (response *ExecuteSqlscriptsStatementsResponse) {
	response = &ExecuteSqlscriptsStatementsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}