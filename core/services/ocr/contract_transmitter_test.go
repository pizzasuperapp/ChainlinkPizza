package ocr_test

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/smartcontractkit/chainlink/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/core/services/ocr"
	"github.com/smartcontractkit/libocr/gethwrappers/offchainaggregator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ContractTransmitter_ChainID(t *testing.T) {
	chainID := big.NewInt(42)
	contractABI, err := abi.JSON(strings.NewReader(offchainaggregator.OffchainAggregatorABI))
	require.NoError(t, err)
	ct := ocr.NewOCRContractTransmitter(
		testutils.NewAddress(),
		nil,
		contractABI,
		nil,
		nil,
		nil,
		chainID,
	)

	assert.Equal(t, chainID, ct.ChainID())
}
