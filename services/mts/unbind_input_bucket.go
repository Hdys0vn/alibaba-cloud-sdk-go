package mts

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

// UnbindInputBucket invokes the mts.UnbindInputBucket API synchronously
func (client *Client) UnbindInputBucket(request *UnbindInputBucketRequest) (response *UnbindInputBucketResponse, err error) {
	response = CreateUnbindInputBucketResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindInputBucketWithChan invokes the mts.UnbindInputBucket API asynchronously
func (client *Client) UnbindInputBucketWithChan(request *UnbindInputBucketRequest) (<-chan *UnbindInputBucketResponse, <-chan error) {
	responseChan := make(chan *UnbindInputBucketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindInputBucket(request)
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

// UnbindInputBucketWithCallback invokes the mts.UnbindInputBucket API asynchronously
func (client *Client) UnbindInputBucketWithCallback(request *UnbindInputBucketRequest, callback func(response *UnbindInputBucketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindInputBucketResponse
		var err error
		defer close(result)
		response, err = client.UnbindInputBucket(request)
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

// UnbindInputBucketRequest is the request struct for api UnbindInputBucket
type UnbindInputBucketRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Bucket               string           `position:"Query" name:"Bucket"`
	RoleArn              string           `position:"Query" name:"RoleArn"`
}

// UnbindInputBucketResponse is the response struct for api UnbindInputBucket
type UnbindInputBucketResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindInputBucketRequest creates a request to invoke UnbindInputBucket API
func CreateUnbindInputBucketRequest() (request *UnbindInputBucketRequest) {
	request = &UnbindInputBucketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UnbindInputBucket", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindInputBucketResponse creates a response to parse from UnbindInputBucket response
func CreateUnbindInputBucketResponse() (response *UnbindInputBucketResponse) {
	response = &UnbindInputBucketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}