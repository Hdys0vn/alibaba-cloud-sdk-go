package dataworks_public

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

// CreateExportMigration invokes the dataworks_public.CreateExportMigration API synchronously
func (client *Client) CreateExportMigration(request *CreateExportMigrationRequest) (response *CreateExportMigrationResponse, err error) {
	response = CreateCreateExportMigrationResponse()
	err = client.DoAction(request, response)
	return
}

// CreateExportMigrationWithChan invokes the dataworks_public.CreateExportMigration API asynchronously
func (client *Client) CreateExportMigrationWithChan(request *CreateExportMigrationRequest) (<-chan *CreateExportMigrationResponse, <-chan error) {
	responseChan := make(chan *CreateExportMigrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateExportMigration(request)
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

// CreateExportMigrationWithCallback invokes the dataworks_public.CreateExportMigration API asynchronously
func (client *Client) CreateExportMigrationWithCallback(request *CreateExportMigrationRequest, callback func(response *CreateExportMigrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateExportMigrationResponse
		var err error
		defer close(result)
		response, err = client.CreateExportMigration(request)
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

// CreateExportMigrationRequest is the request struct for api CreateExportMigration
type CreateExportMigrationRequest struct {
	*requests.RpcRequest
	IncrementalSince   requests.Integer `position:"Body" name:"IncrementalSince"`
	Description        string           `position:"Body" name:"Description"`
	ExportObjectStatus string           `position:"Body" name:"ExportObjectStatus"`
	ExportMode         string           `position:"Body" name:"ExportMode"`
	Name               string           `position:"Body" name:"Name"`
	ProjectId          requests.Integer `position:"Body" name:"ProjectId"`
}

// CreateExportMigrationResponse is the response struct for api CreateExportMigration
type CreateExportMigrationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      int64  `json:"Data" xml:"Data"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateCreateExportMigrationRequest creates a request to invoke CreateExportMigration API
func CreateCreateExportMigrationRequest() (request *CreateExportMigrationRequest) {
	request = &CreateExportMigrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateExportMigration", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateExportMigrationResponse creates a response to parse from CreateExportMigration response
func CreateCreateExportMigrationResponse() (response *CreateExportMigrationResponse) {
	response = &CreateExportMigrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
