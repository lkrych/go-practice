// In this area, there are sixteen memory banks; each memory bank can hold any number of blocks. 
// The goal of the reallocation routine is to balance the blocks between the memory banks.
// The reallocation routine operates in cycles. In each cycle, it finds the memory bank with 
// the most blocks (ties won by the lowest-numbered memory bank) and redistributes those blocks among the banks. 
// To do this, it removes all of the blocks from the selected bank, then moves to the next (by index) memory bank and inserts one of the blocks. 
// It continues doing this until it runs out of blocks; if it reaches the last memory bank, it wraps around to the first one.
// The debugger would like to know how many redistributions can be done before a blocks-in-banks configuration is produced that has been seen before.
func memoryBank(bank []int) int {
	
}

func ConvertIntSliceToStringSlice(arr []int) []string {
	//this might be useful ;)
}

func memoryBankAdvent() int {
	//read from memory_bank_input.txt
}