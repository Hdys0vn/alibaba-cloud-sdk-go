package baas

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

// UpdateAntChainContractProject invokes the baas.UpdateAntChainContractProject API synchronously
// api document: https://help.aliyun.com/api/baas/updateantchaincontractproject.html
func (client *Client) UpdateAntChainContractProject(request *UpdateAntChainContractProjectRequest) (response *UpdateAntChainContractProjectResponse, err error) {
	response = CreateUpdateAntChainContractProjectResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAntChainContractProjectWithChan invokes the baas.UpdateAntChainContractProject API asynchronously
// api document: https://help.aliyun.com/api/baas/updateantchaincontractproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAntChainContractProjectWithChan(request *UpdateAntChainContractProjectRequest) (<-chan *UpdateAntChainContractProjectResponse, <-chan error) {
	responseChan := make(chan *UpdateAntChainContractProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAntChainContractProject(request)
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

// UpdateAntChainContractProjectWithCallback invokes the baas.UpdateAntChainContractProject API asynchronously
// api document: https://help.aliyun.com/api/baas/updateantchaincontractproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAntChainContractProjectWithCallback(request *UpdateAntChainContractProjectRequest, callback func(response *UpdateAntChainContractProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAntChainContractProjectResponse
		var err error
		defer close(result)
		response, err = client.UpdateAntChainContractProject(request)
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

// UpdateAntChainContractProjectRequest is the request struct for api UpdateAntChainContractProject
type UpdateAntChainContractProjectRequest struct {
	*requests.RpcRequest
	ProjectVersion     string `position:"Body" name:"ProjectVersion"`
	ProjectId          string `position:"Body" name:"ProjectId"`
	ProjectName        string `position:"Body" name:"ProjectName"`
	ProjectDescription string `position:"Body" name:"ProjectDescription"`
}

// UpdateAntChainContractProjectResponse is the response struct for api UpdateAntChainContractProject
type UpdateAntChainContractProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    string `json:"Result" xml:"Result"`
}

// CreateUpdateAntChainContractProjectRequest creates a request to invoke UpdateAntChainContractProject API
func CreateUpdateAntChainContractProjectRequest() (request *UpdateAntChainContractProjectRequest) {
	request = &UpdateAntChainContractProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "UpdateAntChainContractProject", "baas", "openAPI")
	return
}

// CreateUpdateAntChainContractProjectResponse creates a response to parse from UpdateAntChainContractProject response
func CreateUpdateAntChainContractProjectResponse() (response *UpdateAntChainContractProjectResponse) {
	response = &UpdateAntChainContractProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}