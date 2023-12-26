package ddoscoo

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

// CreateTagResources invokes the ddoscoo.CreateTagResources API synchronously
func (client *Client) CreateTagResources(request *CreateTagResourcesRequest) (response *CreateTagResourcesResponse, err error) {
	response = CreateCreateTagResourcesResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTagResourcesWithChan invokes the ddoscoo.CreateTagResources API asynchronously
func (client *Client) CreateTagResourcesWithChan(request *CreateTagResourcesRequest) (<-chan *CreateTagResourcesResponse, <-chan error) {
	responseChan := make(chan *CreateTagResourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTagResources(request)
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

// CreateTagResourcesWithCallback invokes the ddoscoo.CreateTagResources API asynchronously
func (client *Client) CreateTagResourcesWithCallback(request *CreateTagResourcesRequest, callback func(response *CreateTagResourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTagResourcesResponse
		var err error
		defer close(result)
		response, err = client.CreateTagResources(request)
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

// CreateTagResourcesRequest is the request struct for api CreateTagResources
type CreateTagResourcesRequest struct {
	*requests.RpcRequest
	ResourceType    string                    `position:"Query" name:"ResourceType"`
	Tags            *[]CreateTagResourcesTags `position:"Query" name:"Tags"  type:"Repeated"`
	ResourceGroupId string                    `position:"Query" name:"ResourceGroupId"`
	SourceIp        string                    `position:"Query" name:"SourceIp"`
	ResourceIds     *[]string                 `position:"Query" name:"ResourceIds"  type:"Repeated"`
}

// CreateTagResourcesTags is a repeated param struct in CreateTagResourcesRequest
type CreateTagResourcesTags struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateTagResourcesResponse is the response struct for api CreateTagResources
type CreateTagResourcesResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTagResourcesRequest creates a request to invoke CreateTagResources API
func CreateCreateTagResourcesRequest() (request *CreateTagResourcesRequest) {
	request = &CreateTagResourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "CreateTagResources", "ddoscoo", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTagResourcesResponse creates a response to parse from CreateTagResources response
func CreateCreateTagResourcesResponse() (response *CreateTagResourcesResponse) {
	response = &CreateTagResourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
