// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

import (
	"encoding/json"
	v1api20220701s "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_NatGateway_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NatGateway to hub returns original",
		prop.ForAll(RunResourceConversionTestForNatGateway, NatGatewayGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForNatGateway tests if a specific instance of NatGateway round trips to the hub storage version and back losslessly
func RunResourceConversionTestForNatGateway(subject NatGateway) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v1api20220701s.NatGateway
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual NatGateway
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NatGateway_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NatGateway to NatGateway via AssignProperties_To_NatGateway & AssignProperties_From_NatGateway returns original",
		prop.ForAll(RunPropertyAssignmentTestForNatGateway, NatGatewayGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNatGateway tests if a specific instance of NatGateway can be assigned to v1api20220701storage and back losslessly
func RunPropertyAssignmentTestForNatGateway(subject NatGateway) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220701s.NatGateway
	err := copied.AssignProperties_To_NatGateway(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NatGateway
	err = actual.AssignProperties_From_NatGateway(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NatGateway_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGateway via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGateway, NatGatewayGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGateway runs a test to see if a specific instance of NatGateway round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGateway(subject NatGateway) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGateway
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NatGateway instances for property testing - lazily instantiated by NatGatewayGenerator()
var natGatewayGenerator gopter.Gen

// NatGatewayGenerator returns a generator of NatGateway instances for property testing.
func NatGatewayGenerator() gopter.Gen {
	if natGatewayGenerator != nil {
		return natGatewayGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForNatGateway(generators)
	natGatewayGenerator = gen.Struct(reflect.TypeOf(NatGateway{}), generators)

	return natGatewayGenerator
}

// AddRelatedPropertyGeneratorsForNatGateway is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNatGateway(gens map[string]gopter.Gen) {
	gens["Spec"] = NatGateway_SpecGenerator()
	gens["Status"] = NatGateway_STATUSGenerator()
}

func Test_NatGateway_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NatGateway_Spec to NatGateway_Spec via AssignProperties_To_NatGateway_Spec & AssignProperties_From_NatGateway_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForNatGateway_Spec, NatGateway_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNatGateway_Spec tests if a specific instance of NatGateway_Spec can be assigned to v1api20220701storage and back losslessly
func RunPropertyAssignmentTestForNatGateway_Spec(subject NatGateway_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220701s.NatGateway_Spec
	err := copied.AssignProperties_To_NatGateway_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NatGateway_Spec
	err = actual.AssignProperties_From_NatGateway_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NatGateway_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGateway_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGateway_Spec, NatGateway_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGateway_Spec runs a test to see if a specific instance of NatGateway_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGateway_Spec(subject NatGateway_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGateway_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NatGateway_Spec instances for property testing - lazily instantiated by NatGateway_SpecGenerator()
var natGateway_SpecGenerator gopter.Gen

// NatGateway_SpecGenerator returns a generator of NatGateway_Spec instances for property testing.
// We first initialize natGateway_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NatGateway_SpecGenerator() gopter.Gen {
	if natGateway_SpecGenerator != nil {
		return natGateway_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_Spec(generators)
	natGateway_SpecGenerator = gen.Struct(reflect.TypeOf(NatGateway_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_Spec(generators)
	AddRelatedPropertyGeneratorsForNatGateway_Spec(generators)
	natGateway_SpecGenerator = gen.Struct(reflect.TypeOf(NatGateway_Spec{}), generators)

	return natGateway_SpecGenerator
}

// AddIndependentPropertyGeneratorsForNatGateway_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGateway_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNatGateway_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNatGateway_Spec(gens map[string]gopter.Gen) {
	gens["PublicIpAddresses"] = gen.SliceOf(ApplicationGatewaySubResourceGenerator())
	gens["PublicIpPrefixes"] = gen.SliceOf(ApplicationGatewaySubResourceGenerator())
	gens["Sku"] = gen.PtrOf(NatGatewaySkuGenerator())
}

func Test_NatGateway_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NatGateway_STATUS to NatGateway_STATUS via AssignProperties_To_NatGateway_STATUS & AssignProperties_From_NatGateway_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNatGateway_STATUS, NatGateway_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNatGateway_STATUS tests if a specific instance of NatGateway_STATUS can be assigned to v1api20220701storage and back losslessly
func RunPropertyAssignmentTestForNatGateway_STATUS(subject NatGateway_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220701s.NatGateway_STATUS
	err := copied.AssignProperties_To_NatGateway_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NatGateway_STATUS
	err = actual.AssignProperties_From_NatGateway_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NatGateway_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGateway_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGateway_STATUS, NatGateway_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGateway_STATUS runs a test to see if a specific instance of NatGateway_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGateway_STATUS(subject NatGateway_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGateway_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NatGateway_STATUS instances for property testing - lazily instantiated by NatGateway_STATUSGenerator()
var natGateway_STATUSGenerator gopter.Gen

// NatGateway_STATUSGenerator returns a generator of NatGateway_STATUS instances for property testing.
// We first initialize natGateway_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NatGateway_STATUSGenerator() gopter.Gen {
	if natGateway_STATUSGenerator != nil {
		return natGateway_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_STATUS(generators)
	natGateway_STATUSGenerator = gen.Struct(reflect.TypeOf(NatGateway_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGateway_STATUS(generators)
	AddRelatedPropertyGeneratorsForNatGateway_STATUS(generators)
	natGateway_STATUSGenerator = gen.Struct(reflect.TypeOf(NatGateway_STATUS{}), generators)

	return natGateway_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNatGateway_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGateway_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ApplicationGatewayProvisioningState_STATUS_Deleting,
		ApplicationGatewayProvisioningState_STATUS_Failed,
		ApplicationGatewayProvisioningState_STATUS_Succeeded,
		ApplicationGatewayProvisioningState_STATUS_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Zones"] = gen.SliceOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNatGateway_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNatGateway_STATUS(gens map[string]gopter.Gen) {
	gens["PublicIpAddresses"] = gen.SliceOf(ApplicationGatewaySubResource_STATUSGenerator())
	gens["PublicIpPrefixes"] = gen.SliceOf(ApplicationGatewaySubResource_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(NatGatewaySku_STATUSGenerator())
	gens["Subnets"] = gen.SliceOf(ApplicationGatewaySubResource_STATUSGenerator())
}

func Test_ApplicationGatewaySubResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationGatewaySubResource to ApplicationGatewaySubResource via AssignProperties_To_ApplicationGatewaySubResource & AssignProperties_From_ApplicationGatewaySubResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationGatewaySubResource, ApplicationGatewaySubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationGatewaySubResource tests if a specific instance of ApplicationGatewaySubResource can be assigned to v1api20220701storage and back losslessly
func RunPropertyAssignmentTestForApplicationGatewaySubResource(subject ApplicationGatewaySubResource) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220701s.ApplicationGatewaySubResource
	err := copied.AssignProperties_To_ApplicationGatewaySubResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationGatewaySubResource
	err = actual.AssignProperties_From_ApplicationGatewaySubResource(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ApplicationGatewaySubResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationGatewaySubResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationGatewaySubResource, ApplicationGatewaySubResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationGatewaySubResource runs a test to see if a specific instance of ApplicationGatewaySubResource round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationGatewaySubResource(subject ApplicationGatewaySubResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationGatewaySubResource
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ApplicationGatewaySubResource instances for property testing - lazily instantiated by
// ApplicationGatewaySubResourceGenerator()
var applicationGatewaySubResourceGenerator gopter.Gen

// ApplicationGatewaySubResourceGenerator returns a generator of ApplicationGatewaySubResource instances for property testing.
func ApplicationGatewaySubResourceGenerator() gopter.Gen {
	if applicationGatewaySubResourceGenerator != nil {
		return applicationGatewaySubResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	applicationGatewaySubResourceGenerator = gen.Struct(reflect.TypeOf(ApplicationGatewaySubResource{}), generators)

	return applicationGatewaySubResourceGenerator
}

func Test_ApplicationGatewaySubResource_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ApplicationGatewaySubResource_STATUS to ApplicationGatewaySubResource_STATUS via AssignProperties_To_ApplicationGatewaySubResource_STATUS & AssignProperties_From_ApplicationGatewaySubResource_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForApplicationGatewaySubResource_STATUS, ApplicationGatewaySubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForApplicationGatewaySubResource_STATUS tests if a specific instance of ApplicationGatewaySubResource_STATUS can be assigned to v1api20220701storage and back losslessly
func RunPropertyAssignmentTestForApplicationGatewaySubResource_STATUS(subject ApplicationGatewaySubResource_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220701s.ApplicationGatewaySubResource_STATUS
	err := copied.AssignProperties_To_ApplicationGatewaySubResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ApplicationGatewaySubResource_STATUS
	err = actual.AssignProperties_From_ApplicationGatewaySubResource_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ApplicationGatewaySubResource_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ApplicationGatewaySubResource_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForApplicationGatewaySubResource_STATUS, ApplicationGatewaySubResource_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForApplicationGatewaySubResource_STATUS runs a test to see if a specific instance of ApplicationGatewaySubResource_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForApplicationGatewaySubResource_STATUS(subject ApplicationGatewaySubResource_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ApplicationGatewaySubResource_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ApplicationGatewaySubResource_STATUS instances for property testing - lazily instantiated by
// ApplicationGatewaySubResource_STATUSGenerator()
var applicationGatewaySubResource_STATUSGenerator gopter.Gen

// ApplicationGatewaySubResource_STATUSGenerator returns a generator of ApplicationGatewaySubResource_STATUS instances for property testing.
func ApplicationGatewaySubResource_STATUSGenerator() gopter.Gen {
	if applicationGatewaySubResource_STATUSGenerator != nil {
		return applicationGatewaySubResource_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForApplicationGatewaySubResource_STATUS(generators)
	applicationGatewaySubResource_STATUSGenerator = gen.Struct(reflect.TypeOf(ApplicationGatewaySubResource_STATUS{}), generators)

	return applicationGatewaySubResource_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForApplicationGatewaySubResource_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForApplicationGatewaySubResource_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}

func Test_NatGatewaySku_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NatGatewaySku to NatGatewaySku via AssignProperties_To_NatGatewaySku & AssignProperties_From_NatGatewaySku returns original",
		prop.ForAll(RunPropertyAssignmentTestForNatGatewaySku, NatGatewaySkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNatGatewaySku tests if a specific instance of NatGatewaySku can be assigned to v1api20220701storage and back losslessly
func RunPropertyAssignmentTestForNatGatewaySku(subject NatGatewaySku) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220701s.NatGatewaySku
	err := copied.AssignProperties_To_NatGatewaySku(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NatGatewaySku
	err = actual.AssignProperties_From_NatGatewaySku(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NatGatewaySku_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewaySku via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewaySku, NatGatewaySkuGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewaySku runs a test to see if a specific instance of NatGatewaySku round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewaySku(subject NatGatewaySku) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewaySku
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NatGatewaySku instances for property testing - lazily instantiated by NatGatewaySkuGenerator()
var natGatewaySkuGenerator gopter.Gen

// NatGatewaySkuGenerator returns a generator of NatGatewaySku instances for property testing.
func NatGatewaySkuGenerator() gopter.Gen {
	if natGatewaySkuGenerator != nil {
		return natGatewaySkuGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewaySku(generators)
	natGatewaySkuGenerator = gen.Struct(reflect.TypeOf(NatGatewaySku{}), generators)

	return natGatewaySkuGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewaySku is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewaySku(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(NatGatewaySku_Name_Standard))
}

func Test_NatGatewaySku_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from NatGatewaySku_STATUS to NatGatewaySku_STATUS via AssignProperties_To_NatGatewaySku_STATUS & AssignProperties_From_NatGatewaySku_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForNatGatewaySku_STATUS, NatGatewaySku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForNatGatewaySku_STATUS tests if a specific instance of NatGatewaySku_STATUS can be assigned to v1api20220701storage and back losslessly
func RunPropertyAssignmentTestForNatGatewaySku_STATUS(subject NatGatewaySku_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1api20220701s.NatGatewaySku_STATUS
	err := copied.AssignProperties_To_NatGatewaySku_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual NatGatewaySku_STATUS
	err = actual.AssignProperties_From_NatGatewaySku_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_NatGatewaySku_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NatGatewaySku_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNatGatewaySku_STATUS, NatGatewaySku_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNatGatewaySku_STATUS runs a test to see if a specific instance of NatGatewaySku_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNatGatewaySku_STATUS(subject NatGatewaySku_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NatGatewaySku_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of NatGatewaySku_STATUS instances for property testing - lazily instantiated by
// NatGatewaySku_STATUSGenerator()
var natGatewaySku_STATUSGenerator gopter.Gen

// NatGatewaySku_STATUSGenerator returns a generator of NatGatewaySku_STATUS instances for property testing.
func NatGatewaySku_STATUSGenerator() gopter.Gen {
	if natGatewaySku_STATUSGenerator != nil {
		return natGatewaySku_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNatGatewaySku_STATUS(generators)
	natGatewaySku_STATUSGenerator = gen.Struct(reflect.TypeOf(NatGatewaySku_STATUS{}), generators)

	return natGatewaySku_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNatGatewaySku_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNatGatewaySku_STATUS(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.OneConstOf(NatGatewaySku_Name_STATUS_Standard))
}
