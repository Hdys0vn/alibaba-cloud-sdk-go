package quickbi_public

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

// QueryWorksByWorkspace invokes the quickbi_public.QueryWorksByWorkspace API synchronously
func (client *Client) QueryWorksByWorkspace(request *QueryWorksByWorkspaceRequest) (response *QueryWorksByWorkspaceResponse, err error) {
	response = CreateQueryWorksByWorkspaceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryWorksByWorkspaceWithChan invokes the quickbi_public.QueryWorksByWorkspace API asynchronously
func (client *Client) QueryWorksByWorkspaceWithChan(request *QueryWorksByWorkspaceRequest) (<-chan *QueryWorksByWorkspaceResponse, <-chan error) {
	responseChan := make(chan *QueryWorksByWorkspaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryWorksByWorkspace(request)
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

// QueryWorksByWorkspaceWithCallback invokes the quickbi_public.QueryWorksByWorkspace API asynchronously
func (client *Client) QueryWorksByWorkspaceWithCallback(request *QueryWorksByWorkspaceRequest, callback func(response *QueryWorksByWorkspaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryWorksByWorkspaceResponse
		var err error
		defer close(result)
		response, err = client.QueryWorksByWorkspace(request)
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

// QueryWorksByWorkspaceRequest is the request struct for api QueryWorksByWorkspace
type QueryWorksByWorkspaceRequest struct {
	*requests.RpcRequest
	WorksType         string           `position:"Query" name:"WorksType"`
	ThirdPartAuthFlag requests.Integer `position:"Query" name:"ThirdPartAuthFlag"`
	AccessPoint       string           `position:"Query" name:"AccessPoint"`
	SignType          string           `position:"Query" name:"SignType"`
	PageNum           requests.Integer `position:"Query" name:"PageNum"`
	PageSize          requests.Integer `position:"Query" name:"PageSize"`
	WorkspaceId       string           `position:"Query" name:"WorkspaceId"`
	Status            requests.Integer `position:"Query" name:"Status"`
}

// QueryWorksByWorkspaceResponse is the response struct for api QueryWorksByWorkspace
type QueryWorksByWorkspaceResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryWorksByWorkspaceRequest creates a request to invoke QueryWorksByWorkspace API
func CreateQueryWorksByWorkspaceRequest() (request *QueryWorksByWorkspaceRequest) {
	request = &QueryWorksByWorkspaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryWorksByWorkspace", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryWorksByWorkspaceResponse creates a response to parse from QueryWorksByWorkspace response
func CreateQueryWorksByWorkspaceResponse() (response *QueryWorksByWorkspaceResponse) {
	response = &QueryWorksByWorkspaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}