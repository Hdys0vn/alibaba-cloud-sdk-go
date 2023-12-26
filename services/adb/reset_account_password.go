package adb

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

// ResetAccountPassword invokes the adb.ResetAccountPassword API synchronously
func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (response *ResetAccountPasswordResponse, err error) {
	response = CreateResetAccountPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ResetAccountPasswordWithChan invokes the adb.ResetAccountPassword API asynchronously
func (client *Client) ResetAccountPasswordWithChan(request *ResetAccountPasswordRequest) (<-chan *ResetAccountPasswordResponse, <-chan error) {
	responseChan := make(chan *ResetAccountPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetAccountPassword(request)
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

// ResetAccountPasswordWithCallback invokes the adb.ResetAccountPassword API asynchronously
func (client *Client) ResetAccountPasswordWithCallback(request *ResetAccountPasswordRequest, callback func(response *ResetAccountPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetAccountPasswordResponse
		var err error
		defer close(result)
		response, err = client.ResetAccountPassword(request)
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

// ResetAccountPasswordRequest is the request struct for api ResetAccountPassword
type ResetAccountPasswordRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccountType          string           `position:"Query" name:"AccountType"`
	AccountName          string           `position:"Query" name:"AccountName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AccountPassword      string           `position:"Query" name:"AccountPassword"`
}

// ResetAccountPasswordResponse is the response struct for api ResetAccountPassword
type ResetAccountPasswordResponse struct {
	*responses.BaseResponse
	TaskId      int    `json:"TaskId" xml:"TaskId"`
	RequestId   string `json:"RequestId" xml:"RequestId"`
	DBClusterId string `json:"DBClusterId" xml:"DBClusterId"`
}

// CreateResetAccountPasswordRequest creates a request to invoke ResetAccountPassword API
func CreateResetAccountPasswordRequest() (request *ResetAccountPasswordRequest) {
	request = &ResetAccountPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "ResetAccountPassword", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetAccountPasswordResponse creates a response to parse from ResetAccountPassword response
func CreateResetAccountPasswordResponse() (response *ResetAccountPasswordResponse) {
	response = &ResetAccountPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
