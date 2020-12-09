package directvolumemigration

var phaseDescriptions = map[string]string{
	Created:                         "DVM CR has been created",
	Started:                         "DVM Controller is configuring DVM CR",
	Prepare:                         "DVM Controller is preparing the environment for volume migration. Currently it is non-operational",
	CreateDestinationNamespaces:     "Creating target namespaces",
	DestinationNamespacesCreated:    "Checking if the target namespaces have been created, currently it is non-operational",
	CreateDestinationPVCs:           "Creating PVCs in the target namespaces",
	DestinationPVCsCreated:          "Checking whether the created PVCs are bound",
	CreateRsyncRoute:                "Creating one route for each namespace for Rsync on the target cluster",
	CreateRsyncConfig:               "Creating a config map and secrets on both the source and target clusters for Rsync configuration",
	CreateStunnelConfig:             "Creating a config map and secrets for Stunnel to connect to Rsync on the source and target clusters",
	CreatePVProgressCRs:             "Creating a Direct Volume Migration Progress CR to get progress percent and transfer rate",
	CreateRsyncTransferPods:         "Creating Rsync daemon pods on the target cluster",
	WaitForRsyncTransferPodsRunning: "Waiting for the Rsync daemon pod to run",
	CreateStunnelClientPods:         "Creating Stunnel client pods on the source cluster",
	WaitForStunnelClientPodsRunning: "Waiting for the Stunnel client pods to run",
	CreateRsyncClientPods:           "Creating Rsync client pods",
	WaitForRsyncClientPodsCompleted: "Waiting for the Rsync client pods to be completed",
	DeleteRsyncResources:            "Deleting resources created by this migration",
	MigrationFailed:                 "The migration attempt failed, please see errors for more details",
	Completed:                       "Complete",
}
