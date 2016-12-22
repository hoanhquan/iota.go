package giota

import (
	"testing"
)

func TestNewAPI(t *testing.T) {
	_, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}
}

func TestAPIGetNodeInfo(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	resp, err := api.GetNodeInfo()
	if err != nil {
		t.Errorf("GetNodeInfo() expected err to be nil but got %v", err)
	}

	if resp.AppName == "" {
		t.Errorf("GetNodeInfo() returned invalid response: %#v", resp)
	}
}

func TestAPIGetNeighbors(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	_, err = api.GetNeighbors()
	if err != nil {
		t.Errorf("GetNeighbors() expected err to be nil but got %v", err)
	}
}

func TestAPIAddNeighbors(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	anr := &AddNeighborsRequest{URIS: []string{"udp://127.0.0.1:14265/"}}
	resp, err := api.AddNeighbors(anr)
	if err != nil {
		t.Errorf("AddNeighbors([]) expected err to be nil but got %v", err)
	} else if resp.AddedNeighbors != 1 {
		t.Errorf("AddNeighbors([]) expected to add %d got %d", 0, resp.AddedNeighbors)
	}
}

func TestAPIRemoveNeighbors(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	anr := &RemoveNeighborsRequest{URIS: []string{"udp://127.0.0.1:14265/"}}
	resp, err := api.RemoveNeighbors(anr)
	if err != nil {
		t.Errorf("RemoveNeighbors([]) expected err to be nil but got %v", err)
	} else if resp.RemovedNeighbors != 1 {
		t.Errorf("RemoveNeighbors([]) expected to remove %d got %d", 0, resp.RemovedNeighbors)
	}
}

func TestAPIGetTips(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	resp, err := api.GetTips()
	if err != nil {
		t.Errorf("GetTips() expected err to be nil but got %v", err)
	}

	if len(resp.Hashes) < 1 {
		t.Errorf("GetTips() returned less than one tip")
	}
}

func TestAPIFindTransactions(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	ftr := &FindTransactionsRequest{Bundles: &[]string{}}
	resp, err := api.FindTransactions(ftr)
	if err != nil {
		t.Errorf("FindTransactions([]) expected err to be nil but got %v", err)
	}
	t.Logf("FindTransactions() = %#v", resp)
}

func TestAPIGetTrytes(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	anr := &GetTrytesRequest{Hashes: []string{}}
	resp, err := api.GetTrytes(anr)
	if err != nil {
		t.Errorf("GetTrytes([]) expected err to be nil but got %v", err)
	}
	t.Logf("GetTrytes() = %#v", resp)
}

func TestAPIGetInclusionStates(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	anr := &GetInclusionStatesRequest{Transactions: []string{}, Tips: []string{}}
	resp, err := api.GetInclusionStates(anr)
	if err != nil {
		t.Errorf("GetInclusionStates([]) expected err to be nil but got %v", err)
	}
	t.Logf("GetInclusionStates() = %#v", resp)
}

func TestAPIGetTransactionsToApprove(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	anr := &GetTransactionsToApproveRequest{}
	resp, err := api.GetTransactionsToApprove(anr)
	if err != nil {
		t.Errorf("GetTransactionsToApprove() expected err to be nil but got %v", err)
	}
	if resp.BranchTransaction == "" || resp.TrunkTransaction == "" {
		t.Errorf("GetTransactionsToApprove() return empty branch and/or trunk transactions\n%#v", resp)
	}
}

func TestAPIInterruptAttachingToTangle(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	resp, err := api.InterruptAttachingToTangle()
	if err != nil {
		t.Errorf("InterruptAttachingToTangle() expected err to be nil but got %v", err)
	}
	t.Logf("InterruptAttachingToTangle() = %#v", resp)
}

// XXX: The following tests are failing because I'd rather not just
//      constantly attach/broadcast/store the same transaction
func TestAPIAttachToTangle(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	anr := &AttachToTangleRequest{}
	resp, err := api.AttachToTangle(anr)
	if err != nil {
		t.Errorf("AttachToTangle([]) expected err to be nil but got %v", err)
	}
	t.Logf("AttachToTangle() = %#v", resp)
}

func TestAPIBroadcastTransactions(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	anr := &BroadcastTransactionsRequest{}
	resp, err := api.BroadcastTransactions(anr)
	if err != nil {
		t.Errorf("BroadcastTransactions() expected err to be nil but got %v", err)
	}
	t.Logf("BroadcastTransactions() = %#v", resp)
}

func TestAPIStoreTransactions(t *testing.T) {
	api, err := NewAPI("", nil)
	if err != nil {
		t.Errorf("NewAPI(%q, nil) expected err to be nil but got %v", err)
	}

	anr := &StoreTransactionsRequest{}
	resp, err := api.StoreTransactions(anr)
	if err != nil {
		t.Errorf("StoreTransactions() expected err to be nil but got %v", err)
	}
	t.Logf("StoreTransactions() = %#v", resp)
}