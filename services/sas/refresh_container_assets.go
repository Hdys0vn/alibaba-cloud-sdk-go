package sas

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

// RefreshContainerAssets invokes the sas.RefreshContainerAssets API synchronously
func (client *Client) RefreshContainerAssets(request *RefreshContainerAssetsRequest) (response *RefreshContainerAssetsResponse, err error) {
	response = CreateRefreshContainerAssetsResponse()
	err = client.DoAction(request, response)
	return
}

// RefreshContainerAssetsWithChan invokes the sas.RefreshContainerAssets API asynchronously
func (client *Client) RefreshContainerAssetsWithChan(request *RefreshContainerAssetsRequest) (<-chan *RefreshContainerAssetsResponse, <-chan error) {
	responseChan := make(chan *RefreshContainerAssetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RefreshContainerAssets(request)
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

// RefreshContainerAssetsWithCallback invokes the sas.RefreshContainerAssets API asynchronously
func (client *Client) RefreshContainerAssetsWithCallback(request *RefreshContainerAssetsRequest, callback func(response *RefreshContainerAssetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RefreshContainerAssetsResponse
		var err error
		defer close(result)
		response, err = client.RefreshContainerAssets(request)
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

// RefreshContainerAssetsRequest is the request struct for api RefreshContainerAssets
type RefreshContainerAssetsRequest struct {
	*requests.RpcRequest
	SourceIp  string `position:"Query" name:"SourceIp"`
	AssetType string `position:"Query" name:"AssetType"`
}

// RefreshContainerAssetsResponse is the response struct for api RefreshContainerAssets
type RefreshContainerAssetsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRefreshContainerAssetsRequest creates a request to invoke RefreshContainerAssets API
func CreateRefreshContainerAssetsRequest() (request *RefreshContainerAssetsRequest) {
	request = &RefreshContainerAssetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "RefreshContainerAssets", "sas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRefreshContainerAssetsResponse creates a response to parse from RefreshContainerAssets response
func CreateRefreshContainerAssetsResponse() (response *RefreshContainerAssetsResponse) {
	response = &RefreshContainerAssetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
