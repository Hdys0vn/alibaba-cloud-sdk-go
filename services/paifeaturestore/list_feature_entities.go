package paifeaturestore

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

// ListFeatureEntities invokes the paifeaturestore.ListFeatureEntities API synchronously
func (client *Client) ListFeatureEntities(request *ListFeatureEntitiesRequest) (response *ListFeatureEntitiesResponse, err error) {
	response = CreateListFeatureEntitiesResponse()
	err = client.DoAction(request, response)
	return
}

// ListFeatureEntitiesWithChan invokes the paifeaturestore.ListFeatureEntities API asynchronously
func (client *Client) ListFeatureEntitiesWithChan(request *ListFeatureEntitiesRequest) (<-chan *ListFeatureEntitiesResponse, <-chan error) {
	responseChan := make(chan *ListFeatureEntitiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFeatureEntities(request)
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

// ListFeatureEntitiesWithCallback invokes the paifeaturestore.ListFeatureEntities API asynchronously
func (client *Client) ListFeatureEntitiesWithCallback(request *ListFeatureEntitiesRequest, callback func(response *ListFeatureEntitiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFeatureEntitiesResponse
		var err error
		defer close(result)
		response, err = client.ListFeatureEntities(request)
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

// ListFeatureEntitiesRequest is the request struct for api ListFeatureEntities
type ListFeatureEntitiesRequest struct {
	*requests.RoaRequest
	Owner      string           `position:"Query" name:"Owner"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	InstanceId string           `position:"Path" name:"InstanceId"`
	Name       string           `position:"Query" name:"Name"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	SortBy     string           `position:"Query" name:"SortBy"`
	ProjectId  string           `position:"Query" name:"ProjectId"`
	Order      string           `position:"Query" name:"Order"`
}

// ListFeatureEntitiesResponse is the response struct for api ListFeatureEntities
type ListFeatureEntitiesResponse struct {
	*responses.BaseResponse
	RequestId       string                `json:"RequestId" xml:"RequestId"`
	TotalCount      int                   `json:"TotalCount" xml:"TotalCount"`
	FeatureEntities []FeatureEntitiesItem `json:"FeatureEntities" xml:"FeatureEntities"`
}

// CreateListFeatureEntitiesRequest creates a request to invoke ListFeatureEntities API
func CreateListFeatureEntitiesRequest() (request *ListFeatureEntitiesRequest) {
	request = &ListFeatureEntitiesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiFeatureStore", "2023-06-21", "ListFeatureEntities", "/api/v1/instances/[InstanceId]/featureentities", "", "")
	request.Method = requests.GET
	return
}

// CreateListFeatureEntitiesResponse creates a response to parse from ListFeatureEntities response
func CreateListFeatureEntitiesResponse() (response *ListFeatureEntitiesResponse) {
	response = &ListFeatureEntitiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
