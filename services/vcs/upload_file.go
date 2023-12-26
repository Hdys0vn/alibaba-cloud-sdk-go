package vcs

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

// UploadFile invokes the vcs.UploadFile API synchronously
func (client *Client) UploadFile(request *UploadFileRequest) (response *UploadFileResponse, err error) {
	response = CreateUploadFileResponse()
	err = client.DoAction(request, response)
	return
}

// UploadFileWithChan invokes the vcs.UploadFile API asynchronously
func (client *Client) UploadFileWithChan(request *UploadFileRequest) (<-chan *UploadFileResponse, <-chan error) {
	responseChan := make(chan *UploadFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UploadFile(request)
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

// UploadFileWithCallback invokes the vcs.UploadFile API asynchronously
func (client *Client) UploadFileWithCallback(request *UploadFileRequest, callback func(response *UploadFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UploadFileResponse
		var err error
		defer close(result)
		response, err = client.UploadFile(request)
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

// UploadFileRequest is the request struct for api UploadFile
type UploadFileRequest struct {
	*requests.RpcRequest
	FileType      string `position:"Body" name:"FileType"`
	CorpId        string `position:"Body" name:"CorpId"`
	FileAliasName string `position:"Body" name:"FileAliasName"`
	FileName      string `position:"Body" name:"FileName"`
	FilePath      string `position:"Body" name:"FilePath"`
	FileContent   string `position:"Body" name:"FileContent"`
	DataSourceId  string `position:"Body" name:"DataSourceId"`
	MD5           string `position:"Body" name:"MD5"`
}

// UploadFileResponse is the response struct for api UploadFile
type UploadFileResponse struct {
	*responses.BaseResponse
	Code      string           `json:"Code" xml:"Code"`
	Message   string           `json:"Message" xml:"Message"`
	RequestId string           `json:"RequestId" xml:"RequestId"`
	Data      DataInUploadFile `json:"Data" xml:"Data"`
}

// CreateUploadFileRequest creates a request to invoke UploadFile API
func CreateUploadFileRequest() (request *UploadFileRequest) {
	request = &UploadFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "UploadFile", "", "")
	request.Method = requests.POST
	return
}

// CreateUploadFileResponse creates a response to parse from UploadFile response
func CreateUploadFileResponse() (response *UploadFileResponse) {
	response = &UploadFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
