package modelflatteninggroup_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/Azure/Tests/model-flattening"

	. "Tests/model-flattening"
	"Tests/utils"

	chk "gopkg.in/check.v1"
)

func Test(t *testing.T) { chk.TestingT(t) }

type ModelFlatteningSuite struct{}

var _ = chk.Suite(&ModelFlatteningSuite{})

var modelflatteningClient = getmodelflatteningClient()

func getmodelflatteningClient() modelflatteninggroup.ManagementClient {
	c := modelflatteninggroup.New()
	c.BaseURI = utils.GetBaseURI()
	return c
}

func TestMain(m *testing.M) {
	server, err := utils.StartTestServer()
	if err != nil {
		fmt.Println(err)
	}
	exitCode := m.Run()
	server.Process.Kill()
	os.Exit(exitCode)
}

func (s *ModelFlatteningSuite) TestGetArray(c *chk.C) {
	res, err := modelflatteningClient.GetArray()
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.HasLen, 3)

	id, loc, name, typ := "1", "Building 44", "Resource1", "Microsoft.Web/sites"
	tag1, tag2 := "value1", "value3"
	provisioningState, pname, propty := "Succeeded", "Product1", "Flat"
	id1, loc1, name1 := "2", "Building 44", "Resource2"
	id2, name2 := "3", "Resource3"

	c.Assert(*res.Value, chk.DeepEquals, []modelflatteninggroup.FlattenedProduct{
		modelflatteninggroup.FlattenedProduct{
			ID:       &id,
			Location: &loc,
			Tags:     &map[string]*string{"tag1": &tag1, "tag2": &tag2},
			Name:     &name,
			Type:     &typ,
			Properties: &modelflatteninggroup.FlattenedProductProperties{
				ProvisioningState:       &provisioningState,
				ProvisioningStateValues: "OK",
				Pname: &pname,
				Type:  &propty,
			},
		},
		modelflatteninggroup.FlattenedProduct{
			ID:       &id1,
			Location: &loc1,
			Name:     &name1,
		},
		modelflatteninggroup.FlattenedProduct{
			ID:   &id2,
			Name: &name2,
		},
	})
}

func (s *ModelFlatteningSuite) TestGetDictionary(c *chk.C) {
	res, err := modelflatteningClient.GetDictionary()
	c.Assert(err, chk.IsNil)
	c.Assert(*res.Value, chk.HasLen, 3)

	id, loc, name, typ := "1", "Building 44", "Resource1", "Microsoft.Web/sites"
	tag1, tag2 := "value1", "value3"
	provisioningState, pname, propty := "Succeeded", "Product1", "Flat"
	id1, loc1, name1 := "2", "Building 44", "Resource2"
	id2, name2 := "3", "Resource3"

	c.Assert(*res.Value, chk.DeepEquals, map[string]*modelflatteninggroup.FlattenedProduct{
		"Product1": &modelflatteninggroup.FlattenedProduct{
			ID:       &id,
			Location: &loc,
			Tags:     &map[string]*string{"tag1": &tag1, "tag2": &tag2},
			Name:     &name,
			Type:     &typ,
			Properties: &modelflatteninggroup.FlattenedProductProperties{
				ProvisioningState:       &provisioningState,
				ProvisioningStateValues: "OK",
				Pname: &pname,
				Type:  &propty,
			},
		},
		"Product2": &modelflatteninggroup.FlattenedProduct{
			ID:       &id1,
			Location: &loc1,
			Name:     &name1,
		},
		"Product3": &modelflatteninggroup.FlattenedProduct{
			ID:   &id2,
			Name: &name2,
		},
	})
}

func (s *ModelFlatteningSuite) TestGetResourceCollection(c *chk.C) {
	res, err := modelflatteningClient.GetResourceCollection()
	c.Assert(err, chk.IsNil)

	tag1, tag2 := "value1", "value3"

	id, loc, name, typ := "1", "Building 44", "Resource1", "Microsoft.Web/sites"
	provisioningState, pname, propty := "Succeeded", "Product1", "Flat"
	id1, loc1, name1 := "2", "Building 44", "Resource2"
	id2, name2 := "3", "Resource3"

	dictionaryOfResources := map[string]*modelflatteninggroup.FlattenedProduct{
		"Product1": &modelflatteninggroup.FlattenedProduct{
			ID:       &id,
			Location: &loc,
			Tags:     &map[string]*string{"tag1": &tag1, "tag2": &tag2},
			Name:     &name,
			Type:     &typ,
			Properties: &modelflatteninggroup.FlattenedProductProperties{
				ProvisioningState:       &provisioningState,
				ProvisioningStateValues: "OK",
				Pname: &pname,
				Type:  &propty,
			},
		},
		"Product2": &modelflatteninggroup.FlattenedProduct{
			ID:       &id1,
			Location: &loc1,
			Name:     &name1,
		},
		"Product3": &modelflatteninggroup.FlattenedProduct{
			ID:   &id2,
			Name: &name2,
		},
	}

	c.Assert(*res.Dictionaryofresources, chk.DeepEquals, dictionaryOfResources)

	id4, loc4, name4, typ4 := "4", "Building 44", "Resource4", "Microsoft.Web/sites"
	provisioningState4, pname4, propty4 := "Succeeded", "Product4", "Flat"
	id5, loc5, name5 := "5", "Building 44", "Resource5"
	id6, name6 := "6", "Resource6"

	arrayOfResources := []modelflatteninggroup.FlattenedProduct{
		modelflatteninggroup.FlattenedProduct{
			ID:       &id4,
			Location: &loc4,
			Tags:     &map[string]*string{"tag1": &tag1, "tag2": &tag2},
			Name:     &name4,
			Type:     &typ4,
			Properties: &modelflatteninggroup.FlattenedProductProperties{
				ProvisioningState:       &provisioningState4,
				ProvisioningStateValues: "OK",
				Pname: &pname4,
				Type:  &propty4,
			},
		},
		modelflatteninggroup.FlattenedProduct{
			ID:       &id5,
			Location: &loc5,
			Name:     &name5,
		},
		modelflatteninggroup.FlattenedProduct{
			ID:   &id6,
			Name: &name6,
		},
	}

	c.Assert(*res.Arrayofresources, chk.DeepEquals, arrayOfResources)

	id7, loc7, name7 := "7", "Building 44", "Resource7"

	productresource := modelflatteninggroup.FlattenedProduct{
		ID:       &id7,
		Location: &loc7,
		Name:     &name7,
	}

	c.Assert(*res.Productresource, chk.DeepEquals, productresource)
}

