package dataworks_public

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

// ListDataSources invokes the dataworks_public.ListDataSources API synchronously
func (client *Client) ListDataSources(request *ListDataSourcesRequest) (response *ListDataSourcesResponse, err error) {
	response = CreateListDataSourcesResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataSourcesWithChan invokes the dataworks_public.ListDataSources API asynchronously
func (client *Client) ListDataSourcesWithChan(request *ListDataSourcesRequest) (<-chan *ListDataSourcesResponse, <-chan error) {
	responseChan := make(chan *ListDataSourcesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataSources(request)
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

// ListDataSourcesWithCallback invokes the dataworks_public.ListDataSources API asynchronously
func (client *Client) ListDataSourcesWithCallback(request *ListDataSourcesRequest, callback func(response *ListDataSourcesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataSourcesResponse
		var err error
		defer close(result)
		response, err = client.ListDataSources(request)
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

// ListDataSourcesRequest is the request struct for api ListDataSources
type ListDataSourcesRequest struct {
	*requests.RpcRequest
	DataSourceType string           `position:"Query" name:"DataSourceType"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	SubType        string           `position:"Query" name:"SubType"`
	Name           string           `position:"Query" name:"Name"`
	EnvType        requests.Integer `position:"Query" name:"EnvType"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	ProjectId      requests.Integer `position:"Query" name:"ProjectId"`
	Status         string           `position:"Query" name:"Status"`
}

// ListDataSourcesResponse is the response struct for api ListDataSources
type ListDataSourcesResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateListDataSourcesRequest creates a request to invoke ListDataSources API
func CreateListDataSourcesRequest() (request *ListDataSourcesRequest) {
	request = &ListDataSourcesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListDataSources", "", "")
	request.Method = requests.GET
	return
}

// CreateListDataSourcesResponse creates a response to parse from ListDataSources response
func CreateListDataSourcesResponse() (response *ListDataSourcesResponse) {
	response = &ListDataSourcesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
