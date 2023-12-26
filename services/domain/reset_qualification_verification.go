package domain

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

// ResetQualificationVerification invokes the domain.ResetQualificationVerification API synchronously
func (client *Client) ResetQualificationVerification(request *ResetQualificationVerificationRequest) (response *ResetQualificationVerificationResponse, err error) {
	response = CreateResetQualificationVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// ResetQualificationVerificationWithChan invokes the domain.ResetQualificationVerification API asynchronously
func (client *Client) ResetQualificationVerificationWithChan(request *ResetQualificationVerificationRequest) (<-chan *ResetQualificationVerificationResponse, <-chan error) {
	responseChan := make(chan *ResetQualificationVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResetQualificationVerification(request)
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

// ResetQualificationVerificationWithCallback invokes the domain.ResetQualificationVerification API asynchronously
func (client *Client) ResetQualificationVerificationWithCallback(request *ResetQualificationVerificationRequest, callback func(response *ResetQualificationVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResetQualificationVerificationResponse
		var err error
		defer close(result)
		response, err = client.ResetQualificationVerification(request)
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

// ResetQualificationVerificationRequest is the request struct for api ResetQualificationVerification
type ResetQualificationVerificationRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// ResetQualificationVerificationResponse is the response struct for api ResetQualificationVerification
type ResetQualificationVerificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResetQualificationVerificationRequest creates a request to invoke ResetQualificationVerification API
func CreateResetQualificationVerificationRequest() (request *ResetQualificationVerificationRequest) {
	request = &ResetQualificationVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "ResetQualificationVerification", "", "")
	request.Method = requests.POST
	return
}

// CreateResetQualificationVerificationResponse creates a response to parse from ResetQualificationVerification response
func CreateResetQualificationVerificationResponse() (response *ResetQualificationVerificationResponse) {
	response = &ResetQualificationVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
