package proofs

import (
	"github.com/filecoin-project/go-filecoin/proofs"
	"github.com/filecoin-project/go-filecoin/proofs/sectorbuilder"
	"github.com/filecoin-project/go-filecoin/types"
)

// Provides the sector builder for proofs.
type sectorBuilderer interface {
	SectorBuilder() sectorbuilder.SectorBuilder
}

// Proofs is plumbing for proof calculation and verification.
type Proofs struct {
	deps sectorBuilderer
}

// NewProofs returns new proofs plumbing.
func NewProofs(deps sectorBuilderer) *Proofs {
	return &Proofs{deps}
}

// CalculatePost calls the sector builder to compute a proof.
func (p *Proofs) CalculatePost(sortedCommRs proofs.SortedCommRs, seed types.PoStChallengeSeed) ([]types.PoStProof, []uint64, error) {
	req := sectorbuilder.GeneratePoStRequest{
		SortedCommRs:  sortedCommRs,
		ChallengeSeed: seed,
	}
	res, err := p.deps.SectorBuilder().GeneratePoSt(req)
	if err != nil {
		return nil, nil, err
	}
	return res.Proofs, res.Faults, nil
}
