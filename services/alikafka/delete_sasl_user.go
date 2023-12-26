package alikafka

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

// DeleteSaslUser invokes the alikafka.DeleteSaslUser API synchronously
func (client *Client) DeleteSaslUser(request *DeleteSaslUserRequest) (response *DeleteSaslUserResponse, err error) {
	response = CreateDeleteSaslUserResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSaslUserWithChan invokes the alikafka.DeleteSaslUser API asynchronously
func (client *Client) DeleteSaslUserWithChan(request *DeleteSaslUserRequest) (<-chan *DeleteSaslUserResponse, <-chan error) {
	responseChan := make(chan *DeleteSaslUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSaslUser(request)
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

// DeleteSaslUserWithCallback invokes the alikafka.DeleteSaslUser API asynchronously
func (client *Client) DeleteSaslUserWithCallback(request *DeleteSaslUserRequest, callback func(response *DeleteSaslUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSaslUserResponse
		var err error
		defer close(result)
		response, err = client.DeleteSaslUser(request)
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

// DeleteSaslUserRequest is the request struct for api DeleteSaslUser
type DeleteSaslUserRequest struct {
	*requests.RpcRequest
	Type       string `position:"Query" name:"Type"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Username   string `position:"Query" name:"Username"`
}

// DeleteSaslUserResponse is the response struct for api DeleteSaslUser
type DeleteSaslUserResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateDeleteSaslUserRequest creates a request to invoke DeleteSaslUser API
func CreateDeleteSaslUserRequest() (request *DeleteSaslUserRequest) {
	request = &DeleteSaslUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "DeleteSaslUser", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteSaslUserResponse creates a response to parse from DeleteSaslUser response
func CreateDeleteSaslUserResponse() (response *DeleteSaslUserResponse) {
	response = &DeleteSaslUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
