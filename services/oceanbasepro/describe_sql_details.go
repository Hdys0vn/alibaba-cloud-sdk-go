package oceanbasepro

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

// DescribeSQLDetails invokes the oceanbasepro.DescribeSQLDetails API synchronously
func (client *Client) DescribeSQLDetails(request *DescribeSQLDetailsRequest) (response *DescribeSQLDetailsResponse, err error) {
	response = CreateDescribeSQLDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLDetailsWithChan invokes the oceanbasepro.DescribeSQLDetails API asynchronously
func (client *Client) DescribeSQLDetailsWithChan(request *DescribeSQLDetailsRequest) (<-chan *DescribeSQLDetailsResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLDetails(request)
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

// DescribeSQLDetailsWithCallback invokes the oceanbasepro.DescribeSQLDetails API asynchronously
func (client *Client) DescribeSQLDetailsWithCallback(request *DescribeSQLDetailsRequest, callback func(response *DescribeSQLDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLDetailsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLDetails(request)
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

// DescribeSQLDetailsRequest is the request struct for api DescribeSQLDetails
type DescribeSQLDetailsRequest struct {
	*requests.RpcRequest
	SQLId    string `position:"Body" name:"SQLId"`
	TenantId string `position:"Body" name:"TenantId"`
}

// DescribeSQLDetailsResponse is the response struct for api DescribeSQLDetails
type DescribeSQLDetailsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	SQLDetails []Data `json:"SQLDetails" xml:"SQLDetails"`
}

// CreateDescribeSQLDetailsRequest creates a request to invoke DescribeSQLDetails API
func CreateDescribeSQLDetailsRequest() (request *DescribeSQLDetailsRequest) {
	request = &DescribeSQLDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("OceanBasePro", "2019-09-01", "DescribeSQLDetails", "oceanbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSQLDetailsResponse creates a response to parse from DescribeSQLDetails response
func CreateDescribeSQLDetailsResponse() (response *DescribeSQLDetailsResponse) {
	response = &DescribeSQLDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
