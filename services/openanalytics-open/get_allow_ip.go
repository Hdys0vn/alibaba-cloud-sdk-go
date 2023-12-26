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

// GetAllowIP invokes the openanalytics_open.GetAllowIP API synchronously
func (client *Client) GetAllowIP(request *GetAllowIPRequest) (response *GetAllowIPResponse, err error) {
	response = CreateGetAllowIPResponse()
	err = client.DoAction(request, response)
	return
}

// GetAllowIPWithChan invokes the openanalytics_open.GetAllowIP API asynchronously
func (client *Client) GetAllowIPWithChan(request *GetAllowIPRequest) (<-chan *GetAllowIPResponse, <-chan error) {
	responseChan := make(chan *GetAllowIPResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetAllowIP(request)
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

// GetAllowIPWithCallback invokes the openanalytics_open.GetAllowIP API asynchronously
func (client *Client) GetAllowIPWithCallback(request *GetAllowIPRequest, callback func(response *GetAllowIPResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetAllowIPResponse
		var err error
		defer close(result)
		response, err = client.GetAllowIP(request)
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

// GetAllowIPRequest is the request struct for api GetAllowIP
type GetAllowIPRequest struct {
	*requests.RpcRequest
	NetworkType string `position:"Body" name:"NetworkType"`
	Product     string `position:"Body" name:"Product"`
}

// GetAllowIPResponse is the response struct for api GetAllowIP
type GetAllowIPResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	RegionId  string `json:"RegionId" xml:"RegionId"`
	AllowIP   string `json:"AllowIP" xml:"AllowIP"`
}

// CreateGetAllowIPRequest creates a request to invoke GetAllowIP API
func CreateGetAllowIPRequest() (request *GetAllowIPRequest) {
	request = &GetAllowIPRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "GetAllowIP", "openanalytics", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetAllowIPResponse creates a response to parse from GetAllowIP response
func CreateGetAllowIPResponse() (response *GetAllowIPResponse) {
	response = &GetAllowIPResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
