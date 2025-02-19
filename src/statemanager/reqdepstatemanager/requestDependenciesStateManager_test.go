package reqdepstatemanager

import (
	"testing"

	"github.com/nianticlabs/modron/src/pb"
)

func TestSimpleSateManager(t *testing.T) {

	stateManager, err := New()
	if err != nil {
		t.Error(err)
	}

	if state := stateManager.GetCollectState("collect-id-1"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}
	if state := stateManager.GetScanState("scan-id-1"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	collecting := stateManager.AddCollect("collect-id-1", []string{"p1", "p2", "p3"})

	if len(collecting) != 3 {
		t.Errorf("the state manager should be collection 3 resource groups, but is collecting : %v", collecting)
	}

	scanning := stateManager.AddScan("scan-id-1", []string{"p1", "p2", "p3"})

	if len(scanning) != 3 {
		t.Errorf("the state manager should be scanning 3 resource groups, but is collecting : %v", scanning)
	}

	if state := stateManager.GetCollectState("collect-id-1"); state != pb.RequestStatus_RUNNING {
		t.Errorf("expected RUNNING state, got %v", state)
	}
	if state := stateManager.GetCollectState("collect-id-2"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	if state := stateManager.GetScanState("scan-id-1"); state != pb.RequestStatus_RUNNING {
		t.Errorf("expected RUNNING state, got %v", state)
	}
	if state := stateManager.GetScanState("scan-id-2"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	stateManager.EndCollect("collect-id-2", []string{"p1", "p2", "p3"})
	stateManager.EndScan("scan-id-2", []string{"p1", "p2", "p3"})

	stateManager.EndCollect("collect-id-1", []string{"p1", "p2", "p3"})
	stateManager.EndScan("scan-id-1", []string{"p1", "p2", "p3"})

	if state := stateManager.GetCollectState("collect-id-1"); state != pb.RequestStatus_DONE {
		t.Errorf("expected DONE state, got %v", state)
	}
	if state := stateManager.GetCollectState("collect-id-2"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	if state := stateManager.GetScanState("scan-id-1"); state != pb.RequestStatus_DONE {
		t.Errorf("expected DONE state, got %v", state)
	}
	if state := stateManager.GetScanState("scan-id-2"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

}

func TestDepSateManager(t *testing.T) {

	stateManager, err := New()
	if err != nil {
		t.Error(err)
	}

	if state := stateManager.GetCollectState("collect-id-1"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}
	if state := stateManager.GetScanState("scan-id-1"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	if collecting := stateManager.AddCollect("collect-id-1", []string{"p1", "p2", "p3"}); len(collecting) != 3 {
		t.Errorf("the state manager should be collection 3 resource groups, but is collecting : %v", collecting)
	}

	if collecting := stateManager.AddCollect("collect-id-3", []string{"p0", "p1", "p2"}); len(collecting) != 1 {
		t.Errorf("the state manager should be collection 1 resource groups, but is collecting : %v", collecting)
	}

	if scanning := stateManager.AddScan("scan-id-1", []string{"p1", "p2", "p3"}); len(scanning) != 3 {
		t.Errorf("the state manager should be scanning 3 resource groups, but is collecting : %v", scanning)
	}
	if scanning := stateManager.AddScan("scan-id-3", []string{"p1", "p2"}); len(scanning) != 0 {
		t.Errorf("the state manager should be scanning 0 resource groups, but is collecting : %v", scanning)
	}

	if state := stateManager.GetCollectState("collect-id-1"); state != pb.RequestStatus_RUNNING {
		t.Errorf("expected RUNNING state, got %v", state)
	}
	if state := stateManager.GetCollectState("collect-id-3"); state != pb.RequestStatus_RUNNING {
		t.Errorf("expected RUNNING state, got %v", state)
	}
	if state := stateManager.GetCollectState("collect-id-2"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	if state := stateManager.GetScanState("scan-id-1"); state != pb.RequestStatus_RUNNING {
		t.Errorf("expected RUNNING state, got %v", state)
	}
	if state := stateManager.GetScanState("scan-id-3"); state != pb.RequestStatus_RUNNING {
		t.Errorf("expected RUNNING state, got %v", state)
	}
	if state := stateManager.GetScanState("scan-id-2"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	stateManager.EndCollect("collect-id-2", []string{"p1", "p2", "p3"})
	stateManager.EndScan("scan-id-2", []string{"p1", "p2", "p3"})

	stateManager.EndCollect("collect-id-1", []string{"p1", "p2", "p3"})
	stateManager.EndScan("scan-id-1", []string{"p1", "p2", "p3"})

	if state := stateManager.GetCollectState("collect-id-1"); state != pb.RequestStatus_DONE {
		t.Errorf("expected DONE state, got %v", state)
	}
	if state := stateManager.GetCollectState("collect-id-2"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	if state := stateManager.GetScanState("scan-id-1"); state != pb.RequestStatus_DONE {
		t.Errorf("expected DONE state, got %v", state)
	}
	if state := stateManager.GetScanState("scan-id-3"); state != pb.RequestStatus_DONE {
		t.Errorf("expected DONE state, got %v", state)
	}
	if state := stateManager.GetScanState("scan-id-2"); state != pb.RequestStatus_UNKNOWN {
		t.Errorf("expected UNKNOWN state, got %v", state)
	}

	if state := stateManager.GetCollectState("collect-id-3"); state != pb.RequestStatus_RUNNING {
		t.Errorf("expected RUNNING state, got %v", state)
	}
	stateManager.EndCollect("collect-id-3", []string{"p0", "p1", "p2"})
	if state := stateManager.GetCollectState("collect-id-3"); state != pb.RequestStatus_DONE {
		t.Errorf("expected DONE state, got %v", state)
	}

}
