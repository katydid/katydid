//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package tests

func init() {
	BenchValidateProtoNum("ContextPerson", ContextPerson, &Person{})
	BenchValidateProtoNum("ListIndexAddress", ListIndexAddressPerson, &Person{})
	BenchValidateProtoNum("NilName", NilNamePerson, &Person{})
	BenchValidateProtoNum("LenName", LenNamePerson, &Person{})
	BenchValidateProtoNum("EmptyOrNil", EmptyOrNilPerson, &Person{})
	BenchValidateProtoNum("IncorrectNotName", NaiveNotNamePerson, &Person{})
	BenchValidateProtoNum("CorrectNotName", ProperNotNamePerson, &Person{})
	BenchValidateProtoNum("AndNameTelephone", AndNameTelephonePerson, &Person{})
	BenchValidateProtoNum("OrNameTelephone", OrNameTelephonePerson, &Person{})
}
