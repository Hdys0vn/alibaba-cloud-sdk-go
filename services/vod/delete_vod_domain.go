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

// DeleteVodDomain invokes the vod.DeleteVodDomain API synchronously
func (client *Client) DeleteVodDomain(request *DeleteVodDomainRequest) (response *DeleteVodDomainResponse, err error) {
	response = CreateDeleteVodDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVodDomainWithChan invokes the vod.DeleteVodDomain API asynchronously
func (client *Client) DeleteVodDomainWithChan(request *DeleteVodDomainRequest) (<-chan *DeleteVodDomainResponse, <-chan error) {
	responseChan := make(chan *DeleteVodDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVodDomain(request)
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

// DeleteVodDomainWithCallback invokes the vod.DeleteVodDomain API asynchronously
func (client *Client) DeleteVodDomainWithCallback(request *DeleteVodDomainRequest, callback func(response *DeleteVodDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVodDomainResponse
		var err error
		defer close(result)
		response, err = client.DeleteVodDomain(request)
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

// DeleteVodDomainRequest is the request struct for api DeleteVodDomain
type DeleteVodDomainRequest struct {
	*requests.RpcRequest
	OwnerAccount  string           `position:"Query" name:"OwnerAccount"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// DeleteVodDomainResponse is the response struct for api DeleteVodDomain
type DeleteVodDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVodDomainRequest creates a request to invoke DeleteVodDomain API
func CreateDeleteVodDomainRequest() (request *DeleteVodDomainRequest) {
	request = &DeleteVodDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteVodDomain", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVodDomainResponse creates a response to parse from DeleteVodDomain response
func CreateDeleteVodDomainResponse() (response *DeleteVodDomainResponse) {
	response = &DeleteVodDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
