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

// GetUploadAudioDataParams invokes the ccc.GetUploadAudioDataParams API synchronously
func (client *Client) GetUploadAudioDataParams(request *GetUploadAudioDataParamsRequest) (response *GetUploadAudioDataParamsResponse, err error) {
	response = CreateGetUploadAudioDataParamsResponse()
	err = client.DoAction(request, response)
	return
}

// GetUploadAudioDataParamsWithChan invokes the ccc.GetUploadAudioDataParams API asynchronously
func (client *Client) GetUploadAudioDataParamsWithChan(request *GetUploadAudioDataParamsRequest) (<-chan *GetUploadAudioDataParamsResponse, <-chan error) {
	responseChan := make(chan *GetUploadAudioDataParamsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUploadAudioDataParams(request)
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

// GetUploadAudioDataParamsWithCallback invokes the ccc.GetUploadAudioDataParams API asynchronously
func (client *Client) GetUploadAudioDataParamsWithCallback(request *GetUploadAudioDataParamsRequest, callback func(response *GetUploadAudioDataParamsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUploadAudioDataParamsResponse
		var err error
		defer close(result)
		response, err = client.GetUploadAudioDataParams(request)
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

// GetUploadAudioDataParamsRequest is the request struct for api GetUploadAudioDataParams
type GetUploadAudioDataParamsRequest struct {
	*requests.RpcRequest
	ContactId  string `position:"Query" name:"ContactId"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetUploadAudioDataParamsResponse is the response struct for api GetUploadAudioDataParams
type GetUploadAudioDataParamsResponse struct {
	*responses.BaseResponse
	Code           string `json:"Code" xml:"Code"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetUploadAudioDataParamsRequest creates a request to invoke GetUploadAudioDataParams API
func CreateGetUploadAudioDataParamsRequest() (request *GetUploadAudioDataParamsRequest) {
	request = &GetUploadAudioDataParamsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2020-07-01", "GetUploadAudioDataParams", "CCC", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetUploadAudioDataParamsResponse creates a response to parse from GetUploadAudioDataParams response
func CreateGetUploadAudioDataParamsResponse() (response *GetUploadAudioDataParamsResponse) {
	response = &GetUploadAudioDataParamsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
