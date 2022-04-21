// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200801preview

import (
	"encoding/json"
	v20200801ps "github.com/Azure/azure-service-operator/v2/api/authorization/v1beta20200801previewstorage"
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

func Test_RoleAssignment_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleAssignment to hub returns original",
		prop.ForAll(RunResourceConversionTestForRoleAssignment, RoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRoleAssignment tests if a specific instance of RoleAssignment round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRoleAssignment(subject RoleAssignment) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20200801ps.RoleAssignment
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RoleAssignment
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

func Test_RoleAssignment_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleAssignment to RoleAssignment via AssignPropertiesToRoleAssignment & AssignPropertiesFromRoleAssignment returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleAssignment, RoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleAssignment tests if a specific instance of RoleAssignment can be assigned to v1beta20200801previewstorage and back losslessly
func RunPropertyAssignmentTestForRoleAssignment(subject RoleAssignment) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200801ps.RoleAssignment
	err := copied.AssignPropertiesToRoleAssignment(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleAssignment
	err = actual.AssignPropertiesFromRoleAssignment(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleAssignment_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignment, RoleAssignmentGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignment runs a test to see if a specific instance of RoleAssignment round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignment(subject RoleAssignment) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment
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

// Generator of RoleAssignment instances for property testing - lazily instantiated by RoleAssignmentGenerator()
var roleAssignmentGenerator gopter.Gen

// RoleAssignmentGenerator returns a generator of RoleAssignment instances for property testing.
func RoleAssignmentGenerator() gopter.Gen {
	if roleAssignmentGenerator != nil {
		return roleAssignmentGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRoleAssignment(generators)
	roleAssignmentGenerator = gen.Struct(reflect.TypeOf(RoleAssignment{}), generators)

	return roleAssignmentGenerator
}

// AddRelatedPropertyGeneratorsForRoleAssignment is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRoleAssignment(gens map[string]gopter.Gen) {
	gens["Spec"] = RoleAssignmentsSpecGenerator()
	gens["Status"] = RoleAssignmentStatusGenerator()
}

func Test_RoleAssignment_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleAssignment_Status to RoleAssignment_Status via AssignPropertiesToRoleAssignmentStatus & AssignPropertiesFromRoleAssignmentStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleAssignmentStatus, RoleAssignmentStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleAssignmentStatus tests if a specific instance of RoleAssignment_Status can be assigned to v1beta20200801previewstorage and back losslessly
func RunPropertyAssignmentTestForRoleAssignmentStatus(subject RoleAssignment_Status) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200801ps.RoleAssignment_Status
	err := copied.AssignPropertiesToRoleAssignmentStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleAssignment_Status
	err = actual.AssignPropertiesFromRoleAssignmentStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleAssignment_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignment_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentStatus, RoleAssignmentStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentStatus runs a test to see if a specific instance of RoleAssignment_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentStatus(subject RoleAssignment_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignment_Status
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

// Generator of RoleAssignment_Status instances for property testing - lazily instantiated by
//RoleAssignmentStatusGenerator()
var roleAssignmentStatusGenerator gopter.Gen

// RoleAssignmentStatusGenerator returns a generator of RoleAssignment_Status instances for property testing.
func RoleAssignmentStatusGenerator() gopter.Gen {
	if roleAssignmentStatusGenerator != nil {
		return roleAssignmentStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentStatus(generators)
	roleAssignmentStatusGenerator = gen.Struct(reflect.TypeOf(RoleAssignment_Status{}), generators)

	return roleAssignmentStatusGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentStatus(gens map[string]gopter.Gen) {
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["CreatedOn"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentPropertiesStatusPrincipalTypeForeignGroup,
		RoleAssignmentPropertiesStatusPrincipalTypeGroup,
		RoleAssignmentPropertiesStatusPrincipalTypeServicePrincipal,
		RoleAssignmentPropertiesStatusPrincipalTypeUser))
	gens["RoleDefinitionId"] = gen.PtrOf(gen.AlphaString())
	gens["Scope"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedBy"] = gen.PtrOf(gen.AlphaString())
	gens["UpdatedOn"] = gen.PtrOf(gen.AlphaString())
}

func Test_RoleAssignments_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RoleAssignments_Spec to RoleAssignments_Spec via AssignPropertiesToRoleAssignmentsSpec & AssignPropertiesFromRoleAssignmentsSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRoleAssignmentsSpec, RoleAssignmentsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRoleAssignmentsSpec tests if a specific instance of RoleAssignments_Spec can be assigned to v1beta20200801previewstorage and back losslessly
func RunPropertyAssignmentTestForRoleAssignmentsSpec(subject RoleAssignments_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20200801ps.RoleAssignments_Spec
	err := copied.AssignPropertiesToRoleAssignmentsSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RoleAssignments_Spec
	err = actual.AssignPropertiesFromRoleAssignmentsSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual)
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_RoleAssignments_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoleAssignments_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoleAssignmentsSpec, RoleAssignmentsSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoleAssignmentsSpec runs a test to see if a specific instance of RoleAssignments_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRoleAssignmentsSpec(subject RoleAssignments_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoleAssignments_Spec
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

// Generator of RoleAssignments_Spec instances for property testing - lazily instantiated by
//RoleAssignmentsSpecGenerator()
var roleAssignmentsSpecGenerator gopter.Gen

// RoleAssignmentsSpecGenerator returns a generator of RoleAssignments_Spec instances for property testing.
func RoleAssignmentsSpecGenerator() gopter.Gen {
	if roleAssignmentsSpecGenerator != nil {
		return roleAssignmentsSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoleAssignmentsSpec(generators)
	roleAssignmentsSpecGenerator = gen.Struct(reflect.TypeOf(RoleAssignments_Spec{}), generators)

	return roleAssignmentsSpecGenerator
}

// AddIndependentPropertyGeneratorsForRoleAssignmentsSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoleAssignmentsSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Condition"] = gen.PtrOf(gen.AlphaString())
	gens["ConditionVersion"] = gen.PtrOf(gen.AlphaString())
	gens["DelegatedManagedIdentityResourceId"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalId"] = gen.PtrOf(gen.AlphaString())
	gens["PrincipalType"] = gen.PtrOf(gen.OneConstOf(
		RoleAssignmentPropertiesPrincipalTypeForeignGroup,
		RoleAssignmentPropertiesPrincipalTypeGroup,
		RoleAssignmentPropertiesPrincipalTypeServicePrincipal,
		RoleAssignmentPropertiesPrincipalTypeUser))
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}
