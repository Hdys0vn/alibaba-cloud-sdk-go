package avatar

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

// GetVideoTaskInfo invokes the avatar.GetVideoTaskInfo API synchronously
func (client *Client) GetVideoTaskInfo(request *GetVideoTaskInfoRequest) (response *GetVideoTaskInfoResponse, err error) {
	response = CreateGetVideoTaskInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetVideoTaskInfoWithChan invokes the avatar.GetVideoTaskInfo API asynchronously
func (client *Client) GetVideoTaskInfoWithChan(request *GetVideoTaskInfoRequest) (<-chan *GetVideoTaskInfoResponse, <-chan error) {
	responseChan := make(chan *GetVideoTaskInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetVideoTaskInfo(request)
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

// GetVideoTaskInfoWithCallback invokes the avatar.GetVideoTaskInfo API asynchronously
func (client *Client) GetVideoTaskInfoWithCallback(request *GetVideoTaskInfoRequest, callback func(response *GetVideoTaskInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetVideoTaskInfoResponse
		var err error
		defer close(result)
		response, err = client.GetVideoTaskInfo(request)
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

// GetVideoTaskInfoRequest is the request struct for api GetVideoTaskInfo
type GetVideoTaskInfoRequest struct {
	*requests.RpcRequest
	App      GetVideoTaskInfoApp `position:"Query" name:"App"  type:"Struct"`
	TenantId requests.Integer    `position:"Query" name:"TenantId"`
	TaskUuid string              `position:"Query" name:"TaskUuid"`
}

// GetVideoTaskInfoApp is a repeated param struct in GetVideoTaskInfoRequest
type GetVideoTaskInfoApp struct {
	AppId string `name:"AppId"`
}

// GetVideoTaskInfoResponse is the response struct for api GetVideoTaskInfo
type GetVideoTaskInfoResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateGetVideoTaskInfoRequest creates a request to invoke GetVideoTaskInfo API
func CreateGetVideoTaskInfoRequest() (request *GetVideoTaskInfoRequest) {
	request = &GetVideoTaskInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("avatar", "2022-01-30", "GetVideoTaskInfo", "", "")
	request.Method = requests.GET
	return
}

// CreateGetVideoTaskInfoResponse creates a response to parse from GetVideoTaskInfo response
func CreateGetVideoTaskInfoResponse() (response *GetVideoTaskInfoResponse) {
	response = &GetVideoTaskInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}