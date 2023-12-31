package cdrs

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

// SearchAggregateObject invokes the cdrs.SearchAggregateObject API synchronously
func (client *Client) SearchAggregateObject(request *SearchAggregateObjectRequest) (response *SearchAggregateObjectResponse, err error) {
	response = CreateSearchAggregateObjectResponse()
	err = client.DoAction(request, response)
	return
}

// SearchAggregateObjectWithChan invokes the cdrs.SearchAggregateObject API asynchronously
func (client *Client) SearchAggregateObjectWithChan(request *SearchAggregateObjectRequest) (<-chan *SearchAggregateObjectResponse, <-chan error) {
	responseChan := make(chan *SearchAggregateObjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SearchAggregateObject(request)
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

// SearchAggregateObjectWithCallback invokes the cdrs.SearchAggregateObject API asynchronously
func (client *Client) SearchAggregateObjectWithCallback(request *SearchAggregateObjectRequest, callback func(response *SearchAggregateObjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SearchAggregateObjectResponse
		var err error
		defer close(result)
		response, err = client.SearchAggregateObject(request)
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

// SearchAggregateObjectRequest is the request struct for api SearchAggregateObject
type SearchAggregateObjectRequest struct {
	*requests.RpcRequest
	ShotTimeEnd       string           `position:"Body" name:"ShotTimeEnd"`
	CorpId            string           `position:"Body" name:"CorpId"`
	PageNumber        requests.Integer `position:"Body" name:"PageNumber"`
	Feature           string           `position:"Body" name:"Feature"`
	Vendor            string           `position:"Body" name:"Vendor"`
	RequireTotalCount requests.Boolean `position:"Body" name:"RequireTotalCount"`
	PageSize          requests.Integer `position:"Body" name:"PageSize"`
	ImageContent      string           `position:"Body" name:"ImageContent"`
	ObjectType        string           `position:"Body" name:"ObjectType"`
	DeviceList        string           `position:"Body" name:"DeviceList"`
	ImageUrl          string           `position:"Body" name:"ImageUrl"`
	Attributes        string           `position:"Body" name:"Attributes"`
	ShotTimeStart     string           `position:"Body" name:"ShotTimeStart"`
}

// SearchAggregateObjectResponse is the response struct for api SearchAggregateObject
type SearchAggregateObjectResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Total     int    `json:"Total" xml:"Total"`
	PageSize  int64  `json:"PageSize" xml:"PageSize"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSearchAggregateObjectRequest creates a request to invoke SearchAggregateObject API
func CreateSearchAggregateObjectRequest() (request *SearchAggregateObjectRequest) {
	request = &SearchAggregateObjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CDRS", "2020-11-01", "SearchAggregateObject", "", "")
	request.Method = requests.POST
	return
}

// CreateSearchAggregateObjectResponse creates a response to parse from SearchAggregateObject response
func CreateSearchAggregateObjectResponse() (response *SearchAggregateObjectResponse) {
	response = &SearchAggregateObjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
