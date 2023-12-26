package dypnsapi

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

// VerifySmsCode invokes the dypnsapi.VerifySmsCode API synchronously
func (client *Client) VerifySmsCode(request *VerifySmsCodeRequest) (response *VerifySmsCodeResponse, err error) {
	response = CreateVerifySmsCodeResponse()
	err = client.DoAction(request, response)
	return
}

// VerifySmsCodeWithChan invokes the dypnsapi.VerifySmsCode API asynchronously
func (client *Client) VerifySmsCodeWithChan(request *VerifySmsCodeRequest) (<-chan *VerifySmsCodeResponse, <-chan error) {
	responseChan := make(chan *VerifySmsCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VerifySmsCode(request)
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

// VerifySmsCodeWithCallback invokes the dypnsapi.VerifySmsCode API asynchronously
func (client *Client) VerifySmsCodeWithCallback(request *VerifySmsCodeRequest, callback func(response *VerifySmsCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VerifySmsCodeResponse
		var err error
		defer close(result)
		response, err = client.VerifySmsCode(request)
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

// VerifySmsCodeRequest is the request struct for api VerifySmsCode
type VerifySmsCodeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ProductCode          string           `position:"Query" name:"ProductCode"`
	SmsToken             string           `position:"Query" name:"SmsToken"`
	PhoneNumber          string           `position:"Query" name:"PhoneNumber"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmsCode              string           `position:"Query" name:"SmsCode"`
}

// VerifySmsCodeResponse is the response struct for api VerifySmsCode
type VerifySmsCodeResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      bool   `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateVerifySmsCodeRequest creates a request to invoke VerifySmsCode API
func CreateVerifySmsCodeRequest() (request *VerifySmsCodeRequest) {
	request = &VerifySmsCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dypnsapi", "2017-05-25", "VerifySmsCode", "dypnsapi", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVerifySmsCodeResponse creates a response to parse from VerifySmsCode response
func CreateVerifySmsCodeResponse() (response *VerifySmsCodeResponse) {
	response = &VerifySmsCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
