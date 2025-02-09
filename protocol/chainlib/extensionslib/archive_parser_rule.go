package extensionslib

import (
	spectypes "github.com/lavanet/lava/x/spec/types"
)

type ArchiveParserRule struct {
	extension *spectypes.Extension
}

func (apr ArchiveParserRule) isPassingRule(extensionChainMessage ExtensionsChainMessage, latestBlock uint64) bool {
	requestedBlock := extensionChainMessage.RequestedBlock()
	if requestedBlock < 0 {
		// if asking for the latest block, or an api that doesn't have a specific block requested then it's not archive
		return requestedBlock == spectypes.EARLIEST_BLOCK // only earliest should go to archive
	}
	if latestBlock == 0 {
		return true
	}
	if uint64(requestedBlock) >= latestBlock {
		return false
	}
	if apr.extension.Rule != nil && apr.extension.Rule.Block != 0 {
		if latestBlock-apr.extension.Rule.Block > uint64(requestedBlock) {
			return true
		}
	}
	return false
}
