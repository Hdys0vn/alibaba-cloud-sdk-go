package linkwan

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

// GetGatewayTuplesDownloadUrl invokes the linkwan.GetGatewayTuplesDownloadUrl API synchronously
func (client *Client) GetGatewayTuplesDownloadUrl(request *GetGatewayTuplesDownloadUrlRequest) (response *GetGatewayTuplesDownloadUrlResponse, err error) {
	response = CreateGetGatewayTuplesDownloadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetGatewayTuplesDownloadUrlWithChan invokes the linkwan.GetGatewayTuplesDownloadUrl API asynchronously
func (client *Client) GetGatewayTuplesDownloadUrlWithChan(request *GetGatewayTuplesDownloadUrlRequest) (<-chan *GetGatewayTuplesDownloadUrlResponse, <-chan error) {
	responseChan := make(chan *GetGatewayTuplesDownloadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetGatewayTuplesDownloadUrl(request)
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

// GetGatewayTuplesDownloadUrlWithCallback invokes the linkwan.GetGatewayTuplesDownloadUrl API asynchronously
func (client *Client) GetGatewayTuplesDownloadUrlWithCallback(request *GetGatewayTuplesDownloadUrlRequest, callback func(response *GetGatewayTuplesDownloadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetGatewayTuplesDownloadUrlResponse
		var err error
		defer close(result)
		response, err = client.GetGatewayTuplesDownloadUrl(request)
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

// GetGatewayTuplesDownloadUrlRequest is the request struct for api GetGatewayTuplesDownloadUrl
type GetGatewayTuplesDownloadUrlRequest struct {
	*requests.RpcRequest
	OrderId     string `position:"Query" name:"OrderId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// GetGatewayTuplesDownloadUrlResponse is the response struct for api GetGatewayTuplesDownloadUrl
type GetGatewayTuplesDownloadUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetGatewayTuplesDownloadUrlRequest creates a request to invoke GetGatewayTuplesDownloadUrl API
func CreateGetGatewayTuplesDownloadUrlRequest() (request *GetGatewayTuplesDownloadUrlRequest) {
	request = &GetGatewayTuplesDownloadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "GetGatewayTuplesDownloadUrl", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetGatewayTuplesDownloadUrlResponse creates a response to parse from GetGatewayTuplesDownloadUrl response
func CreateGetGatewayTuplesDownloadUrlResponse() (response *GetGatewayTuplesDownloadUrlResponse) {
	response = &GetGatewayTuplesDownloadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
