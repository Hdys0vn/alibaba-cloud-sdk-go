package aegis

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

// DeleteVulWhitelist invokes the aegis.DeleteVulWhitelist API synchronously
// api document: https://help.aliyun.com/api/aegis/deletevulwhitelist.html
func (client *Client) DeleteVulWhitelist(request *DeleteVulWhitelistRequest) (response *DeleteVulWhitelistResponse, err error) {
	response = CreateDeleteVulWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVulWhitelistWithChan invokes the aegis.DeleteVulWhitelist API asynchronously
// api document: https://help.aliyun.com/api/aegis/deletevulwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVulWhitelistWithChan(request *DeleteVulWhitelistRequest) (<-chan *DeleteVulWhitelistResponse, <-chan error) {
	responseChan := make(chan *DeleteVulWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVulWhitelist(request)
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

// DeleteVulWhitelistWithCallback invokes the aegis.DeleteVulWhitelist API asynchronously
// api document: https://help.aliyun.com/api/aegis/deletevulwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVulWhitelistWithCallback(request *DeleteVulWhitelistRequest, callback func(response *DeleteVulWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVulWhitelistResponse
		var err error
		defer close(result)
		response, err = client.DeleteVulWhitelist(request)
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

// DeleteVulWhitelistRequest is the request struct for api DeleteVulWhitelist
type DeleteVulWhitelistRequest struct {
	*requests.RpcRequest
	SourceIp  string `position:"Query" name:"SourceIp"`
	Whitelist string `position:"Query" name:"Whitelist"`
}

// DeleteVulWhitelistResponse is the response struct for api DeleteVulWhitelist
type DeleteVulWhitelistResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVulWhitelistRequest creates a request to invoke DeleteVulWhitelist API
func CreateDeleteVulWhitelistRequest() (request *DeleteVulWhitelistRequest) {
	request = &DeleteVulWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DeleteVulWhitelist", "vipaegis", "openAPI")
	return
}

// CreateDeleteVulWhitelistResponse creates a response to parse from DeleteVulWhitelist response
func CreateDeleteVulWhitelistResponse() (response *DeleteVulWhitelistResponse) {
	response = &DeleteVulWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
