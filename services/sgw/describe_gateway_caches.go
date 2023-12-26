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

// DescribeGatewayCaches invokes the sgw.DescribeGatewayCaches API synchronously
func (client *Client) DescribeGatewayCaches(request *DescribeGatewayCachesRequest) (response *DescribeGatewayCachesResponse, err error) {
	response = CreateDescribeGatewayCachesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGatewayCachesWithChan invokes the sgw.DescribeGatewayCaches API asynchronously
func (client *Client) DescribeGatewayCachesWithChan(request *DescribeGatewayCachesRequest) (<-chan *DescribeGatewayCachesResponse, <-chan error) {
	responseChan := make(chan *DescribeGatewayCachesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGatewayCaches(request)
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

// DescribeGatewayCachesWithCallback invokes the sgw.DescribeGatewayCaches API asynchronously
func (client *Client) DescribeGatewayCachesWithCallback(request *DescribeGatewayCachesRequest, callback func(response *DescribeGatewayCachesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGatewayCachesResponse
		var err error
		defer close(result)
		response, err = client.DescribeGatewayCaches(request)
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

// DescribeGatewayCachesRequest is the request struct for api DescribeGatewayCaches
type DescribeGatewayCachesRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	GatewayId     string `position:"Query" name:"GatewayId"`
}

// DescribeGatewayCachesResponse is the response struct for api DescribeGatewayCaches
type DescribeGatewayCachesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Caches    Caches `json:"Caches" xml:"Caches"`
}

// CreateDescribeGatewayCachesRequest creates a request to invoke DescribeGatewayCaches API
func CreateDescribeGatewayCachesRequest() (request *DescribeGatewayCachesRequest) {
	request = &DescribeGatewayCachesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("sgw", "2018-05-11", "DescribeGatewayCaches", "hcs_sgw", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeGatewayCachesResponse creates a response to parse from DescribeGatewayCaches response
func CreateDescribeGatewayCachesResponse() (response *DescribeGatewayCachesResponse) {
	response = &DescribeGatewayCachesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
