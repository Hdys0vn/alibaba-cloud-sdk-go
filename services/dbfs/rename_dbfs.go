package dbfs

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

// RenameDbfs invokes the dbfs.RenameDbfs API synchronously
func (client *Client) RenameDbfs(request *RenameDbfsRequest) (response *RenameDbfsResponse, err error) {
	response = CreateRenameDbfsResponse()
	err = client.DoAction(request, response)
	return
}

// RenameDbfsWithChan invokes the dbfs.RenameDbfs API asynchronously
func (client *Client) RenameDbfsWithChan(request *RenameDbfsRequest) (<-chan *RenameDbfsResponse, <-chan error) {
	responseChan := make(chan *RenameDbfsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RenameDbfs(request)
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

// RenameDbfsWithCallback invokes the dbfs.RenameDbfs API asynchronously
func (client *Client) RenameDbfsWithCallback(request *RenameDbfsRequest, callback func(response *RenameDbfsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RenameDbfsResponse
		var err error
		defer close(result)
		response, err = client.RenameDbfs(request)
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

// RenameDbfsRequest is the request struct for api RenameDbfs
type RenameDbfsRequest struct {
	*requests.RpcRequest
	FsName string `position:"Query" name:"FsName"`
	FsId   string `position:"Query" name:"FsId"`
}

// RenameDbfsResponse is the response struct for api RenameDbfs
type RenameDbfsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRenameDbfsRequest creates a request to invoke RenameDbfs API
func CreateRenameDbfsRequest() (request *RenameDbfsRequest) {
	request = &RenameDbfsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DBFS", "2020-04-18", "RenameDbfs", "dbfs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRenameDbfsResponse creates a response to parse from RenameDbfs response
func CreateRenameDbfsResponse() (response *RenameDbfsResponse) {
	response = &RenameDbfsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
