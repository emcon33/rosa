/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/osdfleetmgmt/v1

// LabelBuilder contains the data and logic needed to build 'label' objects.
//
// label settings of the cluster.
type LabelBuilder struct {
	bitmap_ uint32
	id      string
	key     string
	value   string
}

// NewLabel creates a new builder of 'label' objects.
func NewLabel() *LabelBuilder {
	return &LabelBuilder{}
}

// Empty returns true if the builder is empty, i.e. no attribute has a value.
func (b *LabelBuilder) Empty() bool {
	return b == nil || b.bitmap_ == 0
}

// Id sets the value of the 'id' attribute to the given value.
func (b *LabelBuilder) Id(value string) *LabelBuilder {
	b.id = value
	b.bitmap_ |= 1
	return b
}

// Key sets the value of the 'key' attribute to the given value.
func (b *LabelBuilder) Key(value string) *LabelBuilder {
	b.key = value
	b.bitmap_ |= 2
	return b
}

// Value sets the value of the 'value' attribute to the given value.
func (b *LabelBuilder) Value(value string) *LabelBuilder {
	b.value = value
	b.bitmap_ |= 4
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *LabelBuilder) Copy(object *Label) *LabelBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.id = object.id
	b.key = object.key
	b.value = object.value
	return b
}

// Build creates a 'label' object using the configuration stored in the builder.
func (b *LabelBuilder) Build() (object *Label, err error) {
	object = new(Label)
	object.bitmap_ = b.bitmap_
	object.id = b.id
	object.key = b.key
	object.value = b.value
	return
}
