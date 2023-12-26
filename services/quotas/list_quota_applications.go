package quotas

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

// ListQuotaApplications invokes the quotas.ListQuotaApplications API synchronously
func (client *Client) ListQuotaApplications(request *ListQuotaApplicationsRequest) (response *ListQuotaApplicationsResponse, err error) {
	response = CreateListQuotaApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListQuotaApplicationsWithChan invokes the quotas.ListQuotaApplications API asynchronously
func (client *Client) ListQuotaApplicationsWithChan(request *ListQuotaApplicationsRequest) (<-chan *ListQuotaApplicationsResponse, <-chan error) {
	responseChan := make(chan *ListQuotaApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListQuotaApplications(request)
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

// ListQuotaApplicationsWithCallback invokes the quotas.ListQuotaApplications API asynchronously
func (client *Client) ListQuotaApplicationsWithCallback(request *ListQuotaApplicationsRequest, callback func(response *ListQuotaApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListQuotaApplicationsResponse
		var err error
		defer close(result)
		response, err = client.ListQuotaApplications(request)
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

// ListQuotaApplicationsRequest is the request struct for api ListQuotaApplications
type ListQuotaApplicationsRequest struct {
	*requests.RpcRequest
	ProductCode     string                             `position:"Body" name:"ProductCode"`
	QuotaActionCode string                             `position:"Body" name:"QuotaActionCode"`
	NextToken       string                             `position:"Body" name:"NextToken"`
	KeyWord         string                             `position:"Body" name:"KeyWord"`
	QuotaCategory   string                             `position:"Body" name:"QuotaCategory"`
	OriginalContext string                             `position:"Body" name:"OriginalContext"`
	MaxResults      requests.Integer                   `position:"Body" name:"MaxResults"`
	Status          string                             `position:"Body" name:"Status"`
	Dimensions      *[]ListQuotaApplicationsDimensions `position:"Body" name:"Dimensions"  type:"Repeated"`
}

// ListQuotaApplicationsDimensions is a repeated param struct in ListQuotaApplicationsRequest
type ListQuotaApplicationsDimensions struct {
	Key   string `name:"Key"`
	Value string `name:"Value"`
}

// ListQuotaApplicationsResponse is the response struct for api ListQuotaApplications
type ListQuotaApplicationsResponse struct {
	*responses.BaseResponse
	TotalCount        int                     `json:"TotalCount" xml:"TotalCount"`
	NextToken         string                  `json:"NextToken" xml:"NextToken"`
	RequestId         string                  `json:"RequestId" xml:"RequestId"`
	MaxResults        int                     `json:"MaxResults" xml:"MaxResults"`
	QuotaApplications []QuotaApplicationsItem `json:"QuotaApplications" xml:"QuotaApplications"`
}

// CreateListQuotaApplicationsRequest creates a request to invoke ListQuotaApplications API
func CreateListQuotaApplicationsRequest() (request *ListQuotaApplicationsRequest) {
	request = &ListQuotaApplicationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("quotas", "2020-05-10", "ListQuotaApplications", "quotas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListQuotaApplicationsResponse creates a response to parse from ListQuotaApplications response
func CreateListQuotaApplicationsResponse() (response *ListQuotaApplicationsResponse) {
	response = &ListQuotaApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
