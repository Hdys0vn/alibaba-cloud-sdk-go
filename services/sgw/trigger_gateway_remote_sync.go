package sgw

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

// TriggerGatewayRemoteSync invokes the sgw.TriggerGatewayRemoteSync API synchronously
func (client *Client) TriggerGatewayRemoteSync(request *TriggerGatewayRemoteSyncRequest) (response *TriggerGatewayRemoteSyncResponse, err error) {
	response = CreateTriggerGatewayRemoteSyncResponse()
	err = client.DoAction(request, response)
	return
}

// TriggerGatewayRemoteSyncWithChan invokes the sgw.TriggerGatewayRemoteSync API asynchronously
func (client *Client) TriggerGatewayRemoteSyncWithChan(request *TriggerGatewayRemoteSyncRequest) (<-chan *TriggerGatewayRemoteSyncResponse, <-chan error) {
	responseChan := make(chan *TriggerGatewayRemoteSyncResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TriggerGatewayRemoteSync(request)
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

// TriggerGatewayRemoteSyncWithCallback invokes the sgw.TriggerGatewayRemoteSync API asynchronously
func (client *Client) TriggerGatewayRemoteSyncWithCallback(request *TriggerGatewayRemoteSyncRequest, callback func(response *TriggerGatewayRemoteSyncResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TriggerGatewayRemoteSyncResponse
		var err error
		defer close(result)
		response, err = client.TriggerGatewayRemoteSync(request)
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

// TriggerGatewayRemoteSyncRequest is the request struct for api TriggerGatewayRemoteSync
type TriggerGatewayRemoteSyncRequest struct {
	*requests.RpcRequest
	Path          string `position:"Query" name:"Path"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	IndexId       string `position:"Query" name:"IndexId"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// TriggerGatewayRemoteSyncResponse is the response struct for api TriggerGatewayRemoteSync
type TriggerGatewayRemoteSyncResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateTriggerGatewayRemoteSyncRequest creates a request to invoke TriggerGatewayRemoteSync API
func CreateTriggerGatewayRemoteSyncRequest() (request *TriggerGatewayRemoteSyncRequest) {
	request = &TriggerGatewayRemoteSyncRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "TriggerGatewayRemoteSync", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTriggerGatewayRemoteSyncResponse creates a response to parse from TriggerGatewayRemoteSync response
func CreateTriggerGatewayRemoteSyncResponse() (response *TriggerGatewayRemoteSyncResponse) {
	response = &TriggerGatewayRemoteSyncResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}