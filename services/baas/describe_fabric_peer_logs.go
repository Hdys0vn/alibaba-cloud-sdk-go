package baas

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

// DescribeFabricPeerLogs invokes the baas.DescribeFabricPeerLogs API synchronously
// api document: https://help.aliyun.com/api/baas/describefabricpeerlogs.html
func (client *Client) DescribeFabricPeerLogs(request *DescribeFabricPeerLogsRequest) (response *DescribeFabricPeerLogsResponse, err error) {
	response = CreateDescribeFabricPeerLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFabricPeerLogsWithChan invokes the baas.DescribeFabricPeerLogs API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricpeerlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricPeerLogsWithChan(request *DescribeFabricPeerLogsRequest) (<-chan *DescribeFabricPeerLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeFabricPeerLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFabricPeerLogs(request)
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

// DescribeFabricPeerLogsWithCallback invokes the baas.DescribeFabricPeerLogs API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricpeerlogs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricPeerLogsWithCallback(request *DescribeFabricPeerLogsRequest, callback func(response *DescribeFabricPeerLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFabricPeerLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFabricPeerLogs(request)
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

// DescribeFabricPeerLogsRequest is the request struct for api DescribeFabricPeerLogs
type DescribeFabricPeerLogsRequest struct {
	*requests.RpcRequest
	PeerName       string `position:"Query" name:"PeerName"`
	Lines          string `position:"Query" name:"Lines"`
	OrganizationId string `position:"Query" name:"OrganizationId"`
}

// DescribeFabricPeerLogsResponse is the response struct for api DescribeFabricPeerLogs
type DescribeFabricPeerLogsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateDescribeFabricPeerLogsRequest creates a request to invoke DescribeFabricPeerLogs API
func CreateDescribeFabricPeerLogsRequest() (request *DescribeFabricPeerLogsRequest) {
	request = &DescribeFabricPeerLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeFabricPeerLogs", "baas", "openAPI")
	return
}

// CreateDescribeFabricPeerLogsResponse creates a response to parse from DescribeFabricPeerLogs response
func CreateDescribeFabricPeerLogsResponse() (response *DescribeFabricPeerLogsResponse) {
	response = &DescribeFabricPeerLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
