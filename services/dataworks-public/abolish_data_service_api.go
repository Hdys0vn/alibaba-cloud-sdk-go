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

// AbolishDataServiceApi invokes the dataworks_public.AbolishDataServiceApi API synchronously
func (client *Client) AbolishDataServiceApi(request *AbolishDataServiceApiRequest) (response *AbolishDataServiceApiResponse, err error) {
	response = CreateAbolishDataServiceApiResponse()
	err = client.DoAction(request, response)
	return
}

// AbolishDataServiceApiWithChan invokes the dataworks_public.AbolishDataServiceApi API asynchronously
func (client *Client) AbolishDataServiceApiWithChan(request *AbolishDataServiceApiRequest) (<-chan *AbolishDataServiceApiResponse, <-chan error) {
	responseChan := make(chan *AbolishDataServiceApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AbolishDataServiceApi(request)
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

// AbolishDataServiceApiWithCallback invokes the dataworks_public.AbolishDataServiceApi API asynchronously
func (client *Client) AbolishDataServiceApiWithCallback(request *AbolishDataServiceApiRequest, callback func(response *AbolishDataServiceApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AbolishDataServiceApiResponse
		var err error
		defer close(result)
		response, err = client.AbolishDataServiceApi(request)
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

// AbolishDataServiceApiRequest is the request struct for api AbolishDataServiceApi
type AbolishDataServiceApiRequest struct {
	*requests.RpcRequest
	TenantId  requests.Integer `position:"Body" name:"TenantId"`
	ProjectId requests.Integer `position:"Body" name:"ProjectId"`
	ApiId     requests.Integer `position:"Body" name:"ApiId"`
}

// AbolishDataServiceApiResponse is the response struct for api AbolishDataServiceApi
type AbolishDataServiceApiResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           bool   `json:"Data" xml:"Data"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateAbolishDataServiceApiRequest creates a request to invoke AbolishDataServiceApi API
func CreateAbolishDataServiceApiRequest() (request *AbolishDataServiceApiRequest) {
	request = &AbolishDataServiceApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "AbolishDataServiceApi", "", "")
	request.Method = requests.POST
	return
}

// CreateAbolishDataServiceApiResponse creates a response to parse from AbolishDataServiceApi response
func CreateAbolishDataServiceApiResponse() (response *AbolishDataServiceApiResponse) {
	response = &AbolishDataServiceApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
