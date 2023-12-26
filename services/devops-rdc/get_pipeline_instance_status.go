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

// GetPipelineInstanceStatus invokes the devops_rdc.GetPipelineInstanceStatus API synchronously
func (client *Client) GetPipelineInstanceStatus(request *GetPipelineInstanceStatusRequest) (response *GetPipelineInstanceStatusResponse, err error) {
	response = CreateGetPipelineInstanceStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetPipelineInstanceStatusWithChan invokes the devops_rdc.GetPipelineInstanceStatus API asynchronously
func (client *Client) GetPipelineInstanceStatusWithChan(request *GetPipelineInstanceStatusRequest) (<-chan *GetPipelineInstanceStatusResponse, <-chan error) {
	responseChan := make(chan *GetPipelineInstanceStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetPipelineInstanceStatus(request)
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

// GetPipelineInstanceStatusWithCallback invokes the devops_rdc.GetPipelineInstanceStatus API asynchronously
func (client *Client) GetPipelineInstanceStatusWithCallback(request *GetPipelineInstanceStatusRequest, callback func(response *GetPipelineInstanceStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetPipelineInstanceStatusResponse
		var err error
		defer close(result)
		response, err = client.GetPipelineInstanceStatus(request)
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

// GetPipelineInstanceStatusRequest is the request struct for api GetPipelineInstanceStatus
type GetPipelineInstanceStatusRequest struct {
	*requests.RpcRequest
	FlowInstanceId requests.Integer `position:"Query" name:"FlowInstanceId"`
	UserPk         string           `position:"Body" name:"UserPk"`
	OrgId          string           `position:"Query" name:"OrgId"`
	PipelineId     requests.Integer `position:"Query" name:"PipelineId"`
}

// GetPipelineInstanceStatusResponse is the response struct for api GetPipelineInstanceStatus
type GetPipelineInstanceStatusResponse struct {
	*responses.BaseResponse
	Success      bool   `json:"Success" xml:"Success"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Object       string `json:"Object" xml:"Object"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
}

// CreateGetPipelineInstanceStatusRequest creates a request to invoke GetPipelineInstanceStatus API
func CreateGetPipelineInstanceStatusRequest() (request *GetPipelineInstanceStatusRequest) {
	request = &GetPipelineInstanceStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "GetPipelineInstanceStatus", "", "")
	request.Method = requests.POST
	return
}

// CreateGetPipelineInstanceStatusResponse creates a response to parse from GetPipelineInstanceStatus response
func CreateGetPipelineInstanceStatusResponse() (response *GetPipelineInstanceStatusResponse) {
	response = &GetPipelineInstanceStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
