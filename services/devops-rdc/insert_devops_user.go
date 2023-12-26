package devops_rdc

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

// InsertDevopsUser invokes the devops_rdc.InsertDevopsUser API synchronously
func (client *Client) InsertDevopsUser(request *InsertDevopsUserRequest) (response *InsertDevopsUserResponse, err error) {
	response = CreateInsertDevopsUserResponse()
	err = client.DoAction(request, response)
	return
}

// InsertDevopsUserWithChan invokes the devops_rdc.InsertDevopsUser API asynchronously
func (client *Client) InsertDevopsUserWithChan(request *InsertDevopsUserRequest) (<-chan *InsertDevopsUserResponse, <-chan error) {
	responseChan := make(chan *InsertDevopsUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertDevopsUser(request)
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

// InsertDevopsUserWithCallback invokes the devops_rdc.InsertDevopsUser API asynchronously
func (client *Client) InsertDevopsUserWithCallback(request *InsertDevopsUserRequest, callback func(response *InsertDevopsUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertDevopsUserResponse
		var err error
		defer close(result)
		response, err = client.InsertDevopsUser(request)
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

// InsertDevopsUserRequest is the request struct for api InsertDevopsUser
type InsertDevopsUserRequest struct {
	*requests.RpcRequest
	Phone    string `position:"Body" name:"Phone"`
	UserPk   string `position:"Body" name:"UserPk"`
	Email    string `position:"Body" name:"Email"`
	UserName string `position:"Body" name:"UserName"`
}

// InsertDevopsUserResponse is the response struct for api InsertDevopsUser
type InsertDevopsUserResponse struct {
	*responses.BaseResponse
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       string `json:"Object" xml:"Object"`
}

// CreateInsertDevopsUserRequest creates a request to invoke InsertDevopsUser API
func CreateInsertDevopsUserRequest() (request *InsertDevopsUserRequest) {
	request = &InsertDevopsUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "InsertDevopsUser", "", "")
	request.Method = requests.POST
	return
}

// CreateInsertDevopsUserResponse creates a response to parse from InsertDevopsUser response
func CreateInsertDevopsUserResponse() (response *InsertDevopsUserResponse) {
	response = &InsertDevopsUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
