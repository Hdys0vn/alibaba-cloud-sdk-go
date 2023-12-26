package linkvisual

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

// DeleteFaceUser invokes the linkvisual.DeleteFaceUser API synchronously
func (client *Client) DeleteFaceUser(request *DeleteFaceUserRequest) (response *DeleteFaceUserResponse, err error) {
	response = CreateDeleteFaceUserResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFaceUserWithChan invokes the linkvisual.DeleteFaceUser API asynchronously
func (client *Client) DeleteFaceUserWithChan(request *DeleteFaceUserRequest) (<-chan *DeleteFaceUserResponse, <-chan error) {
	responseChan := make(chan *DeleteFaceUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFaceUser(request)
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

// DeleteFaceUserWithCallback invokes the linkvisual.DeleteFaceUser API asynchronously
func (client *Client) DeleteFaceUserWithCallback(request *DeleteFaceUserRequest, callback func(response *DeleteFaceUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFaceUserResponse
		var err error
		defer close(result)
		response, err = client.DeleteFaceUser(request)
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

// DeleteFaceUserRequest is the request struct for api DeleteFaceUser
type DeleteFaceUserRequest struct {
	*requests.RpcRequest
	IsolationId string `position:"Query" name:"IsolationId"`
	UserId      string `position:"Query" name:"UserId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// DeleteFaceUserResponse is the response struct for api DeleteFaceUser
type DeleteFaceUserResponse struct {
	*responses.BaseResponse
	Code         string `json:"Code" xml:"Code"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateDeleteFaceUserRequest creates a request to invoke DeleteFaceUser API
func CreateDeleteFaceUserRequest() (request *DeleteFaceUserRequest) {
	request = &DeleteFaceUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Linkvisual", "2018-01-20", "DeleteFaceUser", "Linkvisual", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFaceUserResponse creates a response to parse from DeleteFaceUser response
func CreateDeleteFaceUserResponse() (response *DeleteFaceUserResponse) {
	response = &DeleteFaceUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
