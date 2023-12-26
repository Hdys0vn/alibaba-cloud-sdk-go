package kms

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

// TagResource invokes the kms.TagResource API synchronously
func (client *Client) TagResource(request *TagResourceRequest) (response *TagResourceResponse, err error) {
	response = CreateTagResourceResponse()
	err = client.DoAction(request, response)
	return
}

// TagResourceWithChan invokes the kms.TagResource API asynchronously
func (client *Client) TagResourceWithChan(request *TagResourceRequest) (<-chan *TagResourceResponse, <-chan error) {
	responseChan := make(chan *TagResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TagResource(request)
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

// TagResourceWithCallback invokes the kms.TagResource API asynchronously
func (client *Client) TagResourceWithCallback(request *TagResourceRequest, callback func(response *TagResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TagResourceResponse
		var err error
		defer close(result)
		response, err = client.TagResource(request)
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

// TagResourceRequest is the request struct for api TagResource
type TagResourceRequest struct {
	*requests.RpcRequest
	CertificateId string `position:"Query" name:"CertificateId"`
	KeyId         string `position:"Query" name:"KeyId"`
	Tags          string `position:"Query" name:"Tags"`
	SecretName    string `position:"Query" name:"SecretName"`
}

// TagResourceResponse is the response struct for api TagResource
type TagResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateTagResourceRequest creates a request to invoke TagResource API
func CreateTagResourceRequest() (request *TagResourceRequest) {
	request = &TagResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Kms", "2016-01-20", "TagResource", "kms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateTagResourceResponse creates a response to parse from TagResource response
func CreateTagResourceResponse() (response *TagResourceResponse) {
	response = &TagResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
