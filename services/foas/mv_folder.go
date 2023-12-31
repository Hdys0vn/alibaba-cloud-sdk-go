package foas

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

// MVFolder invokes the foas.MVFolder API synchronously
func (client *Client) MVFolder(request *MVFolderRequest) (response *MVFolderResponse, err error) {
	response = CreateMVFolderResponse()
	err = client.DoAction(request, response)
	return
}

// MVFolderWithChan invokes the foas.MVFolder API asynchronously
func (client *Client) MVFolderWithChan(request *MVFolderRequest) (<-chan *MVFolderResponse, <-chan error) {
	responseChan := make(chan *MVFolderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MVFolder(request)
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

// MVFolderWithCallback invokes the foas.MVFolder API asynchronously
func (client *Client) MVFolderWithCallback(request *MVFolderRequest, callback func(response *MVFolderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MVFolderResponse
		var err error
		defer close(result)
		response, err = client.MVFolder(request)
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

// MVFolderRequest is the request struct for api MVFolder
type MVFolderRequest struct {
	*requests.RoaRequest
	ProjectName string `position:"Path" name:"projectName"`
	SrcPath     string `position:"Body" name:"srcPath"`
	DestPath    string `position:"Body" name:"destPath"`
}

// MVFolderResponse is the response struct for api MVFolder
type MVFolderResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMVFolderRequest creates a request to invoke MVFolder API
func CreateMVFolderRequest() (request *MVFolderRequest) {
	request = &MVFolderRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "MVFolder", "/api/v2/projects/[projectName]/folders", "foas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateMVFolderResponse creates a response to parse from MVFolder response
func CreateMVFolderResponse() (response *MVFolderResponse) {
	response = &MVFolderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
