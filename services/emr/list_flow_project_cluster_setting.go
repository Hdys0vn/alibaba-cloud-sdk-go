package emr

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

// ListFlowProjectClusterSetting invokes the emr.ListFlowProjectClusterSetting API synchronously
func (client *Client) ListFlowProjectClusterSetting(request *ListFlowProjectClusterSettingRequest) (response *ListFlowProjectClusterSettingResponse, err error) {
	response = CreateListFlowProjectClusterSettingResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowProjectClusterSettingWithChan invokes the emr.ListFlowProjectClusterSetting API asynchronously
func (client *Client) ListFlowProjectClusterSettingWithChan(request *ListFlowProjectClusterSettingRequest) (<-chan *ListFlowProjectClusterSettingResponse, <-chan error) {
	responseChan := make(chan *ListFlowProjectClusterSettingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlowProjectClusterSetting(request)
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

// ListFlowProjectClusterSettingWithCallback invokes the emr.ListFlowProjectClusterSetting API asynchronously
func (client *Client) ListFlowProjectClusterSettingWithCallback(request *ListFlowProjectClusterSettingRequest, callback func(response *ListFlowProjectClusterSettingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowProjectClusterSettingResponse
		var err error
		defer close(result)
		response, err = client.ListFlowProjectClusterSetting(request)
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

// ListFlowProjectClusterSettingRequest is the request struct for api ListFlowProjectClusterSetting
type ListFlowProjectClusterSettingRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	ProjectId  string           `position:"Query" name:"ProjectId"`
}

// ListFlowProjectClusterSettingResponse is the response struct for api ListFlowProjectClusterSetting
type ListFlowProjectClusterSettingResponse struct {
	*responses.BaseResponse
	RequestId       string          `json:"RequestId" xml:"RequestId"`
	PageNumber      int             `json:"PageNumber" xml:"PageNumber"`
	PageSize        int             `json:"PageSize" xml:"PageSize"`
	Total           int             `json:"Total" xml:"Total"`
	ClusterSettings ClusterSettings `json:"ClusterSettings" xml:"ClusterSettings"`
}

// CreateListFlowProjectClusterSettingRequest creates a request to invoke ListFlowProjectClusterSetting API
func CreateListFlowProjectClusterSettingRequest() (request *ListFlowProjectClusterSettingRequest) {
	request = &ListFlowProjectClusterSettingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListFlowProjectClusterSetting", "emr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListFlowProjectClusterSettingResponse creates a response to parse from ListFlowProjectClusterSetting response
func CreateListFlowProjectClusterSettingResponse() (response *ListFlowProjectClusterSettingResponse) {
	response = &ListFlowProjectClusterSettingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
