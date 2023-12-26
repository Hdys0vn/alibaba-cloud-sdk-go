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

// AddDcdnIpaDomain invokes the dcdn.AddDcdnIpaDomain API synchronously
func (client *Client) AddDcdnIpaDomain(request *AddDcdnIpaDomainRequest) (response *AddDcdnIpaDomainResponse, err error) {
	response = CreateAddDcdnIpaDomainResponse()
	err = client.DoAction(request, response)
	return
}

// AddDcdnIpaDomainWithChan invokes the dcdn.AddDcdnIpaDomain API asynchronously
func (client *Client) AddDcdnIpaDomainWithChan(request *AddDcdnIpaDomainRequest) (<-chan *AddDcdnIpaDomainResponse, <-chan error) {
	responseChan := make(chan *AddDcdnIpaDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddDcdnIpaDomain(request)
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

// AddDcdnIpaDomainWithCallback invokes the dcdn.AddDcdnIpaDomain API asynchronously
func (client *Client) AddDcdnIpaDomainWithCallback(request *AddDcdnIpaDomainRequest, callback func(response *AddDcdnIpaDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddDcdnIpaDomainResponse
		var err error
		defer close(result)
		response, err = client.AddDcdnIpaDomain(request)
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

// AddDcdnIpaDomainRequest is the request struct for api AddDcdnIpaDomain
type AddDcdnIpaDomainRequest struct {
	*requests.RpcRequest
	Sources         string           `position:"Query" name:"Sources"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	Protocol        string           `position:"Query" name:"Protocol"`
	SecurityToken   string           `position:"Query" name:"SecurityToken"`
	Scope           string           `position:"Query" name:"Scope"`
	TopLevelDomain  string           `position:"Query" name:"TopLevelDomain"`
	OwnerAccount    string           `position:"Query" name:"OwnerAccount"`
	DomainName      string           `position:"Query" name:"DomainName"`
	OwnerId         requests.Integer `position:"Query" name:"OwnerId"`
	CheckUrl        string           `position:"Query" name:"CheckUrl"`
}

// AddDcdnIpaDomainResponse is the response struct for api AddDcdnIpaDomain
type AddDcdnIpaDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddDcdnIpaDomainRequest creates a request to invoke AddDcdnIpaDomain API
func CreateAddDcdnIpaDomainRequest() (request *AddDcdnIpaDomainRequest) {
	request = &AddDcdnIpaDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "AddDcdnIpaDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateAddDcdnIpaDomainResponse creates a response to parse from AddDcdnIpaDomain response
func CreateAddDcdnIpaDomainResponse() (response *AddDcdnIpaDomainResponse) {
	response = &AddDcdnIpaDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
