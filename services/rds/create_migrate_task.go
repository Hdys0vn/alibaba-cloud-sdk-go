package rds

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

// CreateMigrateTask invokes the rds.CreateMigrateTask API synchronously
func (client *Client) CreateMigrateTask(request *CreateMigrateTaskRequest) (response *CreateMigrateTaskResponse, err error) {
	response = CreateCreateMigrateTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMigrateTaskWithChan invokes the rds.CreateMigrateTask API asynchronously
func (client *Client) CreateMigrateTaskWithChan(request *CreateMigrateTaskRequest) (<-chan *CreateMigrateTaskResponse, <-chan error) {
	responseChan := make(chan *CreateMigrateTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMigrateTask(request)
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

// CreateMigrateTaskWithCallback invokes the rds.CreateMigrateTask API asynchronously
func (client *Client) CreateMigrateTaskWithCallback(request *CreateMigrateTaskRequest, callback func(response *CreateMigrateTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMigrateTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateMigrateTask(request)
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

// CreateMigrateTaskRequest is the request struct for api CreateMigrateTask
type CreateMigrateTaskRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	IsOnlineDB           string           `position:"Query" name:"IsOnlineDB"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	MigrateTaskId        string           `position:"Query" name:"MigrateTaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	OssObjectPositions   string           `position:"Query" name:"OssObjectPositions"`
	OSSUrls              string           `position:"Query" name:"OSSUrls"`
	DBName               string           `position:"Query" name:"DBName"`
	BackupMode           string           `position:"Query" name:"BackupMode"`
	CheckDBMode          string           `position:"Query" name:"CheckDBMode"`
}

// CreateMigrateTaskResponse is the response struct for api CreateMigrateTask
type CreateMigrateTaskResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	DBName        string `json:"DBName" xml:"DBName"`
	BackupMode    string `json:"BackupMode" xml:"BackupMode"`
	DBInstanceId  string `json:"DBInstanceId" xml:"DBInstanceId"`
	MigrateTaskId string `json:"MigrateTaskId" xml:"MigrateTaskId"`
	TaskId        string `json:"TaskId" xml:"TaskId"`
}

// CreateCreateMigrateTaskRequest creates a request to invoke CreateMigrateTask API
func CreateCreateMigrateTaskRequest() (request *CreateMigrateTaskRequest) {
	request = &CreateMigrateTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateMigrateTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateMigrateTaskResponse creates a response to parse from CreateMigrateTask response
func CreateCreateMigrateTaskResponse() (response *CreateMigrateTaskResponse) {
	response = &CreateMigrateTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
