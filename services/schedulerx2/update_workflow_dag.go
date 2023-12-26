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

// UpdateWorkflowDag invokes the schedulerx2.UpdateWorkflowDag API synchronously
func (client *Client) UpdateWorkflowDag(request *UpdateWorkflowDagRequest) (response *UpdateWorkflowDagResponse, err error) {
	response = CreateUpdateWorkflowDagResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateWorkflowDagWithChan invokes the schedulerx2.UpdateWorkflowDag API asynchronously
func (client *Client) UpdateWorkflowDagWithChan(request *UpdateWorkflowDagRequest) (<-chan *UpdateWorkflowDagResponse, <-chan error) {
	responseChan := make(chan *UpdateWorkflowDagResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateWorkflowDag(request)
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

// UpdateWorkflowDagWithCallback invokes the schedulerx2.UpdateWorkflowDag API asynchronously
func (client *Client) UpdateWorkflowDagWithCallback(request *UpdateWorkflowDagRequest, callback func(response *UpdateWorkflowDagResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateWorkflowDagResponse
		var err error
		defer close(result)
		response, err = client.UpdateWorkflowDag(request)
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

// UpdateWorkflowDagRequest is the request struct for api UpdateWorkflowDag
type UpdateWorkflowDagRequest struct {
	*requests.RpcRequest
	DagJson         string `position:"Body" name:"DagJson"`
	NamespaceSource string `position:"Body" name:"NamespaceSource"`
	GroupId         string `position:"Body" name:"GroupId"`
	Namespace       string `position:"Body" name:"Namespace"`
	WorkflowId      string `position:"Body" name:"WorkflowId"`
}

// UpdateWorkflowDagResponse is the response struct for api UpdateWorkflowDag
type UpdateWorkflowDagResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Code      int    `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
}

// CreateUpdateWorkflowDagRequest creates a request to invoke UpdateWorkflowDag API
func CreateUpdateWorkflowDagRequest() (request *UpdateWorkflowDagRequest) {
	request = &UpdateWorkflowDagRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("schedulerx2", "2019-04-30", "UpdateWorkflowDag", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateWorkflowDagResponse creates a response to parse from UpdateWorkflowDag response
func CreateUpdateWorkflowDagResponse() (response *UpdateWorkflowDagResponse) {
	response = &UpdateWorkflowDagResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}