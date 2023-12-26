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

// CreateDataServiceFolder invokes the dataworks_public.CreateDataServiceFolder API synchronously
func (client *Client) CreateDataServiceFolder(request *CreateDataServiceFolderRequest) (response *CreateDataServiceFolderResponse, err error) {
	response = CreateCreateDataServiceFolderResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDataServiceFolderWithChan invokes the dataworks_public.CreateDataServiceFolder API asynchronously
func (client *Client) CreateDataServiceFolderWithChan(request *CreateDataServiceFolderRequest) (<-chan *CreateDataServiceFolderResponse, <-chan error) {
	responseChan := make(chan *CreateDataServiceFolderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDataServiceFolder(request)
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

// CreateDataServiceFolderWithCallback invokes the dataworks_public.CreateDataServiceFolder API asynchronously
func (client *Client) CreateDataServiceFolderWithCallback(request *CreateDataServiceFolderRequest, callback func(response *CreateDataServiceFolderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDataServiceFolderResponse
		var err error
		defer close(result)
		response, err = client.CreateDataServiceFolder(request)
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

// CreateDataServiceFolderRequest is the request struct for api CreateDataServiceFolder
type CreateDataServiceFolderRequest struct {
	*requests.RpcRequest
	GroupId    string           `position:"Body" name:"GroupId"`
	FolderName string           `position:"Body" name:"FolderName"`
	ParentId   requests.Integer `position:"Body" name:"ParentId"`
	TenantId   requests.Integer `position:"Body" name:"TenantId"`
	ProjectId  requests.Integer `position:"Body" name:"ProjectId"`
}

// CreateDataServiceFolderResponse is the response struct for api CreateDataServiceFolder
type CreateDataServiceFolderResponse struct {
	*responses.BaseResponse
	FolderId  int64  `json:"FolderId" xml:"FolderId"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateDataServiceFolderRequest creates a request to invoke CreateDataServiceFolder API
func CreateCreateDataServiceFolderRequest() (request *CreateDataServiceFolderRequest) {
	request = &CreateDataServiceFolderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "CreateDataServiceFolder", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDataServiceFolderResponse creates a response to parse from CreateDataServiceFolder response
func CreateCreateDataServiceFolderResponse() (response *CreateDataServiceFolderResponse) {
	response = &CreateDataServiceFolderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
