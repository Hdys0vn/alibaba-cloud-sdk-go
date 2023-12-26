package vpc

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

// DeleteIPv6TranslatorAclList invokes the vpc.DeleteIPv6TranslatorAclList API synchronously
func (client *Client) DeleteIPv6TranslatorAclList(request *DeleteIPv6TranslatorAclListRequest) (response *DeleteIPv6TranslatorAclListResponse, err error) {
	response = CreateDeleteIPv6TranslatorAclListResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteIPv6TranslatorAclListWithChan invokes the vpc.DeleteIPv6TranslatorAclList API asynchronously
func (client *Client) DeleteIPv6TranslatorAclListWithChan(request *DeleteIPv6TranslatorAclListRequest) (<-chan *DeleteIPv6TranslatorAclListResponse, <-chan error) {
	responseChan := make(chan *DeleteIPv6TranslatorAclListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteIPv6TranslatorAclList(request)
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

// DeleteIPv6TranslatorAclListWithCallback invokes the vpc.DeleteIPv6TranslatorAclList API asynchronously
func (client *Client) DeleteIPv6TranslatorAclListWithCallback(request *DeleteIPv6TranslatorAclListRequest, callback func(response *DeleteIPv6TranslatorAclListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteIPv6TranslatorAclListResponse
		var err error
		defer close(result)
		response, err = client.DeleteIPv6TranslatorAclList(request)
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

// DeleteIPv6TranslatorAclListRequest is the request struct for api DeleteIPv6TranslatorAclList
type DeleteIPv6TranslatorAclListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DeleteIPv6TranslatorAclListResponse is the response struct for api DeleteIPv6TranslatorAclList
type DeleteIPv6TranslatorAclListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteIPv6TranslatorAclListRequest creates a request to invoke DeleteIPv6TranslatorAclList API
func CreateDeleteIPv6TranslatorAclListRequest() (request *DeleteIPv6TranslatorAclListRequest) {
	request = &DeleteIPv6TranslatorAclListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DeleteIPv6TranslatorAclList", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteIPv6TranslatorAclListResponse creates a response to parse from DeleteIPv6TranslatorAclList response
func CreateDeleteIPv6TranslatorAclListResponse() (response *DeleteIPv6TranslatorAclListResponse) {
	response = &DeleteIPv6TranslatorAclListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}