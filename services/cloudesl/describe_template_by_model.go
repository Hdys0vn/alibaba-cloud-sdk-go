package cloudesl

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

// DescribeTemplateByModel invokes the cloudesl.DescribeTemplateByModel API synchronously
func (client *Client) DescribeTemplateByModel(request *DescribeTemplateByModelRequest) (response *DescribeTemplateByModelResponse, err error) {
	response = CreateDescribeTemplateByModelResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTemplateByModelWithChan invokes the cloudesl.DescribeTemplateByModel API asynchronously
func (client *Client) DescribeTemplateByModelWithChan(request *DescribeTemplateByModelRequest) (<-chan *DescribeTemplateByModelResponse, <-chan error) {
	responseChan := make(chan *DescribeTemplateByModelResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTemplateByModel(request)
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

// DescribeTemplateByModelWithCallback invokes the cloudesl.DescribeTemplateByModel API asynchronously
func (client *Client) DescribeTemplateByModelWithCallback(request *DescribeTemplateByModelRequest, callback func(response *DescribeTemplateByModelResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTemplateByModelResponse
		var err error
		defer close(result)
		response, err = client.DescribeTemplateByModel(request)
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

// DescribeTemplateByModelRequest is the request struct for api DescribeTemplateByModel
type DescribeTemplateByModelRequest struct {
	*requests.RpcRequest
	EslSize         string           `position:"Body" name:"EslSize"`
	DeviceType      string           `position:"Body" name:"DeviceType"`
	PageNumber      requests.Integer `position:"Body" name:"PageNumber"`
	TemplateVersion string           `position:"Body" name:"TemplateVersion"`
	PageSize        requests.Integer `position:"Body" name:"PageSize"`
}

// DescribeTemplateByModelResponse is the response struct for api DescribeTemplateByModel
type DescribeTemplateByModelResponse struct {
	*responses.BaseResponse
	RequestId      string           `json:"RequestId" xml:"RequestId"`
	ErrorMessage   string           `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool             `json:"Success" xml:"Success"`
	ErrorCode      string           `json:"ErrorCode" xml:"ErrorCode"`
	Code           string           `json:"Code" xml:"Code"`
	Message        string           `json:"Message" xml:"Message"`
	DynamicMessage string           `json:"DynamicMessage" xml:"DynamicMessage"`
	DynamicCode    string           `json:"DynamicCode" xml:"DynamicCode"`
	TotalCount     int              `json:"TotalCount" xml:"TotalCount"`
	PageSize       int              `json:"PageSize" xml:"PageSize"`
	PageNumber     int              `json:"PageNumber" xml:"PageNumber"`
	Items          []SelectItemInfo `json:"Items" xml:"Items"`
}

// CreateDescribeTemplateByModelRequest creates a request to invoke DescribeTemplateByModel API
func CreateDescribeTemplateByModelRequest() (request *DescribeTemplateByModelRequest) {
	request = &DescribeTemplateByModelRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DescribeTemplateByModel", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeTemplateByModelResponse creates a response to parse from DescribeTemplateByModel response
func CreateDescribeTemplateByModelResponse() (response *DescribeTemplateByModelResponse) {
	response = &DescribeTemplateByModelResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
