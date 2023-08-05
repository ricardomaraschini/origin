package legacynodeinvariants

import (
	"context"
	"time"

	"github.com/openshift/origin/pkg/invariantlibrary/platformidentification"

	"github.com/openshift/origin/pkg/invariants"
	"github.com/openshift/origin/pkg/monitor/monitorapi"
	"github.com/openshift/origin/pkg/test/ginkgo/junitapi"
	"k8s.io/client-go/rest"
)

type legacyInvariantTests struct {
}

func NewLegacyTests() invariants.InvariantTest {
	return &legacyInvariantTests{}
}

func (w *legacyInvariantTests) StartCollection(ctx context.Context, adminRESTConfig *rest.Config, recorder monitorapi.RecorderWriter) error {
	return nil
}

func (w *legacyInvariantTests) CollectData(ctx context.Context, storageDir string, beginning, end time.Time) (monitorapi.Intervals, []*junitapi.JUnitTestCase, error) {
	return nil, nil, nil
}

func (*legacyInvariantTests) ConstructComputedIntervals(ctx context.Context, startingIntervals monitorapi.Intervals, recordedResources monitorapi.ResourcesMap, beginning, end time.Time) (monitorapi.Intervals, error) {
	return nil, nil
}

func (w *legacyInvariantTests) EvaluateTestsFromConstructedIntervals(ctx context.Context, finalIntervals monitorapi.Intervals) ([]*junitapi.JUnitTestCase, error) {
	junits := []*junitapi.JUnitTestCase{}
	junits = append(junits, testContainerFailures(finalIntervals)...)
	junits = append(junits, testDeleteGracePeriodZero(finalIntervals)...)
	junits = append(junits, testKubeApiserverProcessOverlap(finalIntervals)...)
	junits = append(junits, testKubeAPIServerGracefulTermination(finalIntervals)...)
	junits = append(junits, testKubeletToAPIServerGracefulTermination(finalIntervals)...)
	junits = append(junits, testPodTransitions(finalIntervals)...)
	junits = append(junits, testErrImagePullConnTimeoutOpenShiftNamespaces(finalIntervals)...)
	junits = append(junits, testErrImagePullConnTimeout(finalIntervals)...)
	junits = append(junits, testErrImagePullQPSExceededOpenShiftNamespaces(finalIntervals)...)
	junits = append(junits, testErrImagePullQPSExceeded(finalIntervals)...)
	junits = append(junits, testErrImagePullManifestUnknownOpenShiftNamespaces(finalIntervals)...)
	junits = append(junits, testErrImagePullManifestUnknown(finalIntervals)...)
	junits = append(junits, testErrImagePullGenericOpenShiftNamespaces(finalIntervals)...)
	junits = append(junits, testErrImagePullGeneric(finalIntervals)...)
	junits = append(junits, testFailedToDeleteCGroupsPath(finalIntervals)...)
	junits = append(junits, testAnonymousCertConnectionFailure(finalIntervals)...)
	junits = append(junits, testHttpConnectionLost(finalIntervals)...)
	junits = append(junits, testErrImagePullUnrecognizedSignatureFormat(finalIntervals)...)
	junits = append(junits, testLeaseUpdateError(finalIntervals)...)
	junits = append(junits, testSystemDTimeout(finalIntervals)...)
	junits = append(junits, testNodeHasNoDiskPressure(finalIntervals)...)
	junits = append(junits, testNodeHasSufficientMemory(finalIntervals)...)
	junits = append(junits, testNodeHasSufficientPID(finalIntervals)...)
	junits = append(junits, testErrorReconcilingNode(finalIntervals)...)
	junits = append(junits, testBackoffPullingRegistryRedhatImage(finalIntervals)...)
	junits = append(junits, testBackoffStartingFailedContainer(finalIntervals)...)
	junits = append(junits, testConfigOperatorReadinessProbe(finalIntervals)...)
	junits = append(junits, testConfigOperatorProbeErrorReadinessProbe(finalIntervals)...)
	junits = append(junits, testConfigOperatorProbeErrorLivenessProbe(finalIntervals)...)
	junits = append(junits, testMasterNodesUpdated(finalIntervals)...)
	junits = append(junits, testMarketplaceStartupProbeFailure(finalIntervals)...)
	junits = append(junits, testFailedScheduling(finalIntervals)...)

	isUpgrade := platformidentification.DidUpgradeHappenDuringCollection(finalIntervals, time.Time{}, time.Time{})
	if isUpgrade {
		junits = append(junits, testNodeUpgradeTransitions(finalIntervals)...)
	}

	return junits, nil
}

func (*legacyInvariantTests) WriteContentToStorage(ctx context.Context, storageDir, timeSuffix string, finalIntervals monitorapi.Intervals, finalResourceState monitorapi.ResourcesMap) error {
	return nil
}

func (*legacyInvariantTests) Cleanup(ctx context.Context) error {
	return nil
}
