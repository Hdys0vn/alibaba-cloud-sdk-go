package dts

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

// DescirbeMigrationJobs invokes the dts.DescirbeMigrationJobs API synchronously
func (client *Client) DescirbeMigrationJobs(request *DescirbeMigrationJobsRequest) (response *DescirbeMigrationJobsResponse, err error) {
	response = CreateDescirbeMigrationJobsResponse()
	err = client.DoAction(request, response)
	return
}

// DescirbeMigrationJobsWithChan invokes the dts.DescirbeMigrationJobs API asynchronously
func (client *Client) DescirbeMigrationJobsWithChan(request *DescirbeMigrationJobsRequest) (<-chan *DescirbeMigrationJobsResponse, <-chan error) {
	responseChan := make(chan *DescirbeMigrationJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescirbeMigrationJobs(request)
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

// DescirbeMigrationJobsWithCallback invokes the dts.DescirbeMigrationJobs API asynchronously
func (client *Client) DescirbeMigrationJobsWithCallback(request *DescirbeMigrationJobsRequest, callback func(response *DescirbeMigrationJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescirbeMigrationJobsResponse
		var err error
		defer close(result)
		response, err = client.DescirbeMigrationJobs(request)
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

// DescirbeMigrationJobsRequest is the request struct for api DescirbeMigrationJobs
type DescirbeMigrationJobsRequest struct {
	*requests.RpcRequest
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	MigrationJobName string           `position:"Query" name:"MigrationJobName"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
}

// DescirbeMigrationJobsResponse is the response struct for api DescirbeMigrationJobs
type DescirbeMigrationJobsResponse struct {
	*responses.BaseResponse
	PageNumber       int           `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int           `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int64         `json:"TotalRecordCount" xml:"TotalRecordCount"`
	MigrationJobs    MigrationJobs `json:"MigrationJobs" xml:"MigrationJobs"`
}

// CreateDescirbeMigrationJobsRequest creates a request to invoke DescirbeMigrationJobs API
func CreateDescirbeMigrationJobsRequest() (request *DescirbeMigrationJobsRequest) {
	request = &DescirbeMigrationJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2017-06-01", "DescirbeMigrationJobs", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescirbeMigrationJobsResponse creates a response to parse from DescirbeMigrationJobs response
func CreateDescirbeMigrationJobsResponse() (response *DescirbeMigrationJobsResponse) {
	response = &DescirbeMigrationJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
