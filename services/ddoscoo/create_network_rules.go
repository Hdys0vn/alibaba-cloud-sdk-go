package ddoscoo

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

// CreateNetworkRules invokes the ddoscoo.CreateNetworkRules API synchronously
func (client *Client) CreateNetworkRules(request *CreateNetworkRulesRequest) (response *CreateNetworkRulesResponse, err error) {
	response = CreateCreateNetworkRulesResponse()
	err = client.DoAction(request, response)
	return
}

// CreateNetworkRulesWithChan invokes the ddoscoo.CreateNetworkRules API asynchronously
func (client *Client) CreateNetworkRulesWithChan(request *CreateNetworkRulesRequest) (<-chan *CreateNetworkRulesResponse, <-chan error) {
	responseChan := make(chan *CreateNetworkRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateNetworkRules(request)
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

// CreateNetworkRulesWithCallback invokes the ddoscoo.CreateNetworkRules API asynchronously
func (client *Client) CreateNetworkRulesWithCallback(request *CreateNetworkRulesRequest, callback func(response *CreateNetworkRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateNetworkRulesResponse
		var err error
		defer close(result)
		response, err = client.CreateNetworkRules(request)
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

// CreateNetworkRulesRequest is the request struct for api CreateNetworkRules
type CreateNetworkRulesRequest struct {
	*requests.RpcRequest
	NetworkRules string `position:"Query" name:"NetworkRules"`
	SourceIp     string `position:"Query" name:"SourceIp"`
}

// CreateNetworkRulesResponse is the response struct for api CreateNetworkRules
type CreateNetworkRulesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateNetworkRulesRequest creates a request to invoke CreateNetworkRules API
func CreateCreateNetworkRulesRequest() (request *CreateNetworkRulesRequest) {
	request = &CreateNetworkRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "CreateNetworkRules", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateNetworkRulesResponse creates a response to parse from CreateNetworkRules response
func CreateCreateNetworkRulesResponse() (response *CreateNetworkRulesResponse) {
	response = &CreateNetworkRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
