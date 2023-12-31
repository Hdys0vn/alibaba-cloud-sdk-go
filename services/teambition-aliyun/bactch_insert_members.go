package teambition_aliyun

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

// BactchInsertMembers invokes the teambition_aliyun.BactchInsertMembers API synchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/bactchinsertmembers.html
func (client *Client) BactchInsertMembers(request *BactchInsertMembersRequest) (response *BactchInsertMembersResponse, err error) {
	response = CreateBactchInsertMembersResponse()
	err = client.DoAction(request, response)
	return
}

// BactchInsertMembersWithChan invokes the teambition_aliyun.BactchInsertMembers API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/bactchinsertmembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BactchInsertMembersWithChan(request *BactchInsertMembersRequest) (<-chan *BactchInsertMembersResponse, <-chan error) {
	responseChan := make(chan *BactchInsertMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BactchInsertMembers(request)
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

// BactchInsertMembersWithCallback invokes the teambition_aliyun.BactchInsertMembers API asynchronously
// api document: https://help.aliyun.com/api/teambition-aliyun/bactchinsertmembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BactchInsertMembersWithCallback(request *BactchInsertMembersRequest, callback func(response *BactchInsertMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BactchInsertMembersResponse
		var err error
		defer close(result)
		response, err = client.BactchInsertMembers(request)
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

// BactchInsertMembersRequest is the request struct for api BactchInsertMembers
type BactchInsertMembersRequest struct {
	*requests.RpcRequest
	Members string `position:"Body" name:"Members"`
	RealPk  string `position:"Body" name:"RealPk"`
	OrgId   string `position:"Body" name:"OrgId"`
}

// BactchInsertMembersResponse is the response struct for api BactchInsertMembers
type BactchInsertMembersResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       bool   `json:"Object" xml:"Object"`
}

// CreateBactchInsertMembersRequest creates a request to invoke BactchInsertMembers API
func CreateBactchInsertMembersRequest() (request *BactchInsertMembersRequest) {
	request = &BactchInsertMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("teambition-aliyun", "2020-02-26", "BactchInsertMembers", "", "")
	request.Method = requests.POST
	return
}

// CreateBactchInsertMembersResponse creates a response to parse from BactchInsertMembers response
func CreateBactchInsertMembersResponse() (response *BactchInsertMembersResponse) {
	response = &BactchInsertMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
