package live

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

// DescribeStudioLayouts invokes the live.DescribeStudioLayouts API synchronously
func (client *Client) DescribeStudioLayouts(request *DescribeStudioLayoutsRequest) (response *DescribeStudioLayoutsResponse, err error) {
	response = CreateDescribeStudioLayoutsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeStudioLayoutsWithChan invokes the live.DescribeStudioLayouts API asynchronously
func (client *Client) DescribeStudioLayoutsWithChan(request *DescribeStudioLayoutsRequest) (<-chan *DescribeStudioLayoutsResponse, <-chan error) {
	responseChan := make(chan *DescribeStudioLayoutsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeStudioLayouts(request)
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

// DescribeStudioLayoutsWithCallback invokes the live.DescribeStudioLayouts API asynchronously
func (client *Client) DescribeStudioLayoutsWithCallback(request *DescribeStudioLayoutsRequest, callback func(response *DescribeStudioLayoutsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeStudioLayoutsResponse
		var err error
		defer close(result)
		response, err = client.DescribeStudioLayouts(request)
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

// DescribeStudioLayoutsRequest is the request struct for api DescribeStudioLayouts
type DescribeStudioLayoutsRequest struct {
	*requests.RpcRequest
	LayoutId string           `position:"Query" name:"LayoutId"`
	CasterId string           `position:"Query" name:"CasterId"`
	OwnerId  requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeStudioLayoutsResponse is the response struct for api DescribeStudioLayouts
type DescribeStudioLayoutsResponse struct {
	*responses.BaseResponse
	Total         int            `json:"Total" xml:"Total"`
	RequestId     string         `json:"RequestId" xml:"RequestId"`
	StudioLayouts []StudioLayout `json:"StudioLayouts" xml:"StudioLayouts"`
}

// CreateDescribeStudioLayoutsRequest creates a request to invoke DescribeStudioLayouts API
func CreateDescribeStudioLayoutsRequest() (request *DescribeStudioLayoutsRequest) {
	request = &DescribeStudioLayoutsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "DescribeStudioLayouts", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeStudioLayoutsResponse creates a response to parse from DescribeStudioLayouts response
func CreateDescribeStudioLayoutsResponse() (response *DescribeStudioLayoutsResponse) {
	response = &DescribeStudioLayoutsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
