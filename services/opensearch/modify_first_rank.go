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

// ModifyFirstRank invokes the opensearch.ModifyFirstRank API synchronously
func (client *Client) ModifyFirstRank(request *ModifyFirstRankRequest) (response *ModifyFirstRankResponse, err error) {
	response = CreateModifyFirstRankResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyFirstRankWithChan invokes the opensearch.ModifyFirstRank API asynchronously
func (client *Client) ModifyFirstRankWithChan(request *ModifyFirstRankRequest) (<-chan *ModifyFirstRankResponse, <-chan error) {
	responseChan := make(chan *ModifyFirstRankResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyFirstRank(request)
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

// ModifyFirstRankWithCallback invokes the opensearch.ModifyFirstRank API asynchronously
func (client *Client) ModifyFirstRankWithCallback(request *ModifyFirstRankRequest, callback func(response *ModifyFirstRankResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyFirstRankResponse
		var err error
		defer close(result)
		response, err = client.ModifyFirstRank(request)
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

// ModifyFirstRankRequest is the request struct for api ModifyFirstRank
type ModifyFirstRankRequest struct {
	*requests.RoaRequest
	DryRun           requests.Boolean `position:"Query" name:"dryRun"`
	AppId            requests.Integer `position:"Path" name:"appId"`
	Name             string           `position:"Path" name:"name"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// ModifyFirstRankResponse is the response struct for api ModifyFirstRank
type ModifyFirstRankResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Result    Result `json:"result" xml:"result"`
}

// CreateModifyFirstRankRequest creates a request to invoke ModifyFirstRank API
func CreateModifyFirstRankRequest() (request *ModifyFirstRankRequest) {
	request = &ModifyFirstRankRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ModifyFirstRank", "/v4/openapi/app-groups/[appGroupIdentity]/apps/[appId]/first-ranks/[name]", "", "")
	request.Method = requests.PUT
	return
}

// CreateModifyFirstRankResponse creates a response to parse from ModifyFirstRank response
func CreateModifyFirstRankResponse() (response *ModifyFirstRankResponse) {
	response = &ModifyFirstRankResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
