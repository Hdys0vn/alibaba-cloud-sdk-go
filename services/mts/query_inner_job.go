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

// QueryInnerJob invokes the mts.QueryInnerJob API synchronously
func (client *Client) QueryInnerJob(request *QueryInnerJobRequest) (response *QueryInnerJobResponse, err error) {
	response = CreateQueryInnerJobResponse()
	err = client.DoAction(request, response)
	return
}

// QueryInnerJobWithChan invokes the mts.QueryInnerJob API asynchronously
func (client *Client) QueryInnerJobWithChan(request *QueryInnerJobRequest) (<-chan *QueryInnerJobResponse, <-chan error) {
	responseChan := make(chan *QueryInnerJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryInnerJob(request)
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

// QueryInnerJobWithCallback invokes the mts.QueryInnerJob API asynchronously
func (client *Client) QueryInnerJobWithCallback(request *QueryInnerJobRequest, callback func(response *QueryInnerJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryInnerJobResponse
		var err error
		defer close(result)
		response, err = client.QueryInnerJob(request)
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

// QueryInnerJobRequest is the request struct for api QueryInnerJob
type QueryInnerJobRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobId                string           `position:"Query" name:"JobId"`
}

// QueryInnerJobResponse is the response struct for api QueryInnerJob
type QueryInnerJobResponse struct {
	*responses.BaseResponse
	Status     string `json:"Status" xml:"Status"`
	Suggestion string `json:"Suggestion" xml:"Suggestion"`
	RequestId  string `json:"RequestId" xml:"RequestId"`
	Video      Video  `json:"Video" xml:"Video"`
	Image      Image  `json:"Image" xml:"Image"`
}

// CreateQueryInnerJobRequest creates a request to invoke QueryInnerJob API
func CreateQueryInnerJobRequest() (request *QueryInnerJobRequest) {
	request = &QueryInnerJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryInnerJob", "mts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryInnerJobResponse creates a response to parse from QueryInnerJob response
func CreateQueryInnerJobResponse() (response *QueryInnerJobResponse) {
	response = &QueryInnerJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
