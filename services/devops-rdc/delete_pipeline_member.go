package devops_rdc

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

// DeletePipelineMember invokes the devops_rdc.DeletePipelineMember API synchronously
func (client *Client) DeletePipelineMember(request *DeletePipelineMemberRequest) (response *DeletePipelineMemberResponse, err error) {
	response = CreateDeletePipelineMemberResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePipelineMemberWithChan invokes the devops_rdc.DeletePipelineMember API asynchronously
func (client *Client) DeletePipelineMemberWithChan(request *DeletePipelineMemberRequest) (<-chan *DeletePipelineMemberResponse, <-chan error) {
	responseChan := make(chan *DeletePipelineMemberResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePipelineMember(request)
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

// DeletePipelineMemberWithCallback invokes the devops_rdc.DeletePipelineMember API asynchronously
func (client *Client) DeletePipelineMemberWithCallback(request *DeletePipelineMemberRequest, callback func(response *DeletePipelineMemberResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePipelineMemberResponse
		var err error
		defer close(result)
		response, err = client.DeletePipelineMember(request)
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

// DeletePipelineMemberRequest is the request struct for api DeletePipelineMember
type DeletePipelineMemberRequest struct {
	*requests.RpcRequest
	UserPk     string           `position:"Body" name:"UserPk"`
	UserId     string           `position:"Body" name:"UserId"`
	OrgId      string           `position:"Query" name:"OrgId"`
	PipelineId requests.Integer `position:"Query" name:"PipelineId"`
}

// DeletePipelineMemberResponse is the response struct for api DeletePipelineMember
type DeletePipelineMemberResponse struct {
	*responses.BaseResponse
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Object       bool   `json:"Object" xml:"Object"`
}

// CreateDeletePipelineMemberRequest creates a request to invoke DeletePipelineMember API
func CreateDeletePipelineMemberRequest() (request *DeletePipelineMemberRequest) {
	request = &DeletePipelineMemberRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "DeletePipelineMember", "", "")
	request.Method = requests.POST
	return
}

// CreateDeletePipelineMemberResponse creates a response to parse from DeletePipelineMember response
func CreateDeletePipelineMemberResponse() (response *DeletePipelineMemberResponse) {
	response = &DeletePipelineMemberResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
