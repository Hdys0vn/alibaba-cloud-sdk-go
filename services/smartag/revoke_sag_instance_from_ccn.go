package smartag

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

// RevokeSagInstanceFromCcn invokes the smartag.RevokeSagInstanceFromCcn API synchronously
func (client *Client) RevokeSagInstanceFromCcn(request *RevokeSagInstanceFromCcnRequest) (response *RevokeSagInstanceFromCcnResponse, err error) {
	response = CreateRevokeSagInstanceFromCcnResponse()
	err = client.DoAction(request, response)
	return
}

// RevokeSagInstanceFromCcnWithChan invokes the smartag.RevokeSagInstanceFromCcn API asynchronously
func (client *Client) RevokeSagInstanceFromCcnWithChan(request *RevokeSagInstanceFromCcnRequest) (<-chan *RevokeSagInstanceFromCcnResponse, <-chan error) {
	responseChan := make(chan *RevokeSagInstanceFromCcnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RevokeSagInstanceFromCcn(request)
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

// RevokeSagInstanceFromCcnWithCallback invokes the smartag.RevokeSagInstanceFromCcn API asynchronously
func (client *Client) RevokeSagInstanceFromCcnWithCallback(request *RevokeSagInstanceFromCcnRequest, callback func(response *RevokeSagInstanceFromCcnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RevokeSagInstanceFromCcnResponse
		var err error
		defer close(result)
		response, err = client.RevokeSagInstanceFromCcn(request)
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

// RevokeSagInstanceFromCcnRequest is the request struct for api RevokeSagInstanceFromCcn
type RevokeSagInstanceFromCcnRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CcnInstanceId        string           `position:"Query" name:"CcnInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
}

// RevokeSagInstanceFromCcnResponse is the response struct for api RevokeSagInstanceFromCcn
type RevokeSagInstanceFromCcnResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRevokeSagInstanceFromCcnRequest creates a request to invoke RevokeSagInstanceFromCcn API
func CreateRevokeSagInstanceFromCcnRequest() (request *RevokeSagInstanceFromCcnRequest) {
	request = &RevokeSagInstanceFromCcnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "RevokeSagInstanceFromCcn", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRevokeSagInstanceFromCcnResponse creates a response to parse from RevokeSagInstanceFromCcn response
func CreateRevokeSagInstanceFromCcnResponse() (response *RevokeSagInstanceFromCcnResponse) {
	response = &RevokeSagInstanceFromCcnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
