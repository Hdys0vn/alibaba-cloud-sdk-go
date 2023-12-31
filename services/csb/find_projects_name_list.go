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

// FindProjectsNameList invokes the csb.FindProjectsNameList API synchronously
// api document: https://help.aliyun.com/api/csb/findprojectsnamelist.html
func (client *Client) FindProjectsNameList(request *FindProjectsNameListRequest) (response *FindProjectsNameListResponse, err error) {
	response = CreateFindProjectsNameListResponse()
	err = client.DoAction(request, response)
	return
}

// FindProjectsNameListWithChan invokes the csb.FindProjectsNameList API asynchronously
// api document: https://help.aliyun.com/api/csb/findprojectsnamelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindProjectsNameListWithChan(request *FindProjectsNameListRequest) (<-chan *FindProjectsNameListResponse, <-chan error) {
	responseChan := make(chan *FindProjectsNameListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FindProjectsNameList(request)
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

// FindProjectsNameListWithCallback invokes the csb.FindProjectsNameList API asynchronously
// api document: https://help.aliyun.com/api/csb/findprojectsnamelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FindProjectsNameListWithCallback(request *FindProjectsNameListRequest, callback func(response *FindProjectsNameListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FindProjectsNameListResponse
		var err error
		defer close(result)
		response, err = client.FindProjectsNameList(request)
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

// FindProjectsNameListRequest is the request struct for api FindProjectsNameList
type FindProjectsNameListRequest struct {
	*requests.RpcRequest
	OperationFlag string           `position:"Query" name:"OperationFlag"`
	CsbId         requests.Integer `position:"Query" name:"CsbId"`
}

// FindProjectsNameListResponse is the response struct for api FindProjectsNameList
type FindProjectsNameListResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateFindProjectsNameListRequest creates a request to invoke FindProjectsNameList API
func CreateFindProjectsNameListRequest() (request *FindProjectsNameListRequest) {
	request = &FindProjectsNameListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CSB", "2017-11-18", "FindProjectsNameList", "", "")
	request.Method = requests.GET
	return
}

// CreateFindProjectsNameListResponse creates a response to parse from FindProjectsNameList response
func CreateFindProjectsNameListResponse() (response *FindProjectsNameListResponse) {
	response = &FindProjectsNameListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
