package migmigration

import (
	"context"
	"errors"
	migapi "github.com/konveyor/mig-controller/pkg/apis/migration/v1alpha1"
	dvmc "github.com/konveyor/mig-controller/pkg/controller/directvolumemigration"
	kapi "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func (t *Task) createDirectVolumeMigration() error {
	dvm := t.buildDirectVolumeMigration()
	if dvm == nil {
		return errors.New("failed to build directvolumeclaim list")
	}
	client, err := t.getDestinationClient()
	if err != nil {
		return err
	}
	err = client.Create(context.TODO(), dvm)
	return err

}

func (t *Task) buildDirectVolumeMigration() *migapi.DirectVolumeMigration {
	// Set correlation labels
	labels := t.Owner.GetCorrelationLabels()
	labels[DirectVolumeMigrationLabel] = t.UID()
	pvcList := t.getDirectVolumeClaimList()
	if pvcList == nil {
		return nil
	}
	dvm := &migapi.DirectVolumeMigration{
		ObjectMeta: metav1.ObjectMeta{
			Labels:       labels,
			GenerateName: t.Owner.GetName() + "-",
			Namespace:    migapi.VeleroNamespace,
		},
		Spec: migapi.DirectVolumeMigrationSpec{
			SrcMigClusterRef:            t.PlanResources.SrcMigCluster.GetObjectReference(),
			DestMigClusterRef:           t.PlanResources.DestMigCluster.GetObjectReference(),
			PersistentVolumeClaims:      *pvcList,
			CreateDestinationNamespaces: true,
		},
	}
	return dvm
}

func (t *Task) getDirectVolumeMigration() (*migapi.DirectVolumeMigration, error) {
	// Get correlation labels
	labels := t.Owner.GetCorrelationLabels()
	labels[DirectVolumeMigrationLabel] = t.UID()
	// Get DVM with label
	client, err := t.getDestinationClient()
	if err != nil {
		return nil, err
	}
	list := migapi.DirectVolumeMigrationList{}
	err = client.List(
		context.TODO(),
		k8sclient.MatchingLabels(labels),
		&list)
	if err != nil {
		return nil, err
	}
	if len(list.Items) > 0 {
		return &list.Items[0], nil
	}

	return nil, nil
}

func (t *Task) hasDirectVolumeMigrationCompleted(dvm *migapi.DirectVolumeMigration) (bool, []string) {
	completed := false
	reasons := []string{}
	progress := []string{}
	switch dvm.Status.Phase {
	case dvmc.Started:
		progress = append(progress, "DVM started")
	case dvmc.Completed:
		progress = append(progress, "DVM completed")
		completed = true
	case dvmc.MigrationFailed:
		reasons = append(reasons, "DVM FAILED")
		completed = true
	default:
		progress = append(progress, "DVM running...")
	}
	return completed, reasons
}

func (t *Task) getDirectVolumeClaimList() *[]migapi.PVCToMigrate {
	pvcList := []migapi.PVCToMigrate{}
	for _, pv := range t.PlanResources.MigPlan.Spec.PersistentVolumes.List {
		if pv.Selection.Action != migapi.PvCopyAction && pv.Selection.CopyMethod != migapi.PvFilesystemCopyMethod {
			continue
		}
		accessModes := pv.PVC.AccessModes
		if pv.Selection.AccessMode != nil {
			accessModes = []kapi.PersistentVolumeAccessMode{pv.Selection.AccessMode}
		}
		pvcList = append(pvcList, migapi.PVCToMigrate{
			Name:               pv.PVC.Name,
			Namespace:          pv.PVC.Namespace,
			TargetStorageClass: pv.Selection.StorageClass,
			TargetAccessModes:  accessModes,
		})
	}
	if len(pvcList) > 0 {
		return &pvcList
	}
	return nil
}
