package computenestsupplier

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

// DeleteArtifact invokes the computenestsupplier.DeleteArtifact API synchronously
func (client *Client) DeleteArtifact(request *DeleteArtifactRequest) (response *DeleteArtifactResponse, err error) {
	response = CreateDeleteArtifactResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteArtifactWithChan invokes the computenestsupplier.DeleteArtifact API asynchronously
func (client *Client) DeleteArtifactWithChan(request *DeleteArtifactRequest) (<-chan *DeleteArtifactResponse, <-chan error) {
	responseChan := make(chan *DeleteArtifactResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteArtifact(request)
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

// DeleteArtifactWithCallback invokes the computenestsupplier.DeleteArtifact API asynchronously
func (client *Client) DeleteArtifactWithCallback(request *DeleteArtifactRequest, callback func(response *DeleteArtifactResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteArtifactResponse
		var err error
		defer close(result)
		response, err = client.DeleteArtifact(request)
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

// DeleteArtifactRequest is the request struct for api DeleteArtifact
type DeleteArtifactRequest struct {
	*requests.RpcRequest
	ArtifactVersion string `position:"Query" name:"ArtifactVersion"`
	ArtifactId      string `position:"Query" name:"ArtifactId"`
}

// DeleteArtifactResponse is the response struct for api DeleteArtifact
type DeleteArtifactResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteArtifactRequest creates a request to invoke DeleteArtifact API
func CreateDeleteArtifactRequest() (request *DeleteArtifactRequest) {
	request = &DeleteArtifactRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ComputeNestSupplier", "2021-05-21", "DeleteArtifact", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteArtifactResponse creates a response to parse from DeleteArtifact response
func CreateDeleteArtifactResponse() (response *DeleteArtifactResponse) {
	response = &DeleteArtifactResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
