package csb

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

// DeleteProjectList invokes the csb.DeleteProjectList API synchronously
// api document: https://help.aliyun.com/api/csb/deleteprojectlist.html
func (client *Client) DeleteProjectList(request *DeleteProjectListRequest) (response *DeleteProjectListResponse, err error) {
	response = CreateDeleteProjectListResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteProjectListWithChan invokes the csb.DeleteProjectList API asynchronously
// api document: https://help.aliyun.com/api/csb/deleteprojectlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteProjectListWithChan(request *DeleteProjectListRequest) (<-chan *DeleteProjectListResponse, <-chan error) {
	responseChan := make(chan *DeleteProjectListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteProjectList(request)
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

// DeleteProjectListWithCallback invokes the csb.DeleteProjectList API asynchronously
// api document: https://help.aliyun.com/api/csb/deleteprojectlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteProjectListWithCallback(request *DeleteProjectListRequest, callback func(response *DeleteProjectListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteProjectListResponse
		var err error
		defer close(result)
		response, err = client.DeleteProjectList(request)
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

// DeleteProjectListRequest is the request struct for api DeleteProjectList
type DeleteProjectListRequest struct {
	*requests.RpcRequest
	Data  string           `position:"Body" name:"Data"`
	CsbId requests.Integer `position:"Query" name:"CsbId"`
}

// DeleteProjectListResponse is the response struct for api DeleteProjectList
type DeleteProjectListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteProjectListRequest creates a request to invoke DeleteProjectList API
func CreateDeleteProjectListRequest() (request *DeleteProjectListRequest) {
	request = &DeleteProjectListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "DeleteProjectList", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteProjectListResponse creates a response to parse from DeleteProjectList response
func CreateDeleteProjectListResponse() (response *DeleteProjectListResponse) {
	response = &DeleteProjectListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
