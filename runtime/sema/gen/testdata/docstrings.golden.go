// Code generated from testdata/docstrings.cdc. DO NOT EDIT.
/*
 * Cadence - The resource-oriented smart contract programming language
 *
 * Copyright Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package sema

import (
	"github.com/onflow/cadence/runtime/ast"
	"github.com/onflow/cadence/runtime/common"
)

const DocstringsTypeOwoFieldName = "owo"

var DocstringsTypeOwoFieldType = IntType

const DocstringsTypeOwoFieldDocString = `
This is a 1-line docstring.
`

const DocstringsTypeUwuFieldName = "uwu"

var DocstringsTypeUwuFieldType = &VariableSizedType{
	Type: IntType,
}

const DocstringsTypeUwuFieldDocString = `
This is a 2-line docstring.
This is the second line.
`

const DocstringsTypeNwnFunctionName = "nwn"

var DocstringsTypeNwnFunctionType = &FunctionType{
	Parameters: []Parameter{
		{
			TypeAnnotation: NewTypeAnnotation(IntType),
		},
	},
	ReturnTypeAnnotation: NewTypeAnnotation(
		&OptionalType{
			Type: StringType,
		},
	),
}

const DocstringsTypeNwnFunctionDocString = `
This is a 3-line docstring for a function.
This is the second line.
And the third line!
`

const DocstringsTypeWithBlanksFieldName = "withBlanks"

var DocstringsTypeWithBlanksFieldType = IntType

const DocstringsTypeWithBlanksFieldDocString = `
This is a multiline docstring.

There should be two newlines before this line!
`

const DocstringsTypeIsSmolBeanFunctionName = "isSmolBean"

var DocstringsTypeIsSmolBeanFunctionType = &FunctionType{
	ReturnTypeAnnotation: NewTypeAnnotation(
		BoolType,
	),
}

const DocstringsTypeIsSmolBeanFunctionDocString = `
The function ` + "`isSmolBean`" + ` has docstrings with backticks.
These should be handled accordingly.
`

const DocstringsTypeRunningOutOfIdeasFunctionName = "runningOutOfIdeas"

var DocstringsTypeRunningOutOfIdeasFunctionType = &FunctionType{
	ReturnTypeAnnotation: NewTypeAnnotation(
		&OptionalType{
			Type: UInt64Type,
		},
	),
}

const DocstringsTypeRunningOutOfIdeasFunctionDocString = `
A function with a docstring.
This docstring is ` + "`cool`" + ` because it has inline backticked expressions.
Look, I did it ` + "`again`" + `, wowie!!
`

const DocstringsTypeName = "Docstrings"

var DocstringsType = &SimpleType{
	Name:          DocstringsTypeName,
	QualifiedName: DocstringsTypeName,
	TypeID:        DocstringsTypeName,
	tag:           DocstringsTypeTag,
	IsResource:    false,
	Storable:      false,
	Equatable:     false,
	Exportable:    false,
	Importable:    false,
	Members: func(t *SimpleType) map[string]MemberResolver {
		return map[string]MemberResolver{
			DocstringsTypeOwoFieldName: {
				Kind: common.DeclarationKindField,
				Resolve: func(memoryGauge common.MemoryGauge,
					identifier string,
					targetRange ast.Range,
					report func(error)) *Member {

					return NewPublicConstantFieldMember(
						memoryGauge,
						t,
						identifier,
						DocstringsTypeOwoFieldType,
						DocstringsTypeOwoFieldDocString,
					)
				},
			},
			DocstringsTypeUwuFieldName: {
				Kind: common.DeclarationKindField,
				Resolve: func(memoryGauge common.MemoryGauge,
					identifier string,
					targetRange ast.Range,
					report func(error)) *Member {

					return NewPublicConstantFieldMember(
						memoryGauge,
						t,
						identifier,
						DocstringsTypeUwuFieldType,
						DocstringsTypeUwuFieldDocString,
					)
				},
			},
			DocstringsTypeNwnFunctionName: {
				Kind: common.DeclarationKindFunction,
				Resolve: func(memoryGauge common.MemoryGauge,
					identifier string,
					targetRange ast.Range,
					report func(error)) *Member {

					return NewPublicFunctionMember(
						memoryGauge,
						t,
						identifier,
						DocstringsTypeNwnFunctionType,
						DocstringsTypeNwnFunctionDocString,
					)
				},
			},
			DocstringsTypeWithBlanksFieldName: {
				Kind: common.DeclarationKindField,
				Resolve: func(memoryGauge common.MemoryGauge,
					identifier string,
					targetRange ast.Range,
					report func(error)) *Member {

					return NewPublicConstantFieldMember(
						memoryGauge,
						t,
						identifier,
						DocstringsTypeWithBlanksFieldType,
						DocstringsTypeWithBlanksFieldDocString,
					)
				},
			},
			DocstringsTypeIsSmolBeanFunctionName: {
				Kind: common.DeclarationKindFunction,
				Resolve: func(memoryGauge common.MemoryGauge,
					identifier string,
					targetRange ast.Range,
					report func(error)) *Member {

					return NewPublicFunctionMember(
						memoryGauge,
						t,
						identifier,
						DocstringsTypeIsSmolBeanFunctionType,
						DocstringsTypeIsSmolBeanFunctionDocString,
					)
				},
			},
			DocstringsTypeRunningOutOfIdeasFunctionName: {
				Kind: common.DeclarationKindFunction,
				Resolve: func(memoryGauge common.MemoryGauge,
					identifier string,
					targetRange ast.Range,
					report func(error)) *Member {

					return NewPublicFunctionMember(
						memoryGauge,
						t,
						identifier,
						DocstringsTypeRunningOutOfIdeasFunctionType,
						DocstringsTypeRunningOutOfIdeasFunctionDocString,
					)
				},
			},
		}
	},
}
