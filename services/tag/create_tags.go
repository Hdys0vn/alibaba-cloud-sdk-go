package tag

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

// CreateTags invokes the tag.CreateTags API synchronously
func (client *Client) CreateTags(request *CreateTagsRequest) (response *CreateTagsResponse, err error) {
	response = CreateCreateTagsResponse()
	err = client.DoAction(request, response)
	return
}

// CreateTagsWithChan invokes the tag.CreateTags API asynchronously
func (client *Client) CreateTagsWithChan(request *CreateTagsRequest) (<-chan *CreateTagsResponse, <-chan error) {
	responseChan := make(chan *CreateTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateTags(request)
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

// CreateTagsWithCallback invokes the tag.CreateTags API asynchronously
func (client *Client) CreateTagsWithCallback(request *CreateTagsRequest, callback func(response *CreateTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateTagsResponse
		var err error
		defer close(result)
		response, err = client.CreateTags(request)
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

// CreateTagsRequest is the request struct for api CreateTags
type CreateTagsRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string                            `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string                            `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer                  `position:"Query" name:"OwnerId"`
	TagKeyValueParamList *[]CreateTagsTagKeyValueParamList `position:"Query" name:"TagKeyValueParamList"  type:"Repeated"`
}

// CreateTagsTagKeyValueParamList is a repeated param struct in CreateTagsRequest
type CreateTagsTagKeyValueParamList struct {
	Key               string                                             `name:"Key"`
	TagValueParamList *[]CreateTagsTagKeyValueParamListTagValueParamList `name:"TagValueParamList" type:"Repeated"`
	Description       string                                             `name:"Description"`
}

// CreateTagsTagKeyValueParamListTagValueParamList is a repeated param struct in CreateTagsRequest
type CreateTagsTagKeyValueParamListTagValueParamList struct {
	Value       string `name:"Value"`
	Description string `name:"Description"`
}

// CreateTagsResponse is the response struct for api CreateTags
type CreateTagsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateTagsRequest creates a request to invoke CreateTags API
func CreateCreateTagsRequest() (request *CreateTagsRequest) {
	request = &CreateTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Tag", "2018-08-28", "CreateTags", "tag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateTagsResponse creates a response to parse from CreateTags response
func CreateCreateTagsResponse() (response *CreateTagsResponse) {
	response = &CreateTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
