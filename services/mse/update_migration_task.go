package mse

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

// UpdateMigrationTask invokes the mse.UpdateMigrationTask API synchronously
func (client *Client) UpdateMigrationTask(request *UpdateMigrationTaskRequest) (response *UpdateMigrationTaskResponse, err error) {
	response = CreateUpdateMigrationTaskResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateMigrationTaskWithChan invokes the mse.UpdateMigrationTask API asynchronously
func (client *Client) UpdateMigrationTaskWithChan(request *UpdateMigrationTaskRequest) (<-chan *UpdateMigrationTaskResponse, <-chan error) {
	responseChan := make(chan *UpdateMigrationTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMigrationTask(request)
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

// UpdateMigrationTaskWithCallback invokes the mse.UpdateMigrationTask API asynchronously
func (client *Client) UpdateMigrationTaskWithCallback(request *UpdateMigrationTaskRequest, callback func(response *UpdateMigrationTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMigrationTaskResponse
		var err error
		defer close(result)
		response, err = client.UpdateMigrationTask(request)
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

// UpdateMigrationTaskRequest is the request struct for api UpdateMigrationTask
type UpdateMigrationTaskRequest struct {
	*requests.RpcRequest
	MseSessionId            string `position:"Query" name:"MseSessionId"`
	TargetClusterUrl        string `position:"Query" name:"TargetClusterUrl"`
	OriginInstanceAddress   string `position:"Query" name:"OriginInstanceAddress"`
	RequestPars             string `position:"Query" name:"RequestPars"`
	Id                      string `position:"Query" name:"Id"`
	OriginInstanceName      string `position:"Query" name:"OriginInstanceName"`
	ProjectDesc             string `position:"Query" name:"ProjectDesc"`
	OriginInstanceNamespace string `position:"Query" name:"OriginInstanceNamespace"`
	ClusterType             string `position:"Query" name:"ClusterType"`
	TargetInstanceId        string `position:"Query" name:"TargetInstanceId"`
	TargetClusterName       string `position:"Query" name:"TargetClusterName"`
	AcceptLanguage          string `position:"Query" name:"AcceptLanguage"`
}

// UpdateMigrationTaskResponse is the response struct for api UpdateMigrationTask
type UpdateMigrationTaskResponse struct {
	*responses.BaseResponse
	HttpCode  string `json:"HttpCode" xml:"HttpCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateUpdateMigrationTaskRequest creates a request to invoke UpdateMigrationTask API
func CreateUpdateMigrationTaskRequest() (request *UpdateMigrationTaskRequest) {
	request = &UpdateMigrationTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "UpdateMigrationTask", "mse", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateMigrationTaskResponse creates a response to parse from UpdateMigrationTask response
func CreateUpdateMigrationTaskResponse() (response *UpdateMigrationTaskResponse) {
	response = &UpdateMigrationTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}