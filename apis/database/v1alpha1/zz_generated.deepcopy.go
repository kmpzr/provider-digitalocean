// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseCluster) DeepCopyInto(out *DODatabaseCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseCluster.
func (in *DODatabaseCluster) DeepCopy() *DODatabaseCluster {
	if in == nil {
		return nil
	}
	out := new(DODatabaseCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DODatabaseCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseClusterConnection) DeepCopyInto(out *DODatabaseClusterConnection) {
	*out = *in
	if in.URI != nil {
		in, out := &in.URI, &out.URI
		*out = new(string)
		**out = **in
	}
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
	if in.User != nil {
		in, out := &in.User, &out.User
		*out = new(string)
		**out = **in
	}
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseClusterConnection.
func (in *DODatabaseClusterConnection) DeepCopy() *DODatabaseClusterConnection {
	if in == nil {
		return nil
	}
	out := new(DODatabaseClusterConnection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseClusterList) DeepCopyInto(out *DODatabaseClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DODatabaseCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseClusterList.
func (in *DODatabaseClusterList) DeepCopy() *DODatabaseClusterList {
	if in == nil {
		return nil
	}
	out := new(DODatabaseClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DODatabaseClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseClusterMaintenanceWindow) DeepCopyInto(out *DODatabaseClusterMaintenanceWindow) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseClusterMaintenanceWindow.
func (in *DODatabaseClusterMaintenanceWindow) DeepCopy() *DODatabaseClusterMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := new(DODatabaseClusterMaintenanceWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseClusterObservation) DeepCopyInto(out *DODatabaseClusterObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DbNames != nil {
		in, out := &in.DbNames, &out.DbNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Connection.DeepCopyInto(&out.Connection)
	in.PrivateConnection.DeepCopyInto(&out.PrivateConnection)
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]DODatabaseClusterUser, len(*in))
		copy(*out, *in)
	}
	in.MaintenanceWindow.DeepCopyInto(&out.MaintenanceWindow)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseClusterObservation.
func (in *DODatabaseClusterObservation) DeepCopy() *DODatabaseClusterObservation {
	if in == nil {
		return nil
	}
	out := new(DODatabaseClusterObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseClusterParameters) DeepCopyInto(out *DODatabaseClusterParameters) {
	*out = *in
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.PrivateNetworkUUID != nil {
		in, out := &in.PrivateNetworkUUID, &out.PrivateNetworkUUID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseClusterParameters.
func (in *DODatabaseClusterParameters) DeepCopy() *DODatabaseClusterParameters {
	if in == nil {
		return nil
	}
	out := new(DODatabaseClusterParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseClusterSpec) DeepCopyInto(out *DODatabaseClusterSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseClusterSpec.
func (in *DODatabaseClusterSpec) DeepCopy() *DODatabaseClusterSpec {
	if in == nil {
		return nil
	}
	out := new(DODatabaseClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseClusterStatus) DeepCopyInto(out *DODatabaseClusterStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseClusterStatus.
func (in *DODatabaseClusterStatus) DeepCopy() *DODatabaseClusterStatus {
	if in == nil {
		return nil
	}
	out := new(DODatabaseClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseClusterUser) DeepCopyInto(out *DODatabaseClusterUser) {
	*out = *in
	out.MySQLSettings = in.MySQLSettings
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseClusterUser.
func (in *DODatabaseClusterUser) DeepCopy() *DODatabaseClusterUser {
	if in == nil {
		return nil
	}
	out := new(DODatabaseClusterUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DODatabaseUserMySQLSettings) DeepCopyInto(out *DODatabaseUserMySQLSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DODatabaseUserMySQLSettings.
func (in *DODatabaseUserMySQLSettings) DeepCopy() *DODatabaseUserMySQLSettings {
	if in == nil {
		return nil
	}
	out := new(DODatabaseUserMySQLSettings)
	in.DeepCopyInto(out)
	return out
}
