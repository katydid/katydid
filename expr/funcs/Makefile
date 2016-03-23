# Copyright 2015 Walter Schulze
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

regenerate:
	go install github.com/katydid/katydid/expr/funcs/funcs-gen
	funcs-gen ./

nuke:
	rm -rf funcs/compare.gen.go
	rm -rf funcs/newfunc.gen.go
	rm -rf funcs/const.gen.go
	rm -rf funcs/elem.gen.go
	rm -rf funcs/length.gen.go
	rm -rf funcs/list.gen.go
	rm -rf funcs/print.gen.go