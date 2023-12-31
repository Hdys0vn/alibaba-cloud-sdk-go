package cloudcallcenter

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

// GetUserByExtension invokes the cloudcallcenter.GetUserByExtension API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getuserbyextension.html
func (client *Client) GetUserByExtension(request *GetUserByExtensionRequest) (response *GetUserByExtensionResponse, err error) {
	response = CreateGetUserByExtensionResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserByExtensionWithChan invokes the cloudcallcenter.GetUserByExtension API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getuserbyextension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserByExtensionWithChan(request *GetUserByExtensionRequest) (<-chan *GetUserByExtensionResponse, <-chan error) {
	responseChan := make(chan *GetUserByExtensionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserByExtension(request)
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

// GetUserByExtensionWithCallback invokes the cloudcallcenter.GetUserByExtension API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getuserbyextension.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserByExtensionWithCallback(request *GetUserByExtensionRequest, callback func(response *GetUserByExtensionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserByExtensionResponse
		var err error
		defer close(result)
		response, err = client.GetUserByExtension(request)
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

// GetUserByExtensionRequest is the request struct for api GetUserByExtension
type GetUserByExtensionRequest struct {
	*requests.RpcRequest
	Extension  string `position:"Query" name:"Extension"`
	InstanceId string `position:"Query" name:"InstanceId"`
}

// GetUserByExtensionResponse is the response struct for api GetUserByExtension
type GetUserByExtensionResponse struct {
	*responses.BaseResponse
	RequestId      string                   `json:"RequestId" xml:"RequestId"`
	Success        bool                     `json:"Success" xml:"Success"`
	Code           string                   `json:"Code" xml:"Code"`
	Message        string                   `json:"Message" xml:"Message"`
	HttpStatusCode int                      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	User           UserInGetUserByExtension `json:"User" xml:"User"`
}

// CreateGetUserByExtensionRequest creates a request to invoke GetUserByExtension API
func CreateGetUserByExtensionRequest() (request *GetUserByExtensionRequest) {
	request = &GetUserByExtensionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetUserByExtension", "", "")
	request.Method = requests.POST
	return
}

// CreateGetUserByExtensionResponse creates a response to parse from GetUserByExtension response
func CreateGetUserByExtensionResponse() (response *GetUserByExtensionResponse) {
	response = &GetUserByExtensionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
