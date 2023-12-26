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

// GetNodeTuplesDownloadUrl invokes the linkwan.GetNodeTuplesDownloadUrl API synchronously
func (client *Client) GetNodeTuplesDownloadUrl(request *GetNodeTuplesDownloadUrlRequest) (response *GetNodeTuplesDownloadUrlResponse, err error) {
	response = CreateGetNodeTuplesDownloadUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetNodeTuplesDownloadUrlWithChan invokes the linkwan.GetNodeTuplesDownloadUrl API asynchronously
func (client *Client) GetNodeTuplesDownloadUrlWithChan(request *GetNodeTuplesDownloadUrlRequest) (<-chan *GetNodeTuplesDownloadUrlResponse, <-chan error) {
	responseChan := make(chan *GetNodeTuplesDownloadUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetNodeTuplesDownloadUrl(request)
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

// GetNodeTuplesDownloadUrlWithCallback invokes the linkwan.GetNodeTuplesDownloadUrl API asynchronously
func (client *Client) GetNodeTuplesDownloadUrlWithCallback(request *GetNodeTuplesDownloadUrlRequest, callback func(response *GetNodeTuplesDownloadUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetNodeTuplesDownloadUrlResponse
		var err error
		defer close(result)
		response, err = client.GetNodeTuplesDownloadUrl(request)
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

// GetNodeTuplesDownloadUrlRequest is the request struct for api GetNodeTuplesDownloadUrl
type GetNodeTuplesDownloadUrlRequest struct {
	*requests.RpcRequest
	OrderId     string `position:"Query" name:"OrderId"`
	ApiProduct  string `position:"Body" name:"ApiProduct"`
	ApiRevision string `position:"Body" name:"ApiRevision"`
}

// GetNodeTuplesDownloadUrlResponse is the response struct for api GetNodeTuplesDownloadUrl
type GetNodeTuplesDownloadUrlResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetNodeTuplesDownloadUrlRequest creates a request to invoke GetNodeTuplesDownloadUrl API
func CreateGetNodeTuplesDownloadUrlRequest() (request *GetNodeTuplesDownloadUrlRequest) {
	request = &GetNodeTuplesDownloadUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("LinkWAN", "2019-03-01", "GetNodeTuplesDownloadUrl", "linkwan", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetNodeTuplesDownloadUrlResponse creates a response to parse from GetNodeTuplesDownloadUrl response
func CreateGetNodeTuplesDownloadUrlResponse() (response *GetNodeTuplesDownloadUrlResponse) {
	response = &GetNodeTuplesDownloadUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}