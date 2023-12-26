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

// DescribeFabricConsortiumChaincodes invokes the baas.DescribeFabricConsortiumChaincodes API synchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiumchaincodes.html
func (client *Client) DescribeFabricConsortiumChaincodes(request *DescribeFabricConsortiumChaincodesRequest) (response *DescribeFabricConsortiumChaincodesResponse, err error) {
	response = CreateDescribeFabricConsortiumChaincodesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFabricConsortiumChaincodesWithChan invokes the baas.DescribeFabricConsortiumChaincodes API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiumchaincodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricConsortiumChaincodesWithChan(request *DescribeFabricConsortiumChaincodesRequest) (<-chan *DescribeFabricConsortiumChaincodesResponse, <-chan error) {
	responseChan := make(chan *DescribeFabricConsortiumChaincodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFabricConsortiumChaincodes(request)
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

// DescribeFabricConsortiumChaincodesWithCallback invokes the baas.DescribeFabricConsortiumChaincodes API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiumchaincodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricConsortiumChaincodesWithCallback(request *DescribeFabricConsortiumChaincodesRequest, callback func(response *DescribeFabricConsortiumChaincodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFabricConsortiumChaincodesResponse
		var err error
		defer close(result)
		response, err = client.DescribeFabricConsortiumChaincodes(request)
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

// DescribeFabricConsortiumChaincodesRequest is the request struct for api DescribeFabricConsortiumChaincodes
type DescribeFabricConsortiumChaincodesRequest struct {
	*requests.RpcRequest
	Location     string `position:"Body" name:"Location"`
	ConsortiumId string `position:"Body" name:"ConsortiumId"`
}

// DescribeFabricConsortiumChaincodesResponse is the response struct for api DescribeFabricConsortiumChaincodes
type DescribeFabricConsortiumChaincodesResponse struct {
	*responses.BaseResponse
	RequestId string        `json:"RequestId" xml:"RequestId"`
	Success   bool          `json:"Success" xml:"Success"`
	ErrorCode int           `json:"ErrorCode" xml:"ErrorCode"`
	Result    []ChaincodeVO `json:"Result" xml:"Result"`
}

// CreateDescribeFabricConsortiumChaincodesRequest creates a request to invoke DescribeFabricConsortiumChaincodes API
func CreateDescribeFabricConsortiumChaincodesRequest() (request *DescribeFabricConsortiumChaincodesRequest) {
	request = &DescribeFabricConsortiumChaincodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeFabricConsortiumChaincodes", "baas", "openAPI")
	return
}

// CreateDescribeFabricConsortiumChaincodesResponse creates a response to parse from DescribeFabricConsortiumChaincodes response
func CreateDescribeFabricConsortiumChaincodesResponse() (response *DescribeFabricConsortiumChaincodesResponse) {
	response = &DescribeFabricConsortiumChaincodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}