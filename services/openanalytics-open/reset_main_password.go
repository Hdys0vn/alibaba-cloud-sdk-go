package openanalytics_open

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

// ResetMainPassword invokes the openanalytics_open.ResetMainPassword API synchronously
func (client *Client) ResetMainPassword(request *ResetMainPasswordRequest) (response *ResetMainPasswordResponse, err error) {
	response = CreateResetMainPasswordResponse()
	err = client.DoAction(request, response)
	return
}

// ResetMainPasswordWithChan invokes the openanalytics_open.ResetMainPassword API asynchronously
func (client *Client) ResetMainPasswordWithChan(request *ResetMainPasswordRequest) (<-chan *ResetMainPasswordResponse, <-chan error) {
	responseChan := make(chan *ResetMainPasswordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetMainPassword(request)
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

// ResetMainPasswordWithCallback invokes the openanalytics_open.ResetMainPassword API asynchronously
func (client *Client) ResetMainPasswordWithCallback(request *ResetMainPasswordRequest, callback func(response *ResetMainPasswordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetMainPasswordResponse
		var err error
		defer close(result)
		response, err = client.ResetMainPassword(request)
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

// ResetMainPasswordRequest is the request struct for api ResetMainPassword
type ResetMainPasswordRequest struct {
	*requests.RpcRequest
	ExternalUid          string           `position:"Body" name:"ExternalUid"`
	InitPassword         string           `position:"Body" name:"InitPassword"`
	ExternalAliyunUid    string           `position:"Body" name:"ExternalAliyunUid"`
	UseRandomPassword    requests.Boolean `position:"Body" name:"UseRandomPassword"`
	EnableKMS            requests.Boolean `position:"Body" name:"EnableKMS"`
	ExternalBizAliyunUid string           `position:"Body" name:"ExternalBizAliyunUid"`
}

// ResetMainPasswordResponse is the response struct for api ResetMainPassword
type ResetMainPasswordResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	RegionId  string  `json:"RegionId" xml:"RegionId"`
	Account   Account `json:"Account" xml:"Account"`
}

// CreateResetMainPasswordRequest creates a request to invoke ResetMainPassword API
func CreateResetMainPasswordRequest() (request *ResetMainPasswordRequest) {
	request = &ResetMainPasswordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "ResetMainPassword", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResetMainPasswordResponse creates a response to parse from ResetMainPassword response
func CreateResetMainPasswordResponse() (response *ResetMainPasswordResponse) {
	response = &ResetMainPasswordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
