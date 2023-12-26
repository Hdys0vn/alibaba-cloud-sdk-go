package linkwan

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

// CountGateways invokes the linkwan.CountGateways API synchronously
func (client *Client) CountGateways(request *CountGatewaysRequest) (response *CountGatewaysResponse, err error) {
	response = CreateCountGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// CountGatewaysWithChan invokes the linkwan.CountGateways API asynchronously
func (client *Client) CountGatewaysWithChan(request *CountGatewaysRequest) (<-chan *CountGatewaysResponse, <-chan error) {
	responseChan := make(chan *CountGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CountGateways(request)
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

// CountGatewaysWithCallback invokes the linkwan.CountGateways API asynchronously
func (client *Client) CountGatewaysWithCallback(request *CountGatewaysRequest, callback func(response *CountGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CountGatewaysResponse
		var err error
		defer close(result)
		response, err = client.CountGateways(request)
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

// CountGatewaysRequest is the request struct for api CountGateways
type CountGatewaysRequest struct {
	*requests.RpcRequest
	RealTenantId        string           `position:"Query" name:"RealTenantId"`
	RealTripartiteKey   string           `position:"Query" name:"RealTripartiteKey"`
	FuzzyGwEui          string           `position:"Query" name:"FuzzyGwEui"`
	IotInstanceId       string           `position:"Query" name:"IotInstanceId"`
	FuzzyCity           string           `position:"Query" name:"FuzzyCity"`
	OnlineState         string           `position:"Query" name:"OnlineState"`
	IsEnabled           requests.Boolean `position:"Query" name:"IsEnabled"`
	FuzzyName           string           `position:"Query" name:"FuzzyName"`
	FreqBandPlanGroupId requests.Integer `position:"Query" name:"FreqBandPlanGroupId"`
	ApiProduct          string           `position:"Body" name:"ApiProduct"`
	ApiRevision         string           `position:"Body" name:"ApiRevision"`
}

// CountGatewaysResponse is the response struct for api CountGateways
type CountGatewaysResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      int64  `json:"Data" xml:"Data"`
}

// CreateCountGatewaysRequest creates a request to invoke CountGateways API
func CreateCountGatewaysRequest() (request *CountGatewaysRequest) {
	request = &CountGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "CountGateways", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCountGatewaysResponse creates a response to parse from CountGateways response
func CreateCountGatewaysResponse() (response *CountGatewaysResponse) {
	response = &CountGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
