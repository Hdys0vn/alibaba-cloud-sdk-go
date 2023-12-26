package emr

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

// DeleteFlowCategory invokes the emr.DeleteFlowCategory API synchronously
func (client *Client) DeleteFlowCategory(request *DeleteFlowCategoryRequest) (response *DeleteFlowCategoryResponse, err error) {
	response = CreateDeleteFlowCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteFlowCategoryWithChan invokes the emr.DeleteFlowCategory API asynchronously
func (client *Client) DeleteFlowCategoryWithChan(request *DeleteFlowCategoryRequest) (<-chan *DeleteFlowCategoryResponse, <-chan error) {
	responseChan := make(chan *DeleteFlowCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteFlowCategory(request)
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

// DeleteFlowCategoryWithCallback invokes the emr.DeleteFlowCategory API asynchronously
func (client *Client) DeleteFlowCategoryWithCallback(request *DeleteFlowCategoryRequest, callback func(response *DeleteFlowCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteFlowCategoryResponse
		var err error
		defer close(result)
		response, err = client.DeleteFlowCategory(request)
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

// DeleteFlowCategoryRequest is the request struct for api DeleteFlowCategory
type DeleteFlowCategoryRequest struct {
	*requests.RpcRequest
	Id        string `position:"Query" name:"Id"`
	ProjectId string `position:"Query" name:"ProjectId"`
}

// DeleteFlowCategoryResponse is the response struct for api DeleteFlowCategory
type DeleteFlowCategoryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      bool   `json:"Data" xml:"Data"`
}

// CreateDeleteFlowCategoryRequest creates a request to invoke DeleteFlowCategory API
func CreateDeleteFlowCategoryRequest() (request *DeleteFlowCategoryRequest) {
	request = &DeleteFlowCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DeleteFlowCategory", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteFlowCategoryResponse creates a response to parse from DeleteFlowCategory response
func CreateDeleteFlowCategoryResponse() (response *DeleteFlowCategoryResponse) {
	response = &DeleteFlowCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
