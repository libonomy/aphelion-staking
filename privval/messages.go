package privval

import (
	amino "github.com/evdatsion/go-amino"
	"github.com/libonomy/aphelion-staking/crypto"
	"github.com/libonomy/aphelion-staking/types"
)

// SignerMessage is sent between Signer Clients and Servers.
type SignerMessage interface{}

func RegisterRemoteSignerMsg(cdc *amino.Codec) {
	cdc.RegisterInterface((*SignerMessage)(nil), nil)
	cdc.RegisterConcrete(&PubKeyRequest{}, "aphelion/remotesigner/PubKeyRequest", nil)
	cdc.RegisterConcrete(&PubKeyResponse{}, "aphelion/remotesigner/PubKeyResponse", nil)
	cdc.RegisterConcrete(&SignVoteRequest{}, "aphelion/remotesigner/SignVoteRequest", nil)
	cdc.RegisterConcrete(&SignedVoteResponse{}, "aphelion/remotesigner/SignedVoteResponse", nil)
	cdc.RegisterConcrete(&SignProposalRequest{}, "aphelion/remotesigner/SignProposalRequest", nil)
	cdc.RegisterConcrete(&SignedProposalResponse{}, "aphelion/remotesigner/SignedProposalResponse", nil)

	cdc.RegisterConcrete(&PingRequest{}, "aphelion/remotesigner/PingRequest", nil)
	cdc.RegisterConcrete(&PingResponse{}, "aphelion/remotesigner/PingResponse", nil)
}

// TODO: Add ChainIDRequest

// PubKeyRequest requests the consensus public key from the remote signer.
type PubKeyRequest struct{}

// PubKeyResponse is a response message containing the public key.
type PubKeyResponse struct {
	PubKey crypto.PubKey
	Error  *RemoteSignerError
}

// SignVoteRequest is a request to sign a vote
type SignVoteRequest struct {
	Vote *types.Vote
}

// SignedVoteResponse is a response containing a signed vote or an error
type SignedVoteResponse struct {
	Vote  *types.Vote
	Error *RemoteSignerError
}

// SignProposalRequest is a request to sign a proposal
type SignProposalRequest struct {
	Proposal *types.Proposal
}

// SignedProposalResponse is response containing a signed proposal or an error
type SignedProposalResponse struct {
	Proposal *types.Proposal
	Error    *RemoteSignerError
}

// PingRequest is a request to confirm that the connection is alive.
type PingRequest struct {
}

// PingResponse is a response to confirm that the connection is alive.
type PingResponse struct {
}
