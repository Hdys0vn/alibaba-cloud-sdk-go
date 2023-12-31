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

// SaveTaskForSubmittingDomainDelete invokes the domain.SaveTaskForSubmittingDomainDelete API synchronously
func (client *Client) SaveTaskForSubmittingDomainDelete(request *SaveTaskForSubmittingDomainDeleteRequest) (response *SaveTaskForSubmittingDomainDeleteResponse, err error) {
	response = CreateSaveTaskForSubmittingDomainDeleteResponse()
	err = client.DoAction(request, response)
	return
}

// SaveTaskForSubmittingDomainDeleteWithChan invokes the domain.SaveTaskForSubmittingDomainDelete API asynchronously
func (client *Client) SaveTaskForSubmittingDomainDeleteWithChan(request *SaveTaskForSubmittingDomainDeleteRequest) (<-chan *SaveTaskForSubmittingDomainDeleteResponse, <-chan error) {
	responseChan := make(chan *SaveTaskForSubmittingDomainDeleteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveTaskForSubmittingDomainDelete(request)
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

// SaveTaskForSubmittingDomainDeleteWithCallback invokes the domain.SaveTaskForSubmittingDomainDelete API asynchronously
func (client *Client) SaveTaskForSubmittingDomainDeleteWithCallback(request *SaveTaskForSubmittingDomainDeleteRequest, callback func(response *SaveTaskForSubmittingDomainDeleteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveTaskForSubmittingDomainDeleteResponse
		var err error
		defer close(result)
		response, err = client.SaveTaskForSubmittingDomainDelete(request)
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

// SaveTaskForSubmittingDomainDeleteRequest is the request struct for api SaveTaskForSubmittingDomainDelete
type SaveTaskForSubmittingDomainDeleteRequest struct {
	*requests.RpcRequest
	InstanceId   string `position:"Query" name:"InstanceId"`
	UserClientIp string `position:"Query" name:"UserClientIp"`
	Lang         string `position:"Query" name:"Lang"`
}

// SaveTaskForSubmittingDomainDeleteResponse is the response struct for api SaveTaskForSubmittingDomainDelete
type SaveTaskForSubmittingDomainDeleteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskNo    string `json:"TaskNo" xml:"TaskNo"`
}

// CreateSaveTaskForSubmittingDomainDeleteRequest creates a request to invoke SaveTaskForSubmittingDomainDelete API
func CreateSaveTaskForSubmittingDomainDeleteRequest() (request *SaveTaskForSubmittingDomainDeleteRequest) {
	request = &SaveTaskForSubmittingDomainDeleteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveTaskForSubmittingDomainDelete", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveTaskForSubmittingDomainDeleteResponse creates a response to parse from SaveTaskForSubmittingDomainDelete response
func CreateSaveTaskForSubmittingDomainDeleteResponse() (response *SaveTaskForSubmittingDomainDeleteResponse) {
	response = &SaveTaskForSubmittingDomainDeleteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
