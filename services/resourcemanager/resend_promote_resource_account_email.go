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

// ResendPromoteResourceAccountEmail invokes the resourcemanager.ResendPromoteResourceAccountEmail API synchronously
func (client *Client) ResendPromoteResourceAccountEmail(request *ResendPromoteResourceAccountEmailRequest) (response *ResendPromoteResourceAccountEmailResponse, err error) {
	response = CreateResendPromoteResourceAccountEmailResponse()
	err = client.DoAction(request, response)
	return
}

// ResendPromoteResourceAccountEmailWithChan invokes the resourcemanager.ResendPromoteResourceAccountEmail API asynchronously
func (client *Client) ResendPromoteResourceAccountEmailWithChan(request *ResendPromoteResourceAccountEmailRequest) (<-chan *ResendPromoteResourceAccountEmailResponse, <-chan error) {
	responseChan := make(chan *ResendPromoteResourceAccountEmailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResendPromoteResourceAccountEmail(request)
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

// ResendPromoteResourceAccountEmailWithCallback invokes the resourcemanager.ResendPromoteResourceAccountEmail API asynchronously
func (client *Client) ResendPromoteResourceAccountEmailWithCallback(request *ResendPromoteResourceAccountEmailRequest, callback func(response *ResendPromoteResourceAccountEmailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResendPromoteResourceAccountEmailResponse
		var err error
		defer close(result)
		response, err = client.ResendPromoteResourceAccountEmail(request)
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

// ResendPromoteResourceAccountEmailRequest is the request struct for api ResendPromoteResourceAccountEmail
type ResendPromoteResourceAccountEmailRequest struct {
	*requests.RpcRequest
	RecordId string `position:"Query" name:"RecordId"`
}

// ResendPromoteResourceAccountEmailResponse is the response struct for api ResendPromoteResourceAccountEmail
type ResendPromoteResourceAccountEmailResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Account   Account `json:"Account" xml:"Account"`
}

// CreateResendPromoteResourceAccountEmailRequest creates a request to invoke ResendPromoteResourceAccountEmail API
func CreateResendPromoteResourceAccountEmailRequest() (request *ResendPromoteResourceAccountEmailRequest) {
	request = &ResendPromoteResourceAccountEmailRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ResendPromoteResourceAccountEmail", "", "")
	request.Method = requests.POST
	return
}

// CreateResendPromoteResourceAccountEmailResponse creates a response to parse from ResendPromoteResourceAccountEmail response
func CreateResendPromoteResourceAccountEmailResponse() (response *ResendPromoteResourceAccountEmailResponse) {
	response = &ResendPromoteResourceAccountEmailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
