package nas

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

// ModifyDataFlowAutoRefresh invokes the nas.ModifyDataFlowAutoRefresh API synchronously
func (client *Client) ModifyDataFlowAutoRefresh(request *ModifyDataFlowAutoRefreshRequest) (response *ModifyDataFlowAutoRefreshResponse, err error) {
	response = CreateModifyDataFlowAutoRefreshResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDataFlowAutoRefreshWithChan invokes the nas.ModifyDataFlowAutoRefresh API asynchronously
func (client *Client) ModifyDataFlowAutoRefreshWithChan(request *ModifyDataFlowAutoRefreshRequest) (<-chan *ModifyDataFlowAutoRefreshResponse, <-chan error) {
	responseChan := make(chan *ModifyDataFlowAutoRefreshResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDataFlowAutoRefresh(request)
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

// ModifyDataFlowAutoRefreshWithCallback invokes the nas.ModifyDataFlowAutoRefresh API asynchronously
func (client *Client) ModifyDataFlowAutoRefreshWithCallback(request *ModifyDataFlowAutoRefreshRequest, callback func(response *ModifyDataFlowAutoRefreshResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDataFlowAutoRefreshResponse
		var err error
		defer close(result)
		response, err = client.ModifyDataFlowAutoRefresh(request)
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

// ModifyDataFlowAutoRefreshRequest is the request struct for api ModifyDataFlowAutoRefresh
type ModifyDataFlowAutoRefreshRequest struct {
	*requests.RpcRequest
	AutoRefreshPolicy   string           `position:"Query" name:"AutoRefreshPolicy"`
	ClientToken         string           `position:"Query" name:"ClientToken"`
	FileSystemId        string           `position:"Query" name:"FileSystemId"`
	DryRun              requests.Boolean `position:"Query" name:"DryRun"`
	DataFlowId          string           `position:"Query" name:"DataFlowId"`
	AutoRefreshInterval requests.Integer `position:"Query" name:"AutoRefreshInterval"`
}

// ModifyDataFlowAutoRefreshResponse is the response struct for api ModifyDataFlowAutoRefresh
type ModifyDataFlowAutoRefreshResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDataFlowAutoRefreshRequest creates a request to invoke ModifyDataFlowAutoRefresh API
func CreateModifyDataFlowAutoRefreshRequest() (request *ModifyDataFlowAutoRefreshRequest) {
	request = &ModifyDataFlowAutoRefreshRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "ModifyDataFlowAutoRefresh", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDataFlowAutoRefreshResponse creates a response to parse from ModifyDataFlowAutoRefresh response
func CreateModifyDataFlowAutoRefreshResponse() (response *ModifyDataFlowAutoRefreshResponse) {
	response = &ModifyDataFlowAutoRefreshResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
