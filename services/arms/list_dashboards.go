package arms

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

// ListDashboards invokes the arms.ListDashboards API synchronously
func (client *Client) ListDashboards(request *ListDashboardsRequest) (response *ListDashboardsResponse, err error) {
	response = CreateListDashboardsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDashboardsWithChan invokes the arms.ListDashboards API asynchronously
func (client *Client) ListDashboardsWithChan(request *ListDashboardsRequest) (<-chan *ListDashboardsResponse, <-chan error) {
	responseChan := make(chan *ListDashboardsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDashboards(request)
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

// ListDashboardsWithCallback invokes the arms.ListDashboards API asynchronously
func (client *Client) ListDashboardsWithCallback(request *ListDashboardsRequest, callback func(response *ListDashboardsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDashboardsResponse
		var err error
		defer close(result)
		response, err = client.ListDashboards(request)
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

// ListDashboardsRequest is the request struct for api ListDashboards
type ListDashboardsRequest struct {
	*requests.RpcRequest
	DashboardName  string           `position:"Query" name:"DashboardName"`
	Product        string           `position:"Query" name:"Product"`
	RecreateSwitch requests.Boolean `position:"Query" name:"RecreateSwitch"`
	Language       string           `position:"Query" name:"Language"`
	ClusterId      string           `position:"Query" name:"ClusterId"`
	ProxyUserId    string           `position:"Query" name:"ProxyUserId"`
	Title          string           `position:"Query" name:"Title"`
	ClusterType    string           `position:"Query" name:"ClusterType"`
}

// ListDashboardsResponse is the response struct for api ListDashboards
type ListDashboardsResponse struct {
	*responses.BaseResponse
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	DashboardVos []DashboardVosItem `json:"DashboardVos" xml:"DashboardVos"`
}

// CreateListDashboardsRequest creates a request to invoke ListDashboards API
func CreateListDashboardsRequest() (request *ListDashboardsRequest) {
	request = &ListDashboardsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "ListDashboards", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDashboardsResponse creates a response to parse from ListDashboards response
func CreateListDashboardsResponse() (response *ListDashboardsResponse) {
	response = &ListDashboardsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
