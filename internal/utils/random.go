package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
)

// RandomGenerator defines the interface for generating random numbers
type RandomGenerator interface {
	GenerateInt(min, max int64) (int64, error)
}

// SystemRandomGenerator uses the system's crypto/rand source
type SystemRandomGenerator struct{}

func (g *SystemRandomGenerator) GenerateInt(min, max int64) (int64, error) {
	if min > max {
		return 0, fmt.Errorf("min (%d) cannot be greater than max (%d)", min, max)
	}
	rangeSize := big.NewInt(max - min + 1)
	n, err := rand.Int(rand.Reader, rangeSize)
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %w", err)
	}
	return min + n.Int64(), nil
}

func NewRandomGenerator() RandomGenerator {
	return &SystemRandomGenerator{}
}

func SampleHybrid(N, k int) ([]int, error) {
	if k < 0 || N < 0 || k > N {
		return nil, errors.New("invalid N or k")
	}
	if k == 0 {
		return []int{}, nil
	}

	ratio := float64(k) / float64(N)
	const (
		smallThreshold = 0.1
		largeThreshold = 0.5
	)

	switch {
	case ratio < smallThreshold:
		return sampleVitterD(N, k)
	case ratio > largeThreshold:
		return sampleShuffle(N, k)
	default:
		return sampleKnuthS(N, k)
	}
}

// Vitter's Algorithm D. https://www.ittc.ku.edu/~jsv/Papers/Vit87.RandomSampling.pdf
func sampleVitterD(N, k int) ([]int, error) {
	result := make([]int, 0, k)
	var nn = N
	var kk = k
	var i = 0
	for kk > 0 {
		vBig, err := rand.Int(rand.Reader, big.NewInt(1<<53))
		if err != nil {
			return nil, fmt.Errorf("failed to generate random number: %w", err)
		}
		V := float64(vBig.Int64()) / float64(1<<53)
		// simple geometric skip (Method A), since we only use when k << N
		gap := int(math.Floor(math.Log(V) / math.Log(1-float64(kk)/float64(nn))))
		i += gap + 1
		if i > N-1 {
			break
		}
		result = append(result, i)
		nn = nn - (gap + 1)
		kk--
	}
	return result, nil
}

// Knuth's Algorithm S. TAOCP Vol.2 Sec.3.4.2
func sampleKnuthS(N, k int) ([]int, error) {
	result := make([]int, 0, k)
	m := 0
	for t := 0; t < N && m < k; t++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(N-t)))
		if err != nil {
			return nil, fmt.Errorf("failed to generate random number: %w", err)
		}
		if nBig.Int64() < int64(k-m) {
			result = append(result, t)
			m++
		}
	}
	return result, nil
}

func sampleShuffle(N, k int) ([]int, error) {
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i
	}

	for i := N - 1; i > 0; i-- {
		jBig, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return nil, fmt.Errorf("failed to generate random number: %w", err)
		}
		j := int(jBig.Int64())
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr[:k], nil
}
