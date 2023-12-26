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

// DeleteUdfArtifact invokes the ververica.DeleteUdfArtifact API synchronously
func (client *Client) DeleteUdfArtifact(request *DeleteUdfArtifactRequest) (response *DeleteUdfArtifactResponse, err error) {
	response = CreateDeleteUdfArtifactResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteUdfArtifactWithChan invokes the ververica.DeleteUdfArtifact API asynchronously
func (client *Client) DeleteUdfArtifactWithChan(request *DeleteUdfArtifactRequest) (<-chan *DeleteUdfArtifactResponse, <-chan error) {
	responseChan := make(chan *DeleteUdfArtifactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteUdfArtifact(request)
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

// DeleteUdfArtifactWithCallback invokes the ververica.DeleteUdfArtifact API asynchronously
func (client *Client) DeleteUdfArtifactWithCallback(request *DeleteUdfArtifactRequest, callback func(response *DeleteUdfArtifactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteUdfArtifactResponse
		var err error
		defer close(result)
		response, err = client.DeleteUdfArtifact(request)
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

// DeleteUdfArtifactRequest is the request struct for api DeleteUdfArtifact
type DeleteUdfArtifactRequest struct {
	*requests.RoaRequest
	Workspace       string `position:"Path" name:"workspace"`
	Namespace       string `position:"Path" name:"namespace"`
	UdfArtifactName string `position:"Path" name:"udfArtifactName"`
}

// DeleteUdfArtifactResponse is the response struct for api DeleteUdfArtifact
type DeleteUdfArtifactResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"success" xml:"success"`
	Data      string `json:"data" xml:"data"`
	RequestId string `json:"requestId" xml:"requestId"`
}

// CreateDeleteUdfArtifactRequest creates a request to invoke DeleteUdfArtifact API
func CreateDeleteUdfArtifactRequest() (request *DeleteUdfArtifactRequest) {
	request = &DeleteUdfArtifactRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ververica", "2020-05-01", "DeleteUdfArtifact", "/pop/workspaces/[workspace]/sql/v1beta1/namespaces/[namespace]/udfartifacts/[udfArtifactName]", "", "")
	request.Method = requests.DELETE
	return
}

// CreateDeleteUdfArtifactResponse creates a response to parse from DeleteUdfArtifact response
func CreateDeleteUdfArtifactResponse() (response *DeleteUdfArtifactResponse) {
	response = &DeleteUdfArtifactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}