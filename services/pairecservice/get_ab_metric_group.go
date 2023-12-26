package pairecservice

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

// GetABMetricGroup invokes the pairecservice.GetABMetricGroup API synchronously
func (client *Client) GetABMetricGroup(request *GetABMetricGroupRequest) (response *GetABMetricGroupResponse, err error) {
	response = CreateGetABMetricGroupResponse()
	err = client.DoAction(request, response)
	return
}

// GetABMetricGroupWithChan invokes the pairecservice.GetABMetricGroup API asynchronously
func (client *Client) GetABMetricGroupWithChan(request *GetABMetricGroupRequest) (<-chan *GetABMetricGroupResponse, <-chan error) {
	responseChan := make(chan *GetABMetricGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetABMetricGroup(request)
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

// GetABMetricGroupWithCallback invokes the pairecservice.GetABMetricGroup API asynchronously
func (client *Client) GetABMetricGroupWithCallback(request *GetABMetricGroupRequest, callback func(response *GetABMetricGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetABMetricGroupResponse
		var err error
		defer close(result)
		response, err = client.GetABMetricGroup(request)
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

// GetABMetricGroupRequest is the request struct for api GetABMetricGroup
type GetABMetricGroupRequest struct {
	*requests.RoaRequest
	InstanceId      string `position:"Query" name:"InstanceId"`
	ABMetricGroupId string `position:"Path" name:"ABMetricGroupId"`
}

// GetABMetricGroupResponse is the response struct for api GetABMetricGroup
type GetABMetricGroupResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Name          string `json:"Name" xml:"Name"`
	SceneId       string `json:"SceneId" xml:"SceneId"`
	Description   string `json:"Description" xml:"Description"`
	Owner         string `json:"Owner" xml:"Owner"`
	ABMetricIds   string `json:"ABMetricIds" xml:"ABMetricIds"`
	ABMetricNames string `json:"ABMetricNames" xml:"ABMetricNames"`
	Realtime      bool   `json:"Realtime" xml:"Realtime"`
}

// CreateGetABMetricGroupRequest creates a request to invoke GetABMetricGroup API
func CreateGetABMetricGroupRequest() (request *GetABMetricGroupRequest) {
	request = &GetABMetricGroupRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "GetABMetricGroup", "/api/v1/abmetricgroups/[ABMetricGroupId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetABMetricGroupResponse creates a response to parse from GetABMetricGroup response
func CreateGetABMetricGroupResponse() (response *GetABMetricGroupResponse) {
	response = &GetABMetricGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}