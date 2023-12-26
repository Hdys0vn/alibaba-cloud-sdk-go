package opensearch

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

// DescribeDataCollction invokes the opensearch.DescribeDataCollction API synchronously
func (client *Client) DescribeDataCollction(request *DescribeDataCollctionRequest) (response *DescribeDataCollctionResponse, err error) {
	response = CreateDescribeDataCollctionResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDataCollctionWithChan invokes the opensearch.DescribeDataCollction API asynchronously
func (client *Client) DescribeDataCollctionWithChan(request *DescribeDataCollctionRequest) (<-chan *DescribeDataCollctionResponse, <-chan error) {
	responseChan := make(chan *DescribeDataCollctionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDataCollction(request)
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

// DescribeDataCollctionWithCallback invokes the opensearch.DescribeDataCollction API asynchronously
func (client *Client) DescribeDataCollctionWithCallback(request *DescribeDataCollctionRequest, callback func(response *DescribeDataCollctionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDataCollctionResponse
		var err error
		defer close(result)
		response, err = client.DescribeDataCollction(request)
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

// DescribeDataCollctionRequest is the request struct for api DescribeDataCollction
type DescribeDataCollctionRequest struct {
	*requests.RoaRequest
	DataCollectionIdentity string `position:"Path" name:"dataCollectionIdentity"`
	AppGroupIdentity       string `position:"Path" name:"appGroupIdentity"`
}

// DescribeDataCollctionResponse is the response struct for api DescribeDataCollction
type DescribeDataCollctionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateDescribeDataCollctionRequest creates a request to invoke DescribeDataCollction API
func CreateDescribeDataCollctionRequest() (request *DescribeDataCollctionRequest) {
	request = &DescribeDataCollctionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "DescribeDataCollction", "/v4/openapi/app-groups/[appGroupIdentity]/data-collections/[dataCollectionIdentity]", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeDataCollctionResponse creates a response to parse from DescribeDataCollction response
func CreateDescribeDataCollctionResponse() (response *DescribeDataCollctionResponse) {
	response = &DescribeDataCollctionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}