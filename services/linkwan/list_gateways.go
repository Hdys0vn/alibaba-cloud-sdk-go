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

// ListGateways invokes the linkwan.ListGateways API synchronously
func (client *Client) ListGateways(request *ListGatewaysRequest) (response *ListGatewaysResponse, err error) {
	response = CreateListGatewaysResponse()
	err = client.DoAction(request, response)
	return
}

// ListGatewaysWithChan invokes the linkwan.ListGateways API asynchronously
func (client *Client) ListGatewaysWithChan(request *ListGatewaysRequest) (<-chan *ListGatewaysResponse, <-chan error) {
	responseChan := make(chan *ListGatewaysResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListGateways(request)
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

// ListGatewaysWithCallback invokes the linkwan.ListGateways API asynchronously
func (client *Client) ListGatewaysWithCallback(request *ListGatewaysRequest, callback func(response *ListGatewaysResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListGatewaysResponse
		var err error
		defer close(result)
		response, err = client.ListGateways(request)
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

// ListGatewaysRequest is the request struct for api ListGateways
type ListGatewaysRequest struct {
	*requests.RpcRequest
	FuzzyGwEui          string           `position:"Query" name:"FuzzyGwEui"`
	IotInstanceId       string           `position:"Query" name:"IotInstanceId"`
	Limit               requests.Integer `position:"Query" name:"Limit"`
	FuzzyCity           string           `position:"Query" name:"FuzzyCity"`
	OnlineState         string           `position:"Query" name:"OnlineState"`
	IsEnabled           requests.Boolean `position:"Query" name:"IsEnabled"`
	FuzzyName           string           `position:"Query" name:"FuzzyName"`
	Offset              requests.Integer `position:"Query" name:"Offset"`
	Ascending           requests.Boolean `position:"Query" name:"Ascending"`
	FreqBandPlanGroupId requests.Integer `position:"Query" name:"FreqBandPlanGroupId"`
	ApiProduct          string           `position:"Body" name:"ApiProduct"`
	ApiRevision         string           `position:"Body" name:"ApiRevision"`
	SortingField        string           `position:"Query" name:"SortingField"`
}

// ListGatewaysResponse is the response struct for api ListGateways
type ListGatewaysResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Success   bool               `json:"Success" xml:"Success"`
	Data      DataInListGateways `json:"Data" xml:"Data"`
}

// CreateListGatewaysRequest creates a request to invoke ListGateways API
func CreateListGatewaysRequest() (request *ListGatewaysRequest) {
	request = &ListGatewaysRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "ListGateways", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListGatewaysResponse creates a response to parse from ListGateways response
func CreateListGatewaysResponse() (response *ListGatewaysResponse) {
	response = &ListGatewaysResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
