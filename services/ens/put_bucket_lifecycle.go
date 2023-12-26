package ens

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

// PutBucketLifecycle invokes the ens.PutBucketLifecycle API synchronously
func (client *Client) PutBucketLifecycle(request *PutBucketLifecycleRequest) (response *PutBucketLifecycleResponse, err error) {
	response = CreatePutBucketLifecycleResponse()
	err = client.DoAction(request, response)
	return
}

// PutBucketLifecycleWithChan invokes the ens.PutBucketLifecycle API asynchronously
func (client *Client) PutBucketLifecycleWithChan(request *PutBucketLifecycleRequest) (<-chan *PutBucketLifecycleResponse, <-chan error) {
	responseChan := make(chan *PutBucketLifecycleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutBucketLifecycle(request)
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

// PutBucketLifecycleWithCallback invokes the ens.PutBucketLifecycle API asynchronously
func (client *Client) PutBucketLifecycleWithCallback(request *PutBucketLifecycleRequest, callback func(response *PutBucketLifecycleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutBucketLifecycleResponse
		var err error
		defer close(result)
		response, err = client.PutBucketLifecycle(request)
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

// PutBucketLifecycleRequest is the request struct for api PutBucketLifecycle
type PutBucketLifecycleRequest struct {
	*requests.RpcRequest
	Prefix                 string           `position:"Query" name:"Prefix"`
	AllowSameActionOverlap string           `position:"Query" name:"AllowSameActionOverlap"`
	ExpirationDays         requests.Integer `position:"Query" name:"ExpirationDays"`
	RuleId                 string           `position:"Query" name:"RuleId"`
	Status                 string           `position:"Query" name:"Status"`
	BucketName             string           `position:"Query" name:"BucketName"`
	CreatedBeforeDate      string           `position:"Query" name:"CreatedBeforeDate"`
}

// PutBucketLifecycleResponse is the response struct for api PutBucketLifecycle
type PutBucketLifecycleResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePutBucketLifecycleRequest creates a request to invoke PutBucketLifecycle API
func CreatePutBucketLifecycleRequest() (request *PutBucketLifecycleRequest) {
	request = &PutBucketLifecycleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "PutBucketLifecycle", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePutBucketLifecycleResponse creates a response to parse from PutBucketLifecycle response
func CreatePutBucketLifecycleResponse() (response *PutBucketLifecycleResponse) {
	response = &PutBucketLifecycleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}