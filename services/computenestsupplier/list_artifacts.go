package computenestsupplier

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

// ListArtifacts invokes the computenestsupplier.ListArtifacts API synchronously
func (client *Client) ListArtifacts(request *ListArtifactsRequest) (response *ListArtifactsResponse, err error) {
	response = CreateListArtifactsResponse()
	err = client.DoAction(request, response)
	return
}

// ListArtifactsWithChan invokes the computenestsupplier.ListArtifacts API asynchronously
func (client *Client) ListArtifactsWithChan(request *ListArtifactsRequest) (<-chan *ListArtifactsResponse, <-chan error) {
	responseChan := make(chan *ListArtifactsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListArtifacts(request)
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

// ListArtifactsWithCallback invokes the computenestsupplier.ListArtifacts API asynchronously
func (client *Client) ListArtifactsWithCallback(request *ListArtifactsRequest, callback func(response *ListArtifactsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListArtifactsResponse
		var err error
		defer close(result)
		response, err = client.ListArtifacts(request)
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

// ListArtifactsRequest is the request struct for api ListArtifacts
type ListArtifactsRequest struct {
	*requests.RpcRequest
	NextToken  string                 `position:"Query" name:"NextToken"`
	Filter     *[]ListArtifactsFilter `position:"Query" name:"Filter"  type:"Repeated"`
	MaxResults string                 `position:"Query" name:"MaxResults"`
}

// ListArtifactsFilter is a repeated param struct in ListArtifactsRequest
type ListArtifactsFilter struct {
	Values *[]string `name:"Values" type:"Repeated"`
	Name   string    `name:"Name"`
}

// ListArtifactsResponse is the response struct for api ListArtifacts
type ListArtifactsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	NextToken  string     `json:"NextToken" xml:"NextToken"`
	MaxResults string     `json:"MaxResults" xml:"MaxResults"`
	TotalCount string     `json:"TotalCount" xml:"TotalCount"`
	Artifacts  []Artifact `json:"Artifacts" xml:"Artifacts"`
}

// CreateListArtifactsRequest creates a request to invoke ListArtifacts API
func CreateListArtifactsRequest() (request *ListArtifactsRequest) {
	request = &ListArtifactsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "ListArtifacts", "", "")
	request.Method = requests.POST
	return
}

// CreateListArtifactsResponse creates a response to parse from ListArtifacts response
func CreateListArtifactsResponse() (response *ListArtifactsResponse) {
	response = &ListArtifactsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
