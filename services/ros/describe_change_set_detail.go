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

// DescribeChangeSetDetail invokes the ros.DescribeChangeSetDetail API synchronously
// api document: https://help.aliyun.com/api/ros/describechangesetdetail.html
func (client *Client) DescribeChangeSetDetail(request *DescribeChangeSetDetailRequest) (response *DescribeChangeSetDetailResponse, err error) {
	response = CreateDescribeChangeSetDetailResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeChangeSetDetailWithChan invokes the ros.DescribeChangeSetDetail API asynchronously
// api document: https://help.aliyun.com/api/ros/describechangesetdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeChangeSetDetailWithChan(request *DescribeChangeSetDetailRequest) (<-chan *DescribeChangeSetDetailResponse, <-chan error) {
	responseChan := make(chan *DescribeChangeSetDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeChangeSetDetail(request)
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

// DescribeChangeSetDetailWithCallback invokes the ros.DescribeChangeSetDetail API asynchronously
// api document: https://help.aliyun.com/api/ros/describechangesetdetail.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeChangeSetDetailWithCallback(request *DescribeChangeSetDetailRequest, callback func(response *DescribeChangeSetDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeChangeSetDetailResponse
		var err error
		defer close(result)
		response, err = client.DescribeChangeSetDetail(request)
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

// DescribeChangeSetDetailRequest is the request struct for api DescribeChangeSetDetail
type DescribeChangeSetDetailRequest struct {
	*requests.RoaRequest
	ChangeSetName string `position:"Path" name:"ChangeSetName"`
	StackId       string `position:"Path" name:"StackId"`
	StackName     string `position:"Path" name:"StackName"`
}

// DescribeChangeSetDetailResponse is the response struct for api DescribeChangeSetDetail
type DescribeChangeSetDetailResponse struct {
	*responses.BaseResponse
	Dummy string `json:"Dummy" xml:"Dummy"`
}

// CreateDescribeChangeSetDetailRequest creates a request to invoke DescribeChangeSetDetail API
func CreateDescribeChangeSetDetailRequest() (request *DescribeChangeSetDetailRequest) {
	request = &DescribeChangeSetDetailRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "DescribeChangeSetDetail", "/stacks/[StackName]/[StackId]/changeSets/[ChangeSetName]", "ROS", "openAPI")
	request.Method = requests.GET
	return
}

// CreateDescribeChangeSetDetailResponse creates a response to parse from DescribeChangeSetDetail response
func CreateDescribeChangeSetDetailResponse() (response *DescribeChangeSetDetailResponse) {
	response = &DescribeChangeSetDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
