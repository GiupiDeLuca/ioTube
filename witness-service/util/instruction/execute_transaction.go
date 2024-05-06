package instruction

import (
	"fmt"

	solcommon "github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/pkg/bincode"
	soltypes "github.com/blocto/solana-go-sdk/types"
)

type ExecuteTransactionParam struct {
	Governance          solcommon.PublicKey
	Proposal            solcommon.PublicKey
	VoteRecord          solcommon.PublicKey
	RecordTranaction    solcommon.PublicKey
	TransactionAccounts []soltypes.AccountMeta
}

func ExecuteTransaction(programID solcommon.PublicKey, param *ExecuteTransactionParam) soltypes.Instruction {
	data, err := bincode.SerializeData(struct {
		Instruction GovernanceInstructionType
	}{
		Instruction: GovernanceInstructionType(16),
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("ExecuteTransaction data: %v\n", data)

	accounts := []soltypes.AccountMeta{
		{PubKey: param.Governance, IsSigner: false, IsWritable: false},
		{PubKey: param.Proposal, IsSigner: false, IsWritable: false},
		{PubKey: param.VoteRecord, IsSigner: false, IsWritable: false},
		{PubKey: param.RecordTranaction, IsSigner: false, IsWritable: true},
	}

	accounts = append(accounts, param.TransactionAccounts...)

	return soltypes.Instruction{
		ProgramID: programID,
		Accounts:  accounts,
		Data:      data,
	}
}
