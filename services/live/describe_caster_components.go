package live

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

// DescribeCasterComponents invokes the live.DescribeCasterComponents API synchronously
func (client *Client) DescribeCasterComponents(request *DescribeCasterComponentsRequest) (response *DescribeCasterComponentsResponse, err error) {
	response = CreateDescribeCasterComponentsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCasterComponentsWithChan invokes the live.DescribeCasterComponents API asynchronously
func (client *Client) DescribeCasterComponentsWithChan(request *DescribeCasterComponentsRequest) (<-chan *DescribeCasterComponentsResponse, <-chan error) {
	responseChan := make(chan *DescribeCasterComponentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCasterComponents(request)
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

// DescribeCasterComponentsWithCallback invokes the live.DescribeCasterComponents API asynchronously
func (client *Client) DescribeCasterComponentsWithCallback(request *DescribeCasterComponentsRequest, callback func(response *DescribeCasterComponentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCasterComponentsResponse
		var err error
		defer close(result)
		response, err = client.DescribeCasterComponents(request)
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

// DescribeCasterComponentsRequest is the request struct for api DescribeCasterComponents
type DescribeCasterComponentsRequest struct {
	*requests.RpcRequest
	ComponentId string           `position:"Query" name:"ComponentId"`
	CasterId    string           `position:"Query" name:"CasterId"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCasterComponentsResponse is the response struct for api DescribeCasterComponents
type DescribeCasterComponentsResponse struct {
	*responses.BaseResponse
	Total      int        `json:"Total" xml:"Total"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	Components Components `json:"Components" xml:"Components"`
}

// CreateDescribeCasterComponentsRequest creates a request to invoke DescribeCasterComponents API
func CreateDescribeCasterComponentsRequest() (request *DescribeCasterComponentsRequest) {
	request = &DescribeCasterComponentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeCasterComponents", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeCasterComponentsResponse creates a response to parse from DescribeCasterComponents response
func CreateDescribeCasterComponentsResponse() (response *DescribeCasterComponentsResponse) {
	response = &DescribeCasterComponentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
