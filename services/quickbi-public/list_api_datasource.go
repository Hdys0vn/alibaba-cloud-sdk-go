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

// ListApiDatasource invokes the quickbi_public.ListApiDatasource API synchronously
func (client *Client) ListApiDatasource(request *ListApiDatasourceRequest) (response *ListApiDatasourceResponse, err error) {
	response = CreateListApiDatasourceResponse()
	err = client.DoAction(request, response)
	return
}

// ListApiDatasourceWithChan invokes the quickbi_public.ListApiDatasource API asynchronously
func (client *Client) ListApiDatasourceWithChan(request *ListApiDatasourceRequest) (<-chan *ListApiDatasourceResponse, <-chan error) {
	responseChan := make(chan *ListApiDatasourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListApiDatasource(request)
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

// ListApiDatasourceWithCallback invokes the quickbi_public.ListApiDatasource API asynchronously
func (client *Client) ListApiDatasourceWithCallback(request *ListApiDatasourceRequest, callback func(response *ListApiDatasourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListApiDatasourceResponse
		var err error
		defer close(result)
		response, err = client.ListApiDatasource(request)
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

// ListApiDatasourceRequest is the request struct for api ListApiDatasource
type ListApiDatasourceRequest struct {
	*requests.RpcRequest
	AccessPoint string           `position:"Query" name:"AccessPoint"`
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	SignType    string           `position:"Query" name:"SignType"`
	PageNum     requests.Integer `position:"Query" name:"PageNum"`
	KeyWord     string           `position:"Query" name:"KeyWord"`
	WorkspaceId string           `position:"Query" name:"WorkspaceId"`
}

// ListApiDatasourceResponse is the response struct for api ListApiDatasource
type ListApiDatasourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateListApiDatasourceRequest creates a request to invoke ListApiDatasource API
func CreateListApiDatasourceRequest() (request *ListApiDatasourceRequest) {
	request = &ListApiDatasourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "ListApiDatasource", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListApiDatasourceResponse creates a response to parse from ListApiDatasource response
func CreateListApiDatasourceResponse() (response *ListApiDatasourceResponse) {
	response = &ListApiDatasourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
