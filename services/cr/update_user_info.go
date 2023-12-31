package cr

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

// UpdateUserInfo invokes the cr.UpdateUserInfo API synchronously
func (client *Client) UpdateUserInfo(request *UpdateUserInfoRequest) (response *UpdateUserInfoResponse, err error) {
	response = CreateUpdateUserInfoResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUserInfoWithChan invokes the cr.UpdateUserInfo API asynchronously
func (client *Client) UpdateUserInfoWithChan(request *UpdateUserInfoRequest) (<-chan *UpdateUserInfoResponse, <-chan error) {
	responseChan := make(chan *UpdateUserInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUserInfo(request)
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

// UpdateUserInfoWithCallback invokes the cr.UpdateUserInfo API asynchronously
func (client *Client) UpdateUserInfoWithCallback(request *UpdateUserInfoRequest, callback func(response *UpdateUserInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUserInfoResponse
		var err error
		defer close(result)
		response, err = client.UpdateUserInfo(request)
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

// UpdateUserInfoRequest is the request struct for api UpdateUserInfo
type UpdateUserInfoRequest struct {
	*requests.RoaRequest
}

// UpdateUserInfoResponse is the response struct for api UpdateUserInfo
type UpdateUserInfoResponse struct {
	*responses.BaseResponse
}

// CreateUpdateUserInfoRequest creates a request to invoke UpdateUserInfo API
func CreateUpdateUserInfoRequest() (request *UpdateUserInfoRequest) {
	request = &UpdateUserInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("cr", "2016-06-07", "UpdateUserInfo", "/users", "acr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateUserInfoResponse creates a response to parse from UpdateUserInfo response
func CreateUpdateUserInfoResponse() (response *UpdateUserInfoResponse) {
	response = &UpdateUserInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
