package cloudwf

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

// UserDataCreate invokes the cloudwf.UserDataCreate API synchronously
// api document: https://help.aliyun.com/api/cloudwf/userdatacreate.html
func (client *Client) UserDataCreate(request *UserDataCreateRequest) (response *UserDataCreateResponse, err error) {
	response = CreateUserDataCreateResponse()
	err = client.DoAction(request, response)
	return
}

// UserDataCreateWithChan invokes the cloudwf.UserDataCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/userdatacreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UserDataCreateWithChan(request *UserDataCreateRequest) (<-chan *UserDataCreateResponse, <-chan error) {
	responseChan := make(chan *UserDataCreateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UserDataCreate(request)
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

// UserDataCreateWithCallback invokes the cloudwf.UserDataCreate API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/userdatacreate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UserDataCreateWithCallback(request *UserDataCreateRequest, callback func(response *UserDataCreateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UserDataCreateResponse
		var err error
		defer close(result)
		response, err = client.UserDataCreate(request)
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

// UserDataCreateRequest is the request struct for api UserDataCreate
type UserDataCreateRequest struct {
	*requests.RpcRequest
	UploadFile string           `position:"Query" name:"UploadFile"`
	Name       string           `position:"Query" name:"Name"`
	Bid        requests.Integer `position:"Query" name:"Bid"`
	Type       string           `position:"Query" name:"Type"`
}

// UserDataCreateResponse is the response struct for api UserDataCreate
type UserDataCreateResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateUserDataCreateRequest creates a request to invoke UserDataCreate API
func CreateUserDataCreateRequest() (request *UserDataCreateRequest) {
	request = &UserDataCreateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "UserDataCreate", "cloudwf", "openAPI")
	return
}

// CreateUserDataCreateResponse creates a response to parse from UserDataCreate response
func CreateUserDataCreateResponse() (response *UserDataCreateResponse) {
	response = &UserDataCreateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
