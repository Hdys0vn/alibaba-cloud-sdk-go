package resourcemanager

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

// CreateFolder invokes the resourcemanager.CreateFolder API synchronously
func (client *Client) CreateFolder(request *CreateFolderRequest) (response *CreateFolderResponse, err error) {
	response = CreateCreateFolderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFolderWithChan invokes the resourcemanager.CreateFolder API asynchronously
func (client *Client) CreateFolderWithChan(request *CreateFolderRequest) (<-chan *CreateFolderResponse, <-chan error) {
	responseChan := make(chan *CreateFolderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFolder(request)
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

// CreateFolderWithCallback invokes the resourcemanager.CreateFolder API asynchronously
func (client *Client) CreateFolderWithCallback(request *CreateFolderRequest, callback func(response *CreateFolderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFolderResponse
		var err error
		defer close(result)
		response, err = client.CreateFolder(request)
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

// CreateFolderRequest is the request struct for api CreateFolder
type CreateFolderRequest struct {
	*requests.RpcRequest
	FolderName     string `position:"Query" name:"FolderName"`
	ParentFolderId string `position:"Query" name:"ParentFolderId"`
}

// CreateFolderResponse is the response struct for api CreateFolder
type CreateFolderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Folder    Folder `json:"Folder" xml:"Folder"`
}

// CreateCreateFolderRequest creates a request to invoke CreateFolder API
func CreateCreateFolderRequest() (request *CreateFolderRequest) {
	request = &CreateFolderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "CreateFolder", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateFolderResponse creates a response to parse from CreateFolder response
func CreateCreateFolderResponse() (response *CreateFolderResponse) {
	response = &CreateFolderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
