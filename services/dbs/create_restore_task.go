package dbs

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

// CreateRestoreTask invokes the dbs.CreateRestoreTask API synchronously
func (client *Client) CreateRestoreTask(request *CreateRestoreTaskRequest) (response *CreateRestoreTaskResponse, err error) {
	response = CreateCreateRestoreTaskResponse()
	err = client.DoAction(request, response)
	return
}

// CreateRestoreTaskWithChan invokes the dbs.CreateRestoreTask API asynchronously
func (client *Client) CreateRestoreTaskWithChan(request *CreateRestoreTaskRequest) (<-chan *CreateRestoreTaskResponse, <-chan error) {
	responseChan := make(chan *CreateRestoreTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateRestoreTask(request)
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

// CreateRestoreTaskWithCallback invokes the dbs.CreateRestoreTask API asynchronously
func (client *Client) CreateRestoreTaskWithCallback(request *CreateRestoreTaskRequest, callback func(response *CreateRestoreTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateRestoreTaskResponse
		var err error
		defer close(result)
		response, err = client.CreateRestoreTask(request)
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

// CreateRestoreTaskRequest is the request struct for api CreateRestoreTask
type CreateRestoreTaskRequest struct {
	*requests.RpcRequest
	BackupGatewayId                 requests.Integer `position:"Query" name:"BackupGatewayId"`
	CrossRoleName                   string           `position:"Query" name:"CrossRoleName"`
	DestinationEndpointUserName     string           `position:"Query" name:"DestinationEndpointUserName"`
	RestoreTaskName                 string           `position:"Query" name:"RestoreTaskName"`
	DestinationEndpointOracleSID    string           `position:"Query" name:"DestinationEndpointOracleSID"`
	DestinationEndpointPort         requests.Integer `position:"Query" name:"DestinationEndpointPort"`
	BackupSetId                     string           `position:"Query" name:"BackupSetId"`
	OwnerId                         string           `position:"Query" name:"OwnerId"`
	RestoreDir                      string           `position:"Query" name:"RestoreDir"`
	DestinationEndpointIP           string           `position:"Query" name:"DestinationEndpointIP"`
	DuplicateConflict               string           `position:"Query" name:"DuplicateConflict"`
	DestinationEndpointInstanceType string           `position:"Query" name:"DestinationEndpointInstanceType"`
	ClientToken                     string           `position:"Query" name:"ClientToken"`
	BackupPlanId                    string           `position:"Query" name:"BackupPlanId"`
	DestinationEndpointRegion       string           `position:"Query" name:"DestinationEndpointRegion"`
	RestoreObjects                  string           `position:"Query" name:"RestoreObjects"`
	RestoreHome                     string           `position:"Query" name:"RestoreHome"`
	RestoreTime                     requests.Integer `position:"Query" name:"RestoreTime"`
	CrossAliyunId                   string           `position:"Query" name:"CrossAliyunId"`
	DestinationEndpointInstanceID   string           `position:"Query" name:"DestinationEndpointInstanceID"`
	DestinationEndpointDatabaseName string           `position:"Query" name:"DestinationEndpointDatabaseName"`
	DestinationEndpointPassword     string           `position:"Query" name:"DestinationEndpointPassword"`
}

// CreateRestoreTaskResponse is the response struct for api CreateRestoreTask
type CreateRestoreTaskResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrCode        string `json:"ErrCode" xml:"ErrCode"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrMessage     string `json:"ErrMessage" xml:"ErrMessage"`
	RestoreTaskId  string `json:"RestoreTaskId" xml:"RestoreTaskId"`
}

// CreateCreateRestoreTaskRequest creates a request to invoke CreateRestoreTask API
func CreateCreateRestoreTaskRequest() (request *CreateRestoreTaskRequest) {
	request = &CreateRestoreTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "CreateRestoreTask", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateRestoreTaskResponse creates a response to parse from CreateRestoreTask response
func CreateCreateRestoreTaskResponse() (response *CreateRestoreTaskResponse) {
	response = &CreateRestoreTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
