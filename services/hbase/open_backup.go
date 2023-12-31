package hbase

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

// OpenBackup invokes the hbase.OpenBackup API synchronously
func (client *Client) OpenBackup(request *OpenBackupRequest) (response *OpenBackupResponse, err error) {
	response = CreateOpenBackupResponse()
	err = client.DoAction(request, response)
	return
}

// OpenBackupWithChan invokes the hbase.OpenBackup API asynchronously
func (client *Client) OpenBackupWithChan(request *OpenBackupRequest) (<-chan *OpenBackupResponse, <-chan error) {
	responseChan := make(chan *OpenBackupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OpenBackup(request)
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

// OpenBackupWithCallback invokes the hbase.OpenBackup API asynchronously
func (client *Client) OpenBackupWithCallback(request *OpenBackupRequest, callback func(response *OpenBackupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OpenBackupResponse
		var err error
		defer close(result)
		response, err = client.OpenBackup(request)
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

// OpenBackupRequest is the request struct for api OpenBackup
type OpenBackupRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
}

// OpenBackupResponse is the response struct for api OpenBackup
type OpenBackupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateOpenBackupRequest creates a request to invoke OpenBackup API
func CreateOpenBackupRequest() (request *OpenBackupRequest) {
	request = &OpenBackupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "OpenBackup", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOpenBackupResponse creates a response to parse from OpenBackup response
func CreateOpenBackupResponse() (response *OpenBackupResponse) {
	response = &OpenBackupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
