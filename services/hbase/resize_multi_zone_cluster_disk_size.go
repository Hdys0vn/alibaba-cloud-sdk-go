package hbase

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

// ResizeMultiZoneClusterDiskSize invokes the hbase.ResizeMultiZoneClusterDiskSize API synchronously
func (client *Client) ResizeMultiZoneClusterDiskSize(request *ResizeMultiZoneClusterDiskSizeRequest) (response *ResizeMultiZoneClusterDiskSizeResponse, err error) {
	response = CreateResizeMultiZoneClusterDiskSizeResponse()
	err = client.DoAction(request, response)
	return
}

// ResizeMultiZoneClusterDiskSizeWithChan invokes the hbase.ResizeMultiZoneClusterDiskSize API asynchronously
func (client *Client) ResizeMultiZoneClusterDiskSizeWithChan(request *ResizeMultiZoneClusterDiskSizeRequest) (<-chan *ResizeMultiZoneClusterDiskSizeResponse, <-chan error) {
	responseChan := make(chan *ResizeMultiZoneClusterDiskSizeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResizeMultiZoneClusterDiskSize(request)
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

// ResizeMultiZoneClusterDiskSizeWithCallback invokes the hbase.ResizeMultiZoneClusterDiskSize API asynchronously
func (client *Client) ResizeMultiZoneClusterDiskSizeWithCallback(request *ResizeMultiZoneClusterDiskSizeRequest, callback func(response *ResizeMultiZoneClusterDiskSizeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResizeMultiZoneClusterDiskSizeResponse
		var err error
		defer close(result)
		response, err = client.ResizeMultiZoneClusterDiskSize(request)
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

// ResizeMultiZoneClusterDiskSizeRequest is the request struct for api ResizeMultiZoneClusterDiskSize
type ResizeMultiZoneClusterDiskSizeRequest struct {
	*requests.RpcRequest
	ClusterId    string           `position:"Query" name:"ClusterId"`
	LogDiskSize  requests.Integer `position:"Query" name:"LogDiskSize"`
	CoreDiskSize requests.Integer `position:"Query" name:"CoreDiskSize"`
}

// ResizeMultiZoneClusterDiskSizeResponse is the response struct for api ResizeMultiZoneClusterDiskSize
type ResizeMultiZoneClusterDiskSizeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateResizeMultiZoneClusterDiskSizeRequest creates a request to invoke ResizeMultiZoneClusterDiskSize API
func CreateResizeMultiZoneClusterDiskSizeRequest() (request *ResizeMultiZoneClusterDiskSizeRequest) {
	request = &ResizeMultiZoneClusterDiskSizeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("HBase", "2019-01-01", "ResizeMultiZoneClusterDiskSize", "hbase", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResizeMultiZoneClusterDiskSizeResponse creates a response to parse from ResizeMultiZoneClusterDiskSize response
func CreateResizeMultiZoneClusterDiskSizeResponse() (response *ResizeMultiZoneClusterDiskSizeResponse) {
	response = &ResizeMultiZoneClusterDiskSizeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}