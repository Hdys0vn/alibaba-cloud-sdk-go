package cloudcallcenter

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

// ResumeService invokes the cloudcallcenter.ResumeService API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/resumeservice.html
func (client *Client) ResumeService(request *ResumeServiceRequest) (response *ResumeServiceResponse, err error) {
	response = CreateResumeServiceResponse()
	err = client.DoAction(request, response)
	return
}

// ResumeServiceWithChan invokes the cloudcallcenter.ResumeService API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/resumeservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeServiceWithChan(request *ResumeServiceRequest) (<-chan *ResumeServiceResponse, <-chan error) {
	responseChan := make(chan *ResumeServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResumeService(request)
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

// ResumeServiceWithCallback invokes the cloudcallcenter.ResumeService API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/resumeservice.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResumeServiceWithCallback(request *ResumeServiceRequest, callback func(response *ResumeServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResumeServiceResponse
		var err error
		defer close(result)
		response, err = client.ResumeService(request)
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

// ResumeServiceRequest is the request struct for api ResumeService
type ResumeServiceRequest struct {
	*requests.RpcRequest
	OwnerUid requests.Integer `position:"Query" name:"OwnerUid"`
}

// ResumeServiceResponse is the response struct for api ResumeService
type ResumeServiceResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateResumeServiceRequest creates a request to invoke ResumeService API
func CreateResumeServiceRequest() (request *ResumeServiceRequest) {
	request = &ResumeServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ResumeService", "", "")
	request.Method = requests.POST
	return
}

// CreateResumeServiceResponse creates a response to parse from ResumeService response
func CreateResumeServiceResponse() (response *ResumeServiceResponse) {
	response = &ResumeServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
