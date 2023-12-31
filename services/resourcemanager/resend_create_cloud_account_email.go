package resourcemanager

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

// ResendCreateCloudAccountEmail invokes the resourcemanager.ResendCreateCloudAccountEmail API synchronously
func (client *Client) ResendCreateCloudAccountEmail(request *ResendCreateCloudAccountEmailRequest) (response *ResendCreateCloudAccountEmailResponse, err error) {
	response = CreateResendCreateCloudAccountEmailResponse()
	err = client.DoAction(request, response)
	return
}

// ResendCreateCloudAccountEmailWithChan invokes the resourcemanager.ResendCreateCloudAccountEmail API asynchronously
func (client *Client) ResendCreateCloudAccountEmailWithChan(request *ResendCreateCloudAccountEmailRequest) (<-chan *ResendCreateCloudAccountEmailResponse, <-chan error) {
	responseChan := make(chan *ResendCreateCloudAccountEmailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResendCreateCloudAccountEmail(request)
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

// ResendCreateCloudAccountEmailWithCallback invokes the resourcemanager.ResendCreateCloudAccountEmail API asynchronously
func (client *Client) ResendCreateCloudAccountEmailWithCallback(request *ResendCreateCloudAccountEmailRequest, callback func(response *ResendCreateCloudAccountEmailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResendCreateCloudAccountEmailResponse
		var err error
		defer close(result)
		response, err = client.ResendCreateCloudAccountEmail(request)
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

// ResendCreateCloudAccountEmailRequest is the request struct for api ResendCreateCloudAccountEmail
type ResendCreateCloudAccountEmailRequest struct {
	*requests.RpcRequest
	RecordId string `position:"Query" name:"RecordId"`
}

// ResendCreateCloudAccountEmailResponse is the response struct for api ResendCreateCloudAccountEmail
type ResendCreateCloudAccountEmailResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Account   Account `json:"Account" xml:"Account"`
}

// CreateResendCreateCloudAccountEmailRequest creates a request to invoke ResendCreateCloudAccountEmail API
func CreateResendCreateCloudAccountEmailRequest() (request *ResendCreateCloudAccountEmailRequest) {
	request = &ResendCreateCloudAccountEmailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ResendCreateCloudAccountEmail", "", "")
	request.Method = requests.POST
	return
}

// CreateResendCreateCloudAccountEmailResponse creates a response to parse from ResendCreateCloudAccountEmail response
func CreateResendCreateCloudAccountEmailResponse() (response *ResendCreateCloudAccountEmailResponse) {
	response = &ResendCreateCloudAccountEmailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
