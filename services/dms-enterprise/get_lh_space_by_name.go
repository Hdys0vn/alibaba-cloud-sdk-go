package dms_enterprise

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

// GetLhSpaceByName invokes the dms_enterprise.GetLhSpaceByName API synchronously
func (client *Client) GetLhSpaceByName(request *GetLhSpaceByNameRequest) (response *GetLhSpaceByNameResponse, err error) {
	response = CreateGetLhSpaceByNameResponse()
	err = client.DoAction(request, response)
	return
}

// GetLhSpaceByNameWithChan invokes the dms_enterprise.GetLhSpaceByName API asynchronously
func (client *Client) GetLhSpaceByNameWithChan(request *GetLhSpaceByNameRequest) (<-chan *GetLhSpaceByNameResponse, <-chan error) {
	responseChan := make(chan *GetLhSpaceByNameResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetLhSpaceByName(request)
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

// GetLhSpaceByNameWithCallback invokes the dms_enterprise.GetLhSpaceByName API asynchronously
func (client *Client) GetLhSpaceByNameWithCallback(request *GetLhSpaceByNameRequest, callback func(response *GetLhSpaceByNameResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetLhSpaceByNameResponse
		var err error
		defer close(result)
		response, err = client.GetLhSpaceByName(request)
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

// GetLhSpaceByNameRequest is the request struct for api GetLhSpaceByName
type GetLhSpaceByNameRequest struct {
	*requests.RpcRequest
	Tid       requests.Integer `position:"Query" name:"Tid"`
	SpaceName string           `position:"Query" name:"SpaceName"`
}

// GetLhSpaceByNameResponse is the response struct for api GetLhSpaceByName
type GetLhSpaceByNameResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	ErrorCode      string         `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string         `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool           `json:"Success" xml:"Success"`
	LakehouseSpace LakehouseSpace `json:"LakehouseSpace" xml:"LakehouseSpace"`
}

// CreateGetLhSpaceByNameRequest creates a request to invoke GetLhSpaceByName API
func CreateGetLhSpaceByNameRequest() (request *GetLhSpaceByNameRequest) {
	request = &GetLhSpaceByNameRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "GetLhSpaceByName", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetLhSpaceByNameResponse creates a response to parse from GetLhSpaceByName response
func CreateGetLhSpaceByNameResponse() (response *GetLhSpaceByNameResponse) {
	response = &GetLhSpaceByNameResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}