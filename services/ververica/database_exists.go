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

// DatabaseExists invokes the ververica.DatabaseExists API synchronously
func (client *Client) DatabaseExists(request *DatabaseExistsRequest) (response *DatabaseExistsResponse, err error) {
	response = CreateDatabaseExistsResponse()
	err = client.DoAction(request, response)
	return
}

// DatabaseExistsWithChan invokes the ververica.DatabaseExists API asynchronously
func (client *Client) DatabaseExistsWithChan(request *DatabaseExistsRequest) (<-chan *DatabaseExistsResponse, <-chan error) {
	responseChan := make(chan *DatabaseExistsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DatabaseExists(request)
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

// DatabaseExistsWithCallback invokes the ververica.DatabaseExists API asynchronously
func (client *Client) DatabaseExistsWithCallback(request *DatabaseExistsRequest, callback func(response *DatabaseExistsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DatabaseExistsResponse
		var err error
		defer close(result)
		response, err = client.DatabaseExists(request)
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

// DatabaseExistsRequest is the request struct for api DatabaseExists
type DatabaseExistsRequest struct {
	*requests.RoaRequest
	Workspace string `position:"Path" name:"workspace"`
	Database  string `position:"Query" name:"database"`
	Cat       string `position:"Path" name:"cat"`
	Namespace string `position:"Path" name:"namespace"`
}

// DatabaseExistsResponse is the response struct for api DatabaseExists
type DatabaseExistsResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateDatabaseExistsRequest creates a request to invoke DatabaseExists API
func CreateDatabaseExistsRequest() (request *DatabaseExistsRequest) {
	request = &DatabaseExistsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "DatabaseExists", "/pop/workspaces/[workspace]/catalog/v1beta2/namespaces/[namespace]/catalogs/[cat]:databaseExists", "", "")
	request.Method = requests.GET
	return
}

// CreateDatabaseExistsResponse creates a response to parse from DatabaseExists response
func CreateDatabaseExistsResponse() (response *DatabaseExistsResponse) {
	response = &DatabaseExistsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
