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

// ModifyQueryProcessor invokes the opensearch.ModifyQueryProcessor API synchronously
func (client *Client) ModifyQueryProcessor(request *ModifyQueryProcessorRequest) (response *ModifyQueryProcessorResponse, err error) {
	response = CreateModifyQueryProcessorResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyQueryProcessorWithChan invokes the opensearch.ModifyQueryProcessor API asynchronously
func (client *Client) ModifyQueryProcessorWithChan(request *ModifyQueryProcessorRequest) (<-chan *ModifyQueryProcessorResponse, <-chan error) {
	responseChan := make(chan *ModifyQueryProcessorResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyQueryProcessor(request)
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

// ModifyQueryProcessorWithCallback invokes the opensearch.ModifyQueryProcessor API asynchronously
func (client *Client) ModifyQueryProcessorWithCallback(request *ModifyQueryProcessorRequest, callback func(response *ModifyQueryProcessorResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyQueryProcessorResponse
		var err error
		defer close(result)
		response, err = client.ModifyQueryProcessor(request)
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

// ModifyQueryProcessorRequest is the request struct for api ModifyQueryProcessor
type ModifyQueryProcessorRequest struct {
	*requests.RoaRequest
	DryRun           requests.Boolean `position:"Query" name:"dryRun"`
	AppId            requests.Integer `position:"Path" name:"appId"`
	Name             string           `position:"Path" name:"name"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// ModifyQueryProcessorResponse is the response struct for api ModifyQueryProcessor
type ModifyQueryProcessorResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateModifyQueryProcessorRequest creates a request to invoke ModifyQueryProcessor API
func CreateModifyQueryProcessorRequest() (request *ModifyQueryProcessorRequest) {
	request = &ModifyQueryProcessorRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ModifyQueryProcessor", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/query-processors/[name]", "", "")
	request.Method = requests.PUT
	return
}

// CreateModifyQueryProcessorResponse creates a response to parse from ModifyQueryProcessor response
func CreateModifyQueryProcessorResponse() (response *ModifyQueryProcessorResponse) {
	response = &ModifyQueryProcessorResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
