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

// GenerateMergedTable invokes the opensearch.GenerateMergedTable API synchronously
func (client *Client) GenerateMergedTable(request *GenerateMergedTableRequest) (response *GenerateMergedTableResponse, err error) {
	response = CreateGenerateMergedTableResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateMergedTableWithChan invokes the opensearch.GenerateMergedTable API asynchronously
func (client *Client) GenerateMergedTableWithChan(request *GenerateMergedTableRequest) (<-chan *GenerateMergedTableResponse, <-chan error) {
	responseChan := make(chan *GenerateMergedTableResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateMergedTable(request)
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

// GenerateMergedTableWithCallback invokes the opensearch.GenerateMergedTable API asynchronously
func (client *Client) GenerateMergedTableWithCallback(request *GenerateMergedTableRequest, callback func(response *GenerateMergedTableResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateMergedTableResponse
		var err error
		defer close(result)
		response, err = client.GenerateMergedTable(request)
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

// GenerateMergedTableRequest is the request struct for api GenerateMergedTable
type GenerateMergedTableRequest struct {
	*requests.RoaRequest
	Spec string `position:"Query" name:"spec"`
}

// GenerateMergedTableResponse is the response struct for api GenerateMergedTable
type GenerateMergedTableResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateGenerateMergedTableRequest creates a request to invoke GenerateMergedTable API
func CreateGenerateMergedTableRequest() (request *GenerateMergedTableRequest) {
	request = &GenerateMergedTableRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "GenerateMergedTable", "/v4/openapi/assist/schema/actions/merge", "", "")
	request.Method = requests.POST
	return
}

// CreateGenerateMergedTableResponse creates a response to parse from GenerateMergedTable response
func CreateGenerateMergedTableResponse() (response *GenerateMergedTableResponse) {
	response = &GenerateMergedTableResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
