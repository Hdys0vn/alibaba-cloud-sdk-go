package cbn

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

// DeleteTransitRouteTableAggregation invokes the cbn.DeleteTransitRouteTableAggregation API synchronously
func (client *Client) DeleteTransitRouteTableAggregation(request *DeleteTransitRouteTableAggregationRequest) (response *DeleteTransitRouteTableAggregationResponse, err error) {
	response = CreateDeleteTransitRouteTableAggregationResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteTransitRouteTableAggregationWithChan invokes the cbn.DeleteTransitRouteTableAggregation API asynchronously
func (client *Client) DeleteTransitRouteTableAggregationWithChan(request *DeleteTransitRouteTableAggregationRequest) (<-chan *DeleteTransitRouteTableAggregationResponse, <-chan error) {
	responseChan := make(chan *DeleteTransitRouteTableAggregationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteTransitRouteTableAggregation(request)
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

// DeleteTransitRouteTableAggregationWithCallback invokes the cbn.DeleteTransitRouteTableAggregation API asynchronously
func (client *Client) DeleteTransitRouteTableAggregationWithCallback(request *DeleteTransitRouteTableAggregationRequest, callback func(response *DeleteTransitRouteTableAggregationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteTransitRouteTableAggregationResponse
		var err error
		defer close(result)
		response, err = client.DeleteTransitRouteTableAggregation(request)
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

// DeleteTransitRouteTableAggregationRequest is the request struct for api DeleteTransitRouteTableAggregation
type DeleteTransitRouteTableAggregationRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                  requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                      string           `position:"Query" name:"ClientToken"`
	DryRun                           requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount             string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                     string           `position:"Query" name:"OwnerAccount"`
	OwnerId                          requests.Integer `position:"Query" name:"OwnerId"`
	Version                          string           `position:"Query" name:"Version"`
	TransitRouteTableId              string           `position:"Query" name:"TransitRouteTableId"`
	TransitRouteTableAggregationCidr string           `position:"Query" name:"TransitRouteTableAggregationCidr"`
}

// DeleteTransitRouteTableAggregationResponse is the response struct for api DeleteTransitRouteTableAggregation
type DeleteTransitRouteTableAggregationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteTransitRouteTableAggregationRequest creates a request to invoke DeleteTransitRouteTableAggregation API
func CreateDeleteTransitRouteTableAggregationRequest() (request *DeleteTransitRouteTableAggregationRequest) {
	request = &DeleteTransitRouteTableAggregationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DeleteTransitRouteTableAggregation", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteTransitRouteTableAggregationResponse creates a response to parse from DeleteTransitRouteTableAggregation response
func CreateDeleteTransitRouteTableAggregationResponse() (response *DeleteTransitRouteTableAggregationResponse) {
	response = &DeleteTransitRouteTableAggregationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}