package main

import (
	"assignment01bca"
	"fmt"
)

func main() {
	//creating a list of blockchain using Block structs
	blockchain := []*assignment01bca.Block{}
	///initial block
	genesisBlock := assignment01bca.NewBlock(0, "!!! Genesis Block Data !!! ", 0, "")
	blockchain = append(blockchain, genesisBlock)

	block1 := assignment01bca.NewBlock(1, "Murtaza sent x to Redcliff", 123, genesisBlock.Cur_Hash)
	blockchain = append(blockchain, block1)

	block2 := assignment01bca.NewBlock(2, "Dexter sent y to Philips", 456, block1.Cur_Hash)
	blockchain = append(blockchain, block2)

	block3 := assignment01bca.NewBlock(3, "Dexter sent z to Dexter", 456, block2.Cur_Hash)
	blockchain = append(blockchain, block3)

	fmt.Printf("                 ===========================\n                 Showing created Blockchain\n                 ===========================\n")
	assignment01bca.ListBlocks(blockchain)
	isValid := assignment01bca.VerifyChain(blockchain)
	if isValid {
		fmt.Println("Blockcha is verified and is untampered.")
	} else {
		fmt.Println("Blockchain verification failed (:'~').")
	}

	fmt.Println("                 ===================\n                  Changing block 2 \n                 ===================")
	assignment01bca.ChangeBlock(block2, "Bob to Dave")

	// Verify the integrity of the blockchain.
	isValid = assignment01bca.VerifyChain(blockchain)
	if isValid {
		fmt.Println("Blockcha is verified and is untampered.")
	} else {
		fmt.Println("Blockchain verification failed (:'~').")
	}

	fmt.Printf("                 ===========================\n                 Showing tampered Blockchain\n                 ===========================\n")
	assignment01bca.ListBlocks(blockchain)

}
