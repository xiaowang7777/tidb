// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package privileges

import (
	mysql "github.com/pingcap/tidb/errno"
	"github.com/pingcap/tidb/util/dbterror"
)

// error definitions.
var (
	errInvalidPrivilegeType = dbterror.ClassPrivilege.NewStd(mysql.ErrInvalidPrivilegeType)
	ErrNonexistingGrant     = dbterror.ClassPrivilege.NewStd(mysql.ErrNonexistingGrant)
	errLoadPrivilege        = dbterror.ClassPrivilege.NewStd(mysql.ErrLoadPrivilege)
	errAccessDenied         = dbterror.ClassPrivilege.NewStd(mysql.ErrAccessDenied)
	errAccountHasBeenLocked = dbterror.ClassPrivilege.NewStd(mysql.ErrAccountHasBeenLocked)
)
