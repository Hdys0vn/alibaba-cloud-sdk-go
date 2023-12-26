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

// CreateUser invokes the cloudcallcenter.CreateUser API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createuser.html
func (client *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
	response = CreateCreateUserResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserWithChan invokes the cloudcallcenter.CreateUser API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserWithChan(request *CreateUserRequest) (<-chan *CreateUserResponse, <-chan error) {
	responseChan := make(chan *CreateUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUser(request)
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

// CreateUserWithCallback invokes the cloudcallcenter.CreateUser API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserWithCallback(request *CreateUserRequest, callback func(response *CreateUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserResponse
		var err error
		defer close(result)
		response, err = client.CreateUser(request)
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

// CreateUserRequest is the request struct for api CreateUser
type CreateUserRequest struct {
	*requests.RpcRequest
	PrivateOutboundNumberId string    `position:"Query" name:"PrivateOutboundNumberId"`
	LoginName               string    `position:"Query" name:"LoginName"`
	RoleId                  *[]string `position:"Query" name:"RoleId"  type:"Repeated"`
	SkillLevel              *[]string `position:"Query" name:"SkillLevel"  type:"Repeated"`
	InstanceId              string    `position:"Query" name:"InstanceId"`
	Phone                   string    `position:"Query" name:"Phone"`
	DisplayName             string    `position:"Query" name:"DisplayName"`
	SkillGroupId            *[]string `position:"Query" name:"SkillGroupId"  type:"Repeated"`
	Email                   string    `position:"Query" name:"Email"`
}

// CreateUserResponse is the response struct for api CreateUser
type CreateUserResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	UserId         string `json:"UserId" xml:"UserId"`
}

// CreateCreateUserRequest creates a request to invoke CreateUser API
func CreateCreateUserRequest() (request *CreateUserRequest) {
	request = &CreateUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateUser", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateUserResponse creates a response to parse from CreateUser response
func CreateCreateUserResponse() (response *CreateUserResponse) {
	response = &CreateUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
