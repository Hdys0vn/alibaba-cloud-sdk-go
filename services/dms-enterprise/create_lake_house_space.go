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

// CreateLakeHouseSpace invokes the dms_enterprise.CreateLakeHouseSpace API synchronously
func (client *Client) CreateLakeHouseSpace(request *CreateLakeHouseSpaceRequest) (response *CreateLakeHouseSpaceResponse, err error) {
	response = CreateCreateLakeHouseSpaceResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLakeHouseSpaceWithChan invokes the dms_enterprise.CreateLakeHouseSpace API asynchronously
func (client *Client) CreateLakeHouseSpaceWithChan(request *CreateLakeHouseSpaceRequest) (<-chan *CreateLakeHouseSpaceResponse, <-chan error) {
	responseChan := make(chan *CreateLakeHouseSpaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLakeHouseSpace(request)
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

// CreateLakeHouseSpaceWithCallback invokes the dms_enterprise.CreateLakeHouseSpace API asynchronously
func (client *Client) CreateLakeHouseSpaceWithCallback(request *CreateLakeHouseSpaceRequest, callback func(response *CreateLakeHouseSpaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLakeHouseSpaceResponse
		var err error
		defer close(result)
		response, err = client.CreateLakeHouseSpace(request)
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

// CreateLakeHouseSpaceRequest is the request struct for api CreateLakeHouseSpace
type CreateLakeHouseSpaceRequest struct {
	*requests.RpcRequest
	Description string           `position:"Query" name:"Description"`
	Tid         requests.Integer `position:"Query" name:"Tid"`
	Mode        string           `position:"Query" name:"Mode"`
	ProdDbId    string           `position:"Query" name:"ProdDbId"`
	DevDbId     string           `position:"Query" name:"DevDbId"`
	SpaceName   string           `position:"Query" name:"SpaceName"`
	DwDbType    string           `position:"Query" name:"DwDbType"`
	SpaceConfig string           `position:"Query" name:"SpaceConfig"`
}

// CreateLakeHouseSpaceResponse is the response struct for api CreateLakeHouseSpace
type CreateLakeHouseSpaceResponse struct {
	*responses.BaseResponse
	SpaceId      int64  `json:"SpaceId" xml:"SpaceId"`
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
}

// CreateCreateLakeHouseSpaceRequest creates a request to invoke CreateLakeHouseSpace API
func CreateCreateLakeHouseSpaceRequest() (request *CreateLakeHouseSpaceRequest) {
	request = &CreateLakeHouseSpaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "CreateLakeHouseSpace", "dms-enterprise", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLakeHouseSpaceResponse creates a response to parse from CreateLakeHouseSpace response
func CreateCreateLakeHouseSpaceResponse() (response *CreateLakeHouseSpaceResponse) {
	response = &CreateLakeHouseSpaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
