package pts

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

// DeletePtsScenes invokes the pts.DeletePtsScenes API synchronously
func (client *Client) DeletePtsScenes(request *DeletePtsScenesRequest) (response *DeletePtsScenesResponse, err error) {
	response = CreateDeletePtsScenesResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePtsScenesWithChan invokes the pts.DeletePtsScenes API asynchronously
func (client *Client) DeletePtsScenesWithChan(request *DeletePtsScenesRequest) (<-chan *DeletePtsScenesResponse, <-chan error) {
	responseChan := make(chan *DeletePtsScenesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePtsScenes(request)
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

// DeletePtsScenesWithCallback invokes the pts.DeletePtsScenes API asynchronously
func (client *Client) DeletePtsScenesWithCallback(request *DeletePtsScenesRequest, callback func(response *DeletePtsScenesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePtsScenesResponse
		var err error
		defer close(result)
		response, err = client.DeletePtsScenes(request)
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

// DeletePtsScenesRequest is the request struct for api DeletePtsScenes
type DeletePtsScenesRequest struct {
	*requests.RpcRequest
	SceneIds string `position:"Query" name:"SceneIds"`
}

// DeletePtsScenesResponse is the response struct for api DeletePtsScenes
type DeletePtsScenesResponse struct {
	*responses.BaseResponse
	Message        string `json:"Message" xml:"Message"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string `json:"Code" xml:"Code"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeletePtsScenesRequest creates a request to invoke DeletePtsScenes API
func CreateDeletePtsScenesRequest() (request *DeletePtsScenesRequest) {
	request = &DeletePtsScenesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "DeletePtsScenes", "", "")
	request.Method = requests.POST
	return
}

// CreateDeletePtsScenesResponse creates a response to parse from DeletePtsScenes response
func CreateDeletePtsScenesResponse() (response *DeletePtsScenesResponse) {
	response = &DeletePtsScenesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
