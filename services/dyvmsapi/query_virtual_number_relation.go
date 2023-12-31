package dyvmsapi

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

// QueryVirtualNumberRelation invokes the dyvmsapi.QueryVirtualNumberRelation API synchronously
func (client *Client) QueryVirtualNumberRelation(request *QueryVirtualNumberRelationRequest) (response *QueryVirtualNumberRelationResponse, err error) {
	response = CreateQueryVirtualNumberRelationResponse()
	err = client.DoAction(request, response)
	return
}

// QueryVirtualNumberRelationWithChan invokes the dyvmsapi.QueryVirtualNumberRelation API asynchronously
func (client *Client) QueryVirtualNumberRelationWithChan(request *QueryVirtualNumberRelationRequest) (<-chan *QueryVirtualNumberRelationResponse, <-chan error) {
	responseChan := make(chan *QueryVirtualNumberRelationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryVirtualNumberRelation(request)
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

// QueryVirtualNumberRelationWithCallback invokes the dyvmsapi.QueryVirtualNumberRelation API asynchronously
func (client *Client) QueryVirtualNumberRelationWithCallback(request *QueryVirtualNumberRelationRequest, callback func(response *QueryVirtualNumberRelationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryVirtualNumberRelationResponse
		var err error
		defer close(result)
		response, err = client.QueryVirtualNumberRelation(request)
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

// QueryVirtualNumberRelationRequest is the request struct for api QueryVirtualNumberRelation
type QueryVirtualNumberRelationRequest struct {
	*requests.RpcRequest
	SpecId               requests.Integer `position:"Query" name:"SpecId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	RouteType            requests.Integer `position:"Query" name:"RouteType"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	RelatedNum           string           `position:"Query" name:"RelatedNum"`
	RegionNameCity       string           `position:"Query" name:"RegionNameCity"`
	QualificationId      requests.Integer `position:"Query" name:"QualificationId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ProdCode             string           `position:"Query" name:"ProdCode"`
	PhoneNum             string           `position:"Query" name:"PhoneNum"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNo               requests.Integer `position:"Query" name:"PageNo"`
}

// QueryVirtualNumberRelationResponse is the response struct for api QueryVirtualNumberRelation
type QueryVirtualNumberRelationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      string `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateQueryVirtualNumberRelationRequest creates a request to invoke QueryVirtualNumberRelation API
func CreateQueryVirtualNumberRelationRequest() (request *QueryVirtualNumberRelationRequest) {
	request = &QueryVirtualNumberRelationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyvmsapi", "2017-05-25", "QueryVirtualNumberRelation", "dyvms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryVirtualNumberRelationResponse creates a response to parse from QueryVirtualNumberRelation response
func CreateQueryVirtualNumberRelationResponse() (response *QueryVirtualNumberRelationResponse) {
	response = &QueryVirtualNumberRelationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
