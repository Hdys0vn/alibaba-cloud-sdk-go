package cloudcallcenter

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

// GetRoutePoint invokes the cloudcallcenter.GetRoutePoint API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getroutepoint.html
func (client *Client) GetRoutePoint(request *GetRoutePointRequest) (response *GetRoutePointResponse, err error) {
	response = CreateGetRoutePointResponse()
	err = client.DoAction(request, response)
	return
}

// GetRoutePointWithChan invokes the cloudcallcenter.GetRoutePoint API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getroutepoint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRoutePointWithChan(request *GetRoutePointRequest) (<-chan *GetRoutePointResponse, <-chan error) {
	responseChan := make(chan *GetRoutePointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetRoutePoint(request)
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

// GetRoutePointWithCallback invokes the cloudcallcenter.GetRoutePoint API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/getroutepoint.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetRoutePointWithCallback(request *GetRoutePointRequest, callback func(response *GetRoutePointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetRoutePointResponse
		var err error
		defer close(result)
		response, err = client.GetRoutePoint(request)
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

// GetRoutePointRequest is the request struct for api GetRoutePoint
type GetRoutePointRequest struct {
	*requests.RpcRequest
	ContactFlowId string `position:"Query" name:"ContactFlowId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
}

// GetRoutePointResponse is the response struct for api GetRoutePoint
type GetRoutePointResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	StatusCode     string `json:"StatusCode" xml:"StatusCode"`
	StatusDesc     string `json:"StatusDesc" xml:"StatusDesc"`
	RoutePoint     string `json:"RoutePoint" xml:"RoutePoint"`
}

// CreateGetRoutePointRequest creates a request to invoke GetRoutePoint API
func CreateGetRoutePointRequest() (request *GetRoutePointRequest) {
	request = &GetRoutePointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "GetRoutePoint", "", "")
	request.Method = requests.POST
	return
}

// CreateGetRoutePointResponse creates a response to parse from GetRoutePoint response
func CreateGetRoutePointResponse() (response *GetRoutePointResponse) {
	response = &GetRoutePointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
