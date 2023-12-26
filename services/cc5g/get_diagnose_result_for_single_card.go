package cc5g

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

// GetDiagnoseResultForSingleCard invokes the cc5g.GetDiagnoseResultForSingleCard API synchronously
func (client *Client) GetDiagnoseResultForSingleCard(request *GetDiagnoseResultForSingleCardRequest) (response *GetDiagnoseResultForSingleCardResponse, err error) {
	response = CreateGetDiagnoseResultForSingleCardResponse()
	err = client.DoAction(request, response)
	return
}

// GetDiagnoseResultForSingleCardWithChan invokes the cc5g.GetDiagnoseResultForSingleCard API asynchronously
func (client *Client) GetDiagnoseResultForSingleCardWithChan(request *GetDiagnoseResultForSingleCardRequest) (<-chan *GetDiagnoseResultForSingleCardResponse, <-chan error) {
	responseChan := make(chan *GetDiagnoseResultForSingleCardResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDiagnoseResultForSingleCard(request)
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

// GetDiagnoseResultForSingleCardWithCallback invokes the cc5g.GetDiagnoseResultForSingleCard API asynchronously
func (client *Client) GetDiagnoseResultForSingleCardWithCallback(request *GetDiagnoseResultForSingleCardRequest, callback func(response *GetDiagnoseResultForSingleCardResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDiagnoseResultForSingleCardResponse
		var err error
		defer close(result)
		response, err = client.GetDiagnoseResultForSingleCard(request)
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

// GetDiagnoseResultForSingleCardRequest is the request struct for api GetDiagnoseResultForSingleCard
type GetDiagnoseResultForSingleCardRequest struct {
	*requests.RpcRequest
	RegionNo       string `position:"Query" name:"RegionNo"`
	DiagnoseTaskId string `position:"Query" name:"DiagnoseTaskId"`
}

// GetDiagnoseResultForSingleCardResponse is the response struct for api GetDiagnoseResultForSingleCard
type GetDiagnoseResultForSingleCardResponse struct {
	*responses.BaseResponse
	RequestId                string             `json:"RequestId" xml:"RequestId"`
	WirelessCloudConnectorId string             `json:"WirelessCloudConnectorId" xml:"WirelessCloudConnectorId"`
	CardIp                   string             `json:"CardIp" xml:"CardIp"`
	IccId                    string             `json:"IccId" xml:"IccId"`
	Destination              string             `json:"Destination" xml:"Destination"`
	DestinationType          string             `json:"DestinationType" xml:"DestinationType"`
	BeginTime                int64              `json:"BeginTime" xml:"BeginTime"`
	EndTime                  int64              `json:"EndTime" xml:"EndTime"`
	Status                   string             `json:"Status" xml:"Status"`
	ErrorResult              []DiagnoseResult   `json:"ErrorResult" xml:"ErrorResult"`
	DiagnoseItem             []DiagnoseItemItem `json:"DiagnoseItem" xml:"DiagnoseItem"`
}

// CreateGetDiagnoseResultForSingleCardRequest creates a request to invoke GetDiagnoseResultForSingleCard API
func CreateGetDiagnoseResultForSingleCardRequest() (request *GetDiagnoseResultForSingleCardRequest) {
	request = &GetDiagnoseResultForSingleCardRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CC5G", "2022-03-14", "GetDiagnoseResultForSingleCard", "fivegcc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDiagnoseResultForSingleCardResponse creates a response to parse from GetDiagnoseResultForSingleCard response
func CreateGetDiagnoseResultForSingleCardResponse() (response *GetDiagnoseResultForSingleCardResponse) {
	response = &GetDiagnoseResultForSingleCardResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}