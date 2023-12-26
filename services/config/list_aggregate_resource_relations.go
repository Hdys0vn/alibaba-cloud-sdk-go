package config

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

// ListAggregateResourceRelations invokes the config.ListAggregateResourceRelations API synchronously
func (client *Client) ListAggregateResourceRelations(request *ListAggregateResourceRelationsRequest) (response *ListAggregateResourceRelationsResponse, err error) {
	response = CreateListAggregateResourceRelationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListAggregateResourceRelationsWithChan invokes the config.ListAggregateResourceRelations API asynchronously
func (client *Client) ListAggregateResourceRelationsWithChan(request *ListAggregateResourceRelationsRequest) (<-chan *ListAggregateResourceRelationsResponse, <-chan error) {
	responseChan := make(chan *ListAggregateResourceRelationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAggregateResourceRelations(request)
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

// ListAggregateResourceRelationsWithCallback invokes the config.ListAggregateResourceRelations API asynchronously
func (client *Client) ListAggregateResourceRelationsWithCallback(request *ListAggregateResourceRelationsRequest, callback func(response *ListAggregateResourceRelationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAggregateResourceRelationsResponse
		var err error
		defer close(result)
		response, err = client.ListAggregateResourceRelations(request)
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

// ListAggregateResourceRelationsRequest is the request struct for api ListAggregateResourceRelations
type ListAggregateResourceRelationsRequest struct {
	*requests.RpcRequest
	TargetResourceId   string           `position:"Query" name:"TargetResourceId"`
	AggregatorId       string           `position:"Query" name:"AggregatorId"`
	RelationType       string           `position:"Query" name:"RelationType"`
	ResourceAccountId  requests.Integer `position:"Query" name:"ResourceAccountId"`
	NextToken          string           `position:"Query" name:"NextToken"`
	TargetResourceType string           `position:"Query" name:"TargetResourceType"`
	ResourceId         string           `position:"Query" name:"ResourceId"`
	ResourceType       string           `position:"Query" name:"ResourceType"`
	MaxResults         requests.Integer `position:"Query" name:"MaxResults"`
	Region             string           `position:"Query" name:"Region"`
}

// ListAggregateResourceRelationsResponse is the response struct for api ListAggregateResourceRelations
type ListAggregateResourceRelationsResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ResourceRelations ResourceRelations `json:"ResourceRelations" xml:"ResourceRelations"`
}

// CreateListAggregateResourceRelationsRequest creates a request to invoke ListAggregateResourceRelations API
func CreateListAggregateResourceRelationsRequest() (request *ListAggregateResourceRelationsRequest) {
	request = &ListAggregateResourceRelationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ListAggregateResourceRelations", "", "")
	request.Method = requests.POST
	return
}

// CreateListAggregateResourceRelationsResponse creates a response to parse from ListAggregateResourceRelations response
func CreateListAggregateResourceRelationsResponse() (response *ListAggregateResourceRelationsResponse) {
	response = &ListAggregateResourceRelationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
