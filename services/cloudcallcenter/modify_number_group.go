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

// ModifyNumberGroup invokes the cloudcallcenter.ModifyNumberGroup API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifynumbergroup.html
func (client *Client) ModifyNumberGroup(request *ModifyNumberGroupRequest) (response *ModifyNumberGroupResponse, err error) {
	response = CreateModifyNumberGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyNumberGroupWithChan invokes the cloudcallcenter.ModifyNumberGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifynumbergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyNumberGroupWithChan(request *ModifyNumberGroupRequest) (<-chan *ModifyNumberGroupResponse, <-chan error) {
	responseChan := make(chan *ModifyNumberGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyNumberGroup(request)
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

// ModifyNumberGroupWithCallback invokes the cloudcallcenter.ModifyNumberGroup API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/modifynumbergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyNumberGroupWithCallback(request *ModifyNumberGroupRequest, callback func(response *ModifyNumberGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyNumberGroupResponse
		var err error
		defer close(result)
		response, err = client.ModifyNumberGroup(request)
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

// ModifyNumberGroupRequest is the request struct for api ModifyNumberGroup
type ModifyNumberGroupRequest struct {
	*requests.RpcRequest
	NumberGroupName string    `position:"Query" name:"NumberGroupName"`
	Description     string    `position:"Query" name:"Description"`
	ToAddNumbers    *[]string `position:"Query" name:"ToAddNumbers"  type:"Repeated"`
	NumberGroupId   string    `position:"Query" name:"NumberGroupId"`
	ToDeleteNumbers *[]string `position:"Query" name:"ToDeleteNumbers"  type:"Repeated"`
}

// ModifyNumberGroupResponse is the response struct for api ModifyNumberGroup
type ModifyNumberGroupResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifyNumberGroupRequest creates a request to invoke ModifyNumberGroup API
func CreateModifyNumberGroupRequest() (request *ModifyNumberGroupRequest) {
	request = &ModifyNumberGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ModifyNumberGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyNumberGroupResponse creates a response to parse from ModifyNumberGroup response
func CreateModifyNumberGroupResponse() (response *ModifyNumberGroupResponse) {
	response = &ModifyNumberGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
