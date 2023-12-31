package vpc

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

// CreateFullNatEntry invokes the vpc.CreateFullNatEntry API synchronously
func (client *Client) CreateFullNatEntry(request *CreateFullNatEntryRequest) (response *CreateFullNatEntryResponse, err error) {
	response = CreateCreateFullNatEntryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFullNatEntryWithChan invokes the vpc.CreateFullNatEntry API asynchronously
func (client *Client) CreateFullNatEntryWithChan(request *CreateFullNatEntryRequest) (<-chan *CreateFullNatEntryResponse, <-chan error) {
	responseChan := make(chan *CreateFullNatEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFullNatEntry(request)
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

// CreateFullNatEntryWithCallback invokes the vpc.CreateFullNatEntry API asynchronously
func (client *Client) CreateFullNatEntryWithCallback(request *CreateFullNatEntryRequest, callback func(response *CreateFullNatEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFullNatEntryResponse
		var err error
		defer close(result)
		response, err = client.CreateFullNatEntry(request)
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

// CreateFullNatEntryRequest is the request struct for api CreateFullNatEntry
type CreateFullNatEntryRequest struct {
	*requests.RpcRequest
	FullNatEntryDescription string           `position:"Query" name:"FullNatEntryDescription"`
	ResourceOwnerId         requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AccessIp                string           `position:"Query" name:"AccessIp"`
	ClientToken             string           `position:"Query" name:"ClientToken"`
	NatIpPort               string           `position:"Query" name:"NatIpPort"`
	FullNatTableId          string           `position:"Query" name:"FullNatTableId"`
	AccessPort              string           `position:"Query" name:"AccessPort"`
	DryRun                  requests.Boolean `position:"Query" name:"DryRun"`
	ResourceOwnerAccount    string           `position:"Query" name:"ResourceOwnerAccount"`
	IpProtocol              string           `position:"Query" name:"IpProtocol"`
	OwnerAccount            string           `position:"Query" name:"OwnerAccount"`
	OwnerId                 requests.Integer `position:"Query" name:"OwnerId"`
	FullNatEntryName        string           `position:"Query" name:"FullNatEntryName"`
	NatIp                   string           `position:"Query" name:"NatIp"`
	NetworkInterfaceId      string           `position:"Query" name:"NetworkInterfaceId"`
}

// CreateFullNatEntryResponse is the response struct for api CreateFullNatEntry
type CreateFullNatEntryResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	FullNatEntryId string `json:"FullNatEntryId" xml:"FullNatEntryId"`
}

// CreateCreateFullNatEntryRequest creates a request to invoke CreateFullNatEntry API
func CreateCreateFullNatEntryRequest() (request *CreateFullNatEntryRequest) {
	request = &CreateFullNatEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "CreateFullNatEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFullNatEntryResponse creates a response to parse from CreateFullNatEntry response
func CreateCreateFullNatEntryResponse() (response *CreateFullNatEntryResponse) {
	response = &CreateFullNatEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
