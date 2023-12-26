package nas

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

// UpgradeFileSystem invokes the nas.UpgradeFileSystem API synchronously
func (client *Client) UpgradeFileSystem(request *UpgradeFileSystemRequest) (response *UpgradeFileSystemResponse, err error) {
	response = CreateUpgradeFileSystemResponse()
	err = client.DoAction(request, response)
	return
}

// UpgradeFileSystemWithChan invokes the nas.UpgradeFileSystem API asynchronously
func (client *Client) UpgradeFileSystemWithChan(request *UpgradeFileSystemRequest) (<-chan *UpgradeFileSystemResponse, <-chan error) {
	responseChan := make(chan *UpgradeFileSystemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpgradeFileSystem(request)
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

// UpgradeFileSystemWithCallback invokes the nas.UpgradeFileSystem API asynchronously
func (client *Client) UpgradeFileSystemWithCallback(request *UpgradeFileSystemRequest, callback func(response *UpgradeFileSystemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpgradeFileSystemResponse
		var err error
		defer close(result)
		response, err = client.UpgradeFileSystem(request)
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

// UpgradeFileSystemRequest is the request struct for api UpgradeFileSystem
type UpgradeFileSystemRequest struct {
	*requests.RpcRequest
	ClientToken  string           `position:"Query" name:"ClientToken"`
	Capacity     requests.Integer `position:"Query" name:"Capacity"`
	FileSystemId string           `position:"Query" name:"FileSystemId"`
	DryRun       requests.Boolean `position:"Query" name:"DryRun"`
}

// UpgradeFileSystemResponse is the response struct for api UpgradeFileSystem
type UpgradeFileSystemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpgradeFileSystemRequest creates a request to invoke UpgradeFileSystem API
func CreateUpgradeFileSystemRequest() (request *UpgradeFileSystemRequest) {
	request = &UpgradeFileSystemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "UpgradeFileSystem", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpgradeFileSystemResponse creates a response to parse from UpgradeFileSystem response
func CreateUpgradeFileSystemResponse() (response *UpgradeFileSystemResponse) {
	response = &UpgradeFileSystemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
