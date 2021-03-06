// Copyright 2016 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package views

import (
	"testing"

	"github.com/npiganeau/yep/yep/tools/xmlutils"
	. "github.com/smartystreets/goconvey/convey"
)

var viewDef1 string = `
<view id="my_id" name="My View" model="Test__User">
	<form>
		<group>
			<field name="UserName"/>
			<field name="Age"/>
		</group>
	</form>
</view>
`

var viewDef2 string = `
<view id="my_other_id" model="Test__Partner" priority="12">
	<form>
		<h1><field name="Name"/></h1>
		<group name="position_info">
			<field name="Function"/>
		</group>
		<group name="contact_data">
			<field name="Email"/>
		</group>
	</form>
</view>
`

var viewDef3 string = `
<view inherit_id="my_other_id">
	<group name="position_info" position="inside">
		<field name="CompanyName"/>
	</group>
	<xpath expr="//field[@name='Email']" position="after">
		<field name="Phone"/>
	</group>
</view>
`

var viewDef4 string = `
<view inherit_id="my_other_id">
	<group name="contact_data" position="before">
		<group>
			<field name="Address"/>
		</group>
		<hr/>
	</group>
	<h1 position="replace">
		<h2><field name="Name"/></h2>
	</group>
</view>
`

var viewDef5 string = `
<view inherit_id="my_other_id">
	<xpath expr="//field[@name='Address']/.." position="attributes">
		<attribute name="name">address</attribute>
		<attribute name="string">Address</attribute>
	</xpath>
</view>
`

var viewDef6 string = `
<view id="my_tree_id" model="Test__User">
	<tree>
		<field name="UserName"/>
		<field name="Age"/>
	</tree>
</view>
`

func TestViews(t *testing.T) {
	Convey("Creating View 1", t, func() {
		LoadFromEtree(xmlutils.XMLToElement(viewDef1))
		So(len(Registry.views), ShouldEqual, 1)
		So(Registry.GetByID("my_id"), ShouldNotBeNil)
		view := Registry.GetByID("my_id")
		So(view.ID, ShouldEqual, "my_id")
		So(view.Name, ShouldEqual, "My View")
		So(view.Model, ShouldEqual, "Test__User")
		So(view.Priority, ShouldEqual, 16)
		So(view.Arch, ShouldEqual,
			`<form>
	<group>
		<field name="UserName"/>
		<field name="Age"/>
	</group>
</form>
`)
	})
	Convey("Creating View 2", t, func() {
		LoadFromEtree(xmlutils.XMLToElement(viewDef2))
		So(len(Registry.views), ShouldEqual, 2)
		So(Registry.GetByID("my_other_id"), ShouldNotBeNil)
		view := Registry.GetByID("my_other_id")
		So(view.ID, ShouldEqual, "my_other_id")
		So(view.Name, ShouldEqual, "my.other.id")
		So(view.Model, ShouldEqual, "Test__Partner")
		So(view.Priority, ShouldEqual, 12)
		So(view.Arch, ShouldEqual,
			`<form>
	<h1>
		<field name="Name"/>
	</h1>
	<group name="position_info">
		<field name="Function"/>
	</group>
	<group name="contact_data">
		<field name="Email"/>
	</group>
</form>
`)
	})
	Convey("Inheriting View 2", t, func() {
		LoadFromEtree(xmlutils.XMLToElement(viewDef3))
		So(len(Registry.views), ShouldEqual, 2)
		So(Registry.GetByID("my_id"), ShouldNotBeNil)
		So(Registry.GetByID("my_other_id"), ShouldNotBeNil)
		view1 := Registry.GetByID("my_id")
		So(view1.Arch, ShouldEqual,
			`<form>
	<group>
		<field name="UserName"/>
		<field name="Age"/>
	</group>
</form>
`)
		view2 := Registry.GetByID("my_other_id")
		So(view2.Arch, ShouldEqual,
			`<form>
	<h1>
		<field name="Name"/>
	</h1>
	<group name="position_info">
		<field name="Function"/>
		<field name="CompanyName"/>
	</group>
	<group name="contact_data">
		<field name="Email"/>
		<field name="Phone"/>
	</group>
</form>
`)
	})
	Convey("More inheritance on View 2", t, func() {
		LoadFromEtree(xmlutils.XMLToElement(viewDef4))
		So(len(Registry.views), ShouldEqual, 2)
		So(Registry.GetByID("my_id"), ShouldNotBeNil)
		So(Registry.GetByID("my_other_id"), ShouldNotBeNil)
		view2 := Registry.GetByID("my_other_id")
		So(view2.Arch, ShouldEqual,
			`<form>
	<h2>
		<field name="Name"/>
	</h2>
	<group name="position_info">
		<field name="Function"/>
		<field name="CompanyName"/>
	</group>
	<group>
		<field name="Address"/>
	</group>
	<hr/>
	<group name="contact_data">
		<field name="Email"/>
		<field name="Phone"/>
	</group>
</form>
`)
	})
	Convey("Modifying inherited modifications on View 2", t, func() {
		LoadFromEtree(xmlutils.XMLToElement(viewDef5))
		So(len(Registry.views), ShouldEqual, 2)
		So(Registry.GetByID("my_id"), ShouldNotBeNil)
		So(Registry.GetByID("my_other_id"), ShouldNotBeNil)
		view2 := Registry.GetByID("my_other_id")
		So(view2.Arch, ShouldEqual,
			`<form>
	<h2>
		<field name="Name"/>
	</h2>
	<group name="position_info">
		<field name="Function"/>
		<field name="CompanyName"/>
	</group>
	<group name="address" string="Address">
		<field name="Address"/>
	</group>
	<hr/>
	<group name="contact_data">
		<field name="Email"/>
		<field name="Phone"/>
	</group>
</form>
`)
	})
	Convey("Bootstrapping views", t, func() {
		LoadFromEtree(xmlutils.XMLToElement(viewDef6))
		BootStrap()
		view1 := Registry.GetByID("my_id")
		view2 := Registry.GetByID("my_other_id")
		view3 := Registry.GetByID("my_tree_id")
		So(view1, ShouldNotBeNil)
		So(view2, ShouldNotBeNil)
		So(view3, ShouldNotBeNil)
		So(view1.Type, ShouldEqual, VIEW_TYPE_FORM)
		So(view2.Type, ShouldEqual, VIEW_TYPE_FORM)
		So(view3.Type, ShouldEqual, VIEW_TYPE_TREE)
	})
}
