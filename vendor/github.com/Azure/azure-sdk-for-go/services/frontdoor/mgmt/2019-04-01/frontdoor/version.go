package frontdoor

import "github.com/Azure/azure-sdk-for-go/version"

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// UserAgent returns the UserAgent string to use when sending http.Requests.
func UserAgent() string {
<<<<<<< HEAD:vendor/github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2019-04-01/frontdoor/version.go
	return "Azure-SDK-For-Go/" + version.Number + " frontdoor/2019-04-01"
=======
	return "Azure-SDK-For-Go/" + Version() + " frontdoor/2020-01-01"
>>>>>>> 5aa20dd53... vendor: bump terraform-provider-azure to version v2.17.0:vendor/github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-01-01/frontdoor/version.go
}

// Version returns the semantic version (see http://semver.org) of the client.
func Version() string {
	return version.Number
}
