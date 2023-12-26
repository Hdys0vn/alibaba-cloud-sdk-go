package schedulerx2

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

// UpdateWorkflow invokes the schedulerx2.UpdateWorkflow API synchronously
func (client *Client) UpdateWorkflow(request *UpdateWorkflowRequest) (response *UpdateWorkflowResponse, err error) {
	response = CreateUpdateWorkflowResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWorkflowWithChan invokes the schedulerx2.UpdateWorkflow API asynchronously
func (client *Client) UpdateWorkflowWithChan(request *UpdateWorkflowRequest) (<-chan *UpdateWorkflowResponse, <-chan error) {
	responseChan := make(chan *UpdateWorkflowResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWorkflow(request)
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

// UpdateWorkflowWithCallback invokes the schedulerx2.UpdateWorkflow API asynchronously
func (client *Client) UpdateWorkflowWithCallback(request *UpdateWorkflowRequest, callback func(response *UpdateWorkflowResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWorkflowResponse
		var err error
		defer close(result)
		response, err = client.UpdateWorkflow(request)
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

// UpdateWorkflowRequest is the request struct for api UpdateWorkflow
type UpdateWorkflowRequest struct {
	*requests.RpcRequest
	NamespaceSource string           `position:"Body" name:"NamespaceSource"`
	Description     string           `position:"Body" name:"Description"`
	WorkflowId      string           `position:"Body" name:"WorkflowId"`
	GroupId         string           `position:"Body" name:"GroupId"`
	TimeExpression  string           `position:"Body" name:"TimeExpression"`
	Namespace       string           `position:"Body" name:"Namespace"`
	Name            string           `position:"Body" name:"Name"`
	TimeType        requests.Integer `position:"Body" name:"TimeType"`
}

// UpdateWorkflowResponse is the response struct for api UpdateWorkflow
type UpdateWorkflowResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateWorkflowRequest creates a request to invoke UpdateWorkflow API
func CreateUpdateWorkflowRequest() (request *UpdateWorkflowRequest) {
	request = &UpdateWorkflowRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "UpdateWorkflow", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateWorkflowResponse creates a response to parse from UpdateWorkflow response
func CreateUpdateWorkflowResponse() (response *UpdateWorkflowResponse) {
	response = &UpdateWorkflowResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
