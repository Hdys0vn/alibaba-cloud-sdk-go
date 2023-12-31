package vod

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

// AddVodDomain invokes the vod.AddVodDomain API synchronously
func (client *Client) AddVodDomain(request *AddVodDomainRequest) (response *AddVodDomainResponse, err error) {
	response = CreateAddVodDomainResponse()
	err = client.DoAction(request, response)
	return
}

// AddVodDomainWithChan invokes the vod.AddVodDomain API asynchronously
func (client *Client) AddVodDomainWithChan(request *AddVodDomainRequest) (<-chan *AddVodDomainResponse, <-chan error) {
	responseChan := make(chan *AddVodDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddVodDomain(request)
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

// AddVodDomainWithCallback invokes the vod.AddVodDomain API asynchronously
func (client *Client) AddVodDomainWithCallback(request *AddVodDomainRequest, callback func(response *AddVodDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddVodDomainResponse
		var err error
		defer close(result)
		response, err = client.AddVodDomain(request)
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

// AddVodDomainRequest is the request struct for api AddVodDomain
type AddVodDomainRequest struct {
	*requests.RpcRequest
	Sources        string           `position:"Query" name:"Sources"`
	SecurityToken  string           `position:"Query" name:"SecurityToken"`
	Scope          string           `position:"Query" name:"Scope"`
	TopLevelDomain string           `position:"Query" name:"TopLevelDomain"`
	OwnerAccount   string           `position:"Query" name:"OwnerAccount"`
	DomainName     string           `position:"Query" name:"DomainName"`
	OwnerId        requests.Integer `position:"Query" name:"OwnerId"`
	CheckUrl       string           `position:"Query" name:"CheckUrl"`
}

// AddVodDomainResponse is the response struct for api AddVodDomain
type AddVodDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddVodDomainRequest creates a request to invoke AddVodDomain API
func CreateAddVodDomainRequest() (request *AddVodDomainRequest) {
	request = &AddVodDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "AddVodDomain", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddVodDomainResponse creates a response to parse from AddVodDomain response
func CreateAddVodDomainResponse() (response *AddVodDomainResponse) {
	response = &AddVodDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
