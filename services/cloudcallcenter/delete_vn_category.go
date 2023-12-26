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

// DeleteVnCategory invokes the cloudcallcenter.DeleteVnCategory API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletevncategory.html
func (client *Client) DeleteVnCategory(request *DeleteVnCategoryRequest) (response *DeleteVnCategoryResponse, err error) {
	response = CreateDeleteVnCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVnCategoryWithChan invokes the cloudcallcenter.DeleteVnCategory API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletevncategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVnCategoryWithChan(request *DeleteVnCategoryRequest) (<-chan *DeleteVnCategoryResponse, <-chan error) {
	responseChan := make(chan *DeleteVnCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVnCategory(request)
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

// DeleteVnCategoryWithCallback invokes the cloudcallcenter.DeleteVnCategory API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/deletevncategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVnCategoryWithCallback(request *DeleteVnCategoryRequest, callback func(response *DeleteVnCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVnCategoryResponse
		var err error
		defer close(result)
		response, err = client.DeleteVnCategory(request)
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

// DeleteVnCategoryRequest is the request struct for api DeleteVnCategory
type DeleteVnCategoryRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	CategoryId string `position:"Query" name:"CategoryId"`
}

// DeleteVnCategoryResponse is the response struct for api DeleteVnCategory
type DeleteVnCategoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteVnCategoryRequest creates a request to invoke DeleteVnCategory API
func CreateDeleteVnCategoryRequest() (request *DeleteVnCategoryRequest) {
	request = &DeleteVnCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "DeleteVnCategory", "", "")
	request.Method = requests.GET
	return
}

// CreateDeleteVnCategoryResponse creates a response to parse from DeleteVnCategory response
func CreateDeleteVnCategoryResponse() (response *DeleteVnCategoryResponse) {
	response = &DeleteVnCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
