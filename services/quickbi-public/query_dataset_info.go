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

// QueryDatasetInfo invokes the quickbi_public.QueryDatasetInfo API synchronously
func (client *Client) QueryDatasetInfo(request *QueryDatasetInfoRequest) (response *QueryDatasetInfoResponse, err error) {
	response = CreateQueryDatasetInfoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDatasetInfoWithChan invokes the quickbi_public.QueryDatasetInfo API asynchronously
func (client *Client) QueryDatasetInfoWithChan(request *QueryDatasetInfoRequest) (<-chan *QueryDatasetInfoResponse, <-chan error) {
	responseChan := make(chan *QueryDatasetInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDatasetInfo(request)
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

// QueryDatasetInfoWithCallback invokes the quickbi_public.QueryDatasetInfo API asynchronously
func (client *Client) QueryDatasetInfoWithCallback(request *QueryDatasetInfoRequest, callback func(response *QueryDatasetInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDatasetInfoResponse
		var err error
		defer close(result)
		response, err = client.QueryDatasetInfo(request)
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

// QueryDatasetInfoRequest is the request struct for api QueryDatasetInfo
type QueryDatasetInfoRequest struct {
	*requests.RpcRequest
	AccessPoint string `position:"Query" name:"AccessPoint"`
	DatasetId   string `position:"Query" name:"DatasetId"`
	SignType    string `position:"Query" name:"SignType"`
}

// QueryDatasetInfoResponse is the response struct for api QueryDatasetInfo
type QueryDatasetInfoResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryDatasetInfoRequest creates a request to invoke QueryDatasetInfo API
func CreateQueryDatasetInfoRequest() (request *QueryDatasetInfoRequest) {
	request = &QueryDatasetInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryDatasetInfo", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDatasetInfoResponse creates a response to parse from QueryDatasetInfo response
func CreateQueryDatasetInfoResponse() (response *QueryDatasetInfoResponse) {
	response = &QueryDatasetInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
