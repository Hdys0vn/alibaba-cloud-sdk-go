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

// GetFootwearPosition invokes the reid.GetFootwearPosition API synchronously
// api document: https://help.aliyun.com/api/reid/getfootwearposition.html
func (client *Client) GetFootwearPosition(request *GetFootwearPositionRequest) (response *GetFootwearPositionResponse, err error) {
	response = CreateGetFootwearPositionResponse()
	err = client.DoAction(request, response)
	return
}

// GetFootwearPositionWithChan invokes the reid.GetFootwearPosition API asynchronously
// api document: https://help.aliyun.com/api/reid/getfootwearposition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFootwearPositionWithChan(request *GetFootwearPositionRequest) (<-chan *GetFootwearPositionResponse, <-chan error) {
	responseChan := make(chan *GetFootwearPositionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFootwearPosition(request)
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

// GetFootwearPositionWithCallback invokes the reid.GetFootwearPosition API asynchronously
// api document: https://help.aliyun.com/api/reid/getfootwearposition.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFootwearPositionWithCallback(request *GetFootwearPositionRequest, callback func(response *GetFootwearPositionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFootwearPositionResponse
		var err error
		defer close(result)
		response, err = client.GetFootwearPosition(request)
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

// GetFootwearPositionRequest is the request struct for api GetFootwearPosition
type GetFootwearPositionRequest struct {
	*requests.RpcRequest
	Date    string           `position:"Body" name:"Date"`
	StoreId requests.Integer `position:"Body" name:"StoreId"`
	SkuId   string           `position:"Body" name:"SkuId"`
}

// GetFootwearPositionResponse is the response struct for api GetFootwearPosition
type GetFootwearPositionResponse struct {
	*responses.BaseResponse
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Message        string `json:"Message" xml:"Message"`
	StartTime      int64  `json:"StartTime" xml:"StartTime"`
	Code           string `json:"Code" xml:"Code"`
	PositionNumber int    `json:"PositionNumber" xml:"PositionNumber"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	SkuId          string `json:"SkuId" xml:"SkuId"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	StoreId        int64  `json:"StoreId" xml:"StoreId"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
}

// CreateGetFootwearPositionRequest creates a request to invoke GetFootwearPosition API
func CreateGetFootwearPositionRequest() (request *GetFootwearPositionRequest) {
	request = &GetFootwearPositionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("reid", "2019-09-28", "GetFootwearPosition", "1.1.8.2", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetFootwearPositionResponse creates a response to parse from GetFootwearPosition response
func CreateGetFootwearPositionResponse() (response *GetFootwearPositionResponse) {
	response = &GetFootwearPositionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
