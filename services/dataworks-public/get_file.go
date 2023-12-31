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

// GetFile invokes the dataworks_public.GetFile API synchronously
func (client *Client) GetFile(request *GetFileRequest) (response *GetFileResponse, err error) {
	response = CreateGetFileResponse()
	err = client.DoAction(request, response)
	return
}

// GetFileWithChan invokes the dataworks_public.GetFile API asynchronously
func (client *Client) GetFileWithChan(request *GetFileRequest) (<-chan *GetFileResponse, <-chan error) {
	responseChan := make(chan *GetFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFile(request)
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

// GetFileWithCallback invokes the dataworks_public.GetFile API asynchronously
func (client *Client) GetFileWithCallback(request *GetFileRequest, callback func(response *GetFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFileResponse
		var err error
		defer close(result)
		response, err = client.GetFile(request)
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

// GetFileRequest is the request struct for api GetFile
type GetFileRequest struct {
	*requests.RpcRequest
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	NodeId            requests.Integer `position:"Body" name:"NodeId"`
	ProjectIdentifier string           `position:"Body" name:"ProjectIdentifier"`
	FileId            requests.Integer `position:"Body" name:"FileId"`
}

// GetFileResponse is the response struct for api GetFile
type GetFileResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetFileRequest creates a request to invoke GetFile API
func CreateGetFileRequest() (request *GetFileRequest) {
	request = &GetFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetFile", "", "")
	request.Method = requests.POST
	return
}

// CreateGetFileResponse creates a response to parse from GetFile response
func CreateGetFileResponse() (response *GetFileResponse) {
	response = &GetFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
