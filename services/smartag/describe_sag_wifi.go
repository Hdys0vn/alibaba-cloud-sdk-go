package smartag

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

// DescribeSagWifi invokes the smartag.DescribeSagWifi API synchronously
func (client *Client) DescribeSagWifi(request *DescribeSagWifiRequest) (response *DescribeSagWifiResponse, err error) {
	response = CreateDescribeSagWifiResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSagWifiWithChan invokes the smartag.DescribeSagWifi API asynchronously
func (client *Client) DescribeSagWifiWithChan(request *DescribeSagWifiRequest) (<-chan *DescribeSagWifiResponse, <-chan error) {
	responseChan := make(chan *DescribeSagWifiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSagWifi(request)
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

// DescribeSagWifiWithCallback invokes the smartag.DescribeSagWifi API asynchronously
func (client *Client) DescribeSagWifiWithCallback(request *DescribeSagWifiRequest, callback func(response *DescribeSagWifiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSagWifiResponse
		var err error
		defer close(result)
		response, err = client.DescribeSagWifi(request)
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

// DescribeSagWifiRequest is the request struct for api DescribeSagWifi
type DescribeSagWifiRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
}

// DescribeSagWifiResponse is the response struct for api DescribeSagWifi
type DescribeSagWifiResponse struct {
	*responses.BaseResponse
	IsEnable           string      `json:"IsEnable" xml:"IsEnable"`
	RequestId          string      `json:"RequestId" xml:"RequestId"`
	IsAuth             string      `json:"IsAuth" xml:"IsAuth"`
	Bandwidth          string      `json:"Bandwidth" xml:"Bandwidth"`
	Channel            string      `json:"Channel" xml:"Channel"`
	Ssid               string      `json:"Ssid" xml:"Ssid"`
	AuthenticationType string      `json:"AuthenticationType" xml:"AuthenticationType"`
	EncryptAlgorithm   string      `json:"EncryptAlgorithm" xml:"EncryptAlgorithm"`
	IsBroadcast        string      `json:"IsBroadcast" xml:"IsBroadcast"`
	TaskStates         []TaskState `json:"TaskStates" xml:"TaskStates"`
}

// CreateDescribeSagWifiRequest creates a request to invoke DescribeSagWifi API
func CreateDescribeSagWifiRequest() (request *DescribeSagWifiRequest) {
	request = &DescribeSagWifiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeSagWifi", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSagWifiResponse creates a response to parse from DescribeSagWifi response
func CreateDescribeSagWifiResponse() (response *DescribeSagWifiResponse) {
	response = &DescribeSagWifiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
