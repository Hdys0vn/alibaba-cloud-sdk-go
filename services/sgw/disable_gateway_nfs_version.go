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

// DisableGatewayNFSVersion invokes the sgw.DisableGatewayNFSVersion API synchronously
func (client *Client) DisableGatewayNFSVersion(request *DisableGatewayNFSVersionRequest) (response *DisableGatewayNFSVersionResponse, err error) {
	response = CreateDisableGatewayNFSVersionResponse()
	err = client.DoAction(request, response)
	return
}

// DisableGatewayNFSVersionWithChan invokes the sgw.DisableGatewayNFSVersion API asynchronously
func (client *Client) DisableGatewayNFSVersionWithChan(request *DisableGatewayNFSVersionRequest) (<-chan *DisableGatewayNFSVersionResponse, <-chan error) {
	responseChan := make(chan *DisableGatewayNFSVersionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableGatewayNFSVersion(request)
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

// DisableGatewayNFSVersionWithCallback invokes the sgw.DisableGatewayNFSVersion API asynchronously
func (client *Client) DisableGatewayNFSVersionWithCallback(request *DisableGatewayNFSVersionRequest, callback func(response *DisableGatewayNFSVersionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableGatewayNFSVersionResponse
		var err error
		defer close(result)
		response, err = client.DisableGatewayNFSVersion(request)
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

// DisableGatewayNFSVersionRequest is the request struct for api DisableGatewayNFSVersion
type DisableGatewayNFSVersionRequest struct {
	*requests.RpcRequest
	NFSVersion    string `position:"Query" name:"NFSVersion"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// DisableGatewayNFSVersionResponse is the response struct for api DisableGatewayNFSVersion
type DisableGatewayNFSVersionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	TaskId    string `json:"TaskId" xml:"TaskId"`
}

// CreateDisableGatewayNFSVersionRequest creates a request to invoke DisableGatewayNFSVersion API
func CreateDisableGatewayNFSVersionRequest() (request *DisableGatewayNFSVersionRequest) {
	request = &DisableGatewayNFSVersionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DisableGatewayNFSVersion", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDisableGatewayNFSVersionResponse creates a response to parse from DisableGatewayNFSVersion response
func CreateDisableGatewayNFSVersionResponse() (response *DisableGatewayNFSVersionResponse) {
	response = &DisableGatewayNFSVersionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
