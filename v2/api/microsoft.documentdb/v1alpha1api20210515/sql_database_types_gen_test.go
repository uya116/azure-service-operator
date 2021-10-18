// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.documentdb/v1alpha1api20210515storage"
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

func Test_SqlDatabase_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabase to SqlDatabase via AssignPropertiesToSqlDatabase & AssignPropertiesFromSqlDatabase returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabase, SqlDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabase tests if a specific instance of SqlDatabase can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabase(subject SqlDatabase) string {
	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlDatabase
	err := subject.AssignPropertiesToSqlDatabase(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabase
	err = actual.AssignPropertiesFromSqlDatabase(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabase, SqlDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabase runs a test to see if a specific instance of SqlDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabase(subject SqlDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabase
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

// Generator of SqlDatabase instances for property testing - lazily instantiated by SqlDatabaseGenerator()
var sqlDatabaseGenerator gopter.Gen

// SqlDatabaseGenerator returns a generator of SqlDatabase instances for property testing.
func SqlDatabaseGenerator() gopter.Gen {
	if sqlDatabaseGenerator != nil {
		return sqlDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabase(generators)
	sqlDatabaseGenerator = gen.Struct(reflect.TypeOf(SqlDatabase{}), generators)

	return sqlDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = DatabaseAccountsSqlDatabasesSpecGenerator()
	gens["Status"] = SqlDatabaseGetResultsStatusGenerator()
}

func Test_DatabaseAccountsSqlDatabases_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from DatabaseAccountsSqlDatabases_Spec to DatabaseAccountsSqlDatabases_Spec via AssignPropertiesToDatabaseAccountsSqlDatabasesSpec & AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec returns original",
		prop.ForAll(RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesSpec, DatabaseAccountsSqlDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesSpec tests if a specific instance of DatabaseAccountsSqlDatabases_Spec can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForDatabaseAccountsSqlDatabasesSpec(subject DatabaseAccountsSqlDatabases_Spec) string {
	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.DatabaseAccountsSqlDatabases_Spec
	err := subject.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual DatabaseAccountsSqlDatabases_Spec
	err = actual.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_DatabaseAccountsSqlDatabases_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseAccountsSqlDatabases_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseAccountsSqlDatabasesSpec, DatabaseAccountsSqlDatabasesSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseAccountsSqlDatabasesSpec runs a test to see if a specific instance of DatabaseAccountsSqlDatabases_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseAccountsSqlDatabasesSpec(subject DatabaseAccountsSqlDatabases_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseAccountsSqlDatabases_Spec
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

// Generator of DatabaseAccountsSqlDatabases_Spec instances for property testing - lazily instantiated by
//DatabaseAccountsSqlDatabasesSpecGenerator()
var databaseAccountsSqlDatabasesSpecGenerator gopter.Gen

// DatabaseAccountsSqlDatabasesSpecGenerator returns a generator of DatabaseAccountsSqlDatabases_Spec instances for property testing.
// We first initialize databaseAccountsSqlDatabasesSpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseAccountsSqlDatabasesSpecGenerator() gopter.Gen {
	if databaseAccountsSqlDatabasesSpecGenerator != nil {
		return databaseAccountsSqlDatabasesSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(generators)
	databaseAccountsSqlDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabases_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(generators)
	AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(generators)
	databaseAccountsSqlDatabasesSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseAccountsSqlDatabases_Spec{}), generators)

	return databaseAccountsSqlDatabasesSpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseAccountsSqlDatabasesSpec(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(CreateUpdateOptionsGenerator())
	gens["Resource"] = SqlDatabaseResourceGenerator()
}

func Test_SqlDatabaseGetResults_Status_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseGetResults_Status to SqlDatabaseGetResults_Status via AssignPropertiesToSqlDatabaseGetResultsStatus & AssignPropertiesFromSqlDatabaseGetResultsStatus returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseGetResultsStatus, SqlDatabaseGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseGetResultsStatus tests if a specific instance of SqlDatabaseGetResults_Status can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseGetResultsStatus(subject SqlDatabaseGetResults_Status) string {
	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlDatabaseGetResults_Status
	err := subject.AssignPropertiesToSqlDatabaseGetResultsStatus(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseGetResults_Status
	err = actual.AssignPropertiesFromSqlDatabaseGetResultsStatus(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseGetResults_Status_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseGetResults_Status via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseGetResultsStatus, SqlDatabaseGetResultsStatusGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseGetResultsStatus runs a test to see if a specific instance of SqlDatabaseGetResults_Status round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseGetResultsStatus(subject SqlDatabaseGetResults_Status) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseGetResults_Status
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

// Generator of SqlDatabaseGetResults_Status instances for property testing - lazily instantiated by
//SqlDatabaseGetResultsStatusGenerator()
var sqlDatabaseGetResultsStatusGenerator gopter.Gen

// SqlDatabaseGetResultsStatusGenerator returns a generator of SqlDatabaseGetResults_Status instances for property testing.
// We first initialize sqlDatabaseGetResultsStatusGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseGetResultsStatusGenerator() gopter.Gen {
	if sqlDatabaseGetResultsStatusGenerator != nil {
		return sqlDatabaseGetResultsStatusGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetResultsStatus(generators)
	sqlDatabaseGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetResults_Status{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetResultsStatus(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseGetResultsStatus(generators)
	sqlDatabaseGetResultsStatusGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetResults_Status{}), generators)

	return sqlDatabaseGetResultsStatusGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseGetResultsStatus is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseGetResultsStatus is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseGetResultsStatus(gens map[string]gopter.Gen) {
	gens["Options"] = gen.PtrOf(OptionsResourceStatusGenerator())
	gens["Resource"] = gen.PtrOf(SqlDatabaseGetPropertiesStatusResourceGenerator())
}

func Test_SqlDatabaseGetProperties_Status_Resource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseGetProperties_Status_Resource to SqlDatabaseGetProperties_Status_Resource via AssignPropertiesToSqlDatabaseGetPropertiesStatusResource & AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseGetPropertiesStatusResource, SqlDatabaseGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseGetPropertiesStatusResource tests if a specific instance of SqlDatabaseGetProperties_Status_Resource can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseGetPropertiesStatusResource(subject SqlDatabaseGetProperties_Status_Resource) string {
	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlDatabaseGetProperties_Status_Resource
	err := subject.AssignPropertiesToSqlDatabaseGetPropertiesStatusResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseGetProperties_Status_Resource
	err = actual.AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseGetProperties_Status_Resource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseGetProperties_Status_Resource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseGetPropertiesStatusResource, SqlDatabaseGetPropertiesStatusResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseGetPropertiesStatusResource runs a test to see if a specific instance of SqlDatabaseGetProperties_Status_Resource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseGetPropertiesStatusResource(subject SqlDatabaseGetProperties_Status_Resource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseGetProperties_Status_Resource
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

// Generator of SqlDatabaseGetProperties_Status_Resource instances for property testing - lazily instantiated by
//SqlDatabaseGetPropertiesStatusResourceGenerator()
var sqlDatabaseGetPropertiesStatusResourceGenerator gopter.Gen

// SqlDatabaseGetPropertiesStatusResourceGenerator returns a generator of SqlDatabaseGetProperties_Status_Resource instances for property testing.
func SqlDatabaseGetPropertiesStatusResourceGenerator() gopter.Gen {
	if sqlDatabaseGetPropertiesStatusResourceGenerator != nil {
		return sqlDatabaseGetPropertiesStatusResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseGetPropertiesStatusResource(generators)
	sqlDatabaseGetPropertiesStatusResourceGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseGetProperties_Status_Resource{}), generators)

	return sqlDatabaseGetPropertiesStatusResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseGetPropertiesStatusResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseGetPropertiesStatusResource(gens map[string]gopter.Gen) {
	gens["Colls"] = gen.PtrOf(gen.AlphaString())
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.AlphaString()
	gens["Rid"] = gen.PtrOf(gen.AlphaString())
	gens["Ts"] = gen.PtrOf(gen.Float64())
	gens["Users"] = gen.PtrOf(gen.AlphaString())
}

func Test_SqlDatabaseResource_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseResource to SqlDatabaseResource via AssignPropertiesToSqlDatabaseResource & AssignPropertiesFromSqlDatabaseResource returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseResource, SqlDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseResource tests if a specific instance of SqlDatabaseResource can be assigned to v1alpha1api20210515storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseResource(subject SqlDatabaseResource) string {
	// Use AssignPropertiesTo() for the first stage of conversion
	var other v1alpha1api20210515storage.SqlDatabaseResource
	err := subject.AssignPropertiesToSqlDatabaseResource(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseResource
	err = actual.AssignPropertiesFromSqlDatabaseResource(&other)
	if err != nil {
		return err.Error()
	}

	//Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SqlDatabaseResource_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseResource via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseResource, SqlDatabaseResourceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseResource runs a test to see if a specific instance of SqlDatabaseResource round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseResource(subject SqlDatabaseResource) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseResource
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

// Generator of SqlDatabaseResource instances for property testing - lazily instantiated by
//SqlDatabaseResourceGenerator()
var sqlDatabaseResourceGenerator gopter.Gen

// SqlDatabaseResourceGenerator returns a generator of SqlDatabaseResource instances for property testing.
func SqlDatabaseResourceGenerator() gopter.Gen {
	if sqlDatabaseResourceGenerator != nil {
		return sqlDatabaseResourceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseResource(generators)
	sqlDatabaseResourceGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseResource{}), generators)

	return sqlDatabaseResourceGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseResource is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseResource(gens map[string]gopter.Gen) {
	gens["Id"] = gen.AlphaString()
}