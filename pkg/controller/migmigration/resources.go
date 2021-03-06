/*
Copyright 2019 Red Hat Inc.

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

package migmigration

import (
	"fmt"

	migapi "github.com/fusor/mig-controller/pkg/apis/migration/v1alpha1"
	velerov1 "github.com/heptio/velero/pkg/apis/velero/v1"
)

// reconcileResources holds the data needed for MigMigration to reconcile.
// At the beginning of a reconcile, this data will be compiled by fetching
// information from each cluster involved in the migration.
type reconcileResources struct {
	migMigration   *migapi.MigMigration
	migPlan        *migapi.MigPlan
	migAssets      *migapi.MigAssetCollection
	srcMigCluster  *migapi.MigCluster
	destMigCluster *migapi.MigCluster

	srcBackup   *velerov1.Backup
	destRestore *velerov1.Restore
}

// getReconcileResources puts together a struct with all resources needed to perform a MigMigration reconcile.
// The slots for Velero Backup and Restore are left empty for later use, as they can't be filled  yet.
func (r *ReconcileMigMigration) getReconcileResources(migMigration *migapi.MigMigration) (*reconcileResources, error) {
	resources := &reconcileResources{}

	// MigMigration
	resources.migMigration = migMigration

	// MigPlan
	migPlan, err := migMigration.GetPlan(r.Client)
	if err != nil {
		log.Info(fmt.Sprintf("[%s] Failed to GET MigPlan referenced by MigMigration [%s/%s]",
			logPrefix, migMigration.Namespace, migMigration.Name))
		return nil, err
	}
	resources.migPlan = migPlan

	// MigAssetCollection
	migAssets, err := migPlan.GetAssetCollection(r.Client)
	if err != nil {
		log.Info(fmt.Sprintf("[%s] Failed to GET MigAssetCollection referenced by MigPlan [%s/%s]",
			logPrefix, migPlan.Namespace, migPlan.Name))
		return nil, err
	}
	resources.migAssets = migAssets

	// SrcMigCluster
	srcMigCluster, err := migPlan.GetSourceCluster(r.Client)
	if err != nil {
		log.Info(fmt.Sprintf("[%s] Failed to GET SrcMigCluster referenced by MigPlan [%s/%s]",
			logPrefix, migPlan.Namespace, migPlan.Name))
		return nil, err
	}
	resources.srcMigCluster = srcMigCluster

	// DestMigCluster
	destMigCluster, err := migPlan.GetDestinationCluster(r.Client)
	if err != nil {
		log.Info(fmt.Sprintf("[%s] Failed to GET DestMigCluster referenced by MigPlan [%s/%s]",
			logPrefix, migPlan.Namespace, migPlan.Name))
		return nil, err
	}
	resources.destMigCluster = destMigCluster

	return resources, nil
}
