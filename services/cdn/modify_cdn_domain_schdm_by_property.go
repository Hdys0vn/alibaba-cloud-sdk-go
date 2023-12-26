package cdn

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

// ModifyCdnDomainSchdmByProperty invokes the cdn.ModifyCdnDomainSchdmByProperty API synchronously
func (client *Client) ModifyCdnDomainSchdmByProperty(request *ModifyCdnDomainSchdmByPropertyRequest) (response *ModifyCdnDomainSchdmByPropertyResponse, err error) {
	response = CreateModifyCdnDomainSchdmByPropertyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCdnDomainSchdmByPropertyWithChan invokes the cdn.ModifyCdnDomainSchdmByProperty API asynchronously
func (client *Client) ModifyCdnDomainSchdmByPropertyWithChan(request *ModifyCdnDomainSchdmByPropertyRequest) (<-chan *ModifyCdnDomainSchdmByPropertyResponse, <-chan error) {
	responseChan := make(chan *ModifyCdnDomainSchdmByPropertyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCdnDomainSchdmByProperty(request)
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

// ModifyCdnDomainSchdmByPropertyWithCallback invokes the cdn.ModifyCdnDomainSchdmByProperty API asynchronously
func (client *Client) ModifyCdnDomainSchdmByPropertyWithCallback(request *ModifyCdnDomainSchdmByPropertyRequest, callback func(response *ModifyCdnDomainSchdmByPropertyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCdnDomainSchdmByPropertyResponse
		var err error
		defer close(result)
		response, err = client.ModifyCdnDomainSchdmByProperty(request)
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

// ModifyCdnDomainSchdmByPropertyRequest is the request struct for api ModifyCdnDomainSchdmByProperty
type ModifyCdnDomainSchdmByPropertyRequest struct {
	*requests.RpcRequest
	DomainName string `position:"Query" name:"DomainName"`
	Property   string `position:"Query" name:"Property"`
}

// ModifyCdnDomainSchdmByPropertyResponse is the response struct for api ModifyCdnDomainSchdmByProperty
type ModifyCdnDomainSchdmByPropertyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCdnDomainSchdmByPropertyRequest creates a request to invoke ModifyCdnDomainSchdmByProperty API
func CreateModifyCdnDomainSchdmByPropertyRequest() (request *ModifyCdnDomainSchdmByPropertyRequest) {
	request = &ModifyCdnDomainSchdmByPropertyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "ModifyCdnDomainSchdmByProperty", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyCdnDomainSchdmByPropertyResponse creates a response to parse from ModifyCdnDomainSchdmByProperty response
func CreateModifyCdnDomainSchdmByPropertyResponse() (response *ModifyCdnDomainSchdmByPropertyResponse) {
	response = &ModifyCdnDomainSchdmByPropertyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
