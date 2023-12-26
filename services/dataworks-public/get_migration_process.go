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

// GetMigrationProcess invokes the dataworks_public.GetMigrationProcess API synchronously
func (client *Client) GetMigrationProcess(request *GetMigrationProcessRequest) (response *GetMigrationProcessResponse, err error) {
	response = CreateGetMigrationProcessResponse()
	err = client.DoAction(request, response)
	return
}

// GetMigrationProcessWithChan invokes the dataworks_public.GetMigrationProcess API asynchronously
func (client *Client) GetMigrationProcessWithChan(request *GetMigrationProcessRequest) (<-chan *GetMigrationProcessResponse, <-chan error) {
	responseChan := make(chan *GetMigrationProcessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMigrationProcess(request)
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

// GetMigrationProcessWithCallback invokes the dataworks_public.GetMigrationProcess API asynchronously
func (client *Client) GetMigrationProcessWithCallback(request *GetMigrationProcessRequest, callback func(response *GetMigrationProcessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMigrationProcessResponse
		var err error
		defer close(result)
		response, err = client.GetMigrationProcess(request)
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

// GetMigrationProcessRequest is the request struct for api GetMigrationProcess
type GetMigrationProcessRequest struct {
	*requests.RpcRequest
	MigrationId requests.Integer `position:"Body" name:"MigrationId"`
	ProjectId   requests.Integer `position:"Body" name:"ProjectId"`
}

// GetMigrationProcessResponse is the response struct for api GetMigrationProcess
type GetMigrationProcessResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string             `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string             `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool               `json:"Success" xml:"Success"`
	ErrorCode      string             `json:"ErrorCode" xml:"ErrorCode"`
	Data           []ProgressTaskItem `json:"Data" xml:"Data"`
}

// CreateGetMigrationProcessRequest creates a request to invoke GetMigrationProcess API
func CreateGetMigrationProcessRequest() (request *GetMigrationProcessRequest) {
	request = &GetMigrationProcessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetMigrationProcess", "", "")
	request.Method = requests.POST
	return
}

// CreateGetMigrationProcessResponse creates a response to parse from GetMigrationProcess response
func CreateGetMigrationProcessResponse() (response *GetMigrationProcessResponse) {
	response = &GetMigrationProcessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}