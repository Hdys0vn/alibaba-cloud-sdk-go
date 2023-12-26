package pairecservice

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

// GetSubCrowd invokes the pairecservice.GetSubCrowd API synchronously
func (client *Client) GetSubCrowd(request *GetSubCrowdRequest) (response *GetSubCrowdResponse, err error) {
	response = CreateGetSubCrowdResponse()
	err = client.DoAction(request, response)
	return
}

// GetSubCrowdWithChan invokes the pairecservice.GetSubCrowd API asynchronously
func (client *Client) GetSubCrowdWithChan(request *GetSubCrowdRequest) (<-chan *GetSubCrowdResponse, <-chan error) {
	responseChan := make(chan *GetSubCrowdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetSubCrowd(request)
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

// GetSubCrowdWithCallback invokes the pairecservice.GetSubCrowd API asynchronously
func (client *Client) GetSubCrowdWithCallback(request *GetSubCrowdRequest, callback func(response *GetSubCrowdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetSubCrowdResponse
		var err error
		defer close(result)
		response, err = client.GetSubCrowd(request)
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

// GetSubCrowdRequest is the request struct for api GetSubCrowd
type GetSubCrowdRequest struct {
	*requests.RoaRequest
	CrowdId    string `position:"Path" name:"CrowdId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	SubCrowdId string `position:"Path" name:"SubCrowdId"`
}

// GetSubCrowdResponse is the response struct for api GetSubCrowd
type GetSubCrowdResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	Source        string `json:"Source" xml:"Source"`
	Users         string `json:"Users" xml:"Users"`
	Quantity      string `json:"Quantity" xml:"Quantity"`
	GmtCreateTime string `json:"GmtCreateTime" xml:"GmtCreateTime"`
}

// CreateGetSubCrowdRequest creates a request to invoke GetSubCrowd API
func CreateGetSubCrowdRequest() (request *GetSubCrowdRequest) {
	request = &GetSubCrowdRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("PaiRecService", "2022-12-13", "GetSubCrowd", "/api/v1/crowds/[CrowdId]/subcrowds/[SubCrowdId]", "", "")
	request.Method = requests.GET
	return
}

// CreateGetSubCrowdResponse creates a response to parse from GetSubCrowd response
func CreateGetSubCrowdResponse() (response *GetSubCrowdResponse) {
	response = &GetSubCrowdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
