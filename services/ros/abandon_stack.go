package ros

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

// AbandonStack invokes the ros.AbandonStack API synchronously
// api document: https://help.aliyun.com/api/ros/abandonstack.html
func (client *Client) AbandonStack(request *AbandonStackRequest) (response *AbandonStackResponse, err error) {
	response = CreateAbandonStackResponse()
	err = client.DoAction(request, response)
	return
}

// AbandonStackWithChan invokes the ros.AbandonStack API asynchronously
// api document: https://help.aliyun.com/api/ros/abandonstack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbandonStackWithChan(request *AbandonStackRequest) (<-chan *AbandonStackResponse, <-chan error) {
	responseChan := make(chan *AbandonStackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AbandonStack(request)
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

// AbandonStackWithCallback invokes the ros.AbandonStack API asynchronously
// api document: https://help.aliyun.com/api/ros/abandonstack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AbandonStackWithCallback(request *AbandonStackRequest, callback func(response *AbandonStackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AbandonStackResponse
		var err error
		defer close(result)
		response, err = client.AbandonStack(request)
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

// AbandonStackRequest is the request struct for api AbandonStack
type AbandonStackRequest struct {
	*requests.RoaRequest
	StackId   string `position:"Path" name:"StackId"`
	StackName string `position:"Path" name:"StackName"`
}

// AbandonStackResponse is the response struct for api AbandonStack
type AbandonStackResponse struct {
	*responses.BaseResponse
}

// CreateAbandonStackRequest creates a request to invoke AbandonStack API
func CreateAbandonStackRequest() (request *AbandonStackRequest) {
	request = &AbandonStackRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "AbandonStack", "/stacks/[StackName]/[StackId]/abandon", "ROS", "openAPI")
	request.Method = requests.DELETE
	return
}

// CreateAbandonStackResponse creates a response to parse from AbandonStack response
func CreateAbandonStackResponse() (response *AbandonStackResponse) {
	response = &AbandonStackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
