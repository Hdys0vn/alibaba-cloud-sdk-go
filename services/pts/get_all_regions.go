package pts

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

// GetAllRegions invokes the pts.GetAllRegions API synchronously
func (client *Client) GetAllRegions(request *GetAllRegionsRequest) (response *GetAllRegionsResponse, err error) {
	response = CreateGetAllRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// GetAllRegionsWithChan invokes the pts.GetAllRegions API asynchronously
func (client *Client) GetAllRegionsWithChan(request *GetAllRegionsRequest) (<-chan *GetAllRegionsResponse, <-chan error) {
	responseChan := make(chan *GetAllRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAllRegions(request)
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

// GetAllRegionsWithCallback invokes the pts.GetAllRegions API asynchronously
func (client *Client) GetAllRegionsWithCallback(request *GetAllRegionsRequest, callback func(response *GetAllRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAllRegionsResponse
		var err error
		defer close(result)
		response, err = client.GetAllRegions(request)
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

// GetAllRegionsRequest is the request struct for api GetAllRegions
type GetAllRegionsRequest struct {
	*requests.RpcRequest
}

// GetAllRegionsResponse is the response struct for api GetAllRegions
type GetAllRegionsResponse struct {
	*responses.BaseResponse
	Message        string                 `json:"Message" xml:"Message"`
	RequestId      string                 `json:"RequestId" xml:"RequestId"`
	AllRegions     map[string]interface{} `json:"AllRegions" xml:"AllRegions"`
	HttpStatusCode int                    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           string                 `json:"Code" xml:"Code"`
	Success        bool                   `json:"Success" xml:"Success"`
}

// CreateGetAllRegionsRequest creates a request to invoke GetAllRegions API
func CreateGetAllRegionsRequest() (request *GetAllRegionsRequest) {
	request = &GetAllRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("PTS", "2020-10-20", "GetAllRegions", "", "")
	request.Method = requests.POST
	return
}

// CreateGetAllRegionsResponse creates a response to parse from GetAllRegions response
func CreateGetAllRegionsResponse() (response *GetAllRegionsResponse) {
	response = &GetAllRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}