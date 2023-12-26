package vs

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

// PutBucket invokes the vs.PutBucket API synchronously
func (client *Client) PutBucket(request *PutBucketRequest) (response *PutBucketResponse, err error) {
	response = CreatePutBucketResponse()
	err = client.DoAction(request, response)
	return
}

// PutBucketWithChan invokes the vs.PutBucket API asynchronously
func (client *Client) PutBucketWithChan(request *PutBucketRequest) (<-chan *PutBucketResponse, <-chan error) {
	responseChan := make(chan *PutBucketResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PutBucket(request)
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

// PutBucketWithCallback invokes the vs.PutBucket API asynchronously
func (client *Client) PutBucketWithCallback(request *PutBucketRequest, callback func(response *PutBucketResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PutBucketResponse
		var err error
		defer close(result)
		response, err = client.PutBucket(request)
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

// PutBucketRequest is the request struct for api PutBucket
type PutBucketRequest struct {
	*requests.RpcRequest
	DataRedundancyType string           `position:"Query" name:"DataRedundancyType"`
	Endpoint           string           `position:"Query" name:"Endpoint"`
	BucketName         string           `position:"Query" name:"BucketName"`
	ShowLog            string           `position:"Query" name:"ShowLog"`
	BucketAcl          string           `position:"Query" name:"BucketAcl"`
	DispatcherType     string           `position:"Query" name:"DispatcherType"`
	OwnerId            requests.Integer `position:"Query" name:"OwnerId"`
	ResourceType       string           `position:"Query" name:"ResourceType"`
	StorageClass       string           `position:"Query" name:"StorageClass"`
	Comment            string           `position:"Query" name:"Comment"`
}

// PutBucketResponse is the response struct for api PutBucket
type PutBucketResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreatePutBucketRequest creates a request to invoke PutBucket API
func CreatePutBucketRequest() (request *PutBucketRequest) {
	request = &PutBucketRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "PutBucket", "", "")
	request.Method = requests.POST
	return
}

// CreatePutBucketResponse creates a response to parse from PutBucket response
func CreatePutBucketResponse() (response *PutBucketResponse) {
	response = &PutBucketResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
