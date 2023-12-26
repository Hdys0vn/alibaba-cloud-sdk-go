package sas

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

// InstallBackupClient invokes the sas.InstallBackupClient API synchronously
func (client *Client) InstallBackupClient(request *InstallBackupClientRequest) (response *InstallBackupClientResponse, err error) {
	response = CreateInstallBackupClientResponse()
	err = client.DoAction(request, response)
	return
}

// InstallBackupClientWithChan invokes the sas.InstallBackupClient API asynchronously
func (client *Client) InstallBackupClientWithChan(request *InstallBackupClientRequest) (<-chan *InstallBackupClientResponse, <-chan error) {
	responseChan := make(chan *InstallBackupClientResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InstallBackupClient(request)
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

// InstallBackupClientWithCallback invokes the sas.InstallBackupClient API asynchronously
func (client *Client) InstallBackupClientWithCallback(request *InstallBackupClientRequest, callback func(response *InstallBackupClientResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InstallBackupClientResponse
		var err error
		defer close(result)
		response, err = client.InstallBackupClient(request)
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

// InstallBackupClientRequest is the request struct for api InstallBackupClient
type InstallBackupClientRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Uuid            string           `position:"Query" name:"Uuid"`
	SourceIp        string           `position:"Query" name:"SourceIp"`
	UuidList        *[]string        `position:"Query" name:"UuidList"  type:"Repeated"`
	PolicyVersion   string           `position:"Query" name:"PolicyVersion"`
}

// InstallBackupClientResponse is the response struct for api InstallBackupClient
type InstallBackupClientResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateInstallBackupClientRequest creates a request to invoke InstallBackupClient API
func CreateInstallBackupClientRequest() (request *InstallBackupClientRequest) {
	request = &InstallBackupClientRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "InstallBackupClient", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInstallBackupClientResponse creates a response to parse from InstallBackupClient response
func CreateInstallBackupClientResponse() (response *InstallBackupClientResponse) {
	response = &InstallBackupClientResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}