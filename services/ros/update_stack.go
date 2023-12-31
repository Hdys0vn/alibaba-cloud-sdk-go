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

// UpdateStack invokes the ros.UpdateStack API synchronously
// api document: https://help.aliyun.com/api/ros/updatestack.html
func (client *Client) UpdateStack(request *UpdateStackRequest) (response *UpdateStackResponse, err error) {
	response = CreateUpdateStackResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateStackWithChan invokes the ros.UpdateStack API asynchronously
// api document: https://help.aliyun.com/api/ros/updatestack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateStackWithChan(request *UpdateStackRequest) (<-chan *UpdateStackResponse, <-chan error) {
	responseChan := make(chan *UpdateStackResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateStack(request)
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

// UpdateStackWithCallback invokes the ros.UpdateStack API asynchronously
// api document: https://help.aliyun.com/api/ros/updatestack.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateStackWithCallback(request *UpdateStackRequest, callback func(response *UpdateStackResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateStackResponse
		var err error
		defer close(result)
		response, err = client.UpdateStack(request)
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

// UpdateStackRequest is the request struct for api UpdateStack
type UpdateStackRequest struct {
	*requests.RoaRequest
	StackId   string `position:"Path" name:"StackId"`
	StackName string `position:"Path" name:"StackName"`
}

// UpdateStackResponse is the response struct for api UpdateStack
type UpdateStackResponse struct {
	*responses.BaseResponse
}

// CreateUpdateStackRequest creates a request to invoke UpdateStack API
func CreateUpdateStackRequest() (request *UpdateStackRequest) {
	request = &UpdateStackRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "UpdateStack", "/stacks/[StackName]/[StackId]", "ROS", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdateStackResponse creates a response to parse from UpdateStack response
func CreateUpdateStackResponse() (response *UpdateStackResponse) {
	response = &UpdateStackResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
