// Copyright 2016 Centimitr

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required` by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package namespace

import "fmt"

type ScopePrefix struct {
	names []string
	// the namespace it belongs to
	namespace *Namespace
}

func getPrefixNamesKey(names []string) string {
	return fmt.Sprint("%+v", names)
}

func stringArrayCopyConcat(a1, a2 []string) []string {
	a := []string{}
	for _, item := range a1 {
		a = append(a, item)
	}
	for _, item := range a2 {
		a = append(a, item)
	}
	return a
}

func (s *ScopePrefix) Extend(extnames ...string) (bool, ScopePrefix) {
	names := stringArrayCopyConcat(s.names, extnames)
	if ok, prefix := s.namespace.NewPrefix(names...); ok {
		return true, prefix
	}
	return false, ScopePrefix{}
}

func (s *ScopePrefix) Apply(extnames ...string) (bool, Scope) {
	names := stringArrayCopyConcat(s.names, extnames)
	prefix := s.namespace.ruleOfPrefixConcat(names...)
	if ok, scope := s.namespace.Apply(prefix); ok {
		return true, scope
	}
	return false, Scope{}
}
