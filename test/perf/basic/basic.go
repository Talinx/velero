/*
Copyright 2021 the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package basic

import (
	"context"
	"fmt"
	"strings"
	"time"

	. "github.com/vmware-tanzu/velero/test"
	. "github.com/vmware-tanzu/velero/test/perf/test"
)

type BasicTest struct {
	TestCase
}

func (b *BasicTest) Init() error {
	b.TestCase.Init()
	b.Ctx, b.CtxCancel = context.WithTimeout(context.Background(), 1*time.Hour)
	b.CaseBaseName = "backuprestore"
	b.BackupName = "backup-" + b.CaseBaseName + "-" + b.UUIDgen
	b.RestoreName = "restore-" + b.CaseBaseName + "-" + b.UUIDgen

	b.BackupArgs = []string{
		"create", "--namespace", VeleroCfg.VeleroNamespace, "backup", b.BackupName,
		"--exclude-namespaces", strings.Join(*b.NSExcluded, ","),
		"--default-volumes-to-fs-backup",
		"--snapshot-volumes=false", "--wait",
	}

	b.RestoreArgs = []string{
		"create", "--namespace", VeleroCfg.VeleroNamespace, "restore", b.RestoreName,
		"--from-backup", b.BackupName, "--wait",
	}

	b.TestMsg = &TestMSG{
		Desc:      "Do backup and restore resources for performance test",
		FailedMSG: "Failed to backup and restore resources",
		Text:      fmt.Sprintf("Should backup and restore resources success"),
	}
	return nil
}
