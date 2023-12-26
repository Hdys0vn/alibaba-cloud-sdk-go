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

// DeleteTableTheme invokes the dataworks_public.DeleteTableTheme API synchronously
func (client *Client) DeleteTableTheme(request *DeleteTableThemeRequest) (response *DeleteTableThemeResponse, err error) {
	response = CreateDeleteTableThemeResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTableThemeWithChan invokes the dataworks_public.DeleteTableTheme API asynchronously
func (client *Client) DeleteTableThemeWithChan(request *DeleteTableThemeRequest) (<-chan *DeleteTableThemeResponse, <-chan error) {
	responseChan := make(chan *DeleteTableThemeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTableTheme(request)
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

// DeleteTableThemeWithCallback invokes the dataworks_public.DeleteTableTheme API asynchronously
func (client *Client) DeleteTableThemeWithCallback(request *DeleteTableThemeRequest, callback func(response *DeleteTableThemeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTableThemeResponse
		var err error
		defer close(result)
		response, err = client.DeleteTableTheme(request)
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

// DeleteTableThemeRequest is the request struct for api DeleteTableTheme
type DeleteTableThemeRequest struct {
	*requests.RpcRequest
	ThemeId   requests.Integer `position:"Query" name:"ThemeId"`
	ProjectId requests.Integer `position:"Query" name:"ProjectId"`
}

// DeleteTableThemeResponse is the response struct for api DeleteTableTheme
type DeleteTableThemeResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
	DeleteResult   bool   `json:"DeleteResult" xml:"DeleteResult"`
}

// CreateDeleteTableThemeRequest creates a request to invoke DeleteTableTheme API
func CreateDeleteTableThemeRequest() (request *DeleteTableThemeRequest) {
	request = &DeleteTableThemeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteTableTheme", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTableThemeResponse creates a response to parse from DeleteTableTheme response
func CreateDeleteTableThemeResponse() (response *DeleteTableThemeResponse) {
	response = &DeleteTableThemeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
