package dcdn

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

// StopDcdnDomain invokes the dcdn.StopDcdnDomain API synchronously
func (client *Client) StopDcdnDomain(request *StopDcdnDomainRequest) (response *StopDcdnDomainResponse, err error) {
	response = CreateStopDcdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// StopDcdnDomainWithChan invokes the dcdn.StopDcdnDomain API asynchronously
func (client *Client) StopDcdnDomainWithChan(request *StopDcdnDomainRequest) (<-chan *StopDcdnDomainResponse, <-chan error) {
	responseChan := make(chan *StopDcdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopDcdnDomain(request)
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

// StopDcdnDomainWithCallback invokes the dcdn.StopDcdnDomain API asynchronously
func (client *Client) StopDcdnDomainWithCallback(request *StopDcdnDomainRequest, callback func(response *StopDcdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopDcdnDomainResponse
		var err error
		defer close(result)
		response, err = client.StopDcdnDomain(request)
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

// StopDcdnDomainRequest is the request struct for api StopDcdnDomain
type StopDcdnDomainRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// StopDcdnDomainResponse is the response struct for api StopDcdnDomain
type StopDcdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopDcdnDomainRequest creates a request to invoke StopDcdnDomain API
func CreateStopDcdnDomainRequest() (request *StopDcdnDomainRequest) {
	request = &StopDcdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "StopDcdnDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateStopDcdnDomainResponse creates a response to parse from StopDcdnDomain response
func CreateStopDcdnDomainResponse() (response *StopDcdnDomainResponse) {
	response = &StopDcdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
