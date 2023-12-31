package emas_appmonitor

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

// SaveAppFile invokes the emas_appmonitor.SaveAppFile API synchronously
func (client *Client) SaveAppFile(request *SaveAppFileRequest) (response *SaveAppFileResponse, err error) {
	response = CreateSaveAppFileResponse()
	err = client.DoAction(request, response)
	return
}

// SaveAppFileWithChan invokes the emas_appmonitor.SaveAppFile API asynchronously
func (client *Client) SaveAppFileWithChan(request *SaveAppFileRequest) (<-chan *SaveAppFileResponse, <-chan error) {
	responseChan := make(chan *SaveAppFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveAppFile(request)
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

// SaveAppFileWithCallback invokes the emas_appmonitor.SaveAppFile API asynchronously
func (client *Client) SaveAppFileWithCallback(request *SaveAppFileRequest, callback func(response *SaveAppFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveAppFileResponse
		var err error
		defer close(result)
		response, err = client.SaveAppFile(request)
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

// SaveAppFileRequest is the request struct for api SaveAppFile
type SaveAppFileRequest struct {
	*requests.RpcRequest
	FileType    string `position:"Body" name:"FileType"`
	FilePath    string `position:"Body" name:"FilePath"`
	UniqueAppId string `position:"Body" name:"UniqueAppId"`
	AppVersion  string `position:"Body" name:"AppVersion"`
	FileName    string `position:"Body" name:"FileName"`
}

// SaveAppFileResponse is the response struct for api SaveAppFile
type SaveAppFileResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	AppFile   AppFile `json:"AppFile" xml:"AppFile"`
}

// CreateSaveAppFileRequest creates a request to invoke SaveAppFile API
func CreateSaveAppFileRequest() (request *SaveAppFileRequest) {
	request = &SaveAppFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("emas-appmonitor", "2019-06-11", "SaveAppFile", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveAppFileResponse creates a response to parse from SaveAppFile response
func CreateSaveAppFileResponse() (response *SaveAppFileResponse) {
	response = &SaveAppFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
