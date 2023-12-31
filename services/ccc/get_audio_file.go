package ccc

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

// GetAudioFile invokes the ccc.GetAudioFile API synchronously
func (client *Client) GetAudioFile(request *GetAudioFileRequest) (response *GetAudioFileResponse, err error) {
	response = CreateGetAudioFileResponse()
	err = client.DoAction(request, response)
	return
}

// GetAudioFileWithChan invokes the ccc.GetAudioFile API asynchronously
func (client *Client) GetAudioFileWithChan(request *GetAudioFileRequest) (<-chan *GetAudioFileResponse, <-chan error) {
	responseChan := make(chan *GetAudioFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAudioFile(request)
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

// GetAudioFileWithCallback invokes the ccc.GetAudioFile API asynchronously
func (client *Client) GetAudioFileWithCallback(request *GetAudioFileRequest, callback func(response *GetAudioFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAudioFileResponse
		var err error
		defer close(result)
		response, err = client.GetAudioFile(request)
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

// GetAudioFileRequest is the request struct for api GetAudioFile
type GetAudioFileRequest struct {
	*requests.RpcRequest
	InstanceId      string `position:"Query" name:"InstanceId"`
	AudioResourceId string `position:"Query" name:"AudioResourceId"`
}

// GetAudioFileResponse is the response struct for api GetAudioFile
type GetAudioFileResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetAudioFileRequest creates a request to invoke GetAudioFile API
func CreateGetAudioFileRequest() (request *GetAudioFileRequest) {
	request = &GetAudioFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "GetAudioFile", "CCC", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAudioFileResponse creates a response to parse from GetAudioFile response
func CreateGetAudioFileResponse() (response *GetAudioFileResponse) {
	response = &GetAudioFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