func (s *ModelFlatteningSuite) TestPostFlattenedSimpleProduct(c *chk.C) {
	id, description, displayName, capacity, odata := "123", "product description", "max name", "Large", "http://foo"
	arg := modelflatteninggroup.SimpleProduct{
		BaseProductID:          &id,
		BaseProductDescription: &description,
		Details: &modelflatteninggroup.SimpleProductProperties{
			MaxProductDisplayName: &displayName,
			MaxProductCapacity:    &capacity,
			MaxProductImage: &modelflatteninggroup.ProductURL{
				Odatavalue: &odata,
			},
		},
	}

	res, err := modelflatteningClient.PostFlattenedSimpleProduct(&arg)
	c.Assert(err, chk.IsNil)
	arg.Response = res.Response
	c.Assert(res, chk.DeepEquals, arg)
}

func (s *ModelFlatteningSuite) TestPutArray(c *chk.C) {
	loc1, loc2, tag1, tag2 := "West US", "Building 44", "value1", "value3"
	_, err := modelflatteningClient.PutArray([]modelflatteninggroup.Resource{
		modelflatteninggroup.Resource{
			Tags:     &map[string]*string{"tag1": &tag1, "tag2": &tag2},
			Location: &loc1,
		},
		modelflatteninggroup.Resource{
			Location: &loc2,
		},
	})
	c.Assert(err, chk.IsNil)
}

func (s *ModelFlatteningSuite) TestPutDictionary(c *chk.C) {
	jsonBlob := `{"Resource1":{"location":"West US", "tags":{"tag1":"value1", "tag2":"value3"},"properties":{"pname":"Product1","type":"Flat"}},
					"Resource2":{"location":"Building 44", "properties":{"pname":"Product2","type":"Flat"}}}`
	type resourceDictionary map[string]modelflatteninggroup.FlattenedProduct
	var r resourceDictionary

	err := json.Unmarshal([]byte(jsonBlob), &r)
	_, err = modelflatteningClient.PutDictionary(r)
	c.Assert(err, chk.IsNil)
}

func (s *ModelFlatteningSuite) TestPutResourceCollection(c *chk.C) {
	dictionaryBlob := `{"Resource1":{"location":"West US", "tags":{"tag1":"value1", "tag2":"value3"},"properties":{"pname":"Product1","type":"Flat"}},
					"Resource2":{"location":"Building 44", "properties":{"pname":"Product2","type":"Flat"}}}`
	jsonBlob := `{"arrayofresources":[{"location":"West US", "tags":{"tag1":"value1", "tag2":"value3"}, "properties":{"pname":"Product1","type":"Flat"}},
                            {"location":"East US", "properties":{"pname":"Product2","type":"Flat"}}],
                            "dictionaryofresources": ` + dictionaryBlob + `,"productresource":{"location":"India", "properties":{"pname":"Azure","type":"Flat"}}}`
	var r modelflatteninggroup.ResourceCollection

	err := json.Unmarshal([]byte(jsonBlob), &r)
	_, err = modelflatteningClient.PutResourceCollection(&r)
	c.Assert(err, chk.IsNil)
}

func (s *ModelFlatteningSuite) TestPutSimpleProduct(c *chk.C) {
	id, description, displayName, capacity, odata := "123", "product description", "max name", "Large", "http://foo"
	arg := &modelflatteninggroup.SimpleProduct{
		BaseProductID:          &id,
		BaseProductDescription: &description,
		Details: &modelflatteninggroup.SimpleProductProperties{
			MaxProductDisplayName: &displayName,
			MaxProductCapacity:    &capacity,
			MaxProductImage: &modelflatteninggroup.ProductURL{
				Odatavalue: &odata,
			},
		},
	}

	res, err := modelflatteningClient.PutSimpleProduct(arg)
	c.Assert(err, chk.IsNil)
	arg.Response = res.Response
	c.Assert(res, chk.DeepEquals, arg)
}

func (s *ModelFlatteningSuite) TestPutSimpleProductWithGrouping(c *chk.C) {
	id, description, displayName, capacity, odata := "123", "product description", "max name", "Large", "http://foo"
	arg := &modelflatteninggroup.SimpleProduct{
		BaseProductID:          &id,
		BaseProductDescription: &description,
		Details: &modelflatteninggroup.SimpleProductProperties{
			MaxProductDisplayName: &displayName,
			MaxProductCapacity:    &capacity,
			MaxProductImage: &modelflatteninggroup.ProductURL{
				Odatavalue: &odata,
			},
		},
	}
	res, err := modelflatteningClient.PutSimpleProductWithGrouping("groupproduct", arg)
	c.Assert(err, chk.IsNil)
	c.Assert(*res.BaseProductID, chk.DeepEquals, *arg.BaseProductID)
	c.Assert(*res.BaseProductDescription, chk.DeepEquals, *arg.BaseProductDescription)
	c.Assert(*res.Details, chk.DeepEquals, *arg.Details)
}
