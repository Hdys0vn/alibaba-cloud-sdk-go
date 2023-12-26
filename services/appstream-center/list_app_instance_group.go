package appstream_center

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

// ListAppInstanceGroup invokes the appstream_center.ListAppInstanceGroup API synchronously
func (client *Client) ListAppInstanceGroup(request *ListAppInstanceGroupRequest) (response *ListAppInstanceGroupResponse, err error) {
	response = CreateListAppInstanceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// ListAppInstanceGroupWithChan invokes the appstream_center.ListAppInstanceGroup API asynchronously
func (client *Client) ListAppInstanceGroupWithChan(request *ListAppInstanceGroupRequest) (<-chan *ListAppInstanceGroupResponse, <-chan error) {
	responseChan := make(chan *ListAppInstanceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListAppInstanceGroup(request)
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

// ListAppInstanceGroupWithCallback invokes the appstream_center.ListAppInstanceGroup API asynchronously
func (client *Client) ListAppInstanceGroupWithCallback(request *ListAppInstanceGroupRequest, callback func(response *ListAppInstanceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListAppInstanceGroupResponse
		var err error
		defer close(result)
		response, err = client.ListAppInstanceGroup(request)
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

// ListAppInstanceGroupRequest is the request struct for api ListAppInstanceGroup
type ListAppInstanceGroupRequest struct {
	*requests.RpcRequest
	BizRegionId          string           `position:"Query" name:"BizRegionId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ProductType          string           `position:"Query" name:"ProductType"`
	AppCenterImageId     string           `position:"Query" name:"AppCenterImageId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	NodeInstanceType     string           `position:"Query" name:"NodeInstanceType"`
	AppInstanceGroupName string           `position:"Query" name:"AppInstanceGroupName"`
	AppInstanceGroupId   string           `position:"Query" name:"AppInstanceGroupId"`
	Status               *[]string        `position:"Body" name:"Status"  type:"Repeated"`
}

// ListAppInstanceGroupResponse is the response struct for api ListAppInstanceGroup
type ListAppInstanceGroupResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	TotalCount             int    `json:"TotalCount" xml:"TotalCount"`
	PageSize               int    `json:"PageSize" xml:"PageSize"`
	PageNumber             int    `json:"PageNumber" xml:"PageNumber"`
	AppInstanceGroupModels []Data `json:"AppInstanceGroupModels" xml:"AppInstanceGroupModels"`
}

// CreateListAppInstanceGroupRequest creates a request to invoke ListAppInstanceGroup API
func CreateListAppInstanceGroupRequest() (request *ListAppInstanceGroupRequest) {
	request = &ListAppInstanceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("appstream-center", "2021-09-01", "ListAppInstanceGroup", "", "")
	request.Method = requests.POST
	return
}

// CreateListAppInstanceGroupResponse creates a response to parse from ListAppInstanceGroup response
func CreateListAppInstanceGroupResponse() (response *ListAppInstanceGroupResponse) {
	response = &ListAppInstanceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
