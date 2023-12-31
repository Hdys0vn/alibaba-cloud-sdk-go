package nas

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

// AddTags invokes the nas.AddTags API synchronously
func (client *Client) AddTags(request *AddTagsRequest) (response *AddTagsResponse, err error) {
	response = CreateAddTagsResponse()
	err = client.DoAction(request, response)
	return
}

// AddTagsWithChan invokes the nas.AddTags API asynchronously
func (client *Client) AddTagsWithChan(request *AddTagsRequest) (<-chan *AddTagsResponse, <-chan error) {
	responseChan := make(chan *AddTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddTags(request)
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

// AddTagsWithCallback invokes the nas.AddTags API asynchronously
func (client *Client) AddTagsWithCallback(request *AddTagsRequest, callback func(response *AddTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddTagsResponse
		var err error
		defer close(result)
		response, err = client.AddTags(request)
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

// AddTagsRequest is the request struct for api AddTags
type AddTagsRequest struct {
	*requests.RpcRequest
	Tag          *[]AddTagsTag `position:"Query" name:"Tag"  type:"Repeated"`
	FileSystemId string        `position:"Query" name:"FileSystemId"`
}

// AddTagsTag is a repeated param struct in AddTagsRequest
type AddTagsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// AddTagsResponse is the response struct for api AddTags
type AddTagsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddTagsRequest creates a request to invoke AddTags API
func CreateAddTagsRequest() (request *AddTagsRequest) {
	request = &AddTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("NAS", "2017-06-26", "AddTags", "nas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddTagsResponse creates a response to parse from AddTags response
func CreateAddTagsResponse() (response *AddTagsResponse) {
	response = &AddTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
