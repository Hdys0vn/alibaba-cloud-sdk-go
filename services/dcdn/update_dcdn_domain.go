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

// UpdateDcdnDomain invokes the dcdn.UpdateDcdnDomain API synchronously
func (client *Client) UpdateDcdnDomain(request *UpdateDcdnDomainRequest) (response *UpdateDcdnDomainResponse, err error) {
	response = CreateUpdateDcdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDcdnDomainWithChan invokes the dcdn.UpdateDcdnDomain API asynchronously
func (client *Client) UpdateDcdnDomainWithChan(request *UpdateDcdnDomainRequest) (<-chan *UpdateDcdnDomainResponse, <-chan error) {
	responseChan := make(chan *UpdateDcdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDcdnDomain(request)
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

// UpdateDcdnDomainWithCallback invokes the dcdn.UpdateDcdnDomain API asynchronously
func (client *Client) UpdateDcdnDomainWithCallback(request *UpdateDcdnDomainRequest, callback func(response *UpdateDcdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDcdnDomainResponse
		var err error
		defer close(result)
		response, err = client.UpdateDcdnDomain(request)
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

// UpdateDcdnDomainRequest is the request struct for api UpdateDcdnDomain
type UpdateDcdnDomainRequest struct {
	*requests.RpcRequest
	Sources         string           `position:"Query" name:"Sources"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	TopLevelDomain  string           `position:"Query" name:"TopLevelDomain"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
}

// UpdateDcdnDomainResponse is the response struct for api UpdateDcdnDomain
type UpdateDcdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateDcdnDomainRequest creates a request to invoke UpdateDcdnDomain API
func CreateUpdateDcdnDomainRequest() (request *UpdateDcdnDomainRequest) {
	request = &UpdateDcdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "UpdateDcdnDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateDcdnDomainResponse creates a response to parse from UpdateDcdnDomain response
func CreateUpdateDcdnDomainResponse() (response *UpdateDcdnDomainResponse) {
	response = &UpdateDcdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
