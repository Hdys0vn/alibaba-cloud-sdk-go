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

// QueryDataService invokes the quickbi_public.QueryDataService API synchronously
func (client *Client) QueryDataService(request *QueryDataServiceRequest) (response *QueryDataServiceResponse, err error) {
	response = CreateQueryDataServiceResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDataServiceWithChan invokes the quickbi_public.QueryDataService API asynchronously
func (client *Client) QueryDataServiceWithChan(request *QueryDataServiceRequest) (<-chan *QueryDataServiceResponse, <-chan error) {
	responseChan := make(chan *QueryDataServiceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDataService(request)
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

// QueryDataServiceWithCallback invokes the quickbi_public.QueryDataService API asynchronously
func (client *Client) QueryDataServiceWithCallback(request *QueryDataServiceRequest, callback func(response *QueryDataServiceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDataServiceResponse
		var err error
		defer close(result)
		response, err = client.QueryDataService(request)
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

// QueryDataServiceRequest is the request struct for api QueryDataService
type QueryDataServiceRequest struct {
	*requests.RpcRequest
	ReturnFields string `position:"Query" name:"ReturnFields"`
	AccessPoint  string `position:"Query" name:"AccessPoint"`
	SignType     string `position:"Query" name:"SignType"`
	Conditions   string `position:"Query" name:"Conditions"`
	ApiId        string `position:"Query" name:"ApiId"`
}

// QueryDataServiceResponse is the response struct for api QueryDataService
type QueryDataServiceResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateQueryDataServiceRequest creates a request to invoke QueryDataService API
func CreateQueryDataServiceRequest() (request *QueryDataServiceRequest) {
	request = &QueryDataServiceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quickbi-public", "2022-01-01", "QueryDataService", "2.2.0", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDataServiceResponse creates a response to parse from QueryDataService response
func CreateQueryDataServiceResponse() (response *QueryDataServiceResponse) {
	response = &QueryDataServiceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
