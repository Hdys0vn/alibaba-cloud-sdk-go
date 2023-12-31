package cloudwf

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

// GetApgroupSsidConfigProgress invokes the cloudwf.GetApgroupSsidConfigProgress API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getapgroupssidconfigprogress.html
func (client *Client) GetApgroupSsidConfigProgress(request *GetApgroupSsidConfigProgressRequest) (response *GetApgroupSsidConfigProgressResponse, err error) {
	response = CreateGetApgroupSsidConfigProgressResponse()
	err = client.DoAction(request, response)
	return
}

// GetApgroupSsidConfigProgressWithChan invokes the cloudwf.GetApgroupSsidConfigProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getapgroupssidconfigprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApgroupSsidConfigProgressWithChan(request *GetApgroupSsidConfigProgressRequest) (<-chan *GetApgroupSsidConfigProgressResponse, <-chan error) {
	responseChan := make(chan *GetApgroupSsidConfigProgressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetApgroupSsidConfigProgress(request)
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

// GetApgroupSsidConfigProgressWithCallback invokes the cloudwf.GetApgroupSsidConfigProgress API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getapgroupssidconfigprogress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetApgroupSsidConfigProgressWithCallback(request *GetApgroupSsidConfigProgressRequest, callback func(response *GetApgroupSsidConfigProgressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetApgroupSsidConfigProgressResponse
		var err error
		defer close(result)
		response, err = client.GetApgroupSsidConfigProgress(request)
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

// GetApgroupSsidConfigProgressRequest is the request struct for api GetApgroupSsidConfigProgress
type GetApgroupSsidConfigProgressRequest struct {
	*requests.RpcRequest
	Id requests.Integer `position:"Query" name:"Id"`
}

// GetApgroupSsidConfigProgressResponse is the response struct for api GetApgroupSsidConfigProgress
type GetApgroupSsidConfigProgressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetApgroupSsidConfigProgressRequest creates a request to invoke GetApgroupSsidConfigProgress API
func CreateGetApgroupSsidConfigProgressRequest() (request *GetApgroupSsidConfigProgressRequest) {
	request = &GetApgroupSsidConfigProgressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetApgroupSsidConfigProgress", "cloudwf", "openAPI")
	return
}

// CreateGetApgroupSsidConfigProgressResponse creates a response to parse from GetApgroupSsidConfigProgress response
func CreateGetApgroupSsidConfigProgressResponse() (response *GetApgroupSsidConfigProgressResponse) {
	response = &GetApgroupSsidConfigProgressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
