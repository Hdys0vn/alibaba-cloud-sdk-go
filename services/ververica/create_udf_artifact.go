package ververica

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

// CreateUdfArtifact invokes the ververica.CreateUdfArtifact API synchronously
func (client *Client) CreateUdfArtifact(request *CreateUdfArtifactRequest) (response *CreateUdfArtifactResponse, err error) {
	response = CreateCreateUdfArtifactResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUdfArtifactWithChan invokes the ververica.CreateUdfArtifact API asynchronously
func (client *Client) CreateUdfArtifactWithChan(request *CreateUdfArtifactRequest) (<-chan *CreateUdfArtifactResponse, <-chan error) {
	responseChan := make(chan *CreateUdfArtifactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUdfArtifact(request)
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

// CreateUdfArtifactWithCallback invokes the ververica.CreateUdfArtifact API asynchronously
func (client *Client) CreateUdfArtifactWithCallback(request *CreateUdfArtifactRequest, callback func(response *CreateUdfArtifactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUdfArtifactResponse
		var err error
		defer close(result)
		response, err = client.CreateUdfArtifact(request)
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

// CreateUdfArtifactRequest is the request struct for api CreateUdfArtifact
type CreateUdfArtifactRequest struct {
	*requests.RoaRequest
	Workspace  string `position:"Path" name:"workspace"`
	ParamsJson string `position:"Body" name:"paramsJson"`
	Namespace  string `position:"Path" name:"namespace"`
}

// CreateUdfArtifactResponse is the response struct for api CreateUdfArtifact
type CreateUdfArtifactResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateCreateUdfArtifactRequest creates a request to invoke CreateUdfArtifact API
func CreateCreateUdfArtifactRequest() (request *CreateUdfArtifactRequest) {
	request = &CreateUdfArtifactRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "CreateUdfArtifact", "/pop/workspaces/[workspace]/sql/v1beta1/namespaces/[namespace]/udfartifacts", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateUdfArtifactResponse creates a response to parse from CreateUdfArtifact response
func CreateCreateUdfArtifactResponse() (response *CreateUdfArtifactResponse) {
	response = &CreateUdfArtifactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
