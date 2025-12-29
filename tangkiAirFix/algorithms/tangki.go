package algorithms

// Asumsi: setiap lantai mengonsumsi 10 liter air
const waterPerFloor = 10

// ================= ITERATIF =================
func TotalWaterIterative(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += waterPerFloor
	}
	return total
}

// ================= REKURSIF =================
func TotalWaterRecursive(n int) int {
	if n == 0 {
		return 0
	}
	return waterPerFloor + TotalWaterRecursive(n-1)
}
