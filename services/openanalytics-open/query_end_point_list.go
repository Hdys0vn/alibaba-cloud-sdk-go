package openanalytics_open

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

// QueryEndPointList invokes the openanalytics_open.QueryEndPointList API synchronously
func (client *Client) QueryEndPointList(request *QueryEndPointListRequest) (response *QueryEndPointListResponse, err error) {
	response = CreateQueryEndPointListResponse()
	err = client.DoAction(request, response)
	return
}

// QueryEndPointListWithChan invokes the openanalytics_open.QueryEndPointList API asynchronously
func (client *Client) QueryEndPointListWithChan(request *QueryEndPointListRequest) (<-chan *QueryEndPointListResponse, <-chan error) {
	responseChan := make(chan *QueryEndPointListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryEndPointList(request)
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

// QueryEndPointListWithCallback invokes the openanalytics_open.QueryEndPointList API asynchronously
func (client *Client) QueryEndPointListWithCallback(request *QueryEndPointListRequest, callback func(response *QueryEndPointListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryEndPointListResponse
		var err error
		defer close(result)
		response, err = client.QueryEndPointList(request)
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

// QueryEndPointListRequest is the request struct for api QueryEndPointList
type QueryEndPointListRequest struct {
	*requests.RpcRequest
}

// QueryEndPointListResponse is the response struct for api QueryEndPointList
type QueryEndPointListResponse struct {
	*responses.BaseResponse
	RequestId    string             `json:"RequestId" xml:"RequestId"`
	RegionId     string             `json:"RegionId" xml:"RegionId"`
	EndPointList []EndPointListItem `json:"EndPointList" xml:"EndPointList"`
}

// CreateQueryEndPointListRequest creates a request to invoke QueryEndPointList API
func CreateQueryEndPointListRequest() (request *QueryEndPointListRequest) {
	request = &QueryEndPointListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "QueryEndPointList", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryEndPointListResponse creates a response to parse from QueryEndPointList response
func CreateQueryEndPointListResponse() (response *QueryEndPointListResponse) {
	response = &QueryEndPointListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
