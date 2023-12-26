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

// DescribeMigrateTaskById invokes the rds.DescribeMigrateTaskById API synchronously
func (client *Client) DescribeMigrateTaskById(request *DescribeMigrateTaskByIdRequest) (response *DescribeMigrateTaskByIdResponse, err error) {
	response = CreateDescribeMigrateTaskByIdResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeMigrateTaskByIdWithChan invokes the rds.DescribeMigrateTaskById API asynchronously
func (client *Client) DescribeMigrateTaskByIdWithChan(request *DescribeMigrateTaskByIdRequest) (<-chan *DescribeMigrateTaskByIdResponse, <-chan error) {
	responseChan := make(chan *DescribeMigrateTaskByIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeMigrateTaskById(request)
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

// DescribeMigrateTaskByIdWithCallback invokes the rds.DescribeMigrateTaskById API asynchronously
func (client *Client) DescribeMigrateTaskByIdWithCallback(request *DescribeMigrateTaskByIdRequest, callback func(response *DescribeMigrateTaskByIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeMigrateTaskByIdResponse
		var err error
		defer close(result)
		response, err = client.DescribeMigrateTaskById(request)
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

// DescribeMigrateTaskByIdRequest is the request struct for api DescribeMigrateTaskById
type DescribeMigrateTaskByIdRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MigrateTaskId        string           `position:"Query" name:"MigrateTaskId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
}

// DescribeMigrateTaskByIdResponse is the response struct for api DescribeMigrateTaskById
type DescribeMigrateTaskByIdResponse struct {
	*responses.BaseResponse
	Status         string `json:"Status" xml:"Status"`
	EndTime        string `json:"EndTime" xml:"EndTime"`
	DBInstanceName string `json:"DBInstanceName" xml:"DBInstanceName"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Description    string `json:"Description" xml:"Description"`
	CreateTime     string `json:"CreateTime" xml:"CreateTime"`
	DBName         string `json:"DBName" xml:"DBName"`
	BackupMode     string `json:"BackupMode" xml:"BackupMode"`
	MigrateTaskId  string `json:"MigrateTaskId" xml:"MigrateTaskId"`
	IsDBReplaced   string `json:"IsDBReplaced" xml:"IsDBReplaced"`
}

// CreateDescribeMigrateTaskByIdRequest creates a request to invoke DescribeMigrateTaskById API
func CreateDescribeMigrateTaskByIdRequest() (request *DescribeMigrateTaskByIdRequest) {
	request = &DescribeMigrateTaskByIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeMigrateTaskById", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeMigrateTaskByIdResponse creates a response to parse from DescribeMigrateTaskById response
func CreateDescribeMigrateTaskByIdResponse() (response *DescribeMigrateTaskByIdResponse) {
	response = &DescribeMigrateTaskByIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}