package reid

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

// DescribeHeatMap invokes the reid.DescribeHeatMap API synchronously
// api document: https://help.aliyun.com/api/reid/describeheatmap.html
func (client *Client) DescribeHeatMap(request *DescribeHeatMapRequest) (response *DescribeHeatMapResponse, err error) {
	response = CreateDescribeHeatMapResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeHeatMapWithChan invokes the reid.DescribeHeatMap API asynchronously
// api document: https://help.aliyun.com/api/reid/describeheatmap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHeatMapWithChan(request *DescribeHeatMapRequest) (<-chan *DescribeHeatMapResponse, <-chan error) {
	responseChan := make(chan *DescribeHeatMapResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHeatMap(request)
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

// DescribeHeatMapWithCallback invokes the reid.DescribeHeatMap API asynchronously
// api document: https://help.aliyun.com/api/reid/describeheatmap.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHeatMapWithCallback(request *DescribeHeatMapRequest, callback func(response *DescribeHeatMapResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHeatMapResponse
		var err error
		defer close(result)
		response, err = client.DescribeHeatMap(request)
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

// DescribeHeatMapRequest is the request struct for api DescribeHeatMap
type DescribeHeatMapRequest struct {
	*requests.RpcRequest
	Date    string           `position:"Body" name:"Date"`
	StoreId requests.Integer `position:"Body" name:"StoreId"`
	EmapId  requests.Integer `position:"Body" name:"EmapId"`
}

// DescribeHeatMapResponse is the response struct for api DescribeHeatMap
type DescribeHeatMapResponse struct {
	*responses.BaseResponse
	ErrorCode     string        `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage  string        `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId     string        `json:"RequestId" xml:"RequestId"`
	Success       bool          `json:"Success" xml:"Success"`
	HeatMapPoints HeatMapPoints `json:"HeatMapPoints" xml:"HeatMapPoints"`
}

// CreateDescribeHeatMapRequest creates a request to invoke DescribeHeatMap API
func CreateDescribeHeatMapRequest() (request *DescribeHeatMapRequest) {
	request = &DescribeHeatMapRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "DescribeHeatMap", "1.1.8.2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeHeatMapResponse creates a response to parse from DescribeHeatMap response
func CreateDescribeHeatMapResponse() (response *DescribeHeatMapResponse) {
	response = &DescribeHeatMapResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
