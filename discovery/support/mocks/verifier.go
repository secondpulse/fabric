// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/hyperledger/fabric/discovery/support/acl"
	"github.com/hyperledger/fabric/gossip/api"
	"github.com/hyperledger/fabric/gossip/common"
)

type Verifier struct {
	VerifyByChannelStub        func(chainID common.ChainID, peerIdentity api.PeerIdentityType, signature, message []byte) error
	verifyByChannelMutex       sync.RWMutex
	verifyByChannelArgsForCall []struct {
		chainID      common.ChainID
		peerIdentity api.PeerIdentityType
		signature    []byte
		message      []byte
	}
	verifyByChannelReturns struct {
		result1 error
	}
	verifyByChannelReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Verifier) VerifyByChannel(chainID common.ChainID, peerIdentity api.PeerIdentityType, signature []byte, message []byte) error {
	var signatureCopy []byte
	if signature != nil {
		signatureCopy = make([]byte, len(signature))
		copy(signatureCopy, signature)
	}
	var messageCopy []byte
	if message != nil {
		messageCopy = make([]byte, len(message))
		copy(messageCopy, message)
	}
	fake.verifyByChannelMutex.Lock()
	ret, specificReturn := fake.verifyByChannelReturnsOnCall[len(fake.verifyByChannelArgsForCall)]
	fake.verifyByChannelArgsForCall = append(fake.verifyByChannelArgsForCall, struct {
		chainID      common.ChainID
		peerIdentity api.PeerIdentityType
		signature    []byte
		message      []byte
	}{chainID, peerIdentity, signatureCopy, messageCopy})
	fake.recordInvocation("VerifyByChannel", []interface{}{chainID, peerIdentity, signatureCopy, messageCopy})
	fake.verifyByChannelMutex.Unlock()
	if fake.VerifyByChannelStub != nil {
		return fake.VerifyByChannelStub(chainID, peerIdentity, signature, message)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.verifyByChannelReturns.result1
}

func (fake *Verifier) VerifyByChannelCallCount() int {
	fake.verifyByChannelMutex.RLock()
	defer fake.verifyByChannelMutex.RUnlock()
	return len(fake.verifyByChannelArgsForCall)
}

func (fake *Verifier) VerifyByChannelArgsForCall(i int) (common.ChainID, api.PeerIdentityType, []byte, []byte) {
	fake.verifyByChannelMutex.RLock()
	defer fake.verifyByChannelMutex.RUnlock()
	return fake.verifyByChannelArgsForCall[i].chainID, fake.verifyByChannelArgsForCall[i].peerIdentity, fake.verifyByChannelArgsForCall[i].signature, fake.verifyByChannelArgsForCall[i].message
}

func (fake *Verifier) VerifyByChannelReturns(result1 error) {
	fake.VerifyByChannelStub = nil
	fake.verifyByChannelReturns = struct {
		result1 error
	}{result1}
}

func (fake *Verifier) VerifyByChannelReturnsOnCall(i int, result1 error) {
	fake.VerifyByChannelStub = nil
	if fake.verifyByChannelReturnsOnCall == nil {
		fake.verifyByChannelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.verifyByChannelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Verifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.verifyByChannelMutex.RLock()
	defer fake.verifyByChannelMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Verifier) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ acl.Verifier = new(Verifier)
