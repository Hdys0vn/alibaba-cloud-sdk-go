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

// DescribeFabricConsortiumOrderers invokes the baas.DescribeFabricConsortiumOrderers API synchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiumorderers.html
func (client *Client) DescribeFabricConsortiumOrderers(request *DescribeFabricConsortiumOrderersRequest) (response *DescribeFabricConsortiumOrderersResponse, err error) {
	response = CreateDescribeFabricConsortiumOrderersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFabricConsortiumOrderersWithChan invokes the baas.DescribeFabricConsortiumOrderers API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiumorderers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricConsortiumOrderersWithChan(request *DescribeFabricConsortiumOrderersRequest) (<-chan *DescribeFabricConsortiumOrderersResponse, <-chan error) {
	responseChan := make(chan *DescribeFabricConsortiumOrderersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFabricConsortiumOrderers(request)
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

// DescribeFabricConsortiumOrderersWithCallback invokes the baas.DescribeFabricConsortiumOrderers API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricconsortiumorderers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricConsortiumOrderersWithCallback(request *DescribeFabricConsortiumOrderersRequest, callback func(response *DescribeFabricConsortiumOrderersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFabricConsortiumOrderersResponse
		var err error
		defer close(result)
		response, err = client.DescribeFabricConsortiumOrderers(request)
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

// DescribeFabricConsortiumOrderersRequest is the request struct for api DescribeFabricConsortiumOrderers
type DescribeFabricConsortiumOrderersRequest struct {
	*requests.RpcRequest
	Location     string `position:"Body" name:"Location"`
	ConsortiumId string `position:"Body" name:"ConsortiumId"`
}

// DescribeFabricConsortiumOrderersResponse is the response struct for api DescribeFabricConsortiumOrderers
type DescribeFabricConsortiumOrderersResponse struct {
	*responses.BaseResponse
	RequestId string    `json:"RequestId" xml:"RequestId"`
	Success   bool      `json:"Success" xml:"Success"`
	ErrorCode int       `json:"ErrorCode" xml:"ErrorCode"`
	Result    []Orderer `json:"Result" xml:"Result"`
}

// CreateDescribeFabricConsortiumOrderersRequest creates a request to invoke DescribeFabricConsortiumOrderers API
func CreateDescribeFabricConsortiumOrderersRequest() (request *DescribeFabricConsortiumOrderersRequest) {
	request = &DescribeFabricConsortiumOrderersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeFabricConsortiumOrderers", "baas", "openAPI")
	return
}

// CreateDescribeFabricConsortiumOrderersResponse creates a response to parse from DescribeFabricConsortiumOrderers response
func CreateDescribeFabricConsortiumOrderersResponse() (response *DescribeFabricConsortiumOrderersResponse) {
	response = &DescribeFabricConsortiumOrderersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
